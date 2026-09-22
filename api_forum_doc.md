# 社区论坛接口文档

本文档说明 HotGo 社区论坛接口。所有接口基础路径为 `/api`，统一返回：

```json
{"code":0,"message":"操作成功","timestamp":1234567890,"traceID":"xxx","data":{}}
```

除明确标记为免登录的查询接口外，请在请求头中携带访问令牌：

```text
Authorization: Bearer <token>
```

也兼容项目现有的 `userToken` 请求头和查询参数。

## 数据模型

论坛模块使用以下数据表：

- `hg_forum_category`：论坛板块，保存名称、简介、图标和主题数量。
- `hg_forum_topic`：主题正文、作者、图片和互动计数。
- `hg_forum_comment`：主题评论及回复关系。
- `hg_forum_interaction`：用户对主题或评论的点赞、主题收藏关系，使用 `targetType` 和 `actionType` 区分对象与操作。
- `hg_forum_user_follow`：用户关注关系，保存关注者、被关注用户和关注时间。
- `hg_forum_report`：用户对主题或评论的举报记录。

主题状态由服务端控制。当前公开查询只返回 `normal` 状态的主题；评论列表和回复校验只使用 `normal` 状态的评论。

时间字段统一使用后缀 `_at`，包括 `created_at`、`updated_at`、`deleted_at`、`last_reply_at` 和 `handled_at`。接口以对应的驼峰字段名和标准日期时间字符串返回。主题对象额外返回 `keywords`（关键词）和 `description`（摘要）。主题作者和评论作者以 `author` 对象返回，包含 `id`、`username`、`nickname`、`avatar`。评论对象中的 `children` 为其无限级子回复数组；每层均包含作者和被回复用户信息。

## 分页参数

主题列表和评论列表支持通用分页参数：

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| page | int | 否 | 页码，默认 1 |
| pageSize | int | 否 | 每页数量，默认 10，最大 200 |
| pagination | bool | 否 | 是否分页，默认 true |

分页响应包含 `list`、`page`、`pageSize`、`pageCount`、`totalCount`。

## 免登录接口

### 获取板块列表

- 路径：`GET /api/forum/category/list`
- 鉴权：免登录
- 参数：无
- 结果：返回启用的板块，按 `sort`、`id` 升序排列。

响应示例：

```json
{
  "list": [
    {
      "id": 1,
      "pid": 0,
      "name": "产品交流",
      "description": "分享产品使用经验",
      "image": "https://example.com/forum.png",
      "sort": 1,
      "status": 1,
      "topicCount": 12
    }
  ]
}
```

### 获取主题列表

- 路径：`GET /api/forum/topic/list`
- 鉴权：免登录

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| categoryId | uint64 | 否 | 按板块筛选 |
| keyword | string | 否 | 按主题标题模糊搜索，最多 100 个字符 |
| sort | string | 否 | `latest` 最新，`hot` 热门，`recommend` 推荐，默认 `latest` |

`latest` 按置顶、`lastReplyAt` 和 `createdAt` 排序；`hot` 按置顶、精华、点赞数、评论数和浏览数排序；`recommend` 按置顶、精华、收藏数、点赞数、评论数和浏览数排序。关键词同时匹配主题标题、关键词和摘要。

主题对象包含 `isLiked` 和 `isFavorited`。该接口支持可选登录：未登录时两个字段均为 `false`；请求携带有效登录令牌时，字段分别反映当前用户是否已点赞和收藏主题。

### 获取主题详情

- 路径：`GET /api/forum/topic/detail`
- 鉴权：免登录

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| id | uint64 | 是 | 主题 ID |

成功后主题浏览量增加 1，并返回 `topic` 对象、同板块中按主题 ID 定位的 `prev`（上一篇）和 `next`（下一篇）。上一篇或下一篇不存在时对应字段为 `null`。主题不存在或状态不可见时返回业务错误。

### 获取论坛用户资料

- 路径：`GET /api/forum/user/profile`
- 鉴权：免登录，支持可选登录

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| userId | uint64 | 是 | 用户 ID |

返回用户公开资料、个性签名、加入时间、文章数、文章总阅读量、关注数、粉丝数、收藏主题数和点赞主题数。携带有效登录令牌时，`isFollowed` 表示当前用户是否已关注目标用户，`isSelf` 表示是否为本人。

### 获取论坛用户动态

- 路径：`GET /api/forum/user/activity`
- 鉴权：免登录，支持可选登录

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| userId | uint64 | 是 | 用户 ID |
| type | string | 否 | `topic` 发文、`like` 点赞、`favorite` 收藏、`follow` 关注 |
| page | int | 否 | 页码，默认 1 |
| pageSize | int | 否 | 每页数量，默认 10 |

动态按发生时间倒序合并。`type` 为空时返回全部公开动态；文章、点赞和收藏动态通过 `topic` 返回关联主题，关注动态通过 `user` 返回被关注用户。已删除、隐藏或非正常状态的主题不会出现在动态中。

### 获取主题评论

- 路径：`GET /api/forum/comment/list`
- 鉴权：免登录

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| topicId | uint64 | 是 | 主题 ID |
| page | int | 否 | 页码 |
| pageSize | int | 否 | 每页数量 |
| sort | string | 否 | `latest` 最新，`hot` 最热，默认 `latest` |

只返回正常状态的顶层评论，并按 `sort` 分页。每条顶层评论通过 `children` 递归返回该主题下全部正常状态的子回复；`latest` 按发布时间倒序，`hot` 按点赞数、发布时间倒序。回复用户存在时通过 `replyUser` 返回。评论对象包含 `isLiked`；该接口支持可选登录，未登录时为 `false`，携带有效登录令牌时反映当前用户是否已点赞评论。

## 需要登录的接口

### 发布主题

- 路径：`POST /api/forum/topic/create`
- 鉴权：需要登录
- 内容类型：`application/json`

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| categoryId | uint64 | 是 | 启用的板块 ID |
| title | string | 是 | 标题长度 2 至 255 |
| keywords | string | 否 | 关键词，英文逗号分隔，最多 500 个字符 |
| description | string | 否 | 主题摘要，最多 500 个字符 |
| content | string | 是 | 正文长度 2 至 50000 |
| images | string | 否 | 图片地址，英文逗号分隔，最多 5000 个字符 |

主题创建成功后状态为 `normal`，对应板块的 `topicCount` 在同一事务中增加。

### 发布评论

- 路径：`POST /api/forum/comment/create`
- 鉴权：需要登录

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| topicId | uint64 | 是 | 正常状态的主题 ID |
| parentId | uint64 | 否 | 父评论 ID，用于回复评论 |
| content | string | 是 | 评论长度 1 至 2000 |

父评论必须属于当前主题且处于正常状态。评论创建成功后，主题评论数、最后回复用户和最后回复时间会在同一事务中更新。

### 删除评论

- 路径：`POST /api/forum/comment/delete`
- 鉴权：需要登录

```json
{"commentId":35}
```

仅评论作者可以删除自己的正常状态评论。删除采用软删除，并在同一事务中清除该评论点赞关系、更新主题 `commentCount` 和最新回复信息。存在未删除的子回复时，接口返回业务错误，需先删除子回复。响应：

```json
{"commentCount":5}
```

### 点赞或取消点赞主题

- 路径：`POST /api/forum/topic/toggleLike`
- 鉴权：需要登录

```json
{"topicId": 12}
```

每次请求切换当前用户与主题的点赞关系。响应：

```json
{"liked":true,"likeCount":8}
```

### 收藏或取消收藏主题

- 路径：`POST /api/forum/topic/toggleFavorite`
- 鉴权：需要登录

```json
{"topicId": 12}
```

每次请求切换当前用户与主题的收藏关系。响应：

```json
{"favorited":true,"favoriteCount":3}
```

### 点赞或取消点赞评论

- 路径：`POST /api/forum/comment/toggleLike`
- 鉴权：需要登录

```json
{"commentId": 35}
```

每次请求切换当前用户与评论的点赞关系。响应：

```json
{"liked":true,"likeCount":6}
```

主题点赞、主题收藏和评论点赞的数量与 `hg_forum_interaction` 中对应关系会在同一事务中更新，取消操作不会将计数减到负数。

### 关注或取消关注用户

- 路径：`POST /api/forum/user/toggleFollow`
- 鉴权：需要登录

```json
{"followUserId": 8}
```

每次请求切换当前用户对目标用户的关注关系。不能关注自己，目标用户必须存在。响应：

```json
{"followed":true,"followerCount":12}
```

### 提交举报

- 路径：`POST /api/forum/report/create`
- 鉴权：需要登录

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| targetType | string | 是 | `topic` 或 `comment` |
| targetId | uint64 | 是 | 被举报主题或评论 ID |
| reason | string | 是 | 举报原因，最多 100 个字符 |
| description | string | 否 | 补充说明，最多 1000 个字符 |

举报对象必须存在且处于公开可见状态。同一用户对同一对象存在待处理举报时不能重复提交。新举报状态为 `pending`。

## 主题响应示例

```json
{
  "topic": {
    "id": 12,
    "categoryId": 1,
    "userId": 8,
    "title": "如何配置项目环境？",
    "keywords": "项目,环境,配置",
    "description": "项目环境配置说明",
    "content": "这里是主题正文。",
    "images": "",
    "status": "normal",
    "isTop": 0,
    "isEssence": 0,
    "commentCount": 2,
    "likeCount": 8,
    "favoriteCount": 3,
    "viewCount": 45,
    "createdAt": "2026-09-20 09:00:00",
    "author": {
      "id": 8,
      "username": "demo",
      "nickname": "示例用户",
      "avatar": "https://example.com/avatar.png"
    },
    "isLiked": true,
    "isFavorited": false
  }
}
```
