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
- **接口说明**: 用于前端网站获取网站基本配置信息。

#### 2.1.1 请求参数
*(无必填参数)*

#### 2.1.2 响应数据 (Response)
```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "siteConfig": {
      "siteSystemOpen": 1,
      "siteCloseText": "网站维护中，暂时无法访问！本网站正在进行系统维护和技术升级，网站暂时无法访问，敬请谅解！",
      "siteLogo": "/images/logo.png",
      "siteName": "ThinkCT",
      "siteTitle": "ThinkCT - 专注于企业网站建设、软件开发、数字化转型",
      "siteKeywords": "企业网站建设,软件开发,数字化转型,ThinkCT",
      "siteDescription": "ThinkCT 专注于为企业提供高质量的网站建设、定制软件开发和全面的数字化转型解决方案...",
      "siteCopyRight": "copyright © 2026 ThinkCT all rights reserved.",
      "siteIcp": "京ICP备2025112105号",
      "siteEmail": "66983239@qq.com",
      "sitePhone": "185-1000-3363",
      "siteWechatQrCode": "",
      "siteMpQrCode": ""
    }
  }
}
```

#### 2.1.3 字段数据来源说明
- **`siteConfig`**: 网站配置信息，由后台系统设置动态读取：
  - `basicSystemOpen` (bool): 网站开启状态（`true`: 正常开放，`false`: 网站关闭/维护）
  - `basicCloseText` (string): 网站关闭/维护时的提示文本
  - `logo` (string): 站点 Logo
  - `siteName` (string): 站点名称
  - `siteTitle` (string): 站点标题
  - `siteKeywords` (string): 站点关键词 (逗号隔开)
  - `siteDescription` (string): 站点描述
  - `siteCopyRight` (string): 版权所有信息
  - `siteIcp` (string): ICP 备案号
  - `email` (string): 邮箱地址
  - `phone` (string): 联系电话
  - `wechatQrCode` (string): 微信二维码图片地址
  - `mpQrCode` (string): 公众号二维码图片地址

### 2.2 获取首页聚合数据
- **接口路径**: `/api/site/home` 
- **请求方式**: `GET`
- **接口说明**: 用于前端网站首页展示产品中心、经典案例和新闻中心等数据。

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
  - `limit` (int): 每页数量 (可选，传入栏目ID时生效，默认 10)
- **说明**: 
  1. 若不传 `channel`：返回状态为启用的前台所有栏目数据列表（按排序正序、ID升序排列），统一放在 `list` 节点，可用于渲染前台的导航栏。
  2. 若传入指定的栏目 `channel`，后端会根据该栏目的 `model` 字段自动判断类型（如 `"article"` 或 `"products"`），在返回体中的 `list` 节点输出对应的文章或产品分页数据，同时在 `channel` 节点输出当前访问的栏目详情。
  3. **支持递归包含子频道数据**：当传入指定栏目时，系统会自动递归获取该栏目下的所有子栏目（多级），在查询列表（`list`）、总数（`total`）以及热门列表（`hotList`）时，会将当前栏目及其所有子栏目 ID 集合通过 `channel_id IN (...)` 聚合查询；同时列表项的 `url` 字段会自动根据内容实际所属的子栏目别名生成正确的链接（形如 `/{sub_diyname}/detail/{id}`）。
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
       "total": 100,
       "page": 1,
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
  - `limit` (int, 可选): 每页条数，默认为 16
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
      "total": 109,
      "page": 1,
      "limit": 16
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
