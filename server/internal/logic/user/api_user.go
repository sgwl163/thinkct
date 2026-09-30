package user

import (
	"context"

	v1 "hotgo/api/api/user/v1"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/storager"
	"hotgo/internal/library/token"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

type sApiUser struct{}

func NewApiUser() *sApiUser {
	return &sApiUser{}
}

func init() {
	service.RegisterApiUser(NewApiUser())
}

func (s *sApiUser) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	var user *entity.User
	cols := dao.User.Columns()
	if err = dao.User.Ctx(ctx).
		Where(cols.Username, req.Username).
		WhereOr(cols.Email, req.Username).
		WhereOr(cols.Mobile, req.Username).
		Scan(&user); err != nil {
		return nil, gerror.Wrap(err, consts.ErrorORM)
	}
	if user == nil || user.Password != gmd5.MustEncryptString(req.Password+user.Salt) {
		if user != nil {
			_, _ = dao.User.Ctx(ctx).WherePri(user.Id).Increment(cols.LoginFailure, 1)
			s.writeLoginLog(ctx, user, "用户名或密码错误")
		} else {
			s.writeLoginLog(ctx, &entity.User{Username: req.Username}, "用户名或密码错误")
		}
		return nil, gerror.New("用户名或密码错误")
	}
	if !isUserEnabled(user) {
		s.writeLoginLog(ctx, user, "账号已被禁用")
		return nil, gerror.New("账号已被禁用")
	}

	identity := s.identity(user)
	loginToken, expires, err := token.Login(ctx, identity)
	if err != nil {
		return nil, gerror.Wrap(err, "登录失败，请稍后重试")
	}

	now := gtime.Now()
	request := ghttp.RequestFromCtx(ctx)
	_, _ = dao.User.Ctx(ctx).WherePri(user.Id).Data(g.Map{
		cols.LastLoginAt:  now,
		cols.LastLoginIp:  request.GetClientIp(),
		cols.LoginFailure: 0,
		cols.UpdatedAt:    now,
	}).Update()
	s.writeLoginLog(ctx, user, "")

	return &v1.LoginRes{
		Token: loginToken, Expire: expires, TokenExpireAt: gtime.Timestamp() + expires, Id: int64(user.Id), Username: user.Username,
		Nickname: user.Nickname, Avatar: user.Avatar, Mobile: user.Mobile, Email: user.Email,
		Money: user.Money, Score: user.Score,
	}, nil
}

func (s *sApiUser) Info(ctx context.Context, _ *v1.InfoReq) (res *v1.InfoRes, err error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	return s.infoRes(user), nil
}

func (s *sApiUser) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	cols := dao.User.Columns()
	if err = s.ensureUniqueUser(ctx, cols.Username, req.Username, 0, "用户名已存在"); err != nil {
		return nil, err
	}
	if req.Mobile != "" {
		if err = s.ensureUniqueUser(ctx, cols.Mobile, req.Mobile, 0, "手机号已被使用"); err != nil {
			return nil, err
		}
	}
	if req.Email != "" {
		if err = s.ensureUniqueUser(ctx, cols.Email, req.Email, 0, "邮箱已被使用"); err != nil {
			return nil, err
		}
	}

	now := gtime.Now()
	salt := grand.S(6)
	request := ghttp.RequestFromCtx(ctx)
	userID, err := dao.User.Ctx(ctx).Data(g.Map{
		cols.Username: req.Username, cols.Password: gmd5.MustEncryptString(req.Password + salt),
		cols.Salt: salt, cols.Nickname: req.Nickname, cols.Mobile: req.Mobile, cols.Email: req.Email,
		cols.Status: "enable", cols.CreatedAt: now, cols.UpdatedAt: now,
		cols.JoinedAt: now, cols.JoinIp: request.GetClientIp(),
	}).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "注册失败，请稍后重试")
	}

	user := &entity.User{Id: uint(userID), Username: req.Username, Nickname: req.Nickname, Mobile: req.Mobile, Email: req.Email, Status: "enable"}
	loginToken, expires, err := token.Login(ctx, s.identity(user))
	if err != nil {
		return nil, gerror.Wrap(err, "注册成功后登录失败，请重新登录")
	}
	s.writeLoginLog(ctx, user, "")
	s.writeUserLog(ctx, user, consts.UserLogTypeOperation, consts.UserLogActionRegister, "用户注册", "", "")
	return &v1.RegisterRes{Token: loginToken, Expire: expires, TokenExpireAt: gtime.Timestamp() + expires, Id: userID, Username: req.Username, Nickname: req.Nickname, Mobile: req.Mobile, Email: req.Email}, nil
}

func (s *sApiUser) Logout(ctx context.Context, _ *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if err = token.Logout(ghttp.RequestFromCtx(ctx)); err != nil {
		return nil, err
	}
	s.writeUserLog(ctx, user, consts.UserLogTypeOperation, consts.UserLogActionLogout, "退出登录", "", "")
	return &v1.LogoutRes{}, nil
}

func (s *sApiUser) UpdateInfo(ctx context.Context, req *v1.UpdateInfoReq) (res *v1.UpdateInfoRes, err error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	cols := dao.User.Columns()
	if _, err = dao.User.Ctx(ctx).WherePri(user.Id).Data(g.Map{
		cols.Nickname: req.Nickname, cols.Avatar: req.Avatar, cols.Gender: req.Gender,
		cols.Birthday: req.Birthday, cols.Motto: req.Motto, cols.UpdatedAt: gtime.Now(),
	}).Update(); err != nil {
		return nil, gerror.Wrap(err, "修改用户信息失败，请稍后重试")
	}
	return &v1.UpdateInfoRes{Id: int64(user.Id), Username: user.Username, Nickname: req.Nickname, Avatar: req.Avatar, Mobile: user.Mobile, Email: user.Email, Gender: req.Gender, Birthday: req.Birthday, Motto: req.Motto, Money: user.Money, Score: user.Score}, nil
}

func (s *sApiUser) UpdatePwd(ctx context.Context, req *v1.UpdatePwdReq) (res *v1.UpdatePwdRes, err error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if user.Password != gmd5.MustEncryptString(req.OldPassword+user.Salt) {
		return nil, gerror.New("原密码不正确")
	}
	if user.Password == gmd5.MustEncryptString(req.NewPassword+user.Salt) {
		return nil, gerror.New("新密码不能与原密码相同")
	}
	if err = s.savePassword(ctx, user, req.NewPassword); err != nil {
		return nil, err
	}
	if err = token.RevokeAll(ctx, consts.AppApi, int64(user.Id)); err != nil {
		return nil, gerror.Wrap(err, "密码已修改，但注销其他登录状态失败")
	}
	s.writeUserLog(ctx, user, consts.UserLogTypeOperation, consts.UserLogActionPassword, "修改登录密码", "", "")
	return &v1.UpdatePwdRes{}, nil
}

func (s *sApiUser) SendResetCode(ctx context.Context, req *v1.SendResetCodeReq) (res *v1.SendResetCodeRes, err error) {
	var user *entity.User
	if err = dao.User.Ctx(ctx).Where(dao.User.Columns().Mobile, req.Mobile).Scan(&user); err != nil {
		return nil, gerror.Wrap(err, consts.ErrorORM)
	}
	if user != nil && isUserEnabled(user) {
		if err = service.SysSmsLog().SendCode(ctx, &sysin.SendCodeInp{Event: consts.SmsTemplateResetPwd, Mobile: req.Mobile}); err != nil {
			return nil, err
		}
	}
	return &v1.SendResetCodeRes{}, nil
}

func (s *sApiUser) ResetPwd(ctx context.Context, req *v1.ResetPwdReq) (res *v1.ResetPwdRes, err error) {
	var user *entity.User
	if err = dao.User.Ctx(ctx).Where(dao.User.Columns().Mobile, req.Mobile).Scan(&user); err != nil {
		return nil, gerror.Wrap(err, consts.ErrorORM)
	}
	if user == nil || !isUserEnabled(user) {
		return nil, gerror.New("账号不存在或已被禁用")
	}
	if err = service.SysSmsLog().VerifyCode(ctx, &sysin.VerifyCodeInp{Event: consts.SmsTemplateResetPwd, Mobile: req.Mobile, Code: req.Code}); err != nil {
		return nil, err
	}
	if err = s.savePassword(ctx, user, req.Password); err != nil {
		return nil, err
	}
	if err = token.RevokeAll(ctx, consts.AppApi, int64(user.Id)); err != nil {
		return nil, gerror.Wrap(err, "密码已重置，但注销已有登录状态失败")
	}
	s.writeUserLog(ctx, user, consts.UserLogTypeOperation, consts.UserLogActionPassword, "重置登录密码", "", "")
	return &v1.ResetPwdRes{}, nil
}

func (s *sApiUser) SendMobileBindCode(ctx context.Context, req *v1.SendMobileBindCodeReq) (res *v1.SendMobileBindCodeRes, err error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if req.Mobile == user.Mobile {
		return nil, gerror.New("新手机号不能与当前手机号相同")
	}
	if err = s.ensureUniqueUser(ctx, dao.User.Columns().Mobile, req.Mobile, int64(user.Id), "手机号已被使用"); err != nil {
		return nil, err
	}
	if err = service.SysSmsLog().SendCode(ctx, &sysin.SendCodeInp{Event: consts.SmsTemplateBind, Mobile: req.Mobile}); err != nil {
		return nil, err
	}
	return &v1.SendMobileBindCodeRes{}, nil
}

func (s *sApiUser) BindMobile(ctx context.Context, req *v1.BindMobileReq) (res *v1.BindMobileRes, err error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if user.Password != gmd5.MustEncryptString(req.Password+user.Salt) {
		return nil, gerror.New("登录密码不正确")
	}
	if req.Mobile == user.Mobile {
		return nil, gerror.New("新手机号不能与当前手机号相同")
	}
	if err = s.ensureUniqueUser(ctx, dao.User.Columns().Mobile, req.Mobile, int64(user.Id), "手机号已被使用"); err != nil {
		return nil, err
	}
	if err = service.SysSmsLog().VerifyCode(ctx, &sysin.VerifyCodeInp{Event: consts.SmsTemplateBind, Mobile: req.Mobile, Code: req.Code}); err != nil {
		return nil, err
	}
	if _, err = dao.User.Ctx(ctx).WherePri(user.Id).Data(g.Map{dao.User.Columns().Mobile: req.Mobile, dao.User.Columns().UpdatedAt: gtime.Now()}).Update(); err != nil {
		return nil, gerror.Wrap(err, "换绑手机号失败，请稍后重试")
	}
	return &v1.BindMobileRes{}, nil
}

func (s *sApiUser) SendEmailBindCode(ctx context.Context, req *v1.SendEmailBindCodeReq) (res *v1.SendEmailBindCodeRes, err error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if req.Email == user.Email {
		return nil, gerror.New("新邮箱不能与当前邮箱相同")
	}
	if err = s.ensureUniqueUser(ctx, dao.User.Columns().Email, req.Email, int64(user.Id), "邮箱已被使用"); err != nil {
		return nil, err
	}
	if err = service.SysEmsLog().Send(ctx, &sysin.SendEmsInp{Event: consts.EmsTemplateBind, Email: req.Email}); err != nil {
		return nil, err
	}
	return &v1.SendEmailBindCodeRes{}, nil
}

func (s *sApiUser) BindEmail(ctx context.Context, req *v1.BindEmailReq) (res *v1.BindEmailRes, err error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if user.Password != gmd5.MustEncryptString(req.Password+user.Salt) {
		return nil, gerror.New("登录密码不正确")
	}
	if req.Email == user.Email {
		return nil, gerror.New("新邮箱不能与当前邮箱相同")
	}
	if err = s.ensureUniqueUser(ctx, dao.User.Columns().Email, req.Email, int64(user.Id), "邮箱已被使用"); err != nil {
		return nil, err
	}
	if err = service.SysEmsLog().VerifyCode(ctx, &sysin.VerifyEmsCodeInp{Event: consts.EmsTemplateBind, Email: req.Email, Code: req.Code}); err != nil {
		return nil, err
	}
	if _, err = dao.User.Ctx(ctx).WherePri(user.Id).Data(g.Map{dao.User.Columns().Email: req.Email, dao.User.Columns().UpdatedAt: gtime.Now()}).Update(); err != nil {
		return nil, gerror.Wrap(err, "换绑邮箱失败，请稍后重试")
	}
	return &v1.BindEmailRes{}, nil
}

func (s *sApiUser) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	if _, err = s.currentUser(ctx); err != nil {
		return nil, err
	}
	if !validate.InSlice(storager.KindSlice, req.UploadType) {
		return nil, gerror.New("上传类型无效")
	}
	attachment, err := service.CommonUpload().UploadFile(ctx, req.UploadType, req.File)
	if err != nil {
		return nil, err
	}
	return &v1.UploadRes{SysAttachment: &attachment.SysAttachment, SizeFormat: attachment.SizeFormat}, nil
}

func (s *sApiUser) CreditsLogList(ctx context.Context, req *v1.CreditsLogListReq) (res *v1.CreditsLogListRes, err error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	mod := dao.UserCreditsLog.Ctx(ctx).Where(dao.UserCreditsLog.Columns().UserId, user.Id)
	if req.CreditType != "" {
		mod = mod.Where(dao.UserCreditsLog.Columns().CreditType, req.CreditType)
	}
	total, err := mod.Clone().Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取资产流水失败，请稍后重试")
	}
	list := make([]*entity.UserCreditsLog, 0)
	if total > 0 {
		if err = mod.Page(req.Page, req.PerPage).OrderDesc(dao.UserCreditsLog.Columns().Id).Scan(&list); err != nil {
			return nil, gerror.Wrap(err, "获取资产流水失败，请稍后重试")
		}
	}
	res = &v1.CreditsLogListRes{List: list}
	res.PageRes.Pack(req, total)
	return res, nil
}

func (s *sApiUser) LoginLogList(ctx context.Context, req *v1.LoginLogListReq) (res *v1.LoginLogListRes, err error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	mod := dao.UserLog.Ctx(ctx).
		Where(dao.UserLog.Columns().UserId, user.Id).
		Where(dao.UserLog.Columns().LogType, consts.UserLogTypeLogin)
	total, err := mod.Clone().Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取登录历史失败，请稍后重试")
	}
	list := make([]*entity.UserLog, 0)
	if total > 0 {
		if err = mod.Page(req.Page, req.PerPage).OrderDesc(dao.UserLog.Columns().Id).Scan(&list); err != nil {
			return nil, gerror.Wrap(err, "获取登录历史失败，请稍后重试")
		}
	}
	res = &v1.LoginLogListRes{List: list}
	res.PageRes.Pack(req, total)
	return res, nil
}

func (s *sApiUser) Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if user.Password != gmd5.MustEncryptString(req.Password+user.Salt) {
		return nil, gerror.New("登录密码不正确")
	}
	if _, err = dao.User.Ctx(ctx).WherePri(user.Id).Data(g.Map{dao.User.Columns().Status: "disable", dao.User.Columns().UpdatedAt: gtime.Now()}).Update(); err != nil {
		return nil, gerror.Wrap(err, "注销账号失败，请稍后重试")
	}
	if err = token.RevokeAll(ctx, consts.AppApi, int64(user.Id)); err != nil {
		return nil, gerror.Wrap(err, "账号已注销，但注销已有登录状态失败")
	}
	s.writeUserLog(ctx, user, consts.UserLogTypeOperation, consts.UserLogActionCancel, "注销账号", "", "")
	return &v1.CancelRes{}, nil
}

func (s *sApiUser) currentUser(ctx context.Context) (*entity.User, error) {
	identity := contexts.GetUser(ctx)
	if identity == nil || identity.Id <= 0 {
		return nil, gerror.New("请先登录")
	}
	var user *entity.User
	if err := dao.User.Ctx(ctx).WherePri(identity.Id).Scan(&user); err != nil {
		return nil, gerror.Wrap(err, consts.ErrorORM)
	}
	if user == nil {
		return nil, gerror.New("用户信息不存在")
	}
	if !isUserEnabled(user) {
		return nil, gerror.New("账号已被禁用")
	}
	return user, nil
}

func (s *sApiUser) ensureUniqueUser(ctx context.Context, column, value string, userID int64, message string) error {
	mod := dao.User.Ctx(ctx).Where(column, value)
	if userID > 0 {
		mod = mod.WhereNot(dao.User.Columns().Id, userID)
	}
	count, err := mod.Count()
	if err != nil {
		return gerror.Wrap(err, consts.ErrorORM)
	}
	if count > 0 {
		return gerror.New(message)
	}
	return nil
}

func (s *sApiUser) savePassword(ctx context.Context, user *entity.User, password string) error {
	if _, err := dao.User.Ctx(ctx).WherePri(user.Id).Data(g.Map{
		dao.User.Columns().Password:  gmd5.MustEncryptString(password + user.Salt),
		dao.User.Columns().UpdatedAt: gtime.Now(),
	}).Update(); err != nil {
		return gerror.Wrap(err, "修改密码失败，请稍后重试")
	}
	return nil
}

func (s *sApiUser) identity(user *entity.User) *model.Identity {
	now := gtime.Now()
	return &model.Identity{Id: int64(user.Id), Username: user.Username, RealName: user.Nickname, Avatar: user.Avatar, Email: user.Email, Mobile: user.Mobile, App: consts.AppApi, LoginAt: now, TokenIssuedAt: now.TimestampNano()}
}

func (s *sApiUser) infoRes(user *entity.User) *v1.InfoRes {
	return &v1.InfoRes{
		Id: int64(user.Id), Username: user.Username, Nickname: user.Nickname, Avatar: user.Avatar,
		Mobile: user.Mobile, Email: user.Email, Gender: user.Gender, Birthday: user.Birthday,
		Motto: user.Motto, JoinedAt: user.JoinedAt, LastLoginAt: user.LastLoginAt,
		LastLoginIp: user.LastLoginIp, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt,
		Money: user.Money, Score: user.Score,
	}
}

func (s *sApiUser) writeLoginLog(ctx context.Context, user *entity.User, errMsg string) {
	s.writeUserLog(ctx, user, consts.UserLogTypeLogin, consts.UserLogActionLogin, "用户登录", errMsg, "")
}

func (s *sApiUser) writeUserLog(ctx context.Context, user *entity.User, logType, action, actionName, message, extra string) {
	request := ghttp.RequestFromCtx(ctx)
	if request == nil {
		return
	}
	status := consts.StatusEnabled
	if message != "" {
		status = consts.StatusDisable
	}
	now := gtime.Now()
	if _, err := dao.UserLog.Ctx(ctx).Data(&entity.UserLog{
		UserId: int64(user.Id), Username: user.Username, LogType: logType, Action: action,
		ActionName: actionName, EventAt: now, Ip: request.GetClientIp(), UserAgent: request.UserAgent(),
		Message: message, Extra: extra, Status: status, CreatedAt: now, UpdatedAt: now,
	}).Insert(); err != nil {
		g.Log().Warningf(ctx, "写入用户行为日志失败 userId:%d action:%s err:%+v", user.Id, action, err)
	}
}

func isUserEnabled(user *entity.User) bool {
	return user.Status == "enable" || user.Status == "1"
}
