// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "hotgo/api/api/openai/v1"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/openaiin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IOpenaiApiGroup interface {
		// Model AI接口分组ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取AI接口分组列表
		List(ctx context.Context, in *openaiin.ApiGroupListInp) (list []*openaiin.ApiGroupListModel, totalCount int, err error)
		// Export 导出AI接口分组
		Export(ctx context.Context, in *openaiin.ApiGroupListInp) (err error)
		// Edit 修改/新增AI接口分组
		Edit(ctx context.Context, in *openaiin.ApiGroupEditInp) (err error)
		// Delete 删除AI接口分组
		Delete(ctx context.Context, in *openaiin.ApiGroupDeleteInp) (err error)
		// MaxSort 获取AI接口分组最大排序
		MaxSort(ctx context.Context, in *openaiin.ApiGroupMaxSortInp) (res *openaiin.ApiGroupMaxSortModel, err error)
		// View 获取AI接口分组指定信息
		View(ctx context.Context, in *openaiin.ApiGroupViewInp) (res *openaiin.ApiGroupViewModel, err error)
		// Status 更新AI接口分组状态
		Status(ctx context.Context, in *openaiin.ApiGroupStatusInp) (err error)
	}
	IOpenaiApiModel interface {
		// Model AI模型列表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取AI模型列表列表
		List(ctx context.Context, in *openaiin.ApiModelListInp) (list []*openaiin.ApiModelListModel, totalCount int, err error)
		// Export 导出AI模型列表
		Export(ctx context.Context, in *openaiin.ApiModelListInp) (err error)
		// Edit 修改/新增AI模型列表
		Edit(ctx context.Context, in *openaiin.ApiModelEditInp) (err error)
		// Delete 删除AI模型列表
		Delete(ctx context.Context, in *openaiin.ApiModelDeleteInp) (err error)
		// MaxSort 获取AI模型列表最大排序
		MaxSort(ctx context.Context, in *openaiin.ApiModelMaxSortInp) (res *openaiin.ApiModelMaxSortModel, err error)
		// View 获取AI模型列表指定信息
		View(ctx context.Context, in *openaiin.ApiModelViewInp) (res *openaiin.ApiModelViewModel, err error)
		// Status 更新AI模型列表状态
		Status(ctx context.Context, in *openaiin.ApiModelStatusInp) (err error)
	}
	IOpenaiApp interface {
		// Model AI应用ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取AI应用列表
		List(ctx context.Context, in *openaiin.AppListInp) (list []*openaiin.AppListModel, totalCount int, err error)
		// Export 导出AI应用
		Export(ctx context.Context, in *openaiin.AppListInp) (err error)
		// Edit 修改/新增AI应用
		Edit(ctx context.Context, in *openaiin.AppEditInp) (err error)
		// Delete 删除AI应用
		Delete(ctx context.Context, in *openaiin.AppDeleteInp) (err error)
		// MaxSort 获取AI应用最大排序
		MaxSort(ctx context.Context, in *openaiin.AppMaxSortInp) (res *openaiin.AppMaxSortModel, err error)
		// View 获取AI应用指定信息
		View(ctx context.Context, in *openaiin.AppViewInp) (res *openaiin.AppViewModel, err error)
		// Status 更新AI应用状态
		Status(ctx context.Context, in *openaiin.AppStatusInp) (err error)
	}
	IOpenaiChatSession interface {
		// Model 聊天会话ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取聊天会话列表
		List(ctx context.Context, in *openaiin.ChatSessionListInp) (list []*openaiin.ChatSessionListModel, totalCount int, err error)
		// Export 导出聊天会话
		Export(ctx context.Context, in *openaiin.ChatSessionListInp) (err error)
		// Edit 修改/新增聊天会话
		Edit(ctx context.Context, in *openaiin.ChatSessionEditInp) (err error)
		// Delete 删除聊天会话
		Delete(ctx context.Context, in *openaiin.ChatSessionDeleteInp) (err error)
		// View 获取聊天会话指定信息
		View(ctx context.Context, in *openaiin.ChatSessionViewInp) (res *openaiin.ChatSessionViewModel, err error)
		// Status 更新聊天会话状态
		Status(ctx context.Context, in *openaiin.ChatSessionStatusInp) (err error)
	}
	IOpenai interface {
		DrawSubmit(ctx context.Context, req *v1.DrawTaskSubmitReq) (res *v1.DrawTaskSubmitRes, err error)
		DrawAction(ctx context.Context, req *v1.DrawTaskActionReq) (res *v1.DrawTaskActionRes, err error)
		DrawPromptTranslate(ctx context.Context, req *v1.DrawPromptTranslateReq) (res *v1.DrawPromptTranslateRes, err error)
		DrawPromptEnrich(ctx context.Context, req *v1.DrawPromptEnrichReq) (res *v1.DrawPromptEnrichRes, err error)
		DrawDelete(ctx context.Context, req *v1.DrawTaskDeleteReq) (res *v1.DrawTaskDeleteRes, err error)
		DrawList(ctx context.Context, req *v1.DrawTaskListReq) (res *v1.DrawTaskListRes, err error)
		DrawSquare(ctx context.Context, req *v1.DrawSquareListReq) (res *v1.DrawSquareListRes, err error)
		DrawDetail(ctx context.Context, req *v1.DrawDetailReq) (res *v1.DrawDetailRes, err error)
		MusicSubmit(ctx context.Context, req *v1.MusicTaskSubmitReq) (res *v1.MusicTaskSubmitRes, err error)
		MusicList(ctx context.Context, req *v1.MusicTaskListReq) (res *v1.MusicTaskListRes, err error)
		MusicSquare(ctx context.Context, req *v1.MusicSquareListReq) (res *v1.MusicSquareListRes, err error)
		MusicDetail(ctx context.Context, req *v1.MusicDetailReq) (res *v1.MusicDetailRes, err error)
		MusicDelete(ctx context.Context, req *v1.MusicTaskDeleteReq) (res *v1.MusicTaskDeleteRes, err error)
		FolderList(ctx context.Context, req *v1.FolderListReq) (list []*entity.OpenaiChatFolder, err error)
		FolderEdit(ctx context.Context, req *v1.FolderEditReq) (err error)
		FolderDelete(ctx context.Context, req *v1.FolderDeleteReq) (err error)
		SessionList(ctx context.Context, req *v1.SessionListReq) (list []*entity.OpenaiChatSession, err error)
		SessionEdit(ctx context.Context, req *v1.SessionEditReq) (res *v1.SessionEditRes, err error)
		SessionDelete(ctx context.Context, req *v1.SessionDeleteReq) (err error)
		SessionTop(ctx context.Context, req *v1.SessionTopReq) (err error)
		SessionClear(ctx context.Context, req *v1.SessionClearReq) (err error)
		SessionRename(ctx context.Context, req *v1.SessionRenameReq) (err error)
		MessageList(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error)
		UploadFile(ctx context.Context, req *v1.UploadFileReq) (url string, err error)
		ModelList(ctx context.Context, req *v1.ModelListReq) (list []v1.ModelItem, err error)
		MessageSend(ctx context.Context, req *v1.MessageSendReq) (res *v1.MessageSendRes, err error)
		MessageDelete(ctx context.Context, req *v1.MessageDeleteReq) (err error)
		MessageEdit(ctx context.Context, req *v1.MessageEditReq) (err error)
		MessageFeedback(ctx context.Context, req *v1.MessageFeedbackReq) (err error)
		MessageRegenerate(ctx context.Context, req *v1.MessageRegenerateReq) (reply *entity.OpenaiChatMessage, err error)
		VideoSubmit(ctx context.Context, req *v1.VideoTaskSubmitReq) (res *v1.VideoTaskSubmitRes, err error)
		VideoList(ctx context.Context, req *v1.VideoTaskListReq) (res *v1.VideoTaskListRes, err error)
		VideoSquare(ctx context.Context, req *v1.VideoSquareListReq) (res *v1.VideoSquareListRes, err error)
		VideoDetail(ctx context.Context, req *v1.VideoDetailReq) (res *v1.VideoDetailRes, err error)
		VideoDelete(ctx context.Context, req *v1.VideoTaskDeleteReq) (res *v1.VideoTaskDeleteRes, err error)
	}
	IOpenaiDrawTask interface {
		// Model AI绘画任务ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取AI绘画任务列表
		List(ctx context.Context, in *openaiin.DrawTaskListInp) (list []*openaiin.DrawTaskListModel, totalCount int, err error)
		// Export 导出AI绘画任务
		Export(ctx context.Context, in *openaiin.DrawTaskListInp) (err error)
		// Edit 修改/新增AI绘画任务
		Edit(ctx context.Context, in *openaiin.DrawTaskEditInp) (err error)
		// Delete 删除AI绘画任务
		Delete(ctx context.Context, in *openaiin.DrawTaskDeleteInp) (err error)
		// View 获取AI绘画任务指定信息
		View(ctx context.Context, in *openaiin.DrawTaskViewInp) (res *openaiin.DrawTaskViewModel, err error)
		// Status 更新AI绘画任务状态
		Status(ctx context.Context, in *openaiin.DrawTaskStatusInp) (err error)
	}
	IOpenaiMusicTask interface {
		// Model AI音乐任务ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取AI音乐任务列表
		List(ctx context.Context, in *openaiin.MusicTaskListInp) (list []*openaiin.MusicTaskListModel, totalCount int, err error)
		// Export 导出AI音乐任务
		Export(ctx context.Context, in *openaiin.MusicTaskListInp) (err error)
		// Edit 修改/新增AI音乐任务
		Edit(ctx context.Context, in *openaiin.MusicTaskEditInp) (err error)
		// Delete 删除AI音乐任务
		Delete(ctx context.Context, in *openaiin.MusicTaskDeleteInp) (err error)
		// View 获取AI音乐任务指定信息
		View(ctx context.Context, in *openaiin.MusicTaskViewInp) (res *openaiin.MusicTaskViewModel, err error)
		// Status 更新AI音乐任务状态
		Status(ctx context.Context, in *openaiin.MusicTaskStatusInp) (err error)
	}
	IOpenaiAppCategory interface {
		// Model AI应用分类ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取AI应用分类列表
		List(ctx context.Context, in *openaiin.OpenaiAppCategoryListInp) (list []*openaiin.OpenaiAppCategoryListModel, totalCount int, err error)
		// Export 导出AI应用分类
		Export(ctx context.Context, in *openaiin.OpenaiAppCategoryListInp) (err error)
		// Edit 修改/新增AI应用分类
		Edit(ctx context.Context, in *openaiin.OpenaiAppCategoryEditInp) (err error)
		// Delete 删除AI应用分类
		Delete(ctx context.Context, in *openaiin.OpenaiAppCategoryDeleteInp) (err error)
		// MaxSort 获取AI应用分类最大排序
		MaxSort(ctx context.Context, in *openaiin.OpenaiAppCategoryMaxSortInp) (res *openaiin.OpenaiAppCategoryMaxSortModel, err error)
		// View 获取AI应用分类指定信息
		View(ctx context.Context, in *openaiin.OpenaiAppCategoryViewInp) (res *openaiin.OpenaiAppCategoryViewModel, err error)
		// Status 更新AI应用分类状态
		Status(ctx context.Context, in *openaiin.OpenaiAppCategoryStatusInp) (err error)
	}
	IOpenaiVideoTask interface {
		// Model AI视频任务ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取AI视频任务列表
		List(ctx context.Context, in *openaiin.VideoTaskListInp) (list []*openaiin.VideoTaskListModel, totalCount int, err error)
		// Export 导出AI视频任务
		Export(ctx context.Context, in *openaiin.VideoTaskListInp) (err error)
		// Edit 修改/新增AI视频任务
		Edit(ctx context.Context, in *openaiin.VideoTaskEditInp) (err error)
		// Delete 删除AI视频任务
		Delete(ctx context.Context, in *openaiin.VideoTaskDeleteInp) (err error)
		// View 获取AI视频任务指定信息
		View(ctx context.Context, in *openaiin.VideoTaskViewInp) (res *openaiin.VideoTaskViewModel, err error)
		// Status 更新AI视频任务状态
		Status(ctx context.Context, in *openaiin.VideoTaskStatusInp) (err error)
	}
)

var (
	localOpenaiApiGroup    IOpenaiApiGroup
	localOpenaiApiModel    IOpenaiApiModel
	localOpenaiApp         IOpenaiApp
	localOpenaiChatSession IOpenaiChatSession
	localOpenai            IOpenai
	localOpenaiDrawTask    IOpenaiDrawTask
	localOpenaiMusicTask   IOpenaiMusicTask
	localOpenaiAppCategory IOpenaiAppCategory
	localOpenaiVideoTask   IOpenaiVideoTask
)

func OpenaiApiGroup() IOpenaiApiGroup {
	if localOpenaiApiGroup == nil {
		panic("implement not found for interface IOpenaiApiGroup, forgot register?")
	}
	return localOpenaiApiGroup
}

func RegisterOpenaiApiGroup(i IOpenaiApiGroup) {
	localOpenaiApiGroup = i
}

func OpenaiApiModel() IOpenaiApiModel {
	if localOpenaiApiModel == nil {
		panic("implement not found for interface IOpenaiApiModel, forgot register?")
	}
	return localOpenaiApiModel
}

func RegisterOpenaiApiModel(i IOpenaiApiModel) {
	localOpenaiApiModel = i
}

func OpenaiApp() IOpenaiApp {
	if localOpenaiApp == nil {
		panic("implement not found for interface IOpenaiApp, forgot register?")
	}
	return localOpenaiApp
}

func RegisterOpenaiApp(i IOpenaiApp) {
	localOpenaiApp = i
}

func OpenaiChatSession() IOpenaiChatSession {
	if localOpenaiChatSession == nil {
		panic("implement not found for interface IOpenaiChatSession, forgot register?")
	}
	return localOpenaiChatSession
}

func RegisterOpenaiChatSession(i IOpenaiChatSession) {
	localOpenaiChatSession = i
}

func Openai() IOpenai {
	if localOpenai == nil {
		panic("implement not found for interface IOpenai, forgot register?")
	}
	return localOpenai
}

func RegisterOpenai(i IOpenai) {
	localOpenai = i
}

func OpenaiDrawTask() IOpenaiDrawTask {
	if localOpenaiDrawTask == nil {
		panic("implement not found for interface IOpenaiDrawTask, forgot register?")
	}
	return localOpenaiDrawTask
}

func RegisterOpenaiDrawTask(i IOpenaiDrawTask) {
	localOpenaiDrawTask = i
}

func OpenaiMusicTask() IOpenaiMusicTask {
	if localOpenaiMusicTask == nil {
		panic("implement not found for interface IOpenaiMusicTask, forgot register?")
	}
	return localOpenaiMusicTask
}

func RegisterOpenaiMusicTask(i IOpenaiMusicTask) {
	localOpenaiMusicTask = i
}

func OpenaiAppCategory() IOpenaiAppCategory {
	if localOpenaiAppCategory == nil {
		panic("implement not found for interface IOpenaiAppCategory, forgot register?")
	}
	return localOpenaiAppCategory
}

func RegisterOpenaiAppCategory(i IOpenaiAppCategory) {
	localOpenaiAppCategory = i
}

func OpenaiVideoTask() IOpenaiVideoTask {
	if localOpenaiVideoTask == nil {
		panic("implement not found for interface IOpenaiVideoTask, forgot register?")
	}
	return localOpenaiVideoTask
}

func RegisterOpenaiVideoTask(i IOpenaiVideoTask) {
	localOpenaiVideoTask = i
}
