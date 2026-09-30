package openai

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	v1 "hotgo/api/api/openai/v1"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/storager"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

func (s *sOpenai) VideoSubmit(ctx context.Context, req *v1.VideoTaskSubmitReq) (res *v1.VideoTaskSubmitRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	generateAudio := req.GenerateAudio
	if generateAudio == 0 {
		generateAudio = 1 // 默认生成音频
	}

	task := &entity.OpenaiVideoTask{
		UserId:         userId,
		Prompt:         req.Prompt,
		ReferenceImage: req.ReferenceImage,
		LastFrame:      req.LastFrame,
		VoiceDubUrl:    req.VoiceDubUrl,
		BgMusicUrl:     req.BgMusicUrl,
		GenerateAudio:  generateAudio,
		Model:          req.Model,
		Status:         0, // 0:进行中
		CreatedAt:      gtime.Now(),
		UpdatedAt:      gtime.Now(),
	}

	insertRes, err := dao.OpenaiVideoTask.Ctx(ctx).Data(task).Insert()
	if err != nil {
		return nil, err
	}
	id, _ := insertRes.LastInsertId()

	go s.doVideoTask(context.Background(), id)

	return &v1.VideoTaskSubmitRes{Id: id}, nil
}

func (s *sOpenai) markVideoFailed(ctx context.Context, taskId int64, reason string) {
	dao.OpenaiVideoTask.Ctx(ctx).WherePri(taskId).Data(g.Map{
		"status":      2, // 失败
		"fail_reason": reason,
		"updated_at":  gtime.Now(),
	}).Update()
}

func (s *sOpenai) doVideoTask(ctx context.Context, taskId int64) {
	var task *entity.OpenaiVideoTask
	err := dao.OpenaiVideoTask.Ctx(ctx).WherePri(taskId).Scan(&task)
	if err != nil || task == nil {
		return
	}

	apiKey, baseUrl, err := getModelConfig(ctx)
	if err != nil {
		s.markVideoFailed(ctx, taskId, "读取模型配置失败: "+err.Error())
		return
	}

	if apiKey == "" || baseUrl == "" {
		s.markVideoFailed(ctx, taskId, "配置不完整(apiKey or baseUrl)")
		return
	}

	// Submit URL: e.g. https://api.gpt.ge/task/volces/seedance
	submitUrl := strings.ReplaceAll(baseUrl, "/v1", "") + "/task/volces/seedance"

	contentList := []g.Map{
		{"type": "text", "text": task.Prompt},
	}

	// 首帧参考图
	if task.ReferenceImage != "" {
		contentList = append(contentList, g.Map{
			"type": "image_url",
			"image_url": g.Map{
				"url": task.ReferenceImage,
			},
			"role": "first_frame",
		})
	}

	// 尾帧图片（仅首尾帧模式）
	if task.LastFrame != "" {
		contentList = append(contentList, g.Map{
			"type": "image_url",
			"image_url": g.Map{
				"url": task.LastFrame,
			},
			"role": "last_frame",
		})
	}

	reqData := g.Map{
		"model":   task.Model,
		"content": contentList,
	}

	// generate_audio 仅 Seedance 1.5 pro 系列支持
	if strings.Contains(task.Model, "1-5") {
		reqData["generate_audio"] = task.GenerateAudio == 1
	}

	reqBody, _ := gjson.Encode(reqData)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", submitUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		s.markVideoFailed(ctx, taskId, "构建请求失败: "+err.Error())
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		s.markVideoFailed(ctx, taskId, "请求提交接口异常: "+err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)
	if resp.StatusCode != 200 {
		s.markVideoFailed(ctx, taskId, "接口返回错误: "+body)
		return
	}

	j, err := gjson.DecodeToJson(body)
	if err != nil {
		s.markVideoFailed(ctx, taskId, "解析响应失败: "+err.Error())
		return
	}

	remoteTaskId := j.Get("id").String()
	if remoteTaskId == "" {
		s.markVideoFailed(ctx, taskId, "未获取到远端任务ID: "+body)
		return
	}

	dao.OpenaiVideoTask.Ctx(ctx).WherePri(taskId).Data(g.Map{
		"remote_task_id": remoteTaskId,
		"updated_at":     gtime.Now(),
	}).Update()

	// Polling the result
	fetchUrl := strings.ReplaceAll(baseUrl, "/v1", "") + "/task/" + remoteTaskId

	for i := 0; i < 120; i++ { // Polling up to 10 minutes
		time.Sleep(5 * time.Second)

		fetchReq, _ := http.NewRequestWithContext(ctx, "GET", fetchUrl, nil)
		fetchReq.Header.Set("Authorization", "Bearer "+apiKey)

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
		if status == "succeeded" {
			videoUrl := fj.Get("content.video_url").String()
			localVideoUrl := s.downloadAndSaveVideo(ctx, videoUrl, task.UserId)

			dao.OpenaiVideoTask.Ctx(ctx).WherePri(taskId).Data(g.Map{
				"status":     1, // 成功
				"video_url":  localVideoUrl,
				"updated_at": gtime.Now(),
			}).Update()
			return
		} else if status == "failed" {
			s.markVideoFailed(ctx, taskId, "视频生成失败")
			return
		}
	}

	s.markVideoFailed(ctx, taskId, "视频任务生成超时")
}

func (s *sOpenai) downloadAndSaveVideo(ctx context.Context, videoUrl string, userId int64) string {
	if videoUrl == "" || strings.HasPrefix(videoUrl, "/") || strings.HasPrefix(videoUrl, "http://127.0.0.1") {
		return videoUrl
	}

	resp, err := g.Client().SetTimeout(5*time.Minute).Get(ctx, videoUrl)
	if err != nil {
		g.Log().Warningf(ctx, "download video failed: %v, url: %s", err, videoUrl)
		return videoUrl
	}
	defer resp.Close()

	if resp.StatusCode != 200 {
		return videoUrl
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil || len(bodyBytes) == 0 {
		return videoUrl
	}

	ctype := resp.Header.Get("Content-Type")
	ext := ".mp4"
	if strings.Contains(ctype, "webm") {
		ext = ".webm"
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
		return videoUrl
	}
	part.Write(bodyBytes)
	w.Close()

	reader := multipart.NewReader(&b, w.Boundary())
	form, err := reader.ReadForm(int64(len(bodyBytes)) + 1024)
	if err != nil {
		g.Log().Warningf(ctx, "ReadForm err: %v", err)
		return videoUrl
	}

	fileHeaders := form.File["file"]
	if len(fileHeaders) == 0 {
		return videoUrl
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

	attachment, err := storager.DoUpload(mockCtx, storager.KindVideo, uf)
	if err != nil || attachment == nil {
		g.Log().Warningf(ctx, "DoUpload failed: %v", err)
		return videoUrl
	}

	return storager.LastUrl(ctx, attachment.FileUrl, attachment.Drive)
}

func (s *sOpenai) VideoList(ctx context.Context, req *v1.VideoTaskListReq) (res *v1.VideoTaskListRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	res = &v1.VideoTaskListRes{
		List: make([]*v1.VideoTaskItem, 0),
	}

	m := dao.OpenaiVideoTask.Ctx(ctx).Where(dao.OpenaiVideoTask.Columns().UserId, userId)
	total, err := m.Count()
	if err != nil {
		return nil, err
	}
	res.Total = total

	var tasks []*entity.OpenaiVideoTask
	err = m.Page(req.Page, req.Size).OrderDesc(dao.OpenaiVideoTask.Columns().Id).Scan(&tasks)
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		res.List = append(res.List, toVideoTaskItem(t))
	}

	return res, nil
}

func (s *sOpenai) VideoSquare(ctx context.Context, req *v1.VideoSquareListReq) (res *v1.VideoSquareListRes, err error) {
	res = &v1.VideoSquareListRes{
		List: make([]*v1.VideoTaskItem, 0),
	}

	m := dao.OpenaiVideoTask.Ctx(ctx).Where(dao.OpenaiVideoTask.Columns().Status, 1)

	total, err := m.Count()
	if err != nil {
		return nil, err
	}
	res.Total = total

	var tasks []*entity.OpenaiVideoTask
	err = m.Page(req.Page, req.Size).OrderDesc(dao.OpenaiVideoTask.Columns().Id).Scan(&tasks)
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		res.List = append(res.List, toVideoTaskItem(t))
	}

	return res, nil
}

func (s *sOpenai) VideoDetail(ctx context.Context, req *v1.VideoDetailReq) (res *v1.VideoDetailRes, err error) {
	var task *entity.OpenaiVideoTask
	err = dao.OpenaiVideoTask.Ctx(ctx).Where(dao.OpenaiVideoTask.Columns().Id, req.Id).Scan(&task)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, gerror.New("视频创作任务不存在")
	}

	res = &v1.VideoDetailRes{
		VideoTaskItem: toVideoTaskItem(task),
	}
	return res, nil
}

func (s *sOpenai) VideoDelete(ctx context.Context, req *v1.VideoTaskDeleteReq) (res *v1.VideoTaskDeleteRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	_, err = dao.OpenaiVideoTask.Ctx(ctx).
		Where(dao.OpenaiVideoTask.Columns().Id, req.Id).
		Where(dao.OpenaiVideoTask.Columns().UserId, userId).
		Delete()

	if err != nil {
		return nil, err
	}

	return &v1.VideoTaskDeleteRes{}, nil
}

func toVideoTaskItem(t *entity.OpenaiVideoTask) *v1.VideoTaskItem {
	return &v1.VideoTaskItem{
		Id:             t.Id,
		Prompt:         t.Prompt,
		ReferenceImage: t.ReferenceImage,
		LastFrame:      t.LastFrame,
		VoiceDubUrl:    t.VoiceDubUrl,
		BgMusicUrl:     t.BgMusicUrl,
		GenerateAudio:  t.GenerateAudio,
		Model:          t.Model,
		VideoUrl:       t.VideoUrl,
		Status:         t.Status,
		FailReason:     t.FailReason,
		RemoteTaskId:   t.RemoteTaskId,
		CreatedAt:      t.CreatedAt.Format("Y-m-d H:i:s"),
	}
}
