# 前台用户模块开发文档

本文档描述了 HotGo 项目前台用户模块（API端）的结构、开发规范以及提供的接口信息。主要包含用户注册、登录、获取用户信息和退出登录功能。

## 1. 目录结构

所有与前台用户相关的代码均遵循 GoFrame 和 HotGo 的规范，主要分布在以下目录中：

```text
p:\go\hotgo-2.0\server
├── api
│   └── api
│       └── user
│           ├── v1
│           │   └── user.go       // 定义 API 的请求参数 (Req) 和响应结果 (Res) 结构体
│           └── user.go           // 声明前台用户的 Controller 接口 (IUserV1)
├── internal
│   ├── controller
│   │   └── api
│   │       └── user
│   │           ├── user.go       // 控制器实现，主要负责接收请求并转发给 Service 层
│   │           └── user_new.go   // 控制器实例化文件
│   ├── logic
│   │   └── user
│   │       └── api_user.go       // 具体业务逻辑实现（注册、登录、信息获取、退出等），依赖于底层的 dao 和 library
│   ├── service
│   │   └── user.go               // 定义了供其他模块调用的服务接口 (IApiUser) 和注册方法
│   └── router
│       └── api.go                // 路由注册文件。免鉴权的接口放在中间件之前，需要鉴权的接口放在中间件之后
└── manifest
    └── config
        └── config.yaml           // 配置文件，其中的 `api.exceptLogin` 节点用于配置免登录拦截的路由白名单
```

---

## 2. 接口信息概览

> **说明**：所有接口在请求时如果涉及登录后的操作（例如获取信息、退出登录），均需要在请求头中携带 Token。
> 
> **Token 传递方式**：
> - 请求头 (Header): `Authorization: Bearer <Your_Token>` 或 `userToken: <Your_Token>`
> - URL 参数 (Query): `?authorization=<Your_Token>` 或 `?userToken=<Your_Token>`

### 2.1 用户注册

* **接口路径**：`/api/user/register`
* **请求方式**：`POST`
* **鉴权要求**：无需鉴权（已加入配置白名单）
* **功能描述**：用于新用户注册，注册成功后自动登录并返回 Token 和用户信息。

**请求参数 (Body-JSON)**：
| 字段名 | 类型 | 必填 | 说明 |
| :--- | :--- | :--- | :--- |
| username | string | 是 | 用户名 |
| password | string | 是 | 密码 (6-20位) |
| nickname | string | 否 | 昵称 |
| mobile | string | 否 | 手机号 |
| email | string | 否 | 邮箱 |

**响应结果 (JSON)**：
| 字段名 | 类型 | 说明 |
| :--- | :--- | :--- |
| id | int64 | 用户ID |
| username | string | 用户名 |
| nickname | string | 昵称 |
| avatar | string | 头像 |
| mobile | string | 手机号 |
| email | string | 邮箱 |
| money | uint | 余额 |
| score | uint | 积分 |
| token | string | 登录令牌 |
| expire | int64 | Token 过期时间戳 |

---

### 2.2 用户登录

* **接口路径**：`/api/user/login`
* **请求方式**：`POST`
* **鉴权要求**：无需鉴权（已加入配置白名单）
* **功能描述**：用户使用账号和明文密码登录，验证通过后返回 Token 和详细信息。

**请求参数 (Body-JSON)**：
| 字段名 | 类型 | 必填 | 说明 |
| :--- | :--- | :--- | :--- |
| username | string | 是 | 用户名 |
| password | string | 是 | 密码 (明文) |

**响应结果 (JSON)**：
*(同注册接口的响应字段一致，包含 Token 和用户信息)*

---

### 2.3 获取用户信息

* **接口路径**：`/api/user/info`
* **请求方式**：`GET`
* **鉴权要求**：**必须携带有效 Token**
* **功能描述**：通过 Token 解析用户身份，查询并返回最新的用户信息。

**请求参数**：无

**响应结果 (JSON)**：
| 字段名 | 类型 | 说明 |
| :--- | :--- | :--- |
| id | int64 | 用户ID |
| username | string | 用户名 |
| nickname | string | 昵称 |
| avatar | string | 头像 |
| mobile | string | 手机号 |
| email | string | 邮箱 |
| money | uint | 余额 |
| score | uint | 积分 |

---

### 2.4 修改用户信息

* **接口路径**：`/api/user/update`
* **请求方式**：`POST`
* **鉴权要求**：**必须携带有效 Token**
* **功能描述**：修改当前登录用户的资料信息（昵称、头像、手机号、邮箱、性别、生日、签名）。用户名和密码不通过此接口修改。

**请求参数 (Body-JSON)**：
| 字段名 | 类型 | 必填 | 说明 |
| :--- | :--- | :--- | :--- |
| nickname | string | 是 | 昵称 |
| avatar | string | 是 | 头像 |
| mobile | string | 否 | 手机号（需符合手机号格式，且不能与其他用户重复） |
| email | string | 否 | 邮箱（需符合邮箱格式，且不能与其他用户重复） |
| gender | int | 否 | 性别：0=未知，1=男，2=女 |
| birthday | string | 否 | 生日 |
| motto | string | 否 | 签名 |

**响应结果 (JSON)**：
| 字段名 | 类型 | 说明 |
| :--- | :--- | :--- |
| id | int64 | 用户ID |
| username | string | 用户名 |
| nickname | string | 昵称 |
| avatar | string | 头像 |
| mobile | string | 手机号 |
| email | string | 邮箱 |
| gender | int | 性别 |
| birthday | string | 生日 |
| motto | string | 签名 |
| money | uint | 余额 |
| score | uint | 积分 |

---

### 2.5 退出登录

* **接口路径**：`/api/user/logout`
* **请求方式**：`POST`
* **鉴权要求**：**必须携带有效 Token**
* **功能描述**：注销当前的 Token，使其失效。

**请求参数**：无

**响应结果 (JSON)**：
操作成功则返回标准成功格式，无额外业务数据。

---

## 3. 开发注意事项

1. **密码加密**：当前为了方便 API 调试工具测试，密码验证直接使用了 `MD5(明文密码 + 数据库Salt)`。如果前端进行了加密传输（如 HotGo 默认的 AES 加密），需要在 `api_user.go` 中改回使用 `simple.CheckPassword()` 校验。
2. **状态验证**：登录和获取信息时均对 `user` 表的 `status` 字段进行了校验，要求其值为 `"1"` 或 `"enable"`。
3. **白名单配置**：如果在 `api` 路由组下新增不需要鉴权的接口，必须将其路由地址添加到 [manifest/config/config.yaml](file:///p:/go/hotgo-2.0/server/manifest/config/config.yaml) 文件的 `api.exceptLogin` 列表中。
4. **Token 鉴权机制**：由于在 `token.go` 中对 `GetAuthorization` 进行了调整，目前支持 `Authorization` 或 `userToken` 字段传递令牌。