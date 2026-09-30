package openai

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
	"strings"
	"time"

	v1 "hotgo/api/api/openai/v1"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/storager"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	hgopenai "hotgo/utility/openai"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

func (s *sOpenai) DrawSubmit(ctx context.Context, req *v1.DrawTaskSubmitReq) (res *v1.DrawTaskSubmitRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	task := &entity.OpenaiDrawTask{
		UserId:         userId,
		Prompt:         req.Prompt,
		ReferenceImage: req.ReferenceImage,
		Model:          req.Model,
		Version:        req.Version,
		Ratio:          req.Ratio,
		Raw:            req.Raw,
		NegativePrompt: req.NegativePrompt,
		Stylize:        req.Stylize,
		Chaos:          req.Chaos,
		Seed:           req.Seed,
		Status:         0, // 0:进行中
		Progress:       0,
		CreatedAt:      gtime.Now(),
		UpdatedAt:      gtime.Now(),
	}

	insertRes, err := dao.OpenaiDrawTask.Ctx(ctx).Data(task).Insert()
	if err != nil {
		return nil, err
	}
	id, _ := insertRes.LastInsertId()

	// 异步调用绘画接口
	go s.doDrawTask(context.Background(), id)

	return &v1.DrawTaskSubmitRes{Id: id}, nil
}

type drawRequest struct {
	Model    string        `json:"model"`
	Messages []drawMessage `json:"messages"`
	Stream   bool          `json:"stream"`
}

type drawMessage struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}

func (s *sOpenai) doDrawTask(ctx context.Context, taskId int64) {
	// 标记为进行中
	dao.OpenaiDrawTask.Ctx(ctx).WherePri(taskId).Data(g.Map{"status": 0}).Update()

	var task *entity.OpenaiDrawTask
	err := dao.OpenaiDrawTask.Ctx(ctx).WherePri(taskId).Scan(&task)
	if err != nil || task == nil {
		return
	}

	apiKey, baseUrl, err := getModelConfig(ctx)
	if err != nil {
		s.markDrawFailed(ctx, taskId, "读取模型配置失败: "+err.Error())
		return
	}

	if apiKey == "" || baseUrl == "" {
		s.markDrawFailed(ctx, taskId, "绘画配置不完整(apiKey or baseUrl)")
		return
	}

	// 拼接比例参数
	size := task.Ratio
	if size == "" {
		size = "1:1"
	}

	if strings.HasPrefix(strings.ToLower(task.Model), "midjourney") || strings.HasPrefix(strings.ToLower(task.Model), "mj") {
		s.doMidjourneyTask(ctx, taskId, task, apiKey, baseUrl)
		return
	}

	var messageContent interface{}
	if task.ReferenceImage != "" {
		// 转换参考图为 Base64 以避免网络下载问题
		b64Url := urlToBase64DataURI(ctx, task.ReferenceImage)
		messageContent = []g.Map{
			{"type": "text", "text": task.Prompt},
			{"type": "image_url", "image_url": g.Map{"url": b64Url}},
		}
	} else {
		messageContent = task.Prompt
	}

	// 构建请求参数
	reqData := drawRequest{
		Model: task.Model,
		Messages: []drawMessage{
			{
				Role:    "user",
				Content: messageContent,
			},
		},
		Stream: false,
	}

	url := baseUrl + "/chat/completions"

	reqBody, _ := gjson.Encode(reqData)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		s.markDrawFailed(ctx, taskId, "构建请求失败: "+err.Error())
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * gtime.M,
	}
	resp, err := client.Do(httpReq)

	if err != nil {
		s.markDrawFailed(ctx, taskId, "请求AI接口异常: "+err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)
	if resp.StatusCode != 200 {
		s.markDrawFailed(ctx, taskId, "AI接口返回错误: "+body)
		return
	}

	j, err := gjson.DecodeToJson(body)
	if err != nil {
		s.markDrawFailed(ctx, taskId, "解析AI接口响应失败: "+err.Error())
		return
	}

	// 提取图片URL或内容
	content := j.Get("choices.0.message.content").String()
	if content == "" {
		// 兼容标准 images/generations 响应格式
		content = j.Get("data.0.url").String()
	}

	if content == "" {
		s.markDrawFailed(ctx, taskId, "AI接口未返回图片URL: "+body)
		return
	}

	imgUrl := content
	// 尝试从 Markdown 格式 ![xxx](url) 中提取纯 URL
	if strings.Contains(content, "](") {
		start := strings.Index(content, "](") + 2
		end := strings.Index(content[start:], ")")
		if end != -1 {
			imgUrl = content[start : start+end]
		}
	} else if strings.Contains(content, "http") {
		// 提取字符串中的 http 链接
		start := strings.Index(content, "http")
		end := strings.Index(content[start:], " ")
		if end != -1 {
			imgUrl = content[start : start+end]
		} else {
			imgUrl = content[start:]
		}
		imgUrl = strings.TrimSpace(imgUrl)
	}

	localImgUrl := s.downloadAndSaveImage(ctx, imgUrl, task.UserId)

	// 更新成功状态
	dao.OpenaiDrawTask.Ctx(ctx).WherePri(taskId).Data(g.Map{
		"status":     1, // 成功
		"image_url":  localImgUrl,
		"progress":   100,
		"updated_at": gtime.Now(),
	}).Update()
}

func (s *sOpenai) doMidjourneyTask(ctx context.Context, taskId int64, task *entity.OpenaiDrawTask, apiKey, baseUrl string) {
	// 1. 提交任务
	submitUrl := strings.ReplaceAll(baseUrl, "/v1", "") + "/mj/submit/imagine"
	prompt := task.Prompt
	if task.Ratio != "" {
		prompt += " --ar " + task.Ratio
	}
	if task.NegativePrompt != "" {
		prompt += " --no " + task.NegativePrompt
	}
	if task.Stylize > 0 {
		prompt += fmt.Sprintf(" --s %d", task.Stylize)
	}
	if task.Chaos > 0 {
		prompt += fmt.Sprintf(" --c %d", task.Chaos)
	}
	if task.Seed > 0 {
		prompt += fmt.Sprintf(" --seed %d", task.Seed)
	}
	reqData := g.Map{
		"prompt": prompt,
	}

	if task.ReferenceImage != "" {
		b64Url := urlToBase64DataURI(ctx, task.ReferenceImage)
		reqData["base64Array"] = []string{b64Url}
	}

	reqBody, _ := gjson.Encode(reqData)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", submitUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		s.markDrawFailed(ctx, taskId, "构建MJ请求失败: "+err.Error())
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("mj-api-secret", apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		s.markDrawFailed(ctx, taskId, "提交MJ任务异常: "+err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)

	j, err := gjson.DecodeToJson(body)
	if err != nil {
		s.markDrawFailed(ctx, taskId, "解析MJ提交响应失败: "+err.Error())
		return
	}

	// {"code":1,"description":"Submit Success","result":"1784915543848681"}
	code := j.Get("code").Int()
	if code != 1 && code != 22 { // 1=成功, 22=排队中
		s.markDrawFailed(ctx, taskId, "提交MJ任务失败: "+body)
		return
	}

	mjTaskId := j.Get("result").String()
	if mjTaskId == "" {
		s.markDrawFailed(ctx, taskId, "未获取到MJ任务ID: "+body)
		return
	}

	// 2. 轮询结果
	fetchUrl := strings.ReplaceAll(baseUrl, "/v1", "") + "/mj/task/" + mjTaskId + "/fetch"

	for i := 0; i < 60; i++ { // 最多轮询 60 * 5 = 300秒 = 5分钟
		time.Sleep(5 * time.Second)

		fetchReq, _ := http.NewRequestWithContext(ctx, "GET", fetchUrl, nil)
		fetchReq.Header.Set("Authorization", "Bearer "+apiKey)
		fetchReq.Header.Set("mj-api-secret", apiKey)

		fetchResp, err := client.Do(fetchReq)
		if err != nil {
			continue // 偶尔网络波动，继续重试
		}

		fetchBodyBytes, _ := io.ReadAll(fetchResp.Body)
		fetchResp.Body.Close()

		fj, err := gjson.DecodeToJson(string(fetchBodyBytes))
		if err != nil {
			continue
		}

		status := fj.Get("status").String()
		progressStr := fj.Get("progress").String()
		progress := 0
		if progressStr != "" {
			progressStr = strings.ReplaceAll(progressStr, "%", "")
			if p, err := strconv.Atoi(progressStr); err == nil {
				progress = p
			}
		}

		if status == "SUCCESS" {
			imgUrl := fj.Get("imageUrl").String()
			localImgUrl := s.downloadAndSaveImage(ctx, imgUrl, task.UserId)

			// 如果包含按钮数据，将其存为 JSON
			buttons := fj.Get("buttons").String() // 有些代理直接返回 buttons 数组

			// 更新成功状态
			dao.OpenaiDrawTask.Ctx(ctx).WherePri(taskId).Data(g.Map{
				"status":     1, // 成功
				"image_url":  localImgUrl,
				"mj_task_id": mjTaskId,
				"buttons":    buttons,
				"progress":   100,
				"updated_at": gtime.Now(),
			}).Update()
			return
		} else if status == "FAILED" {
			failReason := fj.Get("failReason").String()
			s.markDrawFailed(ctx, taskId, "MJ生成失败: "+failReason)
			return
		}

		// 其他状态 (IN_PROGRESS, SUBMITTED) 继续轮询
		if progress > 0 {
			dao.OpenaiDrawTask.Ctx(ctx).WherePri(taskId).Data(g.Map{
				"progress":   progress,
				"mj_task_id": mjTaskId,
				"updated_at": gtime.Now(),
			}).Update()
		}
	}

	s.markDrawFailed(ctx, taskId, "MJ任务生成超时")
}

func (s *sOpenai) DrawAction(ctx context.Context, req *v1.DrawTaskActionReq) (res *v1.DrawTaskActionRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	var task *entity.OpenaiDrawTask
	err = dao.OpenaiDrawTask.Ctx(ctx).
		Where(dao.OpenaiDrawTask.Columns().Id, req.Id).
		Where(dao.OpenaiDrawTask.Columns().UserId, userId).
		Scan(&task)

	if err != nil || task == nil {
		return nil, gerror.New("原绘画任务不存在或无权限")
	}

	if task.MjTaskId == "" {
		return nil, gerror.New("仅支持对 Midjourney 成功生成的任务进行二次操作")
	}

	newTask := &entity.OpenaiDrawTask{
		UserId:         userId,
		Prompt:         task.Prompt,
		ReferenceImage: task.ReferenceImage,
		Model:          task.Model,
		Version:        task.Version,
		Ratio:          task.Ratio,
		Raw:            task.Raw,
		NegativePrompt: task.NegativePrompt,
		Stylize:        task.Stylize,
		Chaos:          task.Chaos,
		Seed:           task.Seed,
		Status:         0, // 0:进行中
		Progress:       0,
		CreatedAt:      gtime.Now(),
		UpdatedAt:      gtime.Now(),
	}

	insertRes, err := dao.OpenaiDrawTask.Ctx(ctx).Data(newTask).Insert()
	if err != nil {
		return nil, err
	}
	newId, _ := insertRes.LastInsertId()

	go s.doMidjourneyActionTask(context.Background(), newId, task.MjTaskId, req.Action, task.Buttons)

	return &v1.DrawTaskActionRes{Id: newId}, nil
}

func (s *sOpenai) doMidjourneyActionTask(ctx context.Context, newTaskId int64, mjTaskId string, action string, buttonsJson string) {
	apiKey, baseUrl, err := getModelConfig(ctx)
	if err != nil {
		s.markDrawFailed(ctx, newTaskId, "读取模型配置失败: "+err.Error())
		return
	}

	if apiKey == "" || baseUrl == "" {
		s.markDrawFailed(ctx, newTaskId, "绘画配置不完整(apiKey or baseUrl)")
		return
	}

	submitUrl := strings.ReplaceAll(baseUrl, "/v1", "") + "/mj/submit/action"

	customId := action
	if buttonsJson != "" {
		j, err := gjson.DecodeToJson(buttonsJson)
		if err == nil {
			for _, btn := range j.Array() {
				btnMap := g.Map(btn.(map[string]interface{}))
				if fmt.Sprint(btnMap["label"]) == action || fmt.Sprint(btnMap["customId"]) == action {
					customId = fmt.Sprint(btnMap["customId"])
					break
				}
			}
		}
	}

	reqData := g.Map{
		"customId": customId,
		"taskId":   mjTaskId,
	}

	reqBody, _ := gjson.Encode(reqData)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", submitUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		s.markDrawFailed(ctx, newTaskId, "构建MJ Action请求失败: "+err.Error())
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("mj-api-secret", apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		s.markDrawFailed(ctx, newTaskId, "提交MJ Action任务异常: "+err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)

	j, err := gjson.DecodeToJson(body)
	if err != nil {
		s.markDrawFailed(ctx, newTaskId, "解析MJ Action响应失败: "+err.Error())
		return
	}

	code := j.Get("code").Int()
	if code != 1 && code != 22 {
		s.markDrawFailed(ctx, newTaskId, "提交MJ Action失败: "+body)
		return
	}

	newMjTaskId := j.Get("result").String()
	if newMjTaskId == "" {
		s.markDrawFailed(ctx, newTaskId, "未获取到新的MJ任务ID: "+body)
		return
	}

	// 轮询结果
	fetchUrl := strings.ReplaceAll(baseUrl, "/v1", "") + "/mj/task/" + newMjTaskId + "/fetch"

	for i := 0; i < 60; i++ {
		time.Sleep(5 * time.Second)

		fetchReq, _ := http.NewRequestWithContext(ctx, "GET", fetchUrl, nil)
		fetchReq.Header.Set("Authorization", "Bearer "+apiKey)
		fetchReq.Header.Set("mj-api-secret", apiKey)

		fetchResp, err := client.Do(fetchReq)
		if err != nil {
			continue
		}

		fetchBodyBytes, _ := io.ReadAll(fetchResp.Body)
		fetchResp.Body.Close()

		fj, err := gjson.DecodeToJson(string(fetchBodyBytes))
		if err != nil {
			continue
		}

		status := fj.Get("status").String()
		progressStr := fj.Get("progress").String()
		progress := 0
		if progressStr != "" {
			progressStr = strings.ReplaceAll(progressStr, "%", "")
			if p, err := strconv.Atoi(progressStr); err == nil {
				progress = p
			}
		}

		if status == "SUCCESS" {
			imgUrl := fj.Get("imageUrl").String()

			// 查找 userId
			var taskUserId int64
			var tempTask *entity.OpenaiDrawTask
			if err := dao.OpenaiDrawTask.Ctx(ctx).WherePri(newTaskId).Scan(&tempTask); err == nil && tempTask != nil {
				taskUserId = tempTask.UserId
			}

			localImgUrl := s.downloadAndSaveImage(ctx, imgUrl, taskUserId)

			buttons := fj.Get("buttons").String()

			dao.OpenaiDrawTask.Ctx(ctx).WherePri(newTaskId).Data(g.Map{
				"status":     1,
				"image_url":  localImgUrl,
				"mj_task_id": newMjTaskId,
				"buttons":    buttons,
				"progress":   100,
				"updated_at": gtime.Now(),
			}).Update()
			return
		} else if status == "FAILED" {
			failReason := fj.Get("failReason").String()
			s.markDrawFailed(ctx, newTaskId, "MJ Action生成失败: "+failReason)
			return
		}

		if progress > 0 {
			dao.OpenaiDrawTask.Ctx(ctx).WherePri(newTaskId).Data(g.Map{
				"progress":   progress,
				"mj_task_id": newMjTaskId,
				"updated_at": gtime.Now(),
			}).Update()
		}
	}

	s.markDrawFailed(ctx, newTaskId, "MJ Action任务生成超时")
}

func (s *sOpenai) DrawPromptTranslate(ctx context.Context, req *v1.DrawPromptTranslateReq) (res *v1.DrawPromptTranslateRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	apiKey, baseUrl, err := getModelConfig(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "读取模型配置失败")
	}

	client := hgopenai.NewClient(apiKey)
	if baseUrl != "" {
		client.SetBaseURL(baseUrl)
	}

	chatReq := hgopenai.ChatCompletionRequest{
		Model: "gpt-4o-mini", // 使用快速模型，若不可用代理会回退或者您可以改为默认的gpt-3.5-turbo
		Messages: []hgopenai.Message{
			{Role: "system", Content: "你是一个专业的AI绘画提示词翻译助手。请将用户输入的中文提示词翻译为对应的英文提示词，只需输出翻译后的英文，不要任何多余的解释。"},
			{Role: "user", Content: req.Prompt},
		},
	}

	chatRes, err := client.ChatCompletion(ctx, chatReq)
	if err != nil {
		// 回退尝试 gpt-3.5-turbo
		chatReq.Model = "gpt-3.5-turbo"
		chatRes, err = client.ChatCompletion(ctx, chatReq)
		if err != nil {
			return nil, gerror.New("翻译请求失败: " + err.Error())
		}
	}

	return &v1.DrawPromptTranslateRes{
		Result: chatRes.Content,
	}, nil
}

func (s *sOpenai) DrawPromptEnrich(ctx context.Context, req *v1.DrawPromptEnrichReq) (res *v1.DrawPromptEnrichRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	apiKey, baseUrl, err := getModelConfig(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "读取模型配置失败")
	}

	client := hgopenai.NewClient(apiKey)
	if baseUrl != "" {
		client.SetBaseURL(baseUrl)
	}

	systemPrompt := `你是一个Midjourney和Stable Diffusion等AI绘画工具的顶级提示词(Prompt)工程师。
用户会给你一个简短的描述，你需要发挥想象力，将其扩写为一个画面丰富、细节生动、高质量的英文提示词。
结构建议：主体描述 + 背景环境 + 光影效果 + 材质/风格 + 渲染参数(如 8k, highly detailed, masterpiece)。
请直接输出扩写后的英文提示词，不要输出多余的解释。`

	chatReq := hgopenai.ChatCompletionRequest{
		Model: "gpt-4o-mini",
		Messages: []hgopenai.Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: req.Prompt},
		},
	}

	chatRes, err := client.ChatCompletion(ctx, chatReq)
	if err != nil {
		chatReq.Model = "gpt-3.5-turbo"
		chatRes, err = client.ChatCompletion(ctx, chatReq)
		if err != nil {
			return nil, gerror.New("联想请求失败: " + err.Error())
		}
	}

	return &v1.DrawPromptEnrichRes{
		Result: chatRes.Content,
	}, nil
}

func (s *sOpenai) markDrawFailed(ctx context.Context, taskId int64, reason string) {
	dao.OpenaiDrawTask.Ctx(ctx).WherePri(taskId).Data(g.Map{
		"status":      2, // 失败
		"fail_reason": reason,
		"updated_at":  gtime.Now(),
	}).Update()
}

// downloadAndSaveImage 下载远端图片保存到本地，并返回相对路径
func (s *sOpenai) downloadAndSaveImage(ctx context.Context, imgUrl string, userId int64) string {
	if imgUrl == "" || strings.HasPrefix(imgUrl, "/") || strings.HasPrefix(imgUrl, "http://127.0.0.1") {
		return imgUrl
	}

	var bodyBytes []byte
	ext := ".png"
	ctype := "image/png"

	if strings.HasPrefix(imgUrl, "data:") {
		commaIdx := strings.Index(imgUrl, ",")
		if commaIdx == -1 {
			return imgUrl // 格式不对
		}

		meta := imgUrl[5:commaIdx]
		parts := strings.Split(meta, ";")
		if len(parts) > 0 {
			ctype = parts[0]
			if strings.Contains(ctype, "jpeg") || strings.Contains(ctype, "jpg") {
				ext = ".jpg"
			} else if strings.Contains(ctype, "gif") {
				ext = ".gif"
			} else if strings.Contains(ctype, "webp") {
				ext = ".webp"
			}
		}

		b64Data := imgUrl[commaIdx+1:]
		decoded, err := base64.StdEncoding.DecodeString(b64Data)
		if err != nil {
			g.Log().Warningf(ctx, "decode base64 image failed: %v", err)
			return imgUrl
		}
		bodyBytes = decoded
	} else {
		resp, err := g.Client().SetTimeout(30*time.Second).Get(ctx, imgUrl)
		if err != nil {
			g.Log().Warningf(ctx, "download image failed: %v, url: %s", err, imgUrl)
			return imgUrl
		}
		defer resp.Close()

		if resp.StatusCode != 200 {
			return imgUrl
		}

		bodyBytes, err = io.ReadAll(resp.Body)
		if err != nil || len(bodyBytes) == 0 {
			return imgUrl
		}

		ctype = resp.Header.Get("Content-Type")
		if strings.Contains(ctype, "jpeg") || strings.Contains(ctype, "jpg") {
			ext = ".jpg"
		} else if strings.Contains(ctype, "gif") {
			ext = ".gif"
		} else if strings.Contains(ctype, "webp") {
			ext = ".webp"
		}
	}

	fileName := grand.S(16) + ext

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName))
	h.Set("Content-Type", ctype)
	part, err := w.CreatePart(h)
	if err != nil {
		g.Log().Warningf(ctx, "CreatePart err: %v", err)
		return imgUrl
	}
	part.Write(bodyBytes)
	w.Close()

	reader := multipart.NewReader(&b, w.Boundary())
	form, err := reader.ReadForm(int64(len(bodyBytes)) + 1024)
	if err != nil {
		g.Log().Warningf(ctx, "ReadForm err: %v", err)
		return imgUrl
	}

	fileHeaders := form.File["file"]
	if len(fileHeaders) == 0 {
		return imgUrl
	}

	uf := &ghttp.UploadFile{
		FileHeader: fileHeaders[0],
	}

	mockCtx := context.WithValue(ctx, consts.ContextHTTPKey, &model.Context{
		User: &model.Identity{
			Id: userId,
		},
		Module: "api",
	})

	attachment, err := storager.DoUpload(mockCtx, storager.KindImg, uf)
	if err != nil || attachment == nil {
		g.Log().Warningf(ctx, "DoUpload failed: %v", err)
		return imgUrl
	}

	return storager.LastUrl(ctx, attachment.FileUrl, attachment.Drive)
}

func (s *sOpenai) DrawDelete(ctx context.Context, req *v1.DrawTaskDeleteReq) (res *v1.DrawTaskDeleteRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	_, err = dao.OpenaiDrawTask.Ctx(ctx).
		Where(dao.OpenaiDrawTask.Columns().Id, req.Id).
		Where(dao.OpenaiDrawTask.Columns().UserId, userId).
		Delete()

	if err != nil {
		return nil, err
	}

	return &v1.DrawTaskDeleteRes{}, nil
}
func (s *sOpenai) DrawList(ctx context.Context, req *v1.DrawTaskListReq) (res *v1.DrawTaskListRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	res = &v1.DrawTaskListRes{
		List: make([]*v1.DrawTaskItem, 0),
	}

	m := dao.OpenaiDrawTask.Ctx(ctx).Where(dao.OpenaiDrawTask.Columns().UserId, userId)
	total, err := m.Count()
	if err != nil {
		return nil, err
	}
	res.Total = total

	var tasks []*entity.OpenaiDrawTask
	err = m.Page(req.Page, req.Size).OrderDesc(dao.OpenaiDrawTask.Columns().Id).Scan(&tasks)
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		res.List = append(res.List, &v1.DrawTaskItem{
			Id:             t.Id,
			Prompt:         t.Prompt,
			ReferenceImage: t.ReferenceImage,
			Model:          t.Model,
			Ratio:          t.Ratio,
			ImageUrl:       t.ImageUrl,
			Status:         t.Status,
			FailReason:     t.FailReason,
			CreatedAt:      t.CreatedAt.Format("Y-m-d H:i:s"),
			MjTaskId:       t.MjTaskId,
			Buttons:        t.Buttons,
			Progress:       t.Progress,
		})
	}

	return res, nil
}

func (s *sOpenai) DrawSquare(ctx context.Context, req *v1.DrawSquareListReq) (res *v1.DrawSquareListRes, err error) {
	res = &v1.DrawSquareListRes{
		List: make([]*v1.DrawTaskItem, 0),
	}

	// 广场只展示成功的作品
	m := dao.OpenaiDrawTask.Ctx(ctx).Where(dao.OpenaiDrawTask.Columns().Status, 1)

	total, err := m.Count()
	if err != nil {
		return nil, err
	}
	res.Total = total

	var tasks []*entity.OpenaiDrawTask
	err = m.Page(req.Page, req.Size).OrderDesc(dao.OpenaiDrawTask.Columns().Id).Scan(&tasks)
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		res.List = append(res.List, &v1.DrawTaskItem{
			Id:             t.Id,
			Prompt:         t.Prompt,
			ReferenceImage: t.ReferenceImage,
			Model:          t.Model,
			Ratio:          t.Ratio,
			ImageUrl:       t.ImageUrl,
			Status:         t.Status,
			FailReason:     t.FailReason,
			CreatedAt:      t.CreatedAt.Format("Y-m-d H:i:s"),
			MjTaskId:       t.MjTaskId,
			Buttons:        t.Buttons,
			Progress:       t.Progress,
		})
	}

	return res, nil
}

func (s *sOpenai) DrawDetail(ctx context.Context, req *v1.DrawDetailReq) (res *v1.DrawDetailRes, err error) {
	var task *entity.OpenaiDrawTask
	err = dao.OpenaiDrawTask.Ctx(ctx).Where(dao.OpenaiDrawTask.Columns().Id, req.Id).Scan(&task)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, gerror.New("绘画任务不存在")
	}

	res = &v1.DrawDetailRes{
		DrawTaskItem: &v1.DrawTaskItem{
			Id:             task.Id,
			Prompt:         task.Prompt,
			ReferenceImage: task.ReferenceImage,
			Model:          task.Model,
			Ratio:          task.Ratio,
			ImageUrl:       task.ImageUrl,
			Status:         task.Status,
			FailReason:     task.FailReason,
			CreatedAt:      task.CreatedAt.Format("Y-m-d H:i:s"),
			MjTaskId:       task.MjTaskId,
			Buttons:        task.Buttons,
			Progress:       task.Progress,
		},
	}
	return res, nil
}
