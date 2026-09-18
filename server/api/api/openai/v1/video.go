package v1

import "github.com/gogf/gf/v2/frame/g"

type VideoTaskSubmitReq struct {
	g.Meta         `path:"/video/submit" method:"post" tags:"AI视频" summary:"提交视频创作任务"`
	Prompt         string `json:"prompt" v:"required#提示词不能为空"`
	ReferenceImage string `json:"referenceImage"`
	LastFrame      string `json:"lastFrame"`
	VoiceDubUrl    string `json:"voiceDubUrl"`
	BgMusicUrl     string `json:"bgMusicUrl"`
	GenerateAudio  int    `json:"generateAudio" d:"1"`
	Model          string `json:"model" v:"required#模型不能为空"`
}

type VideoTaskSubmitRes struct {
	Id int64 `json:"id" dc:"任务ID"`
}

type VideoTaskListReq struct {
	g.Meta `path:"/video/list" method:"get" tags:"AI视频" summary:"获取视频创作任务列表"`
	Page   int `json:"page" d:"1"`
	Size   int `json:"size" d:"10"`
}

type VideoTaskItem struct {
	Id             int64  `json:"id"`
	Prompt         string `json:"prompt"`
	ReferenceImage string `json:"referenceImage"`
	LastFrame      string `json:"lastFrame"`
	VoiceDubUrl    string `json:"voiceDubUrl"`
	BgMusicUrl     string `json:"bgMusicUrl"`
	GenerateAudio  int    `json:"generateAudio"`
	Model          string `json:"model"`
	VideoUrl       string `json:"videoUrl"`
	Status         int    `json:"status"` // 0:进行中 1:成功 2:失败
	FailReason     string `json:"failReason"`
	RemoteTaskId   string `json:"remoteTaskId"`
	CreatedAt      string `json:"createdAt"`
}

type VideoTaskListRes struct {
	List  []*VideoTaskItem `json:"list"`
	Total int              `json:"total"`
}

type VideoTaskDeleteReq struct {
	g.Meta `path:"/video/delete" method:"post" tags:"AI视频" summary:"删除视频创作任务"`
	Id     int64 `json:"id" v:"required#任务ID不能为空"`
}

type VideoTaskDeleteRes struct{}

type VideoSquareListReq struct {
	g.Meta `path:"/video/square" method:"get" tags:"AI视频" summary:"获取视频创作作品广场"`
	Page   int `json:"page" d:"1"`
	Size   int `json:"size" d:"20"`
}

type VideoSquareListRes struct {
	List  []*VideoTaskItem `json:"list"`
	Total int              `json:"total"`
}

type VideoDetailReq struct {
	g.Meta `path:"/video/detail" method:"get" tags:"AI视频" summary:"获取视频创作任务详情"`
	Id     int64 `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

type VideoDetailRes struct {
	*VideoTaskItem
}

