# 即时通讯 (IM) 模块开发文档

本文档描述了 HotGo 项目中基于 WebSocket 协议的即时通讯 (IM) 模块的数据库设计、代码结构、以及消息收发交互流程。

---

## 1. 数据库结构

所有的表前缀均为 `hg_im_`，并强制使用 `utf8mb4` 字符集以支持 Emoji 表情。完整的 SQL 建表语句已保存在项目根目录的 `hg_im.sql` 中。

| 表名 | 作用 | 核心字段说明 |
| :--- | :--- | :--- |
| `hg_im_conversation` | 会话列表表 | 左侧会话列表展示数据，包含未读数、最后消息摘要等。 |
| `hg_im_message` | 聊天记录表 | 右侧对话框的详细消息记录，包含发送方、接收方、消息内容等。 |
| `hg_im_group` | 群组基础表 | 存储群名称、群头像、群主ID、群公告等信息。 |
| `hg_im_group_member` | 群成员关联表 | 记录群内用户角色（群主/管理/普通）及禁言状态。 |
| `hg_im_contact` | 好友关系表 | 记录用户之间的好友关系与备注。 |

---

## 2. 代码目录结构

WebSocket 相关代码和配置按照 HotGo 规范进行了分布：

```text
p:\go\hotgo-2.0\server
├── internal
│   ├── controller
│   │   └── websocket
│   │       └── handler
│   │           └── im
│   │               └── chat.go     // IM 的核心收发逻辑实现（处理发消息、入库、更新会话列表、广播分发）
│   └── router
│       └── websocket.go            // WebSocket 路由注册中心，在这里将 `im/chat/send` 绑定到了 `im.Chat.SendMsg`
├── manifest
│   └── config
│       └── config.yaml             // 生成代码时的配置项（如果您后续使用 HotGo 的代码生成器生成基础的增删改查页面，可在后台使用 `im` 分组模板）
└── hg_im.sql                       // 数据库建表文件
```

---

## 3. WebSocket 连接说明

由于发送和接收 IM 消息需要在持续的长连接上进行，因此不能使用传统的 HTTP POST 请求。

### 3.1 建立连接

* **连接地址**：`ws://127.0.0.1:8000/socket` (默认地址，根据您的配置而定)
* **鉴权要求**：由于 WebSocket 不能通过 Headers 传递 Token，请在建立连接时将 Token 放到 URL 参数中。
* **例如**：`ws://127.0.0.1:8000/socket?authorization=您的Token`
* **断开规则**：如果未登录或 Token 无效，服务端会主动断开连接。

---

## 4. Http API 接口说明

除了 WebSocket 消息收发，IM 模块还提供了一些常规的 HTTP 接口用于获取页面初始化数据。

### 4.1 获取会话列表

* **接口地址**：`GET /api/im/conversation/list`
* **请求头**：必须携带 `Authorization: 您的Token`
* **功能说明**：返回当前登录用户的所有聊天会话，按“是否置顶”及“最新消息时间”排序，用于渲染客户端左侧的聊天列表。

**返回示例：**
```json
{
    "code": 0,
    "message": "操作成功",
    "data": {
        "list": [
            {
                "id": 1,
                "chatType": 1,
                "targetId": 2,
                "targetName": "张三",
                "targetAvatar": "http://example.com/avatar.png",
                "lastMsgId": 105,
                "lastMsgContent": "你好，这是测试消息！",
                "lastMsgTime": "2026-06-09 15:30:00",
                "unreadCount": 1,
                "isTop": 0,
                "isDisturb": 0
            }
        ]
    }
}
```

### 4.2 获取历史聊天记录

* **接口地址**：`GET /api/im/message/list`
* **请求头**：必须携带 `Authorization: 您的Token`
* **请求参数**：
  * `chatType` (必填): 会话类型 `1`=单聊，`2`=群聊
  * `targetId` (必填): 目标ID（好友用户ID或群组ID）
  * `page` (选填): 页码，默认 `1`
  * `pageSize` (选填): 每页数量，默认 `20`
* **功能说明**：返回指定会话的历史聊天记录。支持分页，并且已自动对当前页的数据按照时间**正序排序**（最老的在上面，最新的在下面），方便前端直接渲染到聊天框。

**返回示例：**
```json
{
    "code": 0,
    "message": "操作成功",
    "data": {
        "list": [
            {
                "id": 105,
                "chatType": 1,
                "fromId": 1,
                "fromName": "李四",
                "fromAvatar": "http://example.com/avatar.png",
                "toId": 2,
                "msgType": 1,
                "content": "你好，这是测试消息！",
                "createdAt": "2026-06-09 15:30:00"
            }
        ],
        "total": 1
    }
}
```

### 4.3 添加好友

* **接口地址**：`POST /api/im/friend/add`
* **请求头**：必须携带 `Authorization: 您的Token`
* **请求参数 (JSON)**：
  * `friendId` (必填): 想要添加的好友的 `user_id`。
  * `remark` (选填): 为好友设置的备注名。
* **功能说明**：双向添加好友，自动在 `hg_im_contact` 表中插入两条关系记录。

### 4.4 创建群组

* **接口地址**：`POST /api/im/group/create`
* **请求头**：必须携带 `Authorization: 您的Token`
* **请求参数 (JSON)**：
  * `name` (必填): 群名称。
  * `avatar` (选填): 群头像的 URL。
  * `notice` (选填): 群公告。
* **功能说明**：创建一个新的群组。接口会在事务中执行两步：
  1. 在 `hg_im_group` 表中创建群组，记录当前用户为群主。
  2. 自动将当前用户加入 `hg_im_group_member` 成员表中，并设置角色为 `1`（群主）。

### 4.5 获取联系人列表

* **接口地址**：`GET /api/im/contact/list`
* **请求头**：必须携带 `Authorization: 您的Token`
* **功能说明**：获取当前用户的好友（联系人）列表。返回数据中包含好友ID、昵称、头像、备注名等，用于渲染通讯录列表。

**返回示例：**
```json
{
    "code": 0,
    "message": "操作成功",
    "data": {
        "list": [
            {
                "friendId": 2,
                "nickname": "张三",
                "avatar": "http://...",
                "remark": "张三同学",
                "status": 1
            }
        ]
    }
}
```

### 4.6 获取群组列表

* **接口地址**：`GET /api/im/group/list`
* **请求头**：必须携带 `Authorization: 您的Token`
* **功能说明**：获取当前用户已加入的所有群组列表。

**返回示例：**
```json
{
    "code": 0,
    "message": "操作成功",
    "data": {
        "list": [
            {
                "groupId": 1,
                "name": "技术交流群",
                "avatar": "http://...",
                "role": 1
            }
        ]
    }
}
```

### 4.7 发起会话

* **接口地址**：`POST /api/im/conversation/open`
* **请求头**：必须携带 `Authorization: 您的Token`
* **请求参数 (JSON)**：
  * `chatType` (必填): 会话类型，`1`=单聊，`2`=群聊
  * `targetId` (必填): 目标好友的用户 ID 或目标群组的 ID
* **功能说明**：当在通讯录列表中点击某个好友或群组时，调用此接口可以获取对应的会话 ID（`conversationId`）。如果 `hg_im_conversation` 中尚未存在该会话记录，则会自动在表中插入一条空会话数据以供左侧会话列表展示使用。

**返回示例：**
```json
{
    "code": 0,
    "message": "操作成功",
    "data": {
        "conversationId": 15
    }
}
```

### 4.8 邀请加入群聊

* **接口地址**：`POST /api/im/group/invite`
* **请求头**：必须携带 `Authorization: 您的Token`
* **请求参数 (JSON)**：
  * `groupId` (必填): 目标群组的 ID。
  * `userIds` (必填): 包含想要邀请加入的好友的用户 ID 数组。
* **功能说明**：将一个或多个用户拉入指定的群聊中。后台会自动过滤掉已经存在于该群的用户，避免重复插入。

**请求示例：**
```json
{
    "groupId": 1,
    "userIds": [2, 3]
}
```

### 4.9 获取群成员列表

* **接口地址**：`GET /api/im/group/member/list`
* **请求头**：必须携带 `Authorization: 您的Token`
* **请求参数**：
  * `groupId` (必填): 目标群组的 ID。
* **功能说明**：获取指定群组下的所有群成员。返回数据已按照角色（群主 -> 管理员 -> 普通成员）以及进群时间的先后顺序进行了排序。

**返回示例：**
```json
{
    "code": 0,
    "message": "操作成功",
    "data": {
        "list": [
            {
                "userId": 1,
                "nickname": "张三",
                "avatar": "http://...",
                "role": 1,
                "isMute": 0
            }
        ]
    }
}
```

---

## 5. WebSocket 消息收发交互协议

在连接建立成功后，前后端通过发送和接收 JSON 字符串进行通信。

### 4.1 发送消息 (Client -> Server)

客户端向服务端发送消息时，将 JSON 序列化后通过 WebSocket 发送。

* **Event（事件路由）**：`im/chat/send`

**客户端发送格式示例：**

```json
{
    "event": "im/chat/send",
    "data": {
        "chatType": 1, 
        "toId": 2, 
        "msgType": 1, 
        "content": "你好，这是测试消息！"
    }
}
```

**参数说明：**
* `chatType`: `1`=单聊，`2`=群聊，`3`=系统消息
* `toId`: 接收方的 ID。如果单聊就是目标用户的 `id`；如果群聊就是目标群组的 `group_id`。
* `msgType`: `1`=文本，`2`=图片，`3`=文件，`6`=表情包等
* `content`: 消息内容字符串（图片等类型可以传 URL）

---

### 4.2 接收回执 (Server -> 发送方 Client)

发送方将消息成功发出并被服务端入库处理后，服务端会给发送方返回一条成功的回执。

**服务端回执示例：**

```json
{
    "event": "im/chat/send",
    "code": 0,
    "message": "success",
    "data": {
        "msgId": 105,
        "time": "2026-06-09 15:30:00"
    }
}
```

---

### 4.3 接收新消息 (Server -> 接收方 Client)

当服务端处理完消息后，会主动向目标用户（如果是群聊则向除了发送者以外的所有群成员）推送这条新消息。

* **Event（事件路由）**：`im/chat/receive`

**服务端推送示例：**

```json
{
    "event": "im/chat/receive",
    "data": {
        "id": 105,
        "chatType": 1,
        "fromId": 1,
        "fromName": "张三",
        "fromAvatar": "http://example.com/avatar.png",
        "toId": 2,
        "msgType": 1,
        "content": "你好，这是测试消息！",
        "createdAt": "2026-06-09 15:30:00"
    }
}
```

**参数说明：**
* `fromId`: 发送此消息的用户 ID。客户端收到后可根据此 ID 在左侧会话列表中找到对应的聊天框。

---

## 6. 业务逻辑补充说明

当触发 `im/chat/send` 事件时，服务端会自动执行以下逻辑（参见 `chat.go`）：
1. **数据校验**：验证登录身份是否有效，验证目标ID和内容是否为空。
2. **入库记录**：将完整的消息详情写入 `hg_im_message` 表，生成自增的 `msgId`。
3. **更新会话 (Session) 状态**：
   - 更新**发送方**的 `hg_im_conversation` 表（更新最后一条消息内容、时间）。
   - 更新**接收方**的 `hg_im_conversation` 表（更新最后一条消息内容、时间，并将未读数 `unread_count` 自动加 1）。
4. **实时推送**：利用 HotGo 的 `websocket.SendToUser()` 功能，实时将数据下发至相关用户的客户端上。