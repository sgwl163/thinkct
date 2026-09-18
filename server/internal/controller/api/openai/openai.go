package openai

import (
	"context"

	v1 "hotgo/api/api/openai/v1"
	"hotgo/internal/service"
)

var (
	Openai = cOpenai{}
)

type cOpenai struct{}

func (c *cOpenai) FolderList(ctx context.Context, req *v1.FolderListReq) (res *v1.FolderListRes, err error) {
	list, err := service.Openai().FolderList(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.FolderListRes{List: list}
	return
}

func (c *cOpenai) FolderEdit(ctx context.Context, req *v1.FolderEditReq) (res *v1.FolderEditRes, err error) {
	err = service.Openai().FolderEdit(ctx, req)
	return
}

func (c *cOpenai) FolderDelete(ctx context.Context, req *v1.FolderDeleteReq) (res *v1.FolderDeleteRes, err error) {
	err = service.Openai().FolderDelete(ctx, req)
	return
}

func (c *cOpenai) SessionList(ctx context.Context, req *v1.SessionListReq) (res *v1.SessionListRes, err error) {
	list, err := service.Openai().SessionList(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.SessionListRes{List: list}
	return
}

func (c *cOpenai) SessionEdit(ctx context.Context, req *v1.SessionEditReq) (res *v1.SessionEditRes, err error) {
	editRes, err := service.Openai().SessionEdit(ctx, req)
	if err != nil {
		return nil, err
	}
	return editRes, nil
}

// SessionDelete 删除会话
func (c *cOpenai) SessionDelete(ctx context.Context, req *v1.SessionDeleteReq) (res *v1.SessionDeleteRes, err error) {
	err = service.Openai().SessionDelete(ctx, req)
	return
}

// SessionTop 置顶/取消置顶会话
func (c *cOpenai) SessionTop(ctx context.Context, req *v1.SessionTopReq) (res *v1.SessionTopRes, err error) {
	err = service.Openai().SessionTop(ctx, req)
	return
}

// SessionClear 清空会话上下文
func (c *cOpenai) SessionClear(ctx context.Context, req *v1.SessionClearReq) (res *v1.SessionClearRes, err error) {
	err = service.Openai().SessionClear(ctx, req)
	return
}

// SessionRename 重命名会话
func (c *cOpenai) SessionRename(ctx context.Context, req *v1.SessionRenameReq) (res *v1.SessionRenameRes, err error) {
	err = service.Openai().SessionRename(ctx, req)
	return
}

// MessageList 消息记录列表
func (c *cOpenai) MessageList(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error) {
	res, err = service.Openai().MessageList(ctx, req)
	return
}

// MessageDelete 删除消息
func (c *cOpenai) MessageDelete(ctx context.Context, req *v1.MessageDeleteReq) (res *v1.MessageDeleteRes, err error) {
	err = service.Openai().MessageDelete(ctx, req)
	return
}

// MessageEdit 编辑消息
func (c *cOpenai) MessageEdit(ctx context.Context, req *v1.MessageEditReq) (res *v1.MessageEditRes, err error) {
	err = service.Openai().MessageEdit(ctx, req)
	return
}

// MessageFeedback 消息反馈
func (c *cOpenai) MessageFeedback(ctx context.Context, req *v1.MessageFeedbackReq) (res *v1.MessageFeedbackRes, err error) {
	err = service.Openai().MessageFeedback(ctx, req)
	return
}

// MessageRegenerate 重新生成消息
func (c *cOpenai) MessageRegenerate(ctx context.Context, req *v1.MessageRegenerateReq) (res *v1.MessageRegenerateRes, err error) {
	reply, err := service.Openai().MessageRegenerate(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.MessageRegenerateRes{Reply: reply}
	return
}

func (c *cOpenai) MessageSend(ctx context.Context, req *v1.MessageSendReq) (res *v1.MessageSendRes, err error) {
	sendRes, err := service.Openai().MessageSend(ctx, req)
	if err != nil {
		return nil, err
	}
	return sendRes, nil
}

func (c *cOpenai) ModelList(ctx context.Context, req *v1.ModelListReq) (res *v1.ModelListRes, err error) {
	list, err := service.Openai().ModelList(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.ModelListRes{List: list}
	return
}

func (c *cOpenai) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	url, err := service.Openai().UploadFile(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.UploadFileRes{Url: url}
	return
}
