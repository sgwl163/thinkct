# AI 聊天助手模块开发与接口文档

本文档介绍了基于 HotGo 框架开发的 AI 聊天助手（OpenAI）模块的相关信息，包括数据表结构、API 接口定义以及后续的开发扩展说明。

## 1. 数据库结构

本模块在数据库中创建了 3 张相关表，前缀均为 `hg_openai_`：

1. **`hg_openai_chat_folder` (文件夹表)**
   - 用于管理用户的会话文件夹分类。
   - 核心字段：`id`, `member_id`, `name`, `sort`, `status`。

2. **`hg_openai_chat_session` (会话表)**
   - 用于管理具体的聊天会话。
   - 核心字段：`id`, `member_id`, `folder_id`, `title`, `prompt` (系统提示词), `model` (使用模型)。

3. **`hg_openai_chat_message` (消息表)**
   - 用于存储具体的聊天内容记录。
   - 核心字段：`id`, `member_id`, `session_id`, `role` (user/assistant/system), `content`, `req_tokens`, `res_tokens`。

4. **`hg_openai_app_category` (应用分类表)**
   - 用于管理 AI 应用中心的应用分类。
   - 核心字段：`id`, `name`, `sort`, `status`。

5. **`hg_openai_app` (应用表)**
   - 用于管理 AI 应用中心的具体应用。
   - 核心字段：`id`, `category_id`, `title`, `icon`, `description`, `prompt`, `use_count`, `sort`, `status`。

#### 2.3.5 编辑消息
- **接口路径**: `POST /api/message/edit`
- **请求参数**:
  ```json
  {
    "id": 1,
    "content": "修改后的内容"
  }
  ```

#### 2.3.6 删除消息
- **接口路径**: `POST /api/message/delete`
- **请求参数**:
  ```json
  {
    "id": 1
  }
  ```
- **说明**: 仅软删除指定的单条消息记录。

#### 2.3.7 消息反馈 (喜欢/不喜欢)
- **接口路径**: `POST /api/message/feedback`
- **请求参数**:
  ```json
  {
    "id": 1,
    "feedback": 1
  }
  ```
- **说明**: `feedback` 参数可选值为：`1` (喜欢), `2` (不喜欢), `0` (取消反馈)。

#### 2.3.8 重新生成消息
- **接口路径**: `POST /api/message/regenerate`
- **说明**: 该接口用于重新生成某条由 AI 返回的对话消息。调用此接口后，后端会自动软删除原有的那条 `messageId` 消息记录，并以此节点之前的聊天记录作为上下文向 AI 发起新的生成请求，最终以 SSE 格式流式返回最新生成的回答。
- **请求头**: 需增加 `Accept: text/event-stream`
- **请求参数**:
  ```json
  {
    "sessionId": 1,
    "messageId": 25,       // 必填，要重新生成的那条旧 AI 回复消息的 ID
    "isSearch": false
  }
  ```
- **响应数据**: 标准的 SSE 流格式（同 `2.3.3 发送消息`）。

---

## 2. API 接口文档

所有接口均注册在客户端（前台） API 路由下，前缀为 `/api`。
**注意：** 调用以下所有接口均需要通过 `Token` 鉴权（需要在请求头中携带有效的用户登录 Token）。

### 2.0 模型管理
#### 2.0.1 可用模型列表
- **接口路径**: `GET /api/model/list`
- **功能说明**: 获取后台配置的可用模型列表（含所属分组）。
- **请求参数**: 无
- **响应数据**:
  ```json
  {
    "list": [
      {
        "groupId": 1,
        "groupName": "默认OpenAI组",
        "modelId": 2,
        "name": "GPT-4 Vision",
        "model": "gpt-4-vision-preview",
        "isVision": 1,
        "isFile": 0,
        "isSearch": 0
      }
    ]
  }
  ```

### 2.1 文件夹管理

#### 2.1.1 文件夹列表
- **接口路径**: `GET /api/folder/list`
- **功能说明**: 获取当前登录用户的所有文件夹列表（按 sort 升序、id 降序）。
- **请求参数**: 无
- **响应数据**:
  ```json
  {
    "list": [
      {
        "id": 1,
        "name": "工作常用",
        "sort": 0
      }
    ]
  }
  ```

#### 2.1.2 编辑/新增文件夹
- **接口路径**: `POST /api/folder/edit`
- **功能说明**: 新增或修改文件夹。如果传入了有效的 `id` 则是修改，否则是新增。
- **请求参数**:
  ```json
  {
    "id": 0,          // 选填，大于0表示修改
    "name": "新文件夹", // 必填，文件夹名称
    "sort": 0         // 选填，排序值
  }
  ```

#### 2.1.3 删除文件夹
- **接口路径**: `POST /api/folder/delete`
- **功能说明**: 删除指定的文件夹（软删除）。
- **请求参数**:
  ```json
  {
    "id": 1 // 必填，文件夹ID
  }
  ```

---

### 2.2 会话管理

#### 2.2.1 会话列表
- **接口路径**: `GET /api/session/list`
- **功能说明**: 获取当前登录用户的会话列表，可选择根据文件夹或关键字进行过滤搜索。
- **请求参数**:
  ```json
  {
    "folderId": 0,    // 选填，指定文件夹下的会话，0表示查询所有
    "keyword": "工作" // 选填，支持根据会话的标题(title)进行模糊搜索
  }
  ```

#### 2.2.2 编辑/新增会话
- **接口路径**: `POST /api/session/edit`
- **功能说明**: 新增或修改会话。
- **请求参数**:
  ```json
  {
    "id": 0,
    "folderId": 1,
    "title": "新对话",     // 必填
    "prompt": "你是一个助手", // 选填
    "model": "gpt-4"      // 选填
  }
  ```
- **说明**: 新增会话时，如果没有主动设定温度，后台将默认初始化温度值为 `0.7`。
- **响应数据**:
  ```json
  {
    "id": 1  // 返回会话的实际ID（新增时返回新插入生成的ID，编辑时返回传入的ID）
  }
  ```

#### 2.2.7 删除会话
- **接口路径**: `POST /api/session/delete`
- **请求参数**:
  ```json
  {
    "id": 1
  }
  ```
- **说明**: 删除会话时，也会级联删除其名下的所有历史聊天记录。

#### 2.2.4 置顶/取消置顶会话
- **接口路径**: `POST /api/session/top`
- **请求参数**:
  ```json
  {
    "id": 1,
    "isTop": 1   // 1:置顶 0:取消
  }
  ```

#### 2.2.5 重命名会话
- **接口路径**: `POST /api/session/rename`
- **请求参数**:
  ```json
  {
    "id": 1,
    "title": "新的会话名称"
  }
  ```

#### 2.2.6 清空会话上下文
- **接口路径**: `POST /api/session/clear`
- **请求参数**:
  ```json
  {
    "id": 1
  }
  ```
- **说明**: 软删除该会话下的所有历史消息。

---

### 2.3 消息管理

#### 2.3.1 消息列表
- **接口路径**: `GET /api/message/list`
- **功能说明**: 获取指定会话下的历史消息记录。
- **请求参数**:
  ```json
  {
    "sessionId": 1  // 必填，会话ID
  }
  ```
- **响应数据**:
  ```json
  {
    "list": [
      {
        "id": 1,
        "sessionId": 1,
        "role": "user",
        "content": "你好",
        "model": null,
        "reqTokens": 0,
        "resTokens": 0,
        "totalTokens": 0,
        "feedback": 0,
        "createdAt": "2023-10-01 12:00:00"
      },
      {
        "id": 2,
        "sessionId": 1,
        "role": "assistant",
        "content": "你好！有什么我可以帮你的吗？",
        "model": "gpt-4o",
        "reqTokens": 10,
        "resTokens": 30,
        "totalTokens": 40,
        "feedback": 1,
        "createdAt": "2023-10-01 12:00:01"
      }
    ],
    "model": "gpt-4o",          // 该会话当前绑定的模型标识
    "prompt": "你是一个AI助手",   // 该会话当前绑定的系统提示词设定
    "temperature": 0.7          // 该会话当前绑定的温度值
  }
  ```


#### 2.3.2 上传图片/附件
- **接口路径**: `POST /api/upload/file`
- **功能说明**: 支持发送消息时上传图片及各类文件附件，并返回可供大模型读取的资源外网 URL 链接。
- **请求参数** (FormData 格式):
  - `file`: 二进制文件对象
- **响应数据**:
  ```json
  {
    "url": "https://your-domain.com/uploads/1.png"
  }
  ```

#### 2.3.3 发送消息 (SSE 流式输出)
- **接口路径**: `POST /api/message/send`
- **功能说明**: 用户发送一条消息给 AI，后台处理后通过 Server-Sent Events (SSE) 流式返回 AI 的回复，实现打字机效果。首次有效对话将触发后台异步自动生成并更新会话的短标题。支持传递图片、文件集合调用大模型多模态能力。**支持无会话自动创建**：当 `sessionId` 为 0 或不传时，后台会自动创建新会话，并可通过 `folderId` 可选参数指定该会话所属的文件夹，方便在特定分组下开始新对话。
- **请求头**:
  ```
  Accept: text/event-stream
  Authorization: Bearer <Your_Token>
  ```
- **请求参数**:
  ```json
  {
    "sessionId": 1,         // 当前消息所属会话ID。如果传 0 或不传，则会自动创建一个新会话
    "folderId": 0,          // 选填，指定自动创建会话时所属的文件夹ID（仅 sessionId 为 0 或空时生效）
    "content": "请分析这张图", // 选填（纯图片时可为空或携带指令）
    "images": [             // 选填，上传接口返回的图片URL数组
      "https://your-domain.com/uploads/1.png"
    ],
    "files": [              // 选填，上传接口返回的文件URL数组
      "https://your-domain.com/uploads/doc.pdf"
    ],
    "isSearch": true,       // 选填，是否开启联网搜索（针对支持联网的模型，默认 false）
    "model": "gpt-4o",      // 选填，指定或切换当前对话使用的模型标识（会同步更新会话绑定的模型）
    "prompt": "你是一个助手",  // 选填，指定或切换当前对话使用的系统提示词（会同步更新会话绑定的系统提示词）
    "temperature": 0.7      // 选填，指定或切换当前对话使用的温度值（会同步更新会话绑定的温度值。如果历史会话温度为 0 或空，后台调用时会自动兜底使用 0.7 防止模型回答过于呆板）
  }
  ```
- **响应数据**:
  - 返回流式分块数据，格式为带有 `event` 事件类型的标准 SSE。
  - 对于每次生成的内容，依次推送 `STREAM_CHUNK` 与 `CHUNK_DELTA` 两个事件块，事件载荷为 JSON 格式。
  - 当数据流结束时，将返回 `event: DONE` 标识。
  ```text
  event: STREAM_CHUNK
  data: {"message_id":"47135731852370690","patch_op":[{"patch_object":111,"patch_type":1,"patch_value":{"tts_content":"可以"}}]}
  
  event: CHUNK_DELTA
  data: {"text":"可以"}
  
  event: STREAM_CHUNK
  data: {"message_id":"47135731852370690","patch_op":[{"patch_object":111,"patch_type":1,"patch_value":{"tts_content":"帮你的吗？"}}]}
  
  event: CHUNK_DELTA
  data: {"text":"帮你的吗？"}
  
  event: DONE
  data: {}
  ```
- **最终返回的 JSON 响应数据** (在 DONE 事件之后):
  ```json
  {
    "sessionId": 1,         // 返回当前实际使用的会话ID
    "folderId": 0,          // 返回当前会话所属的文件夹ID
    "reply": {              // 发送成功后最终入库的助手回复实体
      "id": 2,
      "sessionId": 1,
      "role": "assistant",
      "content": "你好！有什么我可以帮你的吗？",
      "model": "gpt-4o",
      "reqTokens": 10,
      "resTokens": 30,
      "totalTokens": 40,
      "createdAt": "2023-10-01 12:00:01"
    }
  }
  ```

---

### 2.4 应用中心管理

#### 2.4.1 获取应用分类列表
- **接口路径**: `GET /api/openaiAppCategory/list`
- **功能说明**: 获取前台展示用的 AI 应用分类列表。
- **请求参数**: 无
- **响应数据**:
  ```json
  {
    "list": [
      {
        "id": 1,
        "name": "通用",
        "sort": 1,
        "status": 1
      }
    ]
  }
  ```

#### 2.4.2 获取应用列表
- **接口路径**: `GET /api/openaiApp/list`
- **功能说明**: 获取前台展示用的 AI 应用列表，支持按分类或关键字查询。
- **请求参数**:
  ```json
  {
    "categoryId": 1,  // 选填，指定分类下的应用，不传则查询所有
    "keyword": "翻译" // 选填，支持根据应用的标题(title)进行模糊搜索
  }
  ```
- **响应数据**:
  ```json
  {
    "list": [
      {
        "id": 1,
        "categoryId": 1,
        "title": "万能翻译官",
        "icon": "/upload/xxx.png",
        "description": "万能翻译，支持多国语言",
        "prompt": "你是一个专业的翻译官...",
        "useCount": 5,
        "sort": 1,
        "status": 1
      }
    ]
  }
  ```

---

## 3. 后续开发与扩展指南

当前 `MessageSend` 接口已实现了 **OpenAI 大模型真实调用** 及 **SSE 流式输出**。

### 3.1 自定义模型与 API 组配置
系统现已支持在数据库中配置不同的模型组与对应的 `apiKey` / `base_url`。在发送对话时，程序将自动根据会话所绑定的模型标识匹配分组参数。

在数据库中对应以下两张表：
- **`hg_openai_api_group` (API分组表)**：用于设置不同服务商的名称、接口地址 (`base_url`) 及 API 密钥 (`api_key`)。
- **`hg_openai_api_model` (可用模型表)**：用于配置支持的具体模型，并通过 `group_id` 关联至具体的接口组。

> 若找不到对应的数据库配置，将兜底使用项目 `manifest/config/config.yaml` 文件中的 `openai` 节点配置。

---

## 4. 二次开发与代码结构说明

为了方便后续对 AI 助手模块进行二次开发，本节详细说明了相关代码的分布位置和整体架构。该模块完全遵循 HotGo (基于 GoFrame v2) 的三层架构规范。

### 4.1 代码目录分布

以下是与 AI 聊天模块相关的核心目录树结构：

```text
p:\go\hotgo-2.0\server\
├── api\
│   └── api\
│       └── openai\
│           └── v1\
│               └── openai.go          # API 入参、出参和路由元数据定义
├── internal\
│   ├── controller\
│   │   └── api\
│   │       └── openai\
│   │           └── openai.go          # 控制器，负责接收请求并转发给 Logic
│   ├── logic\
│   │   └── openai\
│   │       └── openai.go              # 核心业务逻辑层（增删改查、大模型调用、上下文组装等）
│   ├── service\
│   │   └── openai.go                  # 业务逻辑层自动生成的接口文件（gf gen service 生成）
│   ├── router\
│   │   └── api.go                     # API 路由注册（已将 openai 控制器挂载）
│   ├── dao\
│   │   └── internal\                  # 自动生成的数据访问对象（gf gen dao 生成）
│   │       ├── openai_chat_folder.go
│   │       ├── openai_chat_session.go
│   │       ├── openai_chat_message.go
│   │       ├── openai_api_group.go
│   │       └── openai_api_model.go
│   └── model\
│       ├── do\                        # 数据对象，用于 DAO 交互
│       └── entity\                    # 数据实体模型
├── utility\
│   └── openai\
│       └── client.go                  # OpenAI 底层 HTTP 客户端封装（包含 SSE 流式处理）
├── storage\
│   └── data\
│       └── generate\
│           └── openai.sql             # 表结构初始化的 SQL 文件
└── api_openai_doc.md                  # 本接口与开发文档
```

- **API 接口定义层**
  - **路径**：`api/api/openai/v1/openai.go`
  - **作用**：定义所有 HTTP 请求的入参 (`XxxReq`)、出参 (`XxxRes`) 以及接口元数据（路径、方法、Swagger 注解）。添加新接口时，需先在此处声明。

- **控制器层 (Controller)**
  - **路径**：`internal/controller/api/openai/openai.go`
  - **作用**：负责接收客户端请求，调用对应的 Service 接口，并返回响应结果。此层尽量保持轻量，不包含复杂业务逻辑。

- **业务逻辑层 (Logic / Service)**
  - **路径**：`internal/logic/openai/openai.go`
  - **作用**：模块的**核心业务代码**。包括数据库的增删改查、上下文组装、流式回复控制、会话管理等都在这里实现。如果需要修改业务规则（如改变获取上下文的条数限制、修改自动生成标题的 Prompt 等），请修改此文件。每次修改后，如更改了方法签名，需要运行 `gf gen service` 更新接口。

- **数据访问层 (DAO / DO / Entity)**
  - **路径**：`internal/dao/internal/`、`internal/model/do/`、`internal/model/entity/`
  - **作用**：封装数据库操作。当您通过 SQL 修改了相关表结构后，需在终端运行 `gf gen dao`，框架会自动更新这里的模型代码供 Logic 层调用。

- **OpenAI 客户端工具包**
  - **路径**：`utility/openai/client.go`
  - **作用**：底层与大语言模型通信的 HTTP 客户端封装。支持标准的 `ChatCompletion` 以及流式输出 `ChatCompletionStream`。如果您需要对接特定厂商的非标准接口、或增加 Function Calling / Tools 等高级特性，可以在此处进行扩展。

### 4.2 核心业务数据流 (以发送消息为例)

1. **客户端请求**：前端发起 `POST /api/message/send` (带 `Accept: text/event-stream`)。
2. **路由与控制器**：GoFrame 路由拦截，进入 `cOpenai.MessageSend` 控制器。
3. **业务逻辑验证与存储**：进入 `logic.MessageSend`，校验并保存用户的提问内容，更新会话最后操作时间、绑定模型、温度等。
4. **异步任务**：如果为首次对话，触发 `go func()` 后台协程请求大模型生成简短标题并更新。
5. **上下文组装**：从数据库读取最近 20 条消息，将图片/文件/文本组装为多模态 `messages` 数组。
6. **模型配置读取**：根据数据库 `hg_openai_api_model` 和 `hg_openai_api_group` 读取 `baseUrl` 和 `apiKey`。
7. **SSE 流式输出**：调用 `utility/openai` 发起 HTTP 流式请求，通过 `r.Response.Writef` 按 `STREAM_CHUNK` 和 `CHUNK_DELTA` 事件块实时向客户端 flush 数据。
8. **结果落库**：流式读取完毕后，将完整的 AI 响应文本存入数据库，返回最后插入的消息实体。
