package v1

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type LoginReq struct {
	g.Meta   `path:"/user/login" method:"post" tags:"用户" summary:"用户登录"`
	Username string `json:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码"`
}

type LoginRes struct {
	Token    string `json:"token" dc:"登录令牌"`
	Expire   int64  `json:"expire" dc:"过期时间"`
	Id       int64  `json:"id" dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Mobile   string `json:"mobile" dc:"Mobile number"`
	Email    string `json:"email" dc:"邮箱"`
	Money    uint   `json:"money" dc:"余额"`
	Score    uint   `json:"score" dc:"积分"`
}

type InfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户" summary:"获取用户信息"`
}

type InfoRes struct {
	Id       int64       `json:"id" dc:"用户ID"`
	Username string      `json:"username" dc:"用户名"`
	Nickname string      `json:"nickname" dc:"昵称"`
	Avatar   string      `json:"avatar" dc:"头像"`
	Mobile   string      `json:"mobile" dc:"Mobile number"`
	Email    string      `json:"email" dc:"邮箱"`
	Gender   uint        `json:"gender" dc:"性别:0=未知,1=男,2=女"`
	Birthday *gtime.Time `json:"birthday" dc:"生日"`
	Motto    string      `json:"motto" dc:"签名"`
	Money    uint        `json:"money" dc:"余额"`
	Score    uint        `json:"score" dc:"积分"`
}

type RegisterReq struct {
	g.Meta   `path:"/user/register" method:"post" tags:"用户" summary:"用户注册"`
	Username string `json:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `json:"password" v:"required|length:6,20#请输入密码|密码长度需要在6~20位之间" dc:"密码"`
	Nickname string `json:"nickname" dc:"昵称"`
	Mobile   string `json:"mobile" dc:"Mobile number"`
	Email    string `json:"email" dc:"邮箱"`
}

type RegisterRes struct {
	Token    string `json:"token" dc:"登录令牌"`
	Expire   int64  `json:"expire" dc:"过期时间"`
	Id       int64  `json:"id" dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Mobile   string `json:"mobile" dc:"Mobile number"`
	Email    string `json:"email" dc:"邮箱"`
	Money    uint   `json:"money" dc:"余额"`
	Score    uint   `json:"score" dc:"积分"`
}

type LogoutReq struct {
	g.Meta `path:"/user/logout" method:"post" tags:"用户" summary:"退出登录"`
}

type LogoutRes struct{}

type UpdateInfoReq struct {
	g.Meta   `path:"/user/update" method:"post" tags:"用户" summary:"修改用户资料"`
	Nickname string      `json:"nickname" v:"required#昵称不能为空" dc:"昵称"`
	Avatar   string      `json:"avatar" v:"required#头像不能为空" dc:"头像"`
	Gender   int         `json:"gender" dc:"性别:0=未知,1=男,2=女"`
	Birthday *gtime.Time `json:"birthday" dc:"生日"`
	Motto    string      `json:"motto" dc:"签名"`
}

type UpdateInfoRes struct {
	Id       int64       `json:"id" dc:"用户ID"`
	Username string      `json:"username" dc:"用户名"`
	Nickname string      `json:"nickname" dc:"昵称"`
	Avatar   string      `json:"avatar" dc:"头像"`
	Mobile   string      `json:"mobile" dc:"Mobile number"`
	Email    string      `json:"email" dc:"邮箱"`
	Gender   int         `json:"gender" dc:"性别:0=未知,1=男,2=女"`
	Birthday *gtime.Time `json:"birthday" dc:"生日"`
	Motto    string      `json:"motto" dc:"签名"`
	Money    uint        `json:"money" dc:"余额"`
	Score    uint        `json:"score" dc:"积分"`
}

type UpdatePwdReq struct {
	g.Meta      `path:"/user/updatePwd" method:"post" tags:"用户" summary:"修改登录密码"`
	OldPassword string `json:"oldPassword" v:"required#原密码不能为空" dc:"原密码"`
	NewPassword string `json:"newPassword" v:"required|length:6,20#新密码不能为空|密码长度需要在6~20位之间" dc:"新密码"`
}

type UpdatePwdRes struct{}

type SendResetCodeReq struct {
	g.Meta `path:"/user/sendResetCode" method:"post" tags:"用户" summary:"发送密码重置验证码"`
	Mobile string `json:"mobile" v:"required|phone#手机号不能为空|手机号格式不正确" dc:"Mobile number"`
}

type SendResetCodeRes struct{}

type ResetPwdReq struct {
	g.Meta   `path:"/user/resetPwd" method:"post" tags:"用户" summary:"通过验证码重置密码"`
	Mobile   string `json:"mobile" v:"required|phone#手机号不能为空|手机号格式不正确" dc:"Mobile number"`
	Code     string `json:"code" v:"required#验证码不能为空" dc:"验证码"`
	Password string `json:"password" v:"required|length:6,20#新密码不能为空|密码长度需要在6~20位之间" dc:"新密码"`
}

type ResetPwdRes struct{}

type SendMobileBindCodeReq struct {
	g.Meta `path:"/user/sendMobileBindCode" method:"post" tags:"用户" summary:"发送手机号换绑验证码"`
	Mobile string `json:"mobile" v:"required|phone#手机号不能为空|手机号格式不正确" dc:"新手机号"`
}

type SendMobileBindCodeRes struct{}

type BindMobileReq struct {
	g.Meta   `path:"/user/bindMobile" method:"post" tags:"用户" summary:"换绑手机号"`
	Password string `json:"password" v:"required#登录密码不能为空" dc:"登录密码"`
	Mobile   string `json:"mobile" v:"required|phone#手机号不能为空|手机号格式不正确" dc:"新手机号"`
	Code     string `json:"code" v:"required#验证码不能为空" dc:"验证码"`
}

type BindMobileRes struct{}

type SendEmailBindCodeReq struct {
	g.Meta `path:"/user/sendEmailBindCode" method:"post" tags:"用户" summary:"发送邮箱换绑验证码"`
	Email  string `json:"email" v:"required|email#邮箱不能为空|邮箱格式不正确" dc:"新邮箱"`
}

type SendEmailBindCodeRes struct{}

type BindEmailReq struct {
	g.Meta   `path:"/user/bindEmail" method:"post" tags:"用户" summary:"换绑邮箱"`
	Password string `json:"password" v:"required#登录密码不能为空" dc:"登录密码"`
	Email    string `json:"email" v:"required|email#邮箱不能为空|邮箱格式不正确" dc:"新邮箱"`
	Code     string `json:"code" v:"required#验证码不能为空" dc:"验证码"`
}

type BindEmailRes struct{}

type UploadReq struct {
	g.Meta     `path:"/user/upload" method:"post" tags:"用户" summary:"上传用户文件" mime:"multipart/form-data"`
	File       *ghttp.UploadFile `json:"file" type:"file" v:"required#请选择上传文件" dc:"文件"`
	UploadType string            `json:"uploadType" d:"other" dc:"文件类型:image/doc/audio/video/zip/other"`
}

type UploadRes struct {
	*entity.SysAttachment
	SizeFormat string `json:"sizeFormat" dc:"格式化后的文件大小"`
}

type CreditsLogListReq struct {
	g.Meta `path:"/user/creditsLog" method:"get" tags:"用户" summary:"获取资产变动流水"`
	form.PageReq
	CreditType string `json:"creditType" v:"in:money,score#资产类型不正确" dc:"资产类型:money=余额,score=积分"`
}

type CreditsLogListRes struct {
	form.PageRes
	List []*entity.UserCreditsLog `json:"list" dc:"资产流水列表"`
}

type LoginLogListReq struct {
	g.Meta `path:"/user/loginLog" method:"get" tags:"用户" summary:"获取登录历史"`
	form.PageReq
}

type LoginLogListRes struct {
	form.PageRes
	List []*entity.UserLoginLog `json:"list" dc:"登录历史列表"`
}

type CancelReq struct {
	g.Meta   `path:"/user/cancel" method:"post" tags:"用户" summary:"注销当前账号"`
	Password string `json:"password" v:"required#登录密码不能为空" dc:"登录密码"`
}

type CancelRes struct{}
