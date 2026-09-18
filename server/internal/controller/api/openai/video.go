package openai

import (
	"context"

	v1 "hotgo/api/api/openai/v1"
	"hotgo/internal/service"
)

// VideoSubmit 提交视频创作任务
func (c *cOpenai) VideoSubmit(ctx context.Context, req *v1.VideoTaskSubmitReq) (res *v1.VideoTaskSubmitRes, err error) {
	return service.Openai().VideoSubmit(ctx, req)
}

// VideoList 获取视频创作任务列表
func (c *cOpenai) VideoList(ctx context.Context, req *v1.VideoTaskListReq) (res *v1.VideoTaskListRes, err error) {
	return service.Openai().VideoList(ctx, req)
}

// VideoDelete 删除视频创作任务
func (c *cOpenai) VideoDelete(ctx context.Context, req *v1.VideoTaskDeleteReq) (res *v1.VideoTaskDeleteRes, err error) {
	return service.Openai().VideoDelete(ctx, req)
}

// VideoSquare 获取视频创作作品广场
func (c *cOpenai) VideoSquare(ctx context.Context, req *v1.VideoSquareListReq) (res *v1.VideoSquareListRes, err error) {
	return service.Openai().VideoSquare(ctx, req)
}

// VideoDetail 获取视频创作任务详情
func (c *cOpenai) VideoDetail(ctx context.Context, req *v1.VideoDetailReq) (res *v1.VideoDetailRes, err error) {
	return service.Openai().VideoDetail(ctx, req)
}
