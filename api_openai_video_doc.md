# AI 视频创作接口文档

本文档描述了系统集成的 AI 视频创作功能接口规范。基于 OpenAI / 兼容平台（如豆包大模型）的视频生成能力，实现从任务提交、进度轮询到结果展示的完整工作流。

## 1. 提交视频创作任务
- **接口路径**: `POST /api/video/submit`
- **说明**: 异步提交视频生成任务，支持纯文本生成视频和首尾帧图片参考生成视频。任务提交后会立刻返回系统的任务 ID。
- **请求参数 (JSON)**:
  - `prompt` (string): 必填，描述您想生成的视频画面，支持中英文。
  - `model` (string): 必填，使用的视频生成模型。例如：
    - `doubao-seedance-1-5-pro-251215` (1.5全能视频模型)
    - `doubao-seedance-1-0-pro-250528` (1.0全能视频模型)
    - `doubao-seedance-1-0-lite-t2v-250428` (文生视频模型)
    - `doubao-seedance-1-0-lite-i2v-250428` (图生视频模型)
  - `referenceImage` (string): 可选，参考图的 URL，用于图生视频。如果提供了该参数，底层 API 会将其作为 `first_frame` (首帧) 传递给模型。

- **响应数据**:
```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "id": 1001  // 生成的视频任务ID
  }
}
```

## 2. 获取视频创作任务列表
- **接口路径**: `GET /api/video/list`
- **说明**: 分页获取当前用户的视频任务及生成结果。
- **请求参数 (Query)**:
  - `page` (int): 可选，页码，默认 1。
  - `size` (int): 可选，每页数量，默认 10。

- **响应数据**:
```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "list": [
      {
        "id": 1001,
        "prompt": "一个赛博朋克风格的未来城市，霓虹灯闪烁...",
        "referenceImage": "https://example.com/ref.jpg",
        "model": "doubao-seedance-1-5-pro-251215",
        "videoUrl": "https://example.com/video.mp4",
        "status": 1, 
        "failReason": "",
        "remoteTaskId": "volces-task-xxxxxx",
        "createdAt": "2026-08-05 15:00:00"
      }
    ],
    "total": 1
  }
}
```

**关于 `status` 字段说明**:
- `status=0`: 进行中（此时 `videoUrl` 为空，底层正在自动轮询远端任务状态，前端应展示生成中的动画特效）。
- `status=1`: 成功（此时 `videoUrl` 会包含系统自动下载并转存后的视频完整地址）。
- `status=2`: 失败（此时可通过 `failReason` 提示用户失败原因，例如超时或接口报错）。

## 3. 删除视频记录
- **接口路径**: `POST /api/video/delete`
- **说明**: 删除当前用户的指定历史视频任务记录。
- **请求参数 (JSON)**:
  - `id` (int64): 必填，任务 ID。

## 4. 获取视频创作作品广场
- **接口路径**: `GET /api/video/square`
- **说明**: 分页获取全站所有成功生成的视频作品，用于广场展示。
- **请求参数 (Query)**:
  - `page` (int): 可选，页码，默认 1。
  - `size` (int): 可选，每页数量，默认 20。

- **响应数据**:
与 `/api/video/list` 的返回格式完全一致，包含 `list` 数组和 `total` 数量。

## 5. 获取视频任务详情
- **接口路径**: `GET /api/video/detail`
- **说明**: 根据任务 ID 获取单个视频任务的详细信息。
- **请求参数 (Query)**:
  - `id` (int64): 必填，任务 ID。

- **响应数据**:
直接返回任务对象的详情。
```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "id": 1001,
    "prompt": "...",
    "referenceImage": "",
    "model": "doubao-seedance-1-5-pro-251215",
    "videoUrl": "https://...",
    "status": 1,
    "failReason": "",
    "remoteTaskId": "...",
    "createdAt": "2026-08-05 15:00:00"
  }
}
```

## 模型及接口配置机制说明
- **后台异步轮询与下载**：服务端在接收到 `POST /api/video/submit` 任务后，会立即返回系统内部的 `id`，并在后台启动一个协程。该协程会负责调用第三方大模型接口（如 `https://api.gpt.ge/task/volces/seedance`），获取远端 `remoteTaskId` 后，执行最多长达 10 分钟的轮询（每 5 秒一次）。
- **防盗链/资源本地化**：当远端任务成功后，系统会自动请求远端的 `videoUrl`，并将视频流数据保存到本地服务器（或云存储），最终更新数据库的 `videoUrl` 为本地/OSS 的安全链接，避免远端链接过期。
- **配置复用**：调用视频接口时，自动复用系统已有的 `openai` 令牌配置（`manifest/config/config.yaml`）：
```yaml
openai:
  apiKey: "sk-xxxxxx"
  baseUrl: "https://api.gpt.ge/v1"
```
服务端在请求视频接口时会自动将 `/v1` 替换为适配火山引擎接口格式的路由。