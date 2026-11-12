# 前端前台接口开发文档

本文档用于记录前台网站（用户端）的接口开发结构及现有接口说明。

## 1. 分层架构与目录说明
根据 HotGo 的规范，前台 API 的开发严格遵循分层架构，相关代码分布在以下目录：

- **API 契约层 (api)**: `server/api/api/site/v1/`
  *定义 HTTP 路由、Req 请求结构体、Res 响应结构体以及 Swagger 注释。*
- **输入输出层 (model/input)**: `server/internal/model/input/sitein/`
  *定义业务逻辑中流转的 Inp 入参和 Model 出参。*
- **接口声明 (service)**: `server/internal/service/api_site.go`
  *声明业务逻辑接口，供 Controller 调用。*
- **业务逻辑层 (logic)**: `server/internal/logic/site/`
  *实现具体的业务逻辑，如查库、组装数据等，禁止出现 HTTP 相关代码。*
- **控制器层 (controller)**: `server/internal/controller/api/site/`
  *作为薄薄的一层，仅负责接收参数、调用 Service 逻辑并将结果响应给前端。*
- **路由注册 (router)**: `server/internal/router/api.go`
  *统一注册前台路由，视情况使用 `ApiAuth` 中间件进行权限拦截。*

---

## 2. 接口列表

### 2.1 获取首页综合数据
- **接口路径**: `/api/site/index` 
- **请求方式**: `GET`
- **接口说明**: 用于前端网站获取网站基本配置信息与头部多级导航。两者在同一次请求中返回，`navMenu` 与 `siteConfig` 同级，均位于 `data` 下。

#### 2.1.1 请求参数
*(无必填参数)*

#### 2.1.2 响应数据 (Response)
> 以下为节选：`navMenu` 仅保留部分节点用于说明结构，其中包含一条 4 级链路：产品 → 机房与控制系统 → 关键电源管理 → 交流不间断电源（UPS）。

```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "siteConfig": {
      "siteSystemOpen": 1,
      "siteCloseText": "网站维护中，暂时无法访问！本网站正在进行系统维护和技术升级，网站暂时无法访问，敬请谅解！",
      "siteFavicon": "/uploads/favicon.ico",
      "siteLogo": "/images/logo.png",
      "siteName": "ThinkCT",
      "siteTitle": "ThinkCT - 专注于企业网站建设、软件开发、数字化转型",
      "siteKeywords": "企业网站建设,软件开发,数字化转型,ThinkCT",
      "siteDescription": "ThinkCT 专注于为企业提供高质量的网站建设、定制软件开发和全面的数字化转型解决方案...",
      "siteCopyRight": "copyright © 2026 ThinkCT all rights reserved.",
      "siteIcp": "京ICP备2025112105号",
      "sitePoliceRecord": "京公网安备11010802000000号",
      "siteEmail": "66983239@qq.com",
      "sitePhone": "185-1000-3363",
      "siteWechatQrCode": "",
      "siteMpQrCode": ""
    },
    "navMenu": [
      {
        "id": 1,
        "name": "产品",
        "url": "/products",
        "image": "",
        "imageAlt": "",
        "children": [
          {
            "id": 30,
            "name": "机房与控制系统",
            "url": "/products/control-systems",
            "image": "",
            "imageAlt": "",
            "children": [
              {
                "id": 34,
                "name": "关键电源管理",
                "url": "/products",
                "image": "",
                "imageAlt": "",
                "children": [
                  {
                    "id": 35,
                    "name": "交流不间断电源（UPS）",
                    "url": "/ups-power-monitor",
                    "image": "",
                    "imageAlt": "",
                    "children": []
                  }
                ]
              }
            ]
          }
        ]
      },
      {
        "id": 48,
        "name": "服务",
        "url": "/services",
        "image": "//vcdon.oss-cn-beijing.aliyuncs.com/images/services-featured.jpg",
        "imageAlt": "",
        "children": [
          {
            "id": 49,
            "name": "服务内容",
            "url": "/services-details",
            "image": "",
            "imageAlt": "",
            "children": [
              {
                "id": 50,
                "name": "咨询服务",
                "url": "/services-details",
                "image": "",
                "imageAlt": "",
                "children": []
              }
            ]
          }
        ]
      }
    ]
  }
}
```

无可展示导航时，`navMenu` 返回空数组（不会是 `null`）：
```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "siteConfig": { "...": "..." },
    "navMenu": []
  }
}
```

#### 2.1.3 字段数据来源说明
- **`siteConfig`**: 网站配置信息，由后台系统设置动态读取：
  - `basicSystemOpen` (bool): 网站开启状态（`true`: 正常开放，`false`: 网站关闭/维护）
  - `basicCloseText` (string): 网站关闭/维护时的提示文本
  - `siteFavicon` (string): 网站 favicon.ico 图标地址
  - `logo` (string): 站点 Logo
  - `siteName` (string): 站点名称
  - `siteTitle` (string): 站点标题
  - `siteKeywords` (string): 站点关键词 (逗号隔开)
  - `siteDescription` (string): 站点描述
  - `siteCopyRight` (string): 版权所有信息
  - `siteIcp` (string): ICP 备案号
  - `sitePoliceRecord` (string): 公安备案号
  - `email` (string): 邮箱地址
  - `phone` (string): 联系电话
  - `wechatQrCode` (string): 微信二维码图片地址
  - `mpQrCode` (string): 公众号二维码图片地址
- **`navMenu`** (array): 头部多级导航树，与 `siteConfig` 同级，同一次请求返回。
  - 节点字段：
    - `id` (int64): 导航ID
    - `name` (string): 导航名称
    - `url` (string): 链接地址
    - `image` (string): 导航配图地址，未配置时为 `""`
    - `imageAlt` (string): 配图替代文本，未配置时为 `""`
    - `children` (array): 子节点，结构与本节点相同；叶子节点固定为 `[]`，不会省略或为 `null`
  - 数据来源：`hg_web_nav_menu` 表（网站_导航菜单）。导航独立于栏目（`hg_web_channel`）维护，栏目的增删、启用/禁用不会影响导航，反之亦然。目前未提供后台管理页面，导航数据直接在该表维护。
  - 展示规则：
    - 只取 `status = 1` 且未软删除（`deleted_at` 为空）的记录；
    - 只展示从根节点（`pid = 0`）逐级可达的节点：父节点被禁用或软删除时，其整棵子树一并隐藏，不会提升为顶级；
    - 父节点不存在的孤儿节点、互为父子的闭环节点、自引用节点都不展示；
    - 同级排序：`sort` 升序，`sort` 相同时按 `id` 升序；
    - 没有可展示的导航时返回 `[]`；读取导航表失败时整个接口返回错误（提示"获取网站导航失败"）。
  - 初始化：
    - 初始化 SQL：`server/storage/data/generate/web_nav_menu.sql`。可重复执行：表不存在时才创建；只有表完全为空（含已软删除的行）时才写入初始数据，不覆盖已有数据。
    - 初始化脚本：`server/hack/exec_web_nav_menu_schema.go`，在 `server/` 目录下运行：
      - 不带参数：只读巡检，输出表结构、索引、行数、根节点、孤儿/层级/关系树一致性检查及内容摘要，不改动任何数据；
      - `-apply`：执行上述 SQL（建表 + 空表写入初始数据）。执行前会校验 SQL 文件只包含预期的建表与写入两条语句，写入前校验现有列结构一致，否则拒绝执行；
      - `-generate`：仅针对 `hg_web_nav_menu` 调用 `gf gen dao` 生成 dao/entity/do（读取 `hack/config.yaml`），不能与 `-apply` 同时使用。
    - 现有菜单初始数据：来自 `next_thinkct/config/NavMenu.ts` 的全部有效节点（已排除被注释的节点），共 100 项、5 个根节点（产品、服务、解决方案与行业、支持、新闻中心），ID 为 1~100 且唯一（原文件中的重复 ID 已按前序遍历重新分配），`pid`/`level`/`tree` 依新 ID 重建，最深 4 级；`sort` 保留同级原始顺序，URL 保持原样。
  - 前端使用（vcdon 主题）：`next_thinkct/config/Site.ts` 的 `getSiteConfig()` 在同一次 `/api/site/index` 请求中读取 `data.navMenu`，挂到 `siteConfig.navMenu`（非数组时按 `[]` 处理），复用站点配置缓存（开发环境 3 秒、生产环境 120 秒重新验证），因此导航数据修改后生产环境最长约 120 秒生效。vcdon 的头部组件读取 `siteConfig.navMenu` 渲染桌面与移动端菜单，空数组时不渲染菜单项。

### 2.2 获取首页聚合数据
- **接口路径**: `/api/site/home` 
- **请求方式**: `GET`
- **接口说明**: 用于前端网站首页展示首页轮播、产品中心、经典案例和新闻中心等数据。

#### 2.2.1 请求参数
- `productsChannelIds` (string): 产品分类栏目ID组或自定义名称(diyname)，英文逗号分隔 (可选，默认 `7,8,9,10,11`)
- `productsLimit` (int): 每个产品分类显示的产品数量 (可选，默认 `4`)
- `casesChannelIds` (string): 案例栏目ID组或自定义名称(diyname)，英文逗号分隔 (可选，默认 `16` / `cases`)
- `casesLimit` (int): 每个案例分类显示的数量 (可选，默认 `3`)
- `newsChannelIds` (string): 新闻栏目ID组或自定义名称(diyname)，英文逗号分隔 (可选，默认 `2` / `news`)
- `newsLimit` (int): 每个新闻分类显示的数量 (可选，默认 `3`)

#### 2.2.2 响应数据 (Response)
```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "carousels": [
      {
        "id": 1,
        "title": "追求无限可能",
        "description": "通过我们的数据驱动解决方案实现生产率和可持续发展目标。",
        "ctaText": "查看我们的实现方式",
        "ctaLink": "/products",
        "color": "emerald-blue",
        "image": "/images/vcdon/slides1.jpg",
        "alt": "服务器机房中手持 iPad 的女性工作人员",
        "aspectRatio": "16/9"
      },
      {
        "id": 2,
        "title": "第十版年度智能制造现状",
        "description": "为您的数字战略制定基准，探索全球杰出范例，并揭示 2025 年的投资趋势。",
        "ctaText": "获取报告",
        "ctaLink": "/news_hy/detail/1",
        "color": "purple-cyan",
        "image": "/images/vcdon/slides2.jpg",
        "alt": "霓虹蓝绿波",
        "aspectRatio": "3000/884"
      }
    ],
    "productsChannels": [
      {
        "id": 2,
        "name": "UPS电源",
        "model": "products",
        "list": [
          {
            "id": 1,
            "title": "产品标题",
            "image": "产品图片URL"
          }
        ]
      }
    ],
    "casesChannels": [
      {
        "id": 16,
        "name": "应用案例",
        "diyname": "cases",
        "model": "article",
        "list": [
          {
            "id": 8,
            "title": "案例标题",
            "image": "案例图片URL"
          }
        ]
      }
    ],
    "newsChannels": [
      {
        "id": 2,
        "name": "新闻中心",
        "diyname": "news",
        "model": "article",
        "list": [
          {
            "id": 1,
            "title": "新闻标题",
            "publishTime": 1712211234
          }
        ]
      }
    ],
    "links": [
      {
        "id": 1,
        "title": "百度一下",
        "link": "https://www.baidu.com",
        "target": "_blank",
        "logo": "http://.../logo.png",
        "weigh": 1,
        "status": "enable",
        "remark": "",
        "updateTime": 1753173959,
        "createTime": 1753173953
      }
    ]
  }
}
```

#### 2.2.3 字段数据来源说明
- **`carousels`** (array): 首页轮播列表，与 `productsChannels` 等字段同级，始终返回数组；没有可展示的轮播时返回 `[]`，不会为 `null`。上方示例仅列出 2 条，实际按库中启用数据全部返回。
  - 每项仅包含前台展示字段，不返回 `status`、`sort`、`createdAt`、`updatedAt`、`deletedAt` 等管理字段：
    - `id` (int): 轮播ID
    - `title` (string): 标题
    - `description` (string): 描述文案
    - `ctaText` (string): 按钮文字
    - `ctaLink` (string): 按钮链接（站内路径，原样返回）
    - `color` (string): 配色主题标识，如 `emerald-blue`、`purple-cyan`、`orange-pink`
    - `image` (string): 图片地址
    - `alt` (string): 图片替代文本
    - `aspectRatio` (string): 图片宽高比，如 `16/9`、`3000/884`
  - 数据来源：`hg_web_carousel` 表（网站_首页轮播），表字段 `cta_text`/`cta_link`/`aspect_ratio` 分别对应 `ctaText`/`ctaLink`/`aspectRatio`，其余同名。目前未提供后台管理页面，轮播数据直接在该表维护。
  - 展示规则：
    - 只取 `status = 1`（启用）且未软删除（`deleted_at` 为空）的记录；
    - 排序：`sort` 升序，`sort` 相同时按 `id` 升序；
    - 读取轮播表失败时整个接口返回错误（提示"获取首页轮播图失败"）。
  - 初始化：
    - 初始化 SQL：`server/storage/data/generate/web_carousel.sql`。可重复执行：表不存在时才创建；只有表完全为空（含已软删除的行）时才写入初始数据，不覆盖已有数据。
    - 初始化脚本：`server/hack/exec_web_carousel_schema.go`，在 `server/` 目录下运行：
      - `go run hack/exec_web_carousel_schema.go`（不带参数）：只读巡检，输出表是否存在、列结构及按 `sort,id` 排列的现有行，不改动任何数据；
      - `go run hack/exec_web_carousel_schema.go -apply`：执行上述 SQL（建表 + 空表写入初始数据）。执行前会校验 SQL 文件只包含预期的建表与写入两条语句，否则拒绝执行；表中已有数据时跳过写入；
      - `go run hack/exec_web_carousel_schema.go -generate`：仅针对 `hg_web_carousel` 调用 `gf gen dao` 生成 dao/entity/do（读取 `hack/config.yaml`），需与 `-apply` 分开执行，不能同时使用。
    - 初始数据：来自 `next_thinkct/components/templates/vcdon/home/Carousel.tsx` 原静态 `slides`，共 3 条（ID 1~3，`sort` 1~3，均为启用），字段值保持原样。
  - 前端使用（vcdon 主题）：首页复用现有的服务端 `/api/site/home` 请求读取 `data.carousels`，不额外发请求、不额外加缓存；非数组时按空处理。0 条时不渲染轮播区域，1 条时静态展示（不显示进度条与播放控制、不自动切换），多条时每 5 秒自动切换。
- **`productsChannels`**: 取 `web_channel` 表中 ID 包含在 `productsChannelIds` 参数中的分类记录（支持别名），按 `weigh` 与 `id` 排序返回。并包含子频道的关联产品列表数据输出到 `list` 节点中（限制 `productsLimit` 条）。
- **`casesChannels`**: 根据 `casesChannelIds` 查询对应的栏目详情，并循环关联查询其下及子频道带有图片的经典案例文章列表，按 `weigh` 与 `id` 倒序排列（限制 `casesLimit` 条）。
- **`newsChannels`**: 根据 `newsChannelIds` 查询对应的栏目详情，并循环关联查询其下及子频道的最新新闻列表，按发布时间 `publish_time` 与 `id` 倒序排列（限制 `newsLimit` 条）。
- **`links`**: 友情链接列表，取自 `hg_web_links` 表中状态为启用（`status='enable'`）的记录，按权重 `weigh desc, id asc` 排序输出。字段包含：`id`、`title`、`link`、`target`（跳转目标 `_blank`/`_self`）、`logo`、`weigh` 等。


### 2.3 获取栏目列表及详情数据
- **接口路径**: `/api/site/channel`
- **请求方式**: `GET`
- **请求参数**:
  - `channel` (string): 栏目ID或自定义名称(diyname) (可选，传空则只返回栏目导航列表；**若传入栏目ID或自定义名称，则会同时返回该栏目的基本信息及其关联模型的列表数据**)
  - `page` (int): 当前页码 (可选，传入栏目ID时生效，默认 1)
  - `pageSize` (int): 每页数量 (可选，传入栏目ID时生效，默认 10)
- **说明**: 
  1. 若不传 `channel`：返回状态为启用的前台所有栏目数据列表（按排序正序、ID升序排列），统一放在 `list` 节点，可用于渲染前台的导航栏。
  2. 若传入指定的栏目 `channel`，后端会根据该栏目的 `model` 字段自动判断类型（如 `"article"` 或 `"products"`），在返回体中的 `list` 节点输出对应的文章或产品分页数据，同时在 `channel` 节点输出当前访问的栏目详情。
  3. **支持递归包含子频道数据**：当传入指定栏目时，系统会自动递归获取该栏目下的所有子栏目（多级），在查询列表（`list`）、总数（`totalCount`）以及热门列表（`hotList`）时，会将当前栏目及其所有子栏目 ID 集合通过 `channel_id IN (...)` 聚合查询；同时列表项的 `url` 字段会自动根据内容实际所属的子栏目别名生成正确的链接（形如 `/{sub_diyname}/detail/{id}`）。
  4. `breadcrumb` 为从根栏目（顶级栏目）到当前访问栏目的层级数组，便于前端直接渲染面包屑导航路径。
- **响应体示例**:
  ```json
  {
    "code": 0,
    "message": "操作成功",
    "data": {
      "channel": {
        "id": 7,
        "pid": 1,
        "name": "UPS电源",
        "diyname": "ups-power",
        "model": "products",
        "type": "channel",
        "url": "/ups-power"
      },
      "breadcrumb": [
        {
          "id": 1,
          "pid": null,
          "name": "产品中心",
          "diyname": "products",
          "model": "products",
          "type": "channel",
          "url": "/products"
        },
        {
          "id": 7,
          "pid": 1,
          "name": "UPS电源",
          "diyname": "ups-power",
          "model": "products",
          "type": "channel",
          "url": "/ups-power"
        }
      ],
      "list": [ ...这里输出栏目列表，或者指定频道时的文章/产品数据... ],
      "page": 1,
      "pageSize": 10,
      "pagination": false,
      "pageCount": 1,
      "totalCount": 4,
      "relatedList": [ ...同级相关分类列表... ],
       "hotList": [ ...热门文章/产品列表... ]
    }
  }
  ```

### 2.4 获取文章/产品详情内容
- **接口路径**: `/api/site/content`
- **请求方式**: `GET`
- **请求参数**:
  - `id` (int): 内容ID (必填)
  - `channel` (string): 栏目ID或自定义名称(diyname) (必填)
- **说明**:
  1. 通过传入的 `channel` 自动判断内容的类型（`article` 还是 `products`），并根据 `id` 获取对应的内容详情。
  2. 接口会自动让被访问内容的浏览量 (`views`) + 1。
  3. 响应中包含内容主体 `content`，同时也会关联返回其所属的栏目详情 `channel`。文章和产品的 `url` 字段格式为 `/{diyname}/detail/{内容ID}`。
  4. 当内容为文章时，响应中的 `author` 返回作者公开资料，包括 `id`、`username`、`nickname`、`avatar`、`motto`、`isAdmin`、`articleCount`、`viewCount` 和 `followerCount`。`followerCount` 当前固定为 `0`，等待关注功能启用后提供实际统计。产品内容的 `author` 为 `null`。
  5. `breadcrumb` 为从根栏目（顶级栏目）到当前访问栏目的层级数组，便于前端直接渲染面包屑导航路径。
  6. 响应中包含 `prev` (上一篇) 和 `next` (下一篇) 数据。
  7. 响应中包含 `relatedList` (同级分类) 和 `hotList` (热门内容) 列表。
  8. 如果栏目不存在/已禁用，或者内容不存在/被下架，接口会直接返回报错。
- **响应体示例**:
  ```json
  {
    "code": 0,
    "message": "操作成功",
    "data": {
      "content": {
        "id": 1,
        "title": "这是一篇文章",
        "content": "<p>文章的富文本内容...</p>",
        "views": 100,
        "url": "/news/detail/1"
      },
      "author": {
        "id": 1,
        "username": "tigshop",
        "nickname": "Tigshop",
        "avatar": "https://example.com/avatar.png",
        "motto": "",
        "isAdmin": true,
        "articleCount": 239,
        "viewCount": 797100,
        "followerCount": 0
      },
      "channel": {
        "id": 1,
        "name": "公司动态",
        "diyname": "news",
        "model": "article",
        "type": "channel",
        "url": "/news"
      },
      "breadcrumb": [
        {
          "id": 1,
          "pid": null,
          "name": "公司动态",
          "diyname": "news",
          "model": "article",
          "type": "channel",
          "url": "/news"
        }
      ],
      "prev": { ...上一篇信息... },
      "next": { ...下一篇信息... },
      "relatedList": [ ...同级分类... ],
      "hotList": [ ...热门内容... ]
    }
  }
  ```

---

### 2.5 全站内容搜索 (`/api/site/search`)

- **请求方式**: `GET`
- **说明**: 根据关键词全站搜索产品或文章，默认搜索产品（`type=products`）。支持分页，并自动补全关联频道的详情页 `url` 和封面图。
- **请求参数 (Query)**:
  - `keywords` (string, 必须): 搜索关键词，为空时接口返回错误提示 `请输入搜索关键词`
  - `type` (string, 可选): 搜索模型，支持 `products`（产品，默认）或 `article`（文章）
  - `page` (int, 可选): 页码，默认为 1
  - `pageSize` (int, 可选): 每页条数，默认为 16
- **响应示例**:
  ```json
  {
    "code": 0,
    "message": "ok",
    "data": {
      "list": [
        {
          "id": 139,
          "channelId": 14,
          "title": "UPS电源",
          "description": "高品质UPS电源产品介绍...",
          "image": "https://...",
          "url": "/santak-ups/detail/139",
          "weigh": 100,
          "status": 1
        }
      ],
      "page": 1,
      "pageSize": 16,
      "pagination": true,
      "pageCount": 7,
      "totalCount": 109
    }
  }
  ```

---

### 2.6 在线留言接口 (`/api/site/feedback`)

- **请求方式**: `POST`
- **数据格式**: `application/json`
- **说明**: 前台用户提交在线留言咨询信息，自动保存至 `hg_web_feedback` 数据表，并记录用户客户端 IP 及提交时间。
- **请求参数 (JSON Body)**:
  - `name` (string, 必填): 姓名 / 联系人
  - `phone` (string, 必填): 手机号 / 联系电话
  - `email` (string, 可选): 电子邮箱
  - `company` (string, 可选): 公司名称
  - `message_topic` (string, 可选): 留言主题
  - `content` (string, 必填): 留言内容 / 需求描述
  - `expected_response_time` (string, 可选): 期望回复时间（例如：`尽快回复`、`1个工作日内` 等）
- **请求示例**:
  ```json
  {
    "name": "张先生",
    "phone": "13800138000",
    "email": "test@qq.com",
    "company": "某某科技有限公司",
    "message_topic": "网站建设咨询",
    "content": "想咨询一下企业官网定制开发服务与报价",
    "expected_response_time": "尽快回复"
  }
  ```
- **响应示例**:
  ```json
  {
    "code": 0,
    "message": "操作成功",
    "data": {},
    "timestamp": 1789453952,
    "traceID": "44ef5c83e36ad518bba11a1574a2b97b"
  }
  ```

---

## 3. 开发注意事项

3. **安全规范**: 遇到异常不要把详细 DB Error 抛给前端，使用 `gerror.New("友好提示信息")`。
