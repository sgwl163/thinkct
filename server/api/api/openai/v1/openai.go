package v1

import (
	"hotgo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Folder
type FolderListReq struct {
	g.Meta `path:"/folder/list" method:"get" tags:"AI聊天" summary:"文件夹列表"`
}
type FolderListRes struct {
	List []*entity.OpenaiChatFolder `json:"list"` // 文件夹列表数据
}

type FolderEditReq struct {
	g.Meta `path:"/folder/edit" method:"post" tags:"AI聊天" summary:"编辑文件夹"`
	Id     int64  `json:"id"`                          // 文件夹ID，大于0表示修改，等于0表示新增
	Name   string `json:"name" v:"required#文件夹名称不能为空"` // 文件夹名称
	Sort   int    `json:"sort"`                        // 排序值，越小越靠前
}
type FolderEditRes struct{}

type FolderDeleteReq struct {
	g.Meta `path:"/folder/delete" method:"post" tags:"AI聊天" summary:"删除文件夹"`
	Id     int64 `json:"id" v:"required#ID不能为空"` // 要删除的文件夹ID
}
type FolderDeleteRes struct{}

// Session
type SessionListReq struct {
	g.Meta   `path:"/session/list" method:"get" tags:"AI聊天" summary:"会话列表"`
	FolderId int64  `json:"folderId" d:"0"` // 所属文件夹ID，0表示查询所有
	Keyword  string `json:"keyword"`        // 支持根据标题关键字全局搜索
}
type SessionListRes struct {
	List []*entity.OpenaiChatSession `json:"list"` // 会话列表数据
}

type SessionEditReq struct {
	g.Meta   `path:"/session/edit" method:"post" tags:"AI聊天" summary:"编辑会话"`
	Id       int64  `json:"id"`                          // 会话ID，大于0表示修改，等于0表示新增
	FolderId int64  `json:"folderId"`                    // 绑定的文件夹ID
	Title    string `json:"title" v:"required#会话标题不能为空"` // 会话标题
	Prompt   string `json:"prompt"`                      // 会话的系统提示词(System Prompt)
	Model    string `json:"model"`                       // 会话默认使用的模型标识(如 gpt-4o)
}
type SessionEditRes struct {
	Id int64 `json:"id"` // 返回会话ID（新增时返回新插入的ID，编辑时返回原ID）
}

type SessionDeleteReq struct {
	g.Meta `path:"/session/delete" method:"post" tags:"AI聊天" summary:"删除会话"`
	Id     int64 `json:"id" v:"required#ID不能为空"`
}
type SessionDeleteRes struct{}

type SessionTopReq struct {
	g.Meta `path:"/session/top" method:"post" tags:"AI聊天" summary:"置顶/取消置顶会话"`
	Id     int64 `json:"id" v:"required#ID不能为空"`
	IsTop  int   `json:"isTop"` // 1:置顶 0:取消
}
type SessionTopRes struct{}

type SessionClearReq struct {
	g.Meta `path:"/session/clear" method:"post" tags:"AI聊天" summary:"清空会话上下文"`
	Id     int64 `json:"id" v:"required#ID不能为空"`
}
type SessionClearRes struct{}

type SessionRenameReq struct {
	g.Meta `path:"/session/rename" method:"post" tags:"AI聊天" summary:"重命名会话"`
	Id     int64  `json:"id" v:"required#ID不能为空"`
	Title  string `json:"title" v:"required#标题不能为空"`
}
type SessionRenameRes struct{}

// Message
type MessageListReq struct {
	g.Meta    `path:"/message/list" method:"get" tags:"AI聊天" summary:"消息列表"`
	SessionId int64 `json:"sessionId" v:"required#会话ID不能为空"`
}
type MessageListRes struct {
	List        []*entity.OpenaiChatMessage `json:"list"`
	Model       string                      `json:"model"`
	Prompt      string                      `json:"prompt"`
	Temperature float64                     `json:"temperature"`
}

type MessageSendReq struct {
	g.Meta      `path:"/message/send" method:"post" tags:"AI聊天" summary:"发送消息"`
	SessionId   int64    `json:"sessionId"`   // 当前消息所属会话ID。如果传 0 或空，则会自动创建一个新会话
	FolderId    int64    `json:"folderId"`    // 选填：指定会话所属文件夹ID（仅在自动创建新会话时生效）
	Content     string   `json:"content"`     // 文本消息内容
	Images      []string `json:"images"`      // 多模态：图片链接数组
	Files       []string `json:"files"`       // 多模态：普通文件链接数组
	IsSearch    bool     `json:"isSearch"`    // 是否开启联网搜索增强
	Model       string   `json:"model"`       // 会话模型标识，传值会覆盖更新当前会话绑定的模型
	Prompt      *string  `json:"prompt"`      // 系统提示词，传值会覆盖更新当前会话绑定的提示词
	Temperature *float32 `json:"temperature"` // 模型温度值，传值会覆盖更新当前会话绑定的温度
}
type MessageSendRes struct {
	SessionId int64                     `json:"sessionId"` // 返回当前实际使用的会话ID（用于前端新建会话后更新URL）
	FolderId  int64                     `json:"folderId"`  // 返回当前会话所属的文件夹ID
	Reply     *entity.OpenaiChatMessage `json:"reply"`     // 发送成功后最终入库的助手回复实体
}

type MessageDeleteReq struct {
	g.Meta `path:"/message/delete" method:"post" tags:"AI聊天" summary:"删除消息"`
	Id     int64 `json:"id" v:"required#ID不能为空"`
}
type MessageDeleteRes struct{}

type MessageEditReq struct {
	g.Meta  `path:"/message/edit" method:"post" tags:"AI聊天" summary:"编辑消息"`
	Id      int64  `json:"id" v:"required#ID不能为空"`
	Content string `json:"content" v:"required#内容不能为空"`
}
type MessageEditRes struct{}

type MessageFeedbackReq struct {
	g.Meta   `path:"/message/feedback" method:"post" tags:"AI聊天" summary:"消息反馈(喜欢/不喜欢)"`
	Id       int64 `json:"id" v:"required#ID不能为空"`
	Feedback int   `json:"feedback"` // 1:喜欢 2:不喜欢 0:取消反馈
}
type MessageFeedbackRes struct{}

type MessageRegenerateReq struct {
	g.Meta    `path:"/message/regenerate" method:"post" tags:"AI聊天" summary:"重新生成消息"`
	SessionId int64 `json:"sessionId" v:"required#会话ID不能为空"`
	MessageId int64 `json:"messageId" v:"required#要重新生成的消息ID不能为空"` // 该消息会被软删除，并以该节点之前的上下文重新生成
	IsSearch  bool  `json:"isSearch"`
}
type MessageRegenerateRes struct {
	Reply *entity.OpenaiChatMessage `json:"reply"`
}

// Models
type ModelListReq struct {
	g.Meta `path:"/model/list" method:"get" tags:"AI聊天" summary:"获取可用模型列表"`
}
type ModelItem struct {
	GroupId   int64  `json:"groupId"`
	GroupName string `json:"groupName"`
	ModelId   int64  `json:"modelId"`
	Name      string `json:"name"`
	Model     string `json:"model"`
	IsVision  int    `json:"isVision"`
	IsFile    int    `json:"isFile"`
	IsSearch  int    `json:"isSearch"`
}
type ModelListRes struct {
	List []ModelItem `json:"list"`
}

// Upload
type UploadFileReq struct {
	g.Meta `path:"/upload/file" method:"post" tags:"AI聊天" summary:"上传附件" mime:"multipart/form-data"`
	File   *ghttp.UploadFile `json:"file" type:"file" v:"required#请选择上传文件"`
}
type UploadFileRes struct {
	Url string `json:"url"`
}
