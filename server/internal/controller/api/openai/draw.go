package openai

import (
	"context"

	v1 "hotgo/api/api/openai/v1"
	"hotgo/internal/service"
)

// DrawSubmit 提交绘画任务
func (c *cOpenai) DrawSubmit(ctx context.Context, req *v1.DrawTaskSubmitReq) (res *v1.DrawTaskSubmitRes, err error) {
	return service.Openai().DrawSubmit(ctx, req)
}

// DrawList 获取绘画任务列表
func (c *cOpenai) DrawList(ctx context.Context, req *v1.DrawTaskListReq) (res *v1.DrawTaskListRes, err error) {
	return service.Openai().DrawList(ctx, req)
}

// DrawDelete 删除绘画任务
func (c *cOpenai) DrawDelete(ctx context.Context, req *v1.DrawTaskDeleteReq) (res *v1.DrawTaskDeleteRes, err error) {
	return service.Openai().DrawDelete(ctx, req)
}

// DrawAction 提交二次操作(U1/V1等)
func (c *cOpenai) DrawAction(ctx context.Context, req *v1.DrawTaskActionReq) (res *v1.DrawTaskActionRes, err error) {
	return service.Openai().DrawAction(ctx, req)
}

// DrawPromptTranslate 提示词翻译为英文
func (c *cOpenai) DrawPromptTranslate(ctx context.Context, req *v1.DrawPromptTranslateReq) (res *v1.DrawPromptTranslateRes, err error) {
	return service.Openai().DrawPromptTranslate(ctx, req)
}

// DrawPromptEnrich 提示词联想扩写
func (c *cOpenai) DrawPromptEnrich(ctx context.Context, req *v1.DrawPromptEnrichReq) (res *v1.DrawPromptEnrichRes, err error) {
	return service.Openai().DrawPromptEnrich(ctx, req)
}

// DrawSquare 获取绘画作品广场
func (c *cOpenai) DrawSquare(ctx context.Context, req *v1.DrawSquareListReq) (res *v1.DrawSquareListRes, err error) {
	return service.Openai().DrawSquare(ctx, req)
}

// DrawDetail 获取绘画任务详情
func (c *cOpenai) DrawDetail(ctx context.Context, req *v1.DrawDetailReq) (res *v1.DrawDetailRes, err error) {
	return service.Openai().DrawDetail(ctx, req)
}
