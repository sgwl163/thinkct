# 前台用户接口文档

本文档说明 HotGo API 端的用户中心接口。所有接口统一返回：

```json
{"code":0,"message":"操作成功","timestamp":1234567890,"traceID":"xxx","data":{}}
```

接口基础路径为 `/api`。除明确标注为免登录的接口外，请在请求头中携带访问令牌：

```text
Authorization: Bearer <token>
```

也兼容 `userToken: <token>` 请求头，或 `authorization`、`userToken` 查询参数。

## 用户资料字段

用户信息接口中的主要字段如下：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| id | int64 | 用户 ID |
| username | string | 用户名，不支持通过用户中心修改 |
| nickname | string | 昵称 |
| avatar | string | 头像地址 |
| mobile | string | 已绑定手机号 |
| email | string | 已绑定邮箱 |
| gender | uint | 性别：0=未知、1=男、2=女 |
| birthday | datetime | 生日，可为空 |
| motto | string | 签名 |
| money | uint | 余额 |
| score | uint | 积分 |

## 账号与登录

### 注册

- 路径：`POST /api/user/register`
- 鉴权：免登录
- 参数：`username`、`password` 必填；`nickname`、`mobile`、`email` 可选。
- 结果：注册成功后自动登录，返回 `token`、`expire`、`tokenExpireAt` 以及用户基本信息。

`expire` 是初始有效时长（秒），`tokenExpireAt` 是当前绝对过期时间（Unix 秒）。

`password` 长度必须为 6 至 20 位；用户名、手机号和邮箱不能重复。

### 登录

- 路径：`POST /api/user/login`
- 鉴权：免登录
- 参数：`username`、`password`，均必填。
- `username` 支持用户名、手机号或邮箱。
- 结果：返回 `token`、`expire`、`tokenExpireAt` 以及用户基本信息。

后端采用服务端可撤销会话并支持滑动续期，不另发 refresh token。受保护接口响应头 `X-Token-Expire-At` 返回续期后的绝对过期时间（Unix 秒）。

成功、密码错误和账号禁用的登录尝试都会写入用户登录历史。账号状态为 `enable` 或 `1` 时允许登录。

### 获取当前用户信息

- 路径：`GET /api/user/info`
- 鉴权：需要登录
- 参数：无
- 结果：返回最新的用户资料字段。

### 退出登录

- 路径：`POST /api/user/logout`
- 鉴权：需要登录
- 参数：无
- 结果：当前访问令牌失效。

## 用户资料与安全

### 修改用户资料

- 路径：`POST /api/user/update`
- 鉴权：需要登录
- 参数：

| 字段 | 必填 | 说明 |
| --- | --- | --- |
| nickname | 是 | 昵称 |
| avatar | 是 | 头像地址 |
| gender | 否 | 0=未知、1=男、2=女 |
| birthday | 否 | 生日 |
| motto | 否 | 签名 |

- 结果：返回更新后的用户资料。

手机号和邮箱属于敏感资料，不能通过此接口修改，必须使用对应的验证码换绑接口。

### 修改登录密码

- 路径：`POST /api/user/updatePwd`
- 鉴权：需要登录
- 参数：`oldPassword`、`newPassword`，均必填；新密码长度为 6 至 20 位，且不能与旧密码相同。
- 结果：修改成功后当前账号此前签发的所有访问令牌失效，需要重新登录。

### 发送密码重置验证码

- 路径：`POST /api/user/sendResetCode`
- 鉴权：免登录
- 参数：`mobile`，必填且须为合法手机号。
- 结果：对于已绑定且启用的账号发送短信验证码。接口始终返回成功，避免泄露手机号是否已注册。

### 通过验证码重置密码

- 路径：`POST /api/user/resetPwd`
- 鉴权：免登录
- 参数：`mobile`、`code`、`password` 均必填；新密码长度为 6 至 20 位。
- 结果：验证码校验成功后更新密码，并使当前账号此前签发的所有访问令牌失效。

### 发送手机号换绑验证码

- 路径：`POST /api/user/sendMobileBindCode`
- 鉴权：需要登录
- 参数：`mobile`，必填，为待绑定的新手机号。
- 结果：向新手机号发送换绑验证码。新手机号不能与当前手机号相同，也不能已被其他账号使用。

### 换绑手机号

- 路径：`POST /api/user/bindMobile`
- 鉴权：需要登录
- 参数：`password`、`mobile`、`code`，均必填。
- 结果：校验登录密码和新手机号验证码后更新绑定手机号。

### 发送邮箱换绑验证码

- 路径：`POST /api/user/sendEmailBindCode`
- 鉴权：需要登录
- 参数：`email`，必填，为待绑定的新邮箱。
- 结果：向新邮箱发送换绑验证码。新邮箱不能与当前邮箱相同，也不能已被其他账号使用。

### 换绑邮箱

- 路径：`POST /api/user/bindEmail`
- 鉴权：需要登录
- 参数：`password`、`email`、`code`，均必填。
- 结果：校验登录密码和新邮箱验证码后更新绑定邮箱。

### 注销账号

- 路径：`POST /api/user/cancel`
- 鉴权：需要登录
- 参数：`password`，必填。
- 结果：校验密码后将账号状态设为 `disable`，并使当前账号此前签发的所有访问令牌失效。该操作不会物理删除账号数据。

## 文件与历史记录

### 上传用户文件

- 路径：`POST /api/user/upload`
- 鉴权：需要登录
- Content-Type：`multipart/form-data`
- 参数：

| 字段 | 必填 | 说明 |
| --- | --- | --- |
| file | 是 | 待上传的文件 |
| uploadType | 否 | 文件类别：`image`、`doc`、`audio`、`video`、`zip` 或 `other`，默认 `other` |

- 结果：返回附件记录，包括 `id`、`name`、`fileUrl`、`size`、`sizeFormat`、`mimeType`、`ext` 等。

接口复用系统现有存储配置和文件类型校验，不提供匿名上传入口。

### 获取资产变动流水

- 路径：`GET /api/user/creditsLog`
- 鉴权：需要登录
- 查询参数：`page`、`perPage` 分页参数；`creditType` 可选，值为 `money` 或 `score`。
- 结果：返回当前用户的资产流水及分页字段。

每条流水包含 `creditType`、`beforeNum`、`num`、`afterNum`、`remark`、`createdAt`。后台编辑会员余额或积分时，系统会在同一数据库事务中自动创建流水记录，并在备注中保留操作后台用户 ID。

### 用户行为记录表

前台用户相关的登录、注册、退出登录及其他行为统一记录到 `hg_user_log`。其中 `logType` 用于区分记录类型，`action` 用于区分具体操作；登录历史接口只返回 `logType=login` 的记录。

统一记录字段包括：`userId`、`username`、`logType`、`action`、`actionName`、`eventAt`、`ip`、`userAgent`、`message`、`extra`、`status`、`createdAt` 和 `updatedAt`。后续新增用户行为时，复用这张表并新增对应的 `logType`、`action`，无需再创建独立日志表。

## 实现说明

1. 当前项目密码算法保持既有兼容规则：`MD5(明文密码 + Salt)`。
2. 免登录接口已在 `server/manifest/config/config.yaml` 的 `router.api.exceptLogin` 中配置。新增免登录用户接口时，应同步加入该白名单。
3. 用户中心采用 API 契约、控制器、服务、业务逻辑和 DAO 的分层结构。控制器仅做参数透传，存储和短信、邮件能力均复用既有服务。
