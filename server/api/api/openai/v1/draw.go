package v1

import "github.com/gogf/gf/v2/frame/g"

type DrawTaskSubmitReq struct {
	g.Meta         `path:"/draw/submit" method:"post" tags:"AI绘画" summary:"提交绘画任务"`
	Prompt         string `json:"prompt" v:"required#提示词不能为空"`
	ReferenceImage string `json:"referenceImage"`
	Model          string `json:"model" v:"required#模型不能为空"`
	Version        string `json:"version"`
	Ratio          string `json:"ratio"`
	Raw            int    `json:"raw"`
	NegativePrompt string `json:"negativePrompt"`
	Stylize        int    `json:"stylize"`
	Chaos          int    `json:"chaos"`
	Seed           int    `json:"seed"`
}

type DrawTaskSubmitRes struct {
	Id int64 `json:"id" dc:"任务ID"`
}

type DrawTaskListReq struct {
	g.Meta `path:"/draw/list" method:"get" tags:"AI绘画" summary:"获取绘画任务列表"`
	Page   int `json:"page" d:"1"`
	Size   int `json:"size" d:"10"`
}

type DrawTaskItem struct {
	Id             int64  `json:"id"`
	Prompt         string `json:"prompt"`
	ReferenceImage string `json:"referenceImage"`
	Model          string `json:"model"`
	Ratio          string `json:"ratio"`
	ImageUrl       string `json:"imageUrl"`
	Status         int    `json:"status"` // 0:进行中 1:成功 2:失败
	FailReason     string `json:"failReason"`
	CreatedAt      string `json:"createdAt"`
	MjTaskId       string `json:"mjTaskId"`
	Buttons        string `json:"buttons"`
	Progress       int    `json:"progress"`
}

type DrawTaskListRes struct {
	List  []*DrawTaskItem `json:"list"`
	Total int             `json:"total"`
}

type DrawTaskDeleteReq struct {
	g.Meta `path:"/draw/delete" method:"post" tags:"AI绘画" summary:"删除绘画任务"`
	Id     int64 `json:"id" v:"required#任务ID不能为空"`
}

type DrawTaskDeleteRes struct{}

type DrawTaskActionReq struct {
	g.Meta `path:"/draw/action" method:"post" tags:"AI绘画" summary:"MJ等图片操作(放大/变换)"`
	Id     int64  `json:"id" v:"required#原任务ID不能为空"`
	Action string `json:"action" v:"required#操作类型不能为空"` // 例如 U1, V1, U2, V2
}

type DrawTaskActionRes struct {
	Id int64 `json:"id" dc:"新生成的任务ID"`
}

type DrawPromptTranslateReq struct {
	g.Meta `path:"/draw/translate" method:"post" tags:"AI绘画" summary:"绘画提示词翻译为英文"`
	Prompt string `json:"prompt" v:"required#提示词不能为空"`
}

type DrawPromptTranslateRes struct {
	Result string `json:"result" dc:"翻译结果(英文)"`
}

type DrawPromptEnrichReq struct {
	g.Meta `path:"/draw/enrich" method:"post" tags:"AI绘画" summary:"绘画提示词联想扩写"`
	Prompt string `json:"prompt" v:"required#提示词不能为空"`
}

type DrawPromptEnrichRes struct {
	Result string `json:"result" dc:"联想扩写结果"`
}

type DrawSquareListReq struct {
	g.Meta `path:"/draw/square" method:"get" tags:"AI绘画" summary:"获取绘画作品广场"`
	Page   int `json:"page" d:"1"`
	Size   int `json:"size" d:"20"`
}

type DrawSquareListRes struct {
	List  []*DrawTaskItem `json:"list"`
	Total int             `json:"total"`
}

type DrawDetailReq struct {
	g.Meta `path:"/draw/detail" method:"get" tags:"AI绘画" summary:"获取绘画任务详情"`
	Id     int64 `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

type DrawDetailRes struct {
	*DrawTaskItem
}
