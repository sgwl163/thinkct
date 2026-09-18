# AI 音乐创作接口文档

本文档描述了系统集成的 AI 音乐创作功能接口规范。基于 Suno 的音乐生成能力，实现从任务提交、进度轮询到结果展示的完整工作流。

## 1. 提交音乐创作任务
- **接口路径**: `POST /api/music/submit`
- **说明**: 异步提交音乐生成任务，支持灵感模式和定制模式。
- **请求参数 (JSON)**:
  - `prompt` (string): 必填，提示词或歌词。
  - `type` (int): 可选，1=灵感模式，2=定制模式。默认 1。
  - `makeInstrumental` (int): 可选，是否纯音乐（0=否，1=是）。灵感模式专用。
  - `title` (string): 可选，歌名。定制模式专用。
  - `tags` (string): 可选，风格标签。定制模式专用。
  - `model` (string): 必填，使用的音乐生成模型。例如：
    - `suno-v3.5` (默认 3.5 版本)
    - `chirp-v3-5` (Chirp 3.5 版本)
    - `chirp-v3-0` (Chirp 3.0 版本)

- **响应数据**:
```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "id": 1001  // 生成的音乐任务ID
  }
}
```

## 2. 获取音乐创作任务列表
- **接口路径**: `GET /api/music/list`
- **说明**: 分页获取当前用户的音乐任务及生成结果。
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
        "prompt": "一首欢快的流行歌曲",
        "type": 1,
        "makeInstrumental": 0,
        "title": "快乐周末",
        "tags": "pop, happy",
        "model": "suno-v3.5",
        "audioUrl": "https://example.com/audio.mp3",
        "imageUrl": "https://example.com/cover.jpg",
        "videoUrl": "https://example.com/video.mp4",
        "duration": 128.5,
        "status": 1, 
        "failReason": "",
        "remoteTaskId": "task-uuid",
        "remoteClipId": "clip-uuid",
        "createdAt": "2026-08-05 15:00:00"
      }
    ],
    "total": 1
  }
}
```

## 3. 删除音乐记录
- **接口路径**: `POST /api/music/delete`
- **说明**: 删除当前用户的指定历史音乐任务记录。
- **请求参数 (JSON)**:
  - `id` (int64): 必填，任务 ID。

## 4. 获取音乐作品广场
- **接口路径**: `GET /api/music/square`
- **说明**: 分页获取全站所有成功生成的音乐作品，用于广场展示。
- **请求参数 (Query)**:
  - `page` (int): 可选，页码，默认 1。
  - `size` (int): 可选，每页数量，默认 20。

## 5. 获取音乐任务详情
- **接口路径**: `GET /api/music/detail`
- **说明**: 根据任务 ID 获取单个音乐任务的详细信息。
- **请求参数 (Query)**:
  - `id` (int64): 必填，任务 ID。

## 补充说明
- 每次提交音乐生成任务，Suno 模型通常会返回 **两首** 不同版本的歌曲。
- 服务端轮询到成功状态时，会自动将两首歌曲下载并分别保存：
  - 更新原任务（填入第一首歌曲的链接和信息）。
  - 新增一条任务记录（填入第二首歌曲的链接和信息）。
- 因此，用户在界面上会看到一条任务提交后，产生了两个不同版本的音乐记录。
