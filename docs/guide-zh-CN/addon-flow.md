## 模块开发流程

目录

- 创建新插件
- 开发
- 调用主模块服务接口
- 访问路径
- 数据迁移



### 创建新插件

1、HotGo 后台进入 开发工具->插件管理->找到创建新插件，根据引导进行创建即可。

```shell
创建成功后默认情况下会在以下目录中生成插件文件，假设新生成的插件名为：hgexample

1. /server/addons/hgexample/                   # 插件模块目录
2. /server/addons/modules/hgexample.go         # 隐式注册插件文件
3. /server/resource/addons/hgexample           # 静态资源和页面模板目录，属于扩展功能选项，勾选对应选项后才会生成
4. /web/src/api/addons/hgexample               # webApi目录
5. /web/src/views/addons/hgexample             # web页面目录

# 创建插件时只生成文件，不立即写菜单。通过插件管理安装后，会自动创建插件一级菜单，并为已存在的配置页及标准模块入口补齐菜单，详见下方规则。
```


2、创建插件完毕重启服务端后，插件管理中会出现你新创建的插件信息。操作栏有几个按钮，在此进行说明
- 安装：先执行 `server/addons/<name>/main.go` 中开发者定义的 `Install`，再由插件管理业务层初始化菜单。钩子和默认菜单写入共用安装事务，菜单冲突或写入失败会回滚安装记录及事务内数据。钩子中的外部调用、文件操作仍需开发者自行保证可恢复。
  - 创建以插件标签命名的一级目录，固定编码为 `addons_<name>`；可明确识别为该插件独占的现有人工一级目录时优先复用。旧生成器为单个业务模块创建的目录不作为插件总目录，原业务菜单位置保持不变。
  - 前端目录从 `hggen.addon.webViewsPath` 读取：存在 `config/system.vue` 时创建配置页菜单、配置读取权限和保存按钮权限。
  - 存在直接子目录 `<module>/index.vue` 时创建页面菜单；对应代码生成记录可提供菜单名称和列表权限。`edit.vue`、`view.vue`、`BasicSetting.vue`、`components` 等子组件不自动登记。任意命名的页面、嵌套页面及其他操作权限，仍通过菜单管理或代码生成器配置，避免误把组件开放为菜单。
  - 已有页面菜单按组件路径复用，保留人工设置的上级、路由、排序和启停状态；名称、编码或同级路由冲突会明确报错，不覆盖。
  - 菜单只自动创建，不向普通角色自动授权。非超级管理员需在角色权限中分配页面和按钮权限。
  - 自动识别页面需要服务端部署对应前端源码；仅部署后端二进制、未携带前端源码时，只确保一级目录，不猜测不存在的页面。
  - 原有菜单 SQL 不会在安装时再次执行，已有业务表和代码生成配置不会被修改。
```go
// Install 安装模块
func (m *module) Install(ctx context.Context) (err error) {
	// ...
	return
}
```

- 更新：先执行 `Upgrade`，再以相同规则补齐缺失菜单；可用于已有插件补上配置页菜单，无需先卸载。已有菜单不重复创建、不覆盖。
```go
// Upgrade 更新模块
func (m *module) Upgrade(ctx context.Context) (err error) {
	// ...
	return
}
```

- 卸载：会自动执行 server/hgexample/main.go 文件中的UnInstall方法，方法中的具体逻辑默认为空，可以根据实际情况自行配置。如会清除所有的数据表和已安装的信息等。
```go
// UnInstall 卸载模块
func (m *module) UnInstall(ctx context.Context) (err error) {
	// ...
	return
}
```



### 插件与代码生成菜单联动

- 先创建插件、重启后端并安装，再在代码生成中选择对应插件模板和插件名称。
- CRUD 和关系树列表在勾选“生成菜单权限”且上级菜单仍为默认根目录时，保存配置、预览和提交生成都会补为插件一级菜单；手工明确选择的非根上级不变。插件一级菜单缺失时会提示先安装或通过“更新配置”补齐，预览不会自行写菜单。
- 取消“生成菜单权限”时保持原行为，不自动开启。主模块模板不受插件联动影响。
- 业务菜单 SQL 仍走生成器原有流程，保留已有 SQL 默认跳过、代码和 service 完成后才导入菜单等规则；安装不重放历史 SQL。
- 安装、更新、卸载与代码生成/删除生成代码共用项目操作锁，防止同机并发修改插件菜单。
- 安装后仍需重启后端使插件路由生效，再刷新页面加载后端菜单；前端生产构建需包含对应页面。卸载仍由原 `UnInstall` 钩子负责清理，本次不额外删除菜单、业务数据或源码。

### 开发

完全可以根据HotGo正常的开发流程去开发对应的API、控制器、业务逻辑、插件内的应用

### 调用主模块服务接口

这里推荐的方式是在插件input层新建一个结构，继承主模块中的input结构。这样做的目的是为了服务与服务之间的输入/输出关系解耦，便于参数扩展和避免插件模块下使用`gf gen service`时出现`import cycle not allowed`。

一个简单的例子：
> 假设hgexample插件模块要通过主模块的服务接口更新插件配置

插件模块input：\server\addons\hgexample\model\input\sysin\config.go
```go
package sysin

import (
	"hotgo/internal/model/input/sysin"
)

// UpdateConfigInp 更新指定配置
type UpdateConfigInp struct {
	sysin.UpdateAddonsConfigInp
}

```

插件模块业务逻辑：\server\addons\hgexample\logic\sys\config.go
```go
package sys

import (
	"context"
	"hotgo/addons/hgexample/global"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/addons/hgexample/service"
	isc "hotgo/internal/service"
)

type sSysConfig struct{}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

// UpdateConfigByGroup 更新指定分组的配置
func (s *sSysConfig) UpdateConfigByGroup(ctx context.Context, in sysin.UpdateConfigInp) error {
	in.UpdateAddonsConfigInp.AddonName = global.GetSkeleton().Name
	return isc.SysAddonsConfig().UpdateConfigByGroup(ctx, in.UpdateAddonsConfigInp)
}

```

主模块input：\server\internal\model\input\sysin\addons_config.go
```go
package sysin

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UpdateAddonsConfigInp 更新指定插件的配置
type UpdateAddonsConfigInp struct {
	AddonName string `json:"addonName"`
	Group     string `json:"group"`
	List      g.Map  `json:"list"`
}

```

主模块业务逻辑：\server\internal\logic\sys\addons_config.go
```go
package sys

import (
	"context"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

type sSysAddonsConfig struct{}

func NewSysAddonsConfig() *sSysAddonsConfig {
	return &sSysAddonsConfig{}
}

func init() {
	service.RegisterSysAddonsConfig(NewSysAddonsConfig())
}

// UpdateConfigByGroup 更新指定分组的配置
func (s *sSysAddonsConfig) UpdateConfigByGroup(ctx context.Context, in sysin.UpdateAddonsConfigInp) error {
	// ...
	return nil
}

```




### 访问路径

#### 后台插件访问路径

```
// IP+端口或域名/admin/插件名称/API路径
如：127.0.0.1:8000/admin/hgexample/index/test
```

- 对应控制器路径：`server/addons/hgexample/controller/admin/sys/index.go`

#### 前端API插件访问路径

```
// IP+端口或域名/api/插件名称/API路径
如：127.0.0.1:8000/api/hgexample/index/test
```

- 对应控制器路径：`server/addons/hgexample/controller/api/index.go`

#### 前台页面插件访问路径

```
// IP+端口或域名/home/插件名称/API路径
如：127.0.0.1:8000/home/hgexample/index/test
```

- 对应控制器路径：`server/addons/hgexample/controller/home/index.go`
- 对应模板路径：`server/addons/hgexample/resource/public/template`

#### 静态资源插件访问路径

```
// IP+端口或域名/home/插件名称/API路径
如：127.0.0.1:8000/addons/hgexample/default
```

- 对应资源路径：`server/addons/hgexample/resource/public`


#### Websocket插件访问路径

```
// IP+端口或域名/socket/插件名称/API路径
如：127.0.0.1:8000/socket/hgexample/index/test
```

- 对应控制器路径：`server/addons/hgexample/controller/socket/index.go`


### 数据迁移

- 可以将数据迁移逻辑写进server/xxx插件/main.go 文件中的Install方法中，并遵循系统规范进行数据安装
