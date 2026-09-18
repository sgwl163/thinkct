package user

import (
	"context"

	v1 "hotgo/api/api/user/v1"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/token"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"

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

// Login 用户登录
func (s *sApiUser) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	var mb *entity.User
	if err = dao.User.Ctx(ctx).
		Where(dao.User.Columns().Username, req.Username).
		WhereOr(dao.User.Columns().Email, req.Username).
		WhereOr(dao.User.Columns().Mobile, req.Username).
		Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if mb == nil {
		err = gerror.New("用户名或密码错误")
		return
	}

	// entity.User中的Status为string类型，注释说明: enable=启用,disable=禁用
	// 或者可能是 "1" / "2" 根据实际情况，这里兼容 "enable" 和 "1"
	if mb.Status != "enable" && mb.Status != "1" {
		err = gerror.New("账号已被禁用")
		return
	}

	// 这里直接使用明文密码+Salt进行MD5加密校验
	// HotGo默认后台是使用 AES+Base64 传输密码的（simple.CheckPassword），为了方便您直接在接口测试工具中使用明文密码测试，这里改为了直接比对。
	if mb.Password != gmd5.MustEncryptString(req.Password+mb.Salt) {
		dao.User.Ctx(ctx).WherePri(mb.Id).Increment(dao.User.Columns().LoginFailure, 1)
		err = gerror.New("用户名或密码错误")
		return
	}

	user := &model.Identity{
		Id:       int64(mb.Id),
		Pid:      0, // 如果有上级可以在此配置
		Username: mb.Username,
		RealName: mb.Nickname,
		Avatar:   mb.Avatar,
		Email:    mb.Email,
		Mobile:   mb.Mobile,
		App:      consts.AppApi,
		LoginAt:  gtime.Now(),
	}

	lt, expires, err := token.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	res = &v1.LoginRes{
		Token:    lt,
		Expire:   expires,
		Id:       int64(mb.Id),
		Username: mb.Username,
		Nickname: mb.Nickname,
		Avatar:   mb.Avatar,
		Mobile:   mb.Mobile,
		Email:    mb.Email,
		Money:    mb.Money,
		Score:    mb.Score,
	}

	// 记录登录时间和IP
	dao.User.Ctx(ctx).WherePri(mb.Id).Data(g.Map{
		dao.User.Columns().LastLoginTime: gtime.Now().Unix(),
		dao.User.Columns().LastLoginIp:   ghttp.RequestFromCtx(ctx).GetClientIp(),
		dao.User.Columns().LoginFailure:  0,
	}).Update()

	return
}

// Info 获取用户信息
func (s *sApiUser) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	userCtx := contexts.GetUser(ctx)
	if userCtx == nil {
		err = gerror.New("请先登录")
		return
	}

	var mb *entity.User
	if err = dao.User.Ctx(ctx).WherePri(userCtx.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	if mb == nil {
		err = gerror.New("用户信息不存在")
		return
	}

	if mb.Status != "enable" && mb.Status != "1" {
		err = gerror.New("账号已被禁用")
		return
	}

	res = &v1.InfoRes{
		Id:       int64(mb.Id),
		Username: mb.Username,
		Nickname: mb.Nickname,
		Avatar:   mb.Avatar,
		Mobile:   mb.Mobile,
		Email:    mb.Email,
		Money:    mb.Money,
		Score:    mb.Score,
	}
	return
}

// Register 用户注册
func (s *sApiUser) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	// 检查用户名是否存在
	count, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, req.Username).Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	if count > 0 {
		err = gerror.New("用户名已存在")
		return
	}

	salt := grand.S(6)
	pwd := gmd5.MustEncryptString(req.Password + salt)

	// 插入数据并获取自增的ID
	insertId, err := dao.User.Ctx(ctx).Data(g.Map{
		dao.User.Columns().Username:   req.Username,
		dao.User.Columns().Password:   pwd,
		dao.User.Columns().Salt:       salt,
		dao.User.Columns().Nickname:   req.Nickname,
		dao.User.Columns().Mobile:     req.Mobile,
		dao.User.Columns().Email:      req.Email,
		dao.User.Columns().Status:     "enable", // 默认启用
		dao.User.Columns().CreateTime: gtime.Now().Unix(),
		dao.User.Columns().JoinTime:   gtime.Now().Unix(),
		dao.User.Columns().JoinIp:     ghttp.RequestFromCtx(ctx).GetClientIp(),
	}).InsertAndGetId()
	if err != nil {
		err = gerror.Wrap(err, "注册失败")
		return
	}

	// 注册成功后自动执行登录逻辑生成Token
	user := &model.Identity{
		Id:       insertId,
		Pid:      0, // 如果有上级可以在此配置
		Username: req.Username,
		RealName: req.Nickname,
		Email:    req.Email,
		Mobile:   req.Mobile,
		App:      consts.AppApi,
		LoginAt:  gtime.Now(),
	}

	lt, expires, err := token.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	res = &v1.RegisterRes{
		Token:    lt,
		Expire:   expires,
		Id:       insertId,
		Username: req.Username,
		Nickname: req.Nickname,
		Avatar:   "", // 注册时可以默认为空
		Mobile:   req.Mobile,
		Email:    req.Email,
		Money:    0,
		Score:    0,
	}
	return
}

// Logout 退出登录
func (s *sApiUser) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	if err = token.Logout(r); err != nil {
		return
	}
	res = &v1.LogoutRes{}
	return
}

// UpdateInfo 修改用户信息
func (s *sApiUser) UpdateInfo(ctx context.Context, req *v1.UpdateInfoReq) (res *v1.UpdateInfoRes, err error) {
	userCtx := contexts.GetUser(ctx)
	if userCtx == nil {
		err = gerror.New("请先登录")
		return
	}

	var mb *entity.User
	if err = dao.User.Ctx(ctx).WherePri(userCtx.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	if mb == nil {
		err = gerror.New("用户信息不存在")
		return
	}

	if mb.Status != "enable" && mb.Status != "1" {
		err = gerror.New("账号已被禁用")
		return
	}

	// 手机号唯一性校验
	if req.Mobile != "" {
		count, err2 := dao.User.Ctx(ctx).
			Where(dao.User.Columns().Mobile, req.Mobile).
			WhereNot(dao.User.Columns().Id, mb.Id).
			Count()
		if err2 != nil {
			err = gerror.Wrap(err2, consts.ErrorORM)
			return
		}
		if count > 0 {
			err = gerror.New("手机号已被使用")
			return
		}
	}

	// 邮箱唯一性校验
	if req.Email != "" {
		count, err2 := dao.User.Ctx(ctx).
			Where(dao.User.Columns().Email, req.Email).
			WhereNot(dao.User.Columns().Id, mb.Id).
			Count()
		if err2 != nil {
			err = gerror.Wrap(err2, consts.ErrorORM)
			return
		}
		if count > 0 {
			err = gerror.New("邮箱已被使用")
			return
		}
	}

	update := g.Map{
		dao.User.Columns().Nickname:   req.Nickname,
		dao.User.Columns().Avatar:     req.Avatar,
		dao.User.Columns().Mobile:     req.Mobile,
		dao.User.Columns().Email:      req.Email,
		dao.User.Columns().Gender:     req.Gender,
		dao.User.Columns().Birthday:   req.Birthday,
		dao.User.Columns().Motto:      req.Motto,
		dao.User.Columns().UpdateTime: gtime.Now().Unix(),
	}

	if _, err = dao.User.Ctx(ctx).WherePri(mb.Id).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, "修改用户信息失败，请稍后重试！")
		return
	}

	res = &v1.UpdateInfoRes{
		Id:       int64(mb.Id),
		Username: mb.Username,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Gender:   req.Gender,
		Birthday: req.Birthday,
		Motto:    req.Motto,
		Money:    mb.Money,
		Score:    mb.Score,
	}
	return
}
