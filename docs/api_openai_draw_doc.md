# AI 绘画接口文档 (OpenAI)

## 1. 提交绘画任务
- **接口路径**: `POST /api/draw/submit`
- **说明**: 异步提交绘画任务，支持传入文生图和参考图。支持 MJ 等模型的高级参数。任务提交后会立刻返回任务 ID。
- **请求参数 (JSON)**:
  - `prompt` (string): 必填，绘画提示词。
  - `model` (string): 必填，使用的绘画模型（如 `gemini-3-pro-image`, `midjourney`, `gpt-image-2` 等）。
  - `ratio` (string): 可选，图片比例（如 `1:1`, `16:9`, `4:3` 等），默认 `1:1`。
  - `referenceImage` (string): 可选，参考图的 URL，用于图生图（具体支持视模型而定，如 Midjourney 会将图片转换为 Base64 发送）。
  - `version` (string): 可选，模型版本号或特征参数（例如 `6.1` 等）。
  - `raw` (int): 可选，是否开启 RAW 模式（0 否，1 是）。
  - `negativePrompt` (string): 可选，负面提示词。
  - `stylize` (int): 可选，风格化程度(0-1000)。
  - `chaos` (int): 可选，多样性(0-100)。
  - `seed` (int): 可选，随机种子。

- **响应数据**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1001  // 生成的绘画任务ID
  }
}
```

## 2. 获取绘画任务列表
- **接口路径**: `GET /api/draw/list`
- **说明**: 分页获取当前用户的绘画任务及生成结果。
- **请求参数 (Query)**:
  - `page` (int): 可选，页码，默认 1。
  - `size` (int): 可选，每页数量，默认 10。

- **响应数据**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1001,
        "prompt": "A futuristic city with flying cars",
        "referenceImage": "",
        "model": "gemini-3-pro-image",
        "ratio": "16:9",
        "imageUrl": "https://api.vectorengine.ai/.../image.png",
        "status": 1, 
        "progress": 100,
        "failReason": "",
        "createdAt": "2026-06-15 14:00:00",
        "mjTaskId": "1784915543848681",
        "buttons": "[{\"customId\":\"MJ::JOB::upsample::1::xxx\",\"emoji\":\"U1\"}]"
      }
    ],
    "total": 1
  }
}
```

**关于 `status` 和 `progress` 字段说明**:
- `status=0`: 进行中（此时 `imageUrl` 为空，前端应展示加载动画或排队中状态）。结合 `progress` (0-100) 字段可展示 Midjourney 任务的具体生成进度百分比。
- `status=1`: 成功（此时 `imageUrl` 会包含生成的图片完整地址，`progress` 恒定为 100）。
- `status=2`: 失败（此时可通过 `failReason` 提示用户失败原因）。

## 3. 删除绘画记录
- **接口路径**: `POST /api/draw/delete`
- **说明**: 删除历史绘画任务记录。
- **请求参数 (JSON)**:
  - `id` (int64): 必填，任务 ID。

## 4. 提交 MJ 图片交互操作 (U/V 放大变换等)
- **接口路径**: `POST /api/draw/action`
- **说明**: 用于提交针对原 MJ 生成任务的按钮交互操作（比如 U1, V1）。服务端会创建一个新的任务记录，并进行新一轮的轮询出图。
- **请求参数 (JSON)**:
  - `id` (int64): 必填，原画图任务在系统的 ID（前端从列表获取到的 id）。
  - `action` (string): 必填，具体操作类型（如果是代理返回的完整 customId 直接传入，或者类似 U1/V1 代理若兼容也可以直接传入）。
- **响应数据**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1002  // 针对本次操作生成的新任务 ID，前端需轮询该 ID
  }
}
```

## 5. 提示词翻译为英文
- **接口路径**: `POST /api/draw/translate`
- **说明**: 将用户输入的中文绘画描述翻译为英文提示词。
- **请求参数 (JSON)**:
  - `prompt` (string): 必填，需要翻译的中文提示词。
- **响应数据**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "result": "A cute cat sitting on the grass"
  }
}
```

## 6. 提示词联想/扩写
- **接口路径**: `POST /api/draw/enrich`
- **说明**: 接收用户简短的绘画描述，通过 AI 联想并扩写为画面丰富、带有细节和高级渲染参数的高质量英文提示词。
- **请求参数 (JSON)**:
  - `prompt` (string): 必填，简短的原始提示词。
- **响应数据**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "result": "A cute cat sitting on the lush green grass, surrounded by colorful wildflowers, soft sunlight filtering through the leaves, highly detailed, 8k resolution, cinematic lighting, masterpiece."
  }
}
```

## 7. 获取绘画作品广场
- **接口路径**: `GET /api/draw/square`
- **说明**: 分页获取所有用户公开发布或成功生成的绘画作品（广场展示）。
- **请求参数 (Query)**:
  - `page` (int): 可选，页码，默认 1。
  - `size` (int): 可选，每页数量，默认 20。

- **响应数据**:
与 `/api/draw/list` 的返回格式一致，包含 `list` 数组和 `total` 数量。

## 8. 获取绘画任务详情
- **接口路径**: `GET /api/draw/detail`
- **说明**: 根据任务 ID 获取单个绘画任务的详细信息（包含最终生成的 imageUrl、prompt 等）。
- **请求参数 (Query)**:
  - `id` (int64): 必填，任务 ID。

- **响应数据**:
直接返回任务对象的详情。
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1001,
    "prompt": "A futuristic city with flying cars",
    "referenceImage": "",
    "model": "gemini-3-pro-image",
    "ratio": "16:9",
    "imageUrl": "https://api.vectorengine.ai/.../image.png",
    "status": 1, 
    "progress": 100,
    "failReason": "",
    "createdAt": "2026-06-15 14:00:00",
    "mjTaskId": "1784915543848681",
    "buttons": "[{\"customId\":\"MJ::JOB::upsample::1::xxx\",\"emoji\":\"U1\"}]"
  }
}
```

## 9. 上传参考图 (通用文件上传)
- **说明**: 前端如果需要提供图生图的“参考图”功能，可以直接调用 HotGo 系统已有的通用文件上传接口。获取到上传成功的 URL 后，再将该 URL 传给 `/api/draw/submit` 接口的 `referenceImage` 参数。
- **常见接口路径**: `POST /admin/sys/attachment/upload` 或 `/api/common/upload` (具体依赖前端项目中现有的公共上传接口配置)。

## 模型及接口地址配置
- 项目已经在 `manifest/config/config.yaml` 增加了 `openai` 绘画令牌配置：
```yaml
openai:
  apiKey: "sk-AKcoSZWXLytdtoCkWLyWqhfOThqVbZGjiMHy5oiOVIWSVis7"
  baseUrl: "https://api.vectorengine.ai/v1"
```
当前首发对接了 `gemini-3-pro-image` 模型，并且在底层实现了对标准 `/v1/images/generations` 接口的封装。如果后期增加 `gpt-image-2` 或 `即梦ai` 等兼容模型，前端只需改变 `model` 字段传递对应的模型名称即可。