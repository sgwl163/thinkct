package openai

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	v1 "hotgo/api/api/openai/v1"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	hgopenai "hotgo/utility/openai"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

type sOpenai struct{}

func NewOpenai() *sOpenai {
	return &sOpenai{}
}

func init() {
	service.RegisterOpenai(NewOpenai())
}

func getModelConfig(ctx context.Context) (apiKey, baseURL string, err error) {
	conf, err := service.SysConfig().GetOpenai(ctx)
	if err != nil {
		return "", "", err
	}
	if conf == nil || strings.TrimSpace(conf.ApiKey) == "" || strings.TrimSpace(conf.BaseUrl) == "" {
		return "", "", gerror.New("模型配置不完整，请在系统配置的模型配置中填写 API Key 和接口地址")
	}
	return strings.TrimSpace(conf.ApiKey), strings.TrimRight(strings.TrimSpace(conf.BaseUrl), "/"), nil
}

func getDefaultModel(ctx context.Context) string {
	conf, err := service.SysConfig().GetOpenai(ctx)
	if err == nil && conf != nil && strings.TrimSpace(conf.DefaultModel) != "" {
		return strings.TrimSpace(conf.DefaultModel)
	}
	return "gpt-3.5-turbo"
}

func urlToBase64DataURI(ctx context.Context, imgUrl string) string {
	if strings.HasPrefix(imgUrl, "data:image") {
		return imgUrl
	}

	// 使用 storager 获取完整的绝对路径 URL
	fullUrl := storager.LastUrl(ctx, imgUrl, consts.UploadDriveLocal)

	resp, err := g.Client().SetTimeout(10*time.Second).Get(ctx, fullUrl)
	if err != nil {
		g.Log().Warningf(ctx, "fetch image error: %v, url: %s", err, fullUrl)
		return imgUrl
	}
	defer resp.Close()

	if resp.StatusCode != 200 {
		g.Log().Warningf(ctx, "fetch image status error: %d, url: %s", resp.StatusCode, fullUrl)
		return imgUrl
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "image/jpeg"
	}

	body := resp.ReadAll()
	base64Str := base64.StdEncoding.EncodeToString(body)
	return fmt.Sprintf("data:%s;base64,%s", contentType, base64Str)
}

func (s *sOpenai) FolderList(ctx context.Context, req *v1.FolderListReq) (list []*entity.OpenaiChatFolder, err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	err = dao.OpenaiChatFolder.Ctx(ctx).
		Where(dao.OpenaiChatFolder.Columns().UserId, memberId).
		Where(dao.OpenaiChatFolder.Columns().DeletedAt, nil).
		OrderAsc(dao.OpenaiChatFolder.Columns().Sort).
		OrderDesc(dao.OpenaiChatFolder.Columns().Id).
		Scan(&list)
	return
}

func (s *sOpenai) FolderEdit(ctx context.Context, req *v1.FolderEditReq) (err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}

	if req.Id > 0 {
		_, err = dao.OpenaiChatFolder.Ctx(ctx).
			WherePri(req.Id).
			Where(dao.OpenaiChatFolder.Columns().UserId, memberId).
			Data(entity.OpenaiChatFolder{
				Name:      req.Name,
				Sort:      req.Sort,
				UpdatedAt: gtime.Now(),
			}).Update()
	} else {
		_, err = dao.OpenaiChatFolder.Ctx(ctx).Data(entity.OpenaiChatFolder{
			UserId:    memberId,
			Name:      req.Name,
			Sort:      req.Sort,
			Status:    1,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		}).Insert()
	}
	return
}

func (s *sOpenai) FolderDelete(ctx context.Context, req *v1.FolderDeleteReq) (err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}

	_, err = dao.OpenaiChatFolder.Ctx(ctx).
		WherePri(req.Id).
		Where(dao.OpenaiChatFolder.Columns().UserId, memberId).
		Data(g.Map{
			dao.OpenaiChatFolder.Columns().DeletedAt: gtime.Now(),
		}).Update()
	return
}

func (s *sOpenai) SessionList(ctx context.Context, req *v1.SessionListReq) (list []*entity.OpenaiChatSession, err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	m := dao.OpenaiChatSession.Ctx(ctx).
		Where(dao.OpenaiChatSession.Columns().UserId, memberId).
		Where(dao.OpenaiChatSession.Columns().DeletedAt, nil)

	if req.Keyword != "" {
		m = m.WhereLike(dao.OpenaiChatSession.Columns().Title, "%"+req.Keyword+"%")
	} else if req.FolderId > 0 {
		m = m.Where(dao.OpenaiChatSession.Columns().FolderId, req.FolderId)
	}

	err = m.OrderDesc(dao.OpenaiChatSession.Columns().IsTop).OrderDesc(dao.OpenaiChatSession.Columns().Id).Scan(&list)
	return
}

func (s *sOpenai) SessionEdit(ctx context.Context, req *v1.SessionEditReq) (res *v1.SessionEditRes, err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	res = &v1.SessionEditRes{Id: req.Id}

	if req.Id > 0 {
		_, err = dao.OpenaiChatSession.Ctx(ctx).
			WherePri(req.Id).
			Where(dao.OpenaiChatSession.Columns().UserId, memberId).
			Data(entity.OpenaiChatSession{
				FolderId:  req.FolderId,
				Title:     req.Title,
				Prompt:    req.Prompt,
				Model:     req.Model,
				UpdatedAt: gtime.Now(),
			}).Update()
	} else {
		insertRes, err := dao.OpenaiChatSession.Ctx(ctx).Data(entity.OpenaiChatSession{
			UserId:      memberId,
			FolderId:    req.FolderId,
			Title:       req.Title,
			Prompt:      req.Prompt,
			Model:       req.Model,
			Temperature: 0.7, // 设置新建会话时的默认温度值为 0.7
			Status:      1,
			CreatedAt:   gtime.Now(),
			UpdatedAt:   gtime.Now(),
		}).Insert()
		if err == nil {
			newId, _ := insertRes.LastInsertId()
			res.Id = newId
		}
	}
	return
}

func (s *sOpenai) SessionDelete(ctx context.Context, req *v1.SessionDeleteReq) (err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}

	_, err = dao.OpenaiChatSession.Ctx(ctx).
		WherePri(req.Id).
		Where(dao.OpenaiChatSession.Columns().UserId, memberId).
		Data(g.Map{
			dao.OpenaiChatSession.Columns().DeletedAt: gtime.Now(),
		}).Update()

	if err == nil {
		// 删除对话时，将其关联的消息内容也一块软删除
		_, _ = dao.OpenaiChatMessage.Ctx(ctx).
			Where(dao.OpenaiChatMessage.Columns().SessionId, req.Id).
			Where(dao.OpenaiChatMessage.Columns().UserId, memberId).
			Data(g.Map{
				dao.OpenaiChatMessage.Columns().DeletedAt: gtime.Now(),
			}).Update()
	}
	return
}

func (s *sOpenai) SessionTop(ctx context.Context, req *v1.SessionTopReq) (err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}
	_, err = dao.OpenaiChatSession.Ctx(ctx).
		WherePri(req.Id).
		Where(dao.OpenaiChatSession.Columns().UserId, memberId).
		Data(g.Map{
			dao.OpenaiChatSession.Columns().IsTop:     req.IsTop,
			dao.OpenaiChatSession.Columns().UpdatedAt: gtime.Now(),
		}).Update()
	return
}

func (s *sOpenai) SessionClear(ctx context.Context, req *v1.SessionClearReq) (err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}
	_, err = dao.OpenaiChatMessage.Ctx(ctx).
		Where(dao.OpenaiChatMessage.Columns().SessionId, req.Id).
		Where(dao.OpenaiChatMessage.Columns().UserId, memberId).
		Data(g.Map{
			dao.OpenaiChatMessage.Columns().DeletedAt: gtime.Now(),
		}).Update()
	return
}

func (s *sOpenai) SessionRename(ctx context.Context, req *v1.SessionRenameReq) (err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}
	_, err = dao.OpenaiChatSession.Ctx(ctx).
		WherePri(req.Id).
		Where(dao.OpenaiChatSession.Columns().UserId, memberId).
		Data(g.Map{
			dao.OpenaiChatSession.Columns().Title:     req.Title,
			dao.OpenaiChatSession.Columns().UpdatedAt: gtime.Now(),
		}).Update()
	return
}

func (s *sOpenai) MessageList(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	var list []*entity.OpenaiChatMessage
	err = dao.OpenaiChatMessage.Ctx(ctx).
		Where(dao.OpenaiChatMessage.Columns().UserId, memberId).
		Where(dao.OpenaiChatMessage.Columns().SessionId, req.SessionId).
		Where(dao.OpenaiChatMessage.Columns().DeletedAt, nil).
		OrderAsc(dao.OpenaiChatMessage.Columns().Id).
		Scan(&list)
	if err != nil {
		return nil, err
	}

	res = &v1.MessageListRes{
		List: list,
	}

	// 附加会话绑定的模型和系统提示词，方便前端回显
	var session *entity.OpenaiChatSession
	_ = dao.OpenaiChatSession.Ctx(ctx).WherePri(req.SessionId).Scan(&session)
	if session != nil {
		res.Model = session.Model
		res.Prompt = session.Prompt
		res.Temperature = session.Temperature
	}

	return
}

func (s *sOpenai) UploadFile(ctx context.Context, req *v1.UploadFileReq) (url string, err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return "", gerror.New("请先登录")
	}

	if req.File == nil {
		return "", gerror.New("没有找到上传的文件")
	}

	attachment, err := storager.DoUpload(ctx, "default", req.File)
	if err != nil {
		return "", err
	}

	url = storager.LastUrl(ctx, attachment.FileUrl, attachment.Drive)
	return url, nil
}

func (s *sOpenai) ModelList(ctx context.Context, req *v1.ModelListReq) (list []v1.ModelItem, err error) {
	var models []*entity.OpenaiApiModel
	err = dao.OpenaiApiModel.Ctx(ctx).
		Where(dao.OpenaiApiModel.Columns().Status, 1).
		Where(dao.OpenaiApiModel.Columns().DeletedAt, nil).
		OrderAsc(dao.OpenaiApiModel.Columns().Sort).
		Scan(&models)
	if err != nil {
		return nil, err
	}

	var groups []*entity.OpenaiApiGroup
	err = dao.OpenaiApiGroup.Ctx(ctx).
		Where(dao.OpenaiApiGroup.Columns().Status, 1).
		Where(dao.OpenaiApiGroup.Columns().DeletedAt, nil).
		Scan(&groups)
	if err != nil {
		return nil, err
	}

	groupMap := make(map[int64]*entity.OpenaiApiGroup)
	for _, g := range groups {
		groupMap[g.Id] = g
	}

	for _, m := range models {
		groupName := "未知分组"
		if g, ok := groupMap[m.GroupId]; ok {
			groupName = g.Name
		}
		list = append(list, v1.ModelItem{
			GroupId:   m.GroupId,
			GroupName: groupName,
			ModelId:   m.Id,
			Name:      m.Name,
			Model:     m.Model,
			IsVision:  m.IsVision,
			IsFile:    m.IsFile,
			IsSearch:  m.IsSearch,
		})
	}

	return list, nil
}

func (s *sOpenai) MessageSend(ctx context.Context, req *v1.MessageSendReq) (res *v1.MessageSendRes, err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	// =========================================================================
	// Step 0: 检查会话，如果未传 SessionId 则自动创建新会话
	// =========================================================================
	if req.SessionId <= 0 {
		defaultTemp := float64(0.7)
		if req.Temperature != nil {
			defaultTemp = float64(*req.Temperature)
		}
		defaultPrompt := ""
		if req.Prompt != nil {
			defaultPrompt = *req.Prompt
		}

		newSession := &entity.OpenaiChatSession{
			UserId:      memberId,
			FolderId:    req.FolderId, // 使用传入的 FolderId 进行分组，如果不传默认为 0
			Title:       "新对话",
			Prompt:      defaultPrompt,
			Model:       req.Model,
			Temperature: defaultTemp,
			Status:      1,
			CreatedAt:   gtime.Now(),
			UpdatedAt:   gtime.Now(),
		}

		insertRes, err := dao.OpenaiChatSession.Ctx(ctx).Data(newSession).Insert()
		if err != nil {
			return nil, gerror.Wrap(err, "自动创建新会话失败")
		}

		newSessionId, _ := insertRes.LastInsertId()
		req.SessionId = newSessionId
	}

	// =========================================================================
	// Step 1: 保存用户发送的消息 (支持多模态内容组装)
	// =========================================================================
	var dbContent string
	if len(req.Images) > 0 || len(req.Files) > 0 {
		dbContent = gjson.New(g.Map{
			"text":   req.Content,
			"images": req.Images,
			"files":  req.Files,
		}).MustToJsonString()
	} else {
		dbContent = req.Content
	}

	userMsg := &entity.OpenaiChatMessage{
		UserId:    memberId,
		SessionId: req.SessionId,
		Role:      "user",
		Content:   dbContent,
		Status:    1,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}
	_, err = dao.OpenaiChatMessage.Ctx(ctx).Data(userMsg).Insert()
	if err != nil {
		return nil, err
	}

	// =========================================================================
	// Step 2: 更新会话元数据 (模型切换、Prompt、最后更新时间等)
	// =========================================================================
	var session *entity.OpenaiChatSession
	err = dao.OpenaiChatSession.Ctx(ctx).WherePri(req.SessionId).Scan(&session)
	if err == nil && session != nil {
		updateData := g.Map{
			dao.OpenaiChatSession.Columns().UpdatedAt: gtime.Now(),
		}
		// 如果请求中传递了模型标识，则更新会话绑定的模型
		if req.Model != "" && req.Model != session.Model {
			updateData[dao.OpenaiChatSession.Columns().Model] = req.Model
			session.Model = req.Model
		}
		// 如果请求中传递了提示词，则更新会话绑定的提示词
		if req.Prompt != nil && *req.Prompt != session.Prompt {
			updateData[dao.OpenaiChatSession.Columns().Prompt] = *req.Prompt
			session.Prompt = *req.Prompt
		}
		// 如果请求中传递了温度值，则更新会话绑定的温度
		if req.Temperature != nil {
			updateData[dao.OpenaiChatSession.Columns().Temperature] = *req.Temperature
			session.Temperature = float64(*req.Temperature)
		}
		dao.OpenaiChatSession.Ctx(ctx).WherePri(req.SessionId).Data(updateData).Update()
	}

	// =========================================================================
	// Step 3: 首次对话自动生成标题 (后台异步执行，不阻塞当前请求)
	// =========================================================================
	if err == nil && session != nil && (session.Title == "" || session.Title == "新对话") {
		// 查询数据库中的模型和对应分组信息
		modelName := session.Model
		if modelName == "" {
			modelName = getDefaultModel(ctx)
		}
		var titleApiKey, titleBaseUrl string
		if titleApiKey, titleBaseUrl, err = getModelConfig(ctx); err != nil {
			g.Log().Warningf(ctx, "读取标题生成模型配置失败: %v", err)
		}

		if titleApiKey != "" {
			// 异步生成标题，不阻塞当前的对话流
			go func(sessId int64, firstMsg string, apiKey, baseUrl, model string) {
				bgCtx := context.Background() // 必须使用脱离当前请求的新上下文
				titleClient := hgopenai.NewClient(apiKey)
				if baseUrl != "" {
					titleClient.SetBaseURL(baseUrl)
				}

				titlePrompt := "请根据用户的这句话总结一个简短的标题，字数不超过10个字。只返回标题内容，不要带书名号、引号等其他多余字符。用户内容：" + firstMsg
				titleResult, tErr := titleClient.ChatCompletion(bgCtx, hgopenai.ChatCompletionRequest{
					Model: model,
					Messages: []hgopenai.Message{
						{Role: "user", Content: titlePrompt},
					},
				})

				if tErr == nil && titleResult != nil && titleResult.Content != "" {
					dao.OpenaiChatSession.Ctx(bgCtx).WherePri(sessId).Data(g.Map{
						dao.OpenaiChatSession.Columns().Title:     titleResult.Content,
						dao.OpenaiChatSession.Columns().UpdatedAt: gtime.Now(),
					}).Update()
				} else {
					g.Log().Errorf(bgCtx, "自动生成标题失败: %v", tErr)
				}
			}(req.SessionId, req.Content, titleApiKey, titleBaseUrl, modelName)
		}
	}

	// =========================================================================
	// Step 4: 获取历史消息，构建 OpenAI 请求上下文
	// =========================================================================
	var history []*entity.OpenaiChatMessage
	err = dao.OpenaiChatMessage.Ctx(ctx).
		Where(dao.OpenaiChatMessage.Columns().SessionId, req.SessionId).
		Where(dao.OpenaiChatMessage.Columns().Status, 1).
		OrderAsc(dao.OpenaiChatMessage.Columns().Id).
		Limit(20). // 限制获取最近 20 条消息作为上下文，防止Token超限
		Scan(&history)
	if err != nil {
		return nil, err
	}

	var messages []hgopenai.Message
	if session != nil && session.Prompt != "" {
		messages = append(messages, hgopenai.Message{
			Role:    "system",
			Content: session.Prompt,
		})
	}

	// 遍历并解析多模态消息格式
	for _, msg := range history {
		var content interface{}
		if gjson.Valid(msg.Content) {
			j := gjson.New(msg.Content)
			if j.Contains("images") || j.Contains("files") {
				text := j.Get("text").String()
				images := j.Get("images").Strings()
				files := j.Get("files").Strings()

				var contentArr []g.Map
				// 追加文本内容
				if text != "" {
					contentArr = append(contentArr, g.Map{"type": "text", "text": text})
				}
				// 追加普通文件内容（依赖于底层大模型对 `file_url` 块的支持）
				for _, fileUrl := range files {
					contentArr = append(contentArr, g.Map{
						"type":     "file_url",
						"file_url": g.Map{"url": fileUrl},
					})
				}
				// 追加图片内容（用于 GPT-4 Vision / Qwen-VL 等多模态模型）
				for _, imgUrl := range images {
					// 转换为 Base64，避免远端 API 无法访问本地图片导致超时
					b64Url := urlToBase64DataURI(ctx, imgUrl)
					contentArr = append(contentArr, g.Map{
						"type":      "image_url",
						"image_url": g.Map{"url": b64Url},
					})
				}
				content = contentArr
			} else {
				content = msg.Content
			}
		} else {
			content = msg.Content
		}

		messages = append(messages, hgopenai.Message{
			Role:    msg.Role,
			Content: content,
		})
	}

	// =========================================================================
	// Step 5: 读取模型接口组配置并初始化客户端
	// =========================================================================
	modelName := session.Model
	if modelName == "" {
		modelName = getDefaultModel(ctx)
	}

	apiKey, baseUrl, err := getModelConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := hgopenai.NewClient(apiKey)
	if baseUrl != "" {
		client.SetBaseURL(baseUrl)
	}

	// =========================================================================
	// Step 6: 拦截 HTTP 响应进行 SSE 流式输出
	// =========================================================================
	r := g.RequestFromCtx(ctx)
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")

	msgId := grand.Digits(17) // 生成类似 47135731852370690 的 message_id

	// 从会话中读取当前绑定的温度值，如果没有设置过（为0），则默认使用 0.7
	temperature := float32(session.Temperature)
	if temperature <= 0.01 { // 浮点数安全判断，兼容可能为 0 的历史老数据
		temperature = 0.7
	}

	reqData := hgopenai.ChatCompletionRequest{
		Model:       modelName,
		Messages:    messages,
		Temperature: temperature,
	}

	// 开启联网搜索特征支持
	if req.IsSearch {
		reqData.Tools = g.Slice{
			g.Map{
				"type": "web_search",
			},
		}
	}

	aiReplyContent, usage, err := client.ChatCompletionStream(ctx, reqData, func(chunk string) error {
		// 1. 输出 STREAM_CHUNK 事件 (配合前端某些特有的 PATCH OP 格式)
		streamChunkData := gjson.New(g.Map{
			"message_id": msgId,
			"patch_op": g.Slice{
				g.Map{
					"patch_object": 111,
					"patch_type":   1,
					"patch_value": g.Map{
						"tts_content": chunk,
					},
				},
			},
		}).MustToJsonString()
		r.Response.Writef("event: STREAM_CHUNK\ndata: %s\n\n", streamChunkData)

		// 2. 输出 CHUNK_DELTA 事件 (纯文本增量流)
		chunkDeltaData := gjson.New(g.Map{
			"text": chunk,
		}).MustToJsonString()
		r.Response.Writef("event: CHUNK_DELTA\ndata: %s\n\n", chunkDeltaData)

		r.Response.Flush() // 必须主动 Flush 推送到前端
		return nil
	})

	if err != nil {
		errData := gjson.New(g.Map{"error": err.Error()}).MustToJsonString()
		r.Response.Writef("event: ERROR\ndata: %s\n\n", errData)
		r.Response.Flush()
		r.ExitAll() // 终止框架中间件输出
		return nil, gerror.Wrap(err, "请求大模型接口失败")
	}

	// 输出结束标识
	r.Response.Write("event: DONE\ndata: {}\n\n")
	r.Response.Flush()

	// =========================================================================
	// Step 7: 保存完整的 AI 回复消息
	// =========================================================================
	replyMsg := &entity.OpenaiChatMessage{
		UserId:    memberId,
		SessionId: req.SessionId,
		Role:      "assistant",
		Content:   aiReplyContent,
		Model:     modelName,
		Status:    1,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}
	// 将 AI 返回的 Token 用量写入数据库
	if usage != nil {
		replyMsg.ReqTokens = usage.PromptTokens
		replyMsg.ResTokens = usage.CompletionTokens
		replyMsg.TotalTokens = usage.TotalTokens
	}
	insertResMsg, err := dao.OpenaiChatMessage.Ctx(ctx).Data(replyMsg).Insert()
	if err != nil {
		return nil, err
	}

	replyId, _ := insertResMsg.LastInsertId()
	replyMsg.Id = replyId

	// 中断后续的框架常规 JSON 响应输出，保证纯净的 SSE
	r.ExitAll()
	return &v1.MessageSendRes{
		SessionId: req.SessionId,
		FolderId:  session.FolderId,
		Reply:     replyMsg,
	}, nil
}

func (s *sOpenai) MessageDelete(ctx context.Context, req *v1.MessageDeleteReq) (err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}
	_, err = dao.OpenaiChatMessage.Ctx(ctx).
		WherePri(req.Id).
		Where(dao.OpenaiChatMessage.Columns().UserId, memberId).
		Data(g.Map{
			dao.OpenaiChatMessage.Columns().DeletedAt: gtime.Now(),
		}).Update()
	return
}

func (s *sOpenai) MessageEdit(ctx context.Context, req *v1.MessageEditReq) (err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}
	_, err = dao.OpenaiChatMessage.Ctx(ctx).
		WherePri(req.Id).
		Where(dao.OpenaiChatMessage.Columns().UserId, memberId).
		Data(g.Map{
			dao.OpenaiChatMessage.Columns().Content:   req.Content,
			dao.OpenaiChatMessage.Columns().UpdatedAt: gtime.Now(),
		}).Update()
	return
}

func (s *sOpenai) MessageFeedback(ctx context.Context, req *v1.MessageFeedbackReq) (err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}
	_, err = dao.OpenaiChatMessage.Ctx(ctx).
		WherePri(req.Id).
		Where(dao.OpenaiChatMessage.Columns().UserId, memberId).
		Data(g.Map{
			dao.OpenaiChatMessage.Columns().Feedback:  req.Feedback,
			dao.OpenaiChatMessage.Columns().UpdatedAt: gtime.Now(),
		}).Update()
	return
}

func (s *sOpenai) MessageRegenerate(ctx context.Context, req *v1.MessageRegenerateReq) (reply *entity.OpenaiChatMessage, err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	var session *entity.OpenaiChatSession
	err = dao.OpenaiChatSession.Ctx(ctx).WherePri(req.SessionId).Scan(&session)
	if err != nil || session == nil {
		return nil, gerror.New("会话不存在")
	}

	// =========================================================================
	// Step 0: 验证原消息是否存在并将其软删除
	// =========================================================================
	var oldMsg *entity.OpenaiChatMessage
	err = dao.OpenaiChatMessage.Ctx(ctx).WherePri(req.MessageId).Where(dao.OpenaiChatMessage.Columns().UserId, memberId).Scan(&oldMsg)
	if err != nil || oldMsg == nil {
		return nil, gerror.New("要重新生成的消息不存在或已被删除")
	}

	// 将旧的生成消息软删除，防止后续查询历史记录时再次包含它
	_, _ = dao.OpenaiChatMessage.Ctx(ctx).WherePri(req.MessageId).Data(g.Map{
		dao.OpenaiChatMessage.Columns().DeletedAt: gtime.Now(),
	}).Update()

	// =========================================================================
	// Step 1: 获取历史消息，构建 OpenAI 请求上下文
	// =========================================================================
	var history []*entity.OpenaiChatMessage
	err = dao.OpenaiChatMessage.Ctx(ctx).
		Where(dao.OpenaiChatMessage.Columns().SessionId, req.SessionId).
		Where(dao.OpenaiChatMessage.Columns().Status, 1).
		Where(dao.OpenaiChatMessage.Columns().DeletedAt, nil).
		WhereLTE(dao.OpenaiChatMessage.Columns().Id, req.MessageId). // 只取到当前要重新生成的消息之前（包括当前消息对应的用户提问）的历史记录
		OrderAsc(dao.OpenaiChatMessage.Columns().Id).
		Limit(20). // 取最近 20 条消息作为上下文
		Scan(&history)
	if err != nil {
		return nil, err
	}

	var messages []hgopenai.Message
	if session != nil && session.Prompt != "" {
		messages = append(messages, hgopenai.Message{
			Role:    "system",
			Content: session.Prompt,
		})
	}

	// 遍历并解析多模态消息格式
	for _, msg := range history {
		var content interface{}
		if gjson.Valid(msg.Content) {
			j := gjson.New(msg.Content)
			if j.Contains("images") || j.Contains("files") {
				text := j.Get("text").String()
				images := j.Get("images").Strings()
				files := j.Get("files").Strings()

				var contentArr []g.Map
				if text != "" {
					contentArr = append(contentArr, g.Map{"type": "text", "text": text})
				}
				for _, fileUrl := range files {
					contentArr = append(contentArr, g.Map{
						"type":     "file_url",
						"file_url": g.Map{"url": fileUrl},
					})
				}
				for _, imgUrl := range images {
					contentArr = append(contentArr, g.Map{
						"type":      "image_url",
						"image_url": g.Map{"url": imgUrl},
					})
				}
				content = contentArr
			} else {
				content = msg.Content
			}
		} else {
			content = msg.Content
		}

		messages = append(messages, hgopenai.Message{
			Role:    msg.Role,
			Content: content,
		})
	}

	// =========================================================================
	// Step 2: 读取 OpenAI 配置并初始化客户端
	// =========================================================================
	modelName := session.Model
	if modelName == "" {
		modelName = getDefaultModel(ctx)
	}

	apiKey, baseUrl, err := getModelConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := hgopenai.NewClient(apiKey)
	if baseUrl != "" {
		client.SetBaseURL(baseUrl)
	}

	// =========================================================================
	// Step 3: 调用 AI 接口生成流式回复
	// =========================================================================
	r := g.RequestFromCtx(ctx)
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")

	msgId := grand.Digits(17)

	// 从会话中读取当前绑定的温度值，如果没有设置过（为0），则默认使用 0.7
	temperature := float32(session.Temperature)
	if temperature <= 0.01 {
		temperature = 0.7
	}

	reqData := hgopenai.ChatCompletionRequest{
		Model:       modelName,
		Messages:    messages,
		Temperature: temperature,
	}

	// 开启联网搜索特征支持
	if req.IsSearch {
		reqData.Tools = g.Slice{
			g.Map{
				"type": "web_search",
			},
		}
	}

	aiReplyContent, usage, err := client.ChatCompletionStream(ctx, reqData, func(chunk string) error {
		// 1. 输出 STREAM_CHUNK 事件
		streamChunkData := gjson.New(g.Map{
			"message_id": msgId,
			"patch_op": g.Slice{
				g.Map{
					"patch_object": 111,
					"patch_type":   1,
					"patch_value": g.Map{
						"tts_content": chunk,
					},
				},
			},
		}).MustToJsonString()
		r.Response.Writef("event: STREAM_CHUNK\ndata: %s\n\n", streamChunkData)

		// 2. 输出 CHUNK_DELTA 事件
		chunkDeltaData := gjson.New(g.Map{
			"text": chunk,
		}).MustToJsonString()
		r.Response.Writef("event: CHUNK_DELTA\ndata: %s\n\n", chunkDeltaData)

		r.Response.Flush() // 必须主动 Flush 推送到前端
		return nil
	})

	if err != nil {
		errData := gjson.New(g.Map{"error": err.Error()}).MustToJsonString()
		r.Response.Writef("event: ERROR\ndata: %s\n\n", errData)
		r.Response.Flush()
		r.ExitAll() // 终止框架中间件输出
		return nil, gerror.Wrap(err, "请求大模型接口失败")
	}

	// 输出结束标识
	r.Response.Write("event: DONE\ndata: {}\n\n")
	r.Response.Flush()

	// =========================================================================
	// Step 7: 保存AI回复的消息
	// =========================================================================
	replyMsg := &entity.OpenaiChatMessage{
		UserId:    memberId,
		SessionId: req.SessionId,
		Role:      "assistant",
		Content:   aiReplyContent,
		Model:     modelName,
		Status:    1,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}
	// 将 AI 返回的 Token 用量写入数据库
	if usage != nil {
		replyMsg.ReqTokens = usage.PromptTokens
		replyMsg.ResTokens = usage.CompletionTokens
		replyMsg.TotalTokens = usage.TotalTokens
	}
	insertRes, err := dao.OpenaiChatMessage.Ctx(ctx).Data(replyMsg).Insert()
	if err != nil {
		return nil, err
	}

	replyId, _ := insertRes.LastInsertId()
	replyMsg.Id = replyId

	r.ExitAll() // 中断后续的框架常规 JSON 响应输出，保证纯净的 SSE
	return replyMsg, nil
}
