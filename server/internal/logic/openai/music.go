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

func (s *sOpenai) MusicSubmit(ctx context.Context, req *v1.MusicTaskSubmitReq) (res *v1.MusicTaskSubmitRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	if req.Type == 1 && req.Prompt == "" {
		return nil, gerror.New("灵感描述不能为空")
	}
	if req.Type == 2 && req.MakeInstrumental == 0 && req.Prompt == "" {
		return nil, gerror.New("歌词内容不能为空")
	}
	if req.Type == 2 && req.MakeInstrumental == 1 && req.Tags == "" {
		return nil, gerror.New("纯音乐需要提供风格标签")
	}

	task := &entity.OpenaiMusicTask{
		UserId:           userId,
		Prompt:           req.Prompt,
		Type:             req.Type,
		MakeInstrumental: req.MakeInstrumental,
		Title:            req.Title,
		Tags:             req.Tags,
		Model:            req.Model,
		Status:           0, // 0:进行中
		CreatedAt:        gtime.Now(),
		UpdatedAt:        gtime.Now(),
	}

	insertRes, err := dao.OpenaiMusicTask.Ctx(ctx).Data(task).Insert()
	if err != nil {
		return nil, err
	}
	id, _ := insertRes.LastInsertId()

	go s.doMusicTask(context.Background(), id)

	return &v1.MusicTaskSubmitRes{Id: id}, nil
}

func (s *sOpenai) markMusicFailed(ctx context.Context, taskId int64, reason string) {
	dao.OpenaiMusicTask.Ctx(ctx).WherePri(taskId).Data(g.Map{
		"status":      2, // 失败
		"fail_reason": reason,
		"updated_at":  gtime.Now(),
	}).Update()
}

func (s *sOpenai) doMusicTask(ctx context.Context, taskId int64) {
	var task *entity.OpenaiMusicTask
	err := dao.OpenaiMusicTask.Ctx(ctx).WherePri(taskId).Scan(&task)
	if err != nil || task == nil {
		return
	}

	apiKey, baseUrl, err := getModelConfig(ctx)
	if err != nil {
		s.markMusicFailed(ctx, taskId, "读取模型配置失败: "+err.Error())
		return
	}

	if apiKey == "" || baseUrl == "" {
		s.markMusicFailed(ctx, taskId, "配置不完整(apiKey or baseUrl)")
		return
	}

	submitUrl := strings.ReplaceAll(baseUrl, "/v1", "") + "/suno/submit/music"

	reqData := g.Map{}
	if task.Type == 1 { // 灵感模式
		reqData["gpt_description_prompt"] = task.Prompt
		reqData["make_instrumental"] = task.MakeInstrumental == 1
		reqData["mv"] = task.Model
	} else { // 定制模式
		if task.MakeInstrumental == 1 {
			reqData["prompt"] = ""
		} else {
			reqData["prompt"] = task.Prompt
		}
		reqData["tags"] = task.Tags
		reqData["title"] = task.Title
		reqData["mv"] = task.Model
	}

	reqBody, _ := gjson.Encode(reqData)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", submitUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		s.markMusicFailed(ctx, taskId, "构建请求失败: "+err.Error())
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		s.markMusicFailed(ctx, taskId, "请求提交接口异常: "+err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)
	if resp.StatusCode != 200 {
		s.markMusicFailed(ctx, taskId, "接口返回错误: "+body)
		return
	}

	j, err := gjson.DecodeToJson(body)
	if err != nil {
		s.markMusicFailed(ctx, taskId, "解析响应失败: "+err.Error())
		return
	}

	if j.Get("code").String() != "success" {
		s.markMusicFailed(ctx, taskId, "远端报错: "+j.Get("message").String())
		return
	}

	remoteTaskId := j.Get("data").String()
	if remoteTaskId == "" {
		s.markMusicFailed(ctx, taskId, "未获取到远端任务ID: "+body)
		return
	}

	dao.OpenaiMusicTask.Ctx(ctx).WherePri(taskId).Data(g.Map{
		"remote_task_id": remoteTaskId,
		"updated_at":     gtime.Now(),
	}).Update()

	// Polling the result
	fetchUrl := strings.ReplaceAll(baseUrl, "/v1", "") + "/suno/fetch"

	for i := 0; i < 120; i++ { // Polling up to 10 minutes
		time.Sleep(5 * time.Second)

		fetchReqData := g.Map{
			"ids":    []string{remoteTaskId},
			"action": "MUSIC",
		}
		fetchReqBody, _ := gjson.Encode(fetchReqData)
		fetchReq, _ := http.NewRequestWithContext(ctx, "POST", fetchUrl, bytes.NewBuffer(fetchReqBody))
		fetchReq.Header.Set("Authorization", "Bearer "+apiKey)
		fetchReq.Header.Set("Content-Type", "application/json")

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

		status := fj.Get("data.0.status").String()

		if status == "SUCCESS" {
			clips := fj.Get("data.0.data").Array()
			if len(clips) == 0 {
				s.markMusicFailed(ctx, taskId, "生成成功但未返回音频数据")
				return
			}

			for idx, clipObj := range clips {
				clip := gjson.New(clipObj)
				clipId := clip.Get("id").String()
				audioUrl := clip.Get("audio_url").String()
				imageUrl := clip.Get("image_url").String()
				videoUrl := clip.Get("video_url").String()
				title := clip.Get("title").String()
				tags := clip.Get("metadata.tags").String()
				duration := clip.Get("metadata.duration").Float64()

				localAudioUrl := s.downloadAndSaveMedia(ctx, audioUrl, task.UserId, storager.KindAudio)
				localImageUrl := s.downloadAndSaveMedia(ctx, imageUrl, task.UserId, storager.KindImg)
				localVideoUrl := s.downloadAndSaveMedia(ctx, videoUrl, task.UserId, storager.KindVideo)

				if idx == 0 {
					// 更新原任务
					dao.OpenaiMusicTask.Ctx(ctx).WherePri(taskId).Data(g.Map{
						"status":         1, // 成功
						"remote_clip_id": clipId,
						"audio_url":      localAudioUrl,
						"image_url":      localImageUrl,
						"video_url":      localVideoUrl,
						"title":          title,
						"tags":           tags,
						"duration":       duration,
						"updated_at":     gtime.Now(),
					}).Update()
				} else {
					// 新增任务
					newTask := &entity.OpenaiMusicTask{
						UserId:           task.UserId,
						Prompt:           task.Prompt,
						Type:             task.Type,
						MakeInstrumental: task.MakeInstrumental,
						Title:            title,
						Tags:             tags,
						Model:            task.Model,
						Status:           1, // 成功
						RemoteTaskId:     remoteTaskId,
						RemoteClipId:     clipId,
						AudioUrl:         localAudioUrl,
						ImageUrl:         localImageUrl,
						VideoUrl:         localVideoUrl,
						Duration:         duration,
						CreatedAt:        gtime.Now(),
						UpdatedAt:        gtime.Now(),
					}
					dao.OpenaiMusicTask.Ctx(ctx).Data(newTask).Insert()
				}
			}
			return
		} else if status == "FAILED" {
			failReason := fj.Get("data.0.fail_reason").String()
			if failReason == "" {
				failReason = "音乐生成失败"
			}
			s.markMusicFailed(ctx, taskId, failReason)
			return
		}
	}

	s.markMusicFailed(ctx, taskId, "音乐任务生成超时")
}

func (s *sOpenai) downloadAndSaveMedia(ctx context.Context, mediaUrl string, userId int64, kind string) string {
	if mediaUrl == "" || strings.HasPrefix(mediaUrl, "/") || strings.HasPrefix(mediaUrl, "http://127.0.0.1") {
		return mediaUrl
	}

	resp, err := g.Client().SetTimeout(5*time.Minute).Get(ctx, mediaUrl)
	if err != nil {
		g.Log().Warningf(ctx, "download media failed: %v, url: %s", err, mediaUrl)
		return mediaUrl
	}
	defer resp.Close()

	if resp.StatusCode != 200 {
		return mediaUrl
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil || len(bodyBytes) == 0 {
		return mediaUrl
	}

	ctype := resp.Header.Get("Content-Type")
	ext := ".mp3"
	if kind == storager.KindImg {
		ext = ".jpeg"
		if strings.Contains(ctype, "png") {
			ext = ".png"
		}
	} else if kind == storager.KindVideo {
		ext = ".mp4"
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
		return mediaUrl
	}
	part.Write(bodyBytes)
	w.Close()

	reader := multipart.NewReader(&b, w.Boundary())
	form, err := reader.ReadForm(int64(len(bodyBytes)) + 1024)
	if err != nil {
		g.Log().Warningf(ctx, "ReadForm err: %v", err)
		return mediaUrl
	}

	fileHeaders := form.File["file"]
	if len(fileHeaders) == 0 {
		return mediaUrl
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

	attachment, err := storager.DoUpload(mockCtx, kind, uf)
	if err != nil || attachment == nil {
		g.Log().Warningf(ctx, "DoUpload failed: %v", err)
		return mediaUrl
	}

	return storager.LastUrl(ctx, attachment.FileUrl, attachment.Drive)
}

func (s *sOpenai) MusicList(ctx context.Context, req *v1.MusicTaskListReq) (res *v1.MusicTaskListRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	res = &v1.MusicTaskListRes{
		List: make([]*v1.MusicTaskItem, 0),
	}

	m := dao.OpenaiMusicTask.Ctx(ctx).Where(dao.OpenaiMusicTask.Columns().UserId, userId)
	total, err := m.Count()
	if err != nil {
		return nil, err
	}
	res.Total = total

	var tasks []*entity.OpenaiMusicTask
	err = m.Page(req.Page, req.Size).OrderDesc(dao.OpenaiMusicTask.Columns().Id).Scan(&tasks)
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		res.List = append(res.List, &v1.MusicTaskItem{
			Id:               t.Id,
			Prompt:           t.Prompt,
			Type:             t.Type,
			MakeInstrumental: t.MakeInstrumental,
			Title:            t.Title,
			Tags:             t.Tags,
			Model:            t.Model,
			AudioUrl:         t.AudioUrl,
			ImageUrl:         t.ImageUrl,
			VideoUrl:         t.VideoUrl,
			Duration:         t.Duration,
			Status:           t.Status,
			FailReason:       t.FailReason,
			RemoteTaskId:     t.RemoteTaskId,
			RemoteClipId:     t.RemoteClipId,
			CreatedAt:        t.CreatedAt.Format("Y-m-d H:i:s"),
		})
	}

	return res, nil
}

func (s *sOpenai) MusicSquare(ctx context.Context, req *v1.MusicSquareListReq) (res *v1.MusicSquareListRes, err error) {
	res = &v1.MusicSquareListRes{
		List: make([]*v1.MusicTaskItem, 0),
	}

	m := dao.OpenaiMusicTask.Ctx(ctx).Where(dao.OpenaiMusicTask.Columns().Status, 1)

	total, err := m.Count()
	if err != nil {
		return nil, err
	}
	res.Total = total

	var tasks []*entity.OpenaiMusicTask
	err = m.Page(req.Page, req.Size).OrderDesc(dao.OpenaiMusicTask.Columns().Id).Scan(&tasks)
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		res.List = append(res.List, &v1.MusicTaskItem{
			Id:               t.Id,
			Prompt:           t.Prompt,
			Type:             t.Type,
			MakeInstrumental: t.MakeInstrumental,
			Title:            t.Title,
			Tags:             t.Tags,
			Model:            t.Model,
			AudioUrl:         t.AudioUrl,
			ImageUrl:         t.ImageUrl,
			VideoUrl:         t.VideoUrl,
			Duration:         t.Duration,
			Status:           t.Status,
			FailReason:       t.FailReason,
			RemoteTaskId:     t.RemoteTaskId,
			RemoteClipId:     t.RemoteClipId,
			CreatedAt:        t.CreatedAt.Format("Y-m-d H:i:s"),
		})
	}

	return res, nil
}

func (s *sOpenai) MusicDetail(ctx context.Context, req *v1.MusicDetailReq) (res *v1.MusicDetailRes, err error) {
	var task *entity.OpenaiMusicTask
	err = dao.OpenaiMusicTask.Ctx(ctx).Where(dao.OpenaiMusicTask.Columns().Id, req.Id).Scan(&task)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, gerror.New("音乐创作任务不存在")
	}

	res = &v1.MusicDetailRes{
		MusicTaskItem: &v1.MusicTaskItem{
			Id:               task.Id,
			Prompt:           task.Prompt,
			Type:             task.Type,
			MakeInstrumental: task.MakeInstrumental,
			Title:            task.Title,
			Tags:             task.Tags,
			Model:            task.Model,
			AudioUrl:         task.AudioUrl,
			ImageUrl:         task.ImageUrl,
			VideoUrl:         task.VideoUrl,
			Duration:         task.Duration,
			Status:           task.Status,
			FailReason:       task.FailReason,
			RemoteTaskId:     task.RemoteTaskId,
			RemoteClipId:     task.RemoteClipId,
			CreatedAt:        task.CreatedAt.Format("Y-m-d H:i:s"),
		},
	}
	return res, nil
}

func (s *sOpenai) MusicDelete(ctx context.Context, req *v1.MusicTaskDeleteReq) (res *v1.MusicTaskDeleteRes, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	_, err = dao.OpenaiMusicTask.Ctx(ctx).
		Where(dao.OpenaiMusicTask.Columns().Id, req.Id).
		Where(dao.OpenaiMusicTask.Columns().UserId, userId).
		Delete()

	if err != nil {
		return nil, err
	}

	return &v1.MusicTaskDeleteRes{}, nil
}
