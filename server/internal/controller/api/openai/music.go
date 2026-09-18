package openai

import (
	"context"

	v1 "hotgo/api/api/openai/v1"
	"hotgo/internal/service"
)

// MusicSubmit 提交音乐创作任务
func (c *cOpenai) MusicSubmit(ctx context.Context, req *v1.MusicTaskSubmitReq) (res *v1.MusicTaskSubmitRes, err error) {
	return service.Openai().MusicSubmit(ctx, req)
}

// MusicList 获取音乐创作任务列表
func (c *cOpenai) MusicList(ctx context.Context, req *v1.MusicTaskListReq) (res *v1.MusicTaskListRes, err error) {
	return service.Openai().MusicList(ctx, req)
}

// MusicDelete 删除音乐创作任务
func (c *cOpenai) MusicDelete(ctx context.Context, req *v1.MusicTaskDeleteReq) (res *v1.MusicTaskDeleteRes, err error) {
	return service.Openai().MusicDelete(ctx, req)
}

// MusicSquare 获取音乐创作作品广场
func (c *cOpenai) MusicSquare(ctx context.Context, req *v1.MusicSquareListReq) (res *v1.MusicSquareListRes, err error) {
	return service.Openai().MusicSquare(ctx, req)
}

// MusicDetail 获取音乐创作任务详情
func (c *cOpenai) MusicDetail(ctx context.Context, req *v1.MusicDetailReq) (res *v1.MusicDetailRes, err error) {
	return service.Openai().MusicDetail(ctx, req)
}
