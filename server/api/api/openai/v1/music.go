package v1

import "github.com/gogf/gf/v2/frame/g"

type MusicTaskSubmitReq struct {
	g.Meta           `path:"/music/submit" method:"post" tags:"AI音乐" summary:"提交音乐创作任务"`
	Prompt           string `json:"prompt" dc:"提示词"`
	Type             int    `json:"type" d:"1" dc:"1=灵感模式, 2=定制模式"`
	MakeInstrumental int    `json:"makeInstrumental" d:"0" dc:"0=否, 1=是"`
	Title            string `json:"title" dc:"歌名"`
	Tags             string `json:"tags" dc:"风格标签"`
	Model            string `json:"model" v:"required#模型不能为空"`
}

type MusicTaskSubmitRes struct {
	Id int64 `json:"id" dc:"任务ID"`
}

type MusicTaskListReq struct {
	g.Meta `path:"/music/list" method:"get" tags:"AI音乐" summary:"获取音乐创作任务列表"`
	Page   int `json:"page" d:"1"`
	Size   int `json:"size" d:"10"`
}

type MusicTaskItem struct {
	Id               int64   `json:"id"`
	Prompt           string  `json:"prompt"`
	Type             int     `json:"type"`
	MakeInstrumental int     `json:"makeInstrumental"`
	Title            string  `json:"title"`
	Tags             string  `json:"tags"`
	Model            string  `json:"model"`
	AudioUrl         string  `json:"audioUrl"`
	ImageUrl         string  `json:"imageUrl"`
	VideoUrl         string  `json:"videoUrl"`
	Duration         float64 `json:"duration"`
	Status           int     `json:"status"` // 0:进行中 1:成功 2:失败
	FailReason       string  `json:"failReason"`
	RemoteTaskId     string  `json:"remoteTaskId"`
	RemoteClipId     string  `json:"remoteClipId"`
	CreatedAt        string  `json:"createdAt"`
}

type MusicTaskListRes struct {
	List  []*MusicTaskItem `json:"list"`
	Total int              `json:"total"`
}

type MusicTaskDeleteReq struct {
	g.Meta `path:"/music/delete" method:"post" tags:"AI音乐" summary:"删除音乐创作任务"`
	Id     int64 `json:"id" v:"required#任务ID不能为空"`
}

type MusicTaskDeleteRes struct{}

type MusicSquareListReq struct {
	g.Meta `path:"/music/square" method:"get" tags:"AI音乐" summary:"获取音乐创作作品广场"`
	Page   int `json:"page" d:"1"`
	Size   int `json:"size" d:"20"`
}

type MusicSquareListRes struct {
	List  []*MusicTaskItem `json:"list"`
	Total int              `json:"total"`
}

type MusicDetailReq struct {
	g.Meta `path:"/music/detail" method:"get" tags:"AI音乐" summary:"获取音乐创作任务详情"`
	Id     int64 `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

type MusicDetailRes struct {
	*MusicTaskItem
}
