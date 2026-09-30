// Package token
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package token

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/library/cache"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	*model.Identity
	jwt.RegisteredClaims
}

type Token struct {
	ExpireAt     int64 `json:"exp"` // token过期时间
	RefreshAt    int64 `json:"ra"`  // 刷新时间
	RefreshCount int64 `json:"rc"`  // 刷新次数
}

var (
	config          *model.TokenConfig
	errorLogin      = gerror.New("登录身份已失效，请重新登录！")
	errorMultiLogin = gerror.New("账号已在其他地方登录，如非本人操作请及时修改登录密码！")
)

func SetConfig(c *model.TokenConfig) {
	config = c
}

func GetConfig() *model.TokenConfig {
	return config
}

// Login 登录
func Login(ctx context.Context, user *model.Identity) (string, int64, error) {
	if user.TokenIssuedAt <= 0 {
		user.TokenIssuedAt = gtime.TimestampNano()
	}
	claims := Claims{
		user,
		jwt.RegisteredClaims{},
	}

	header, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", 0, err
	}

	var (
		now = gtime.Now()
		// 认证key
		authKey = GetAuthKey(header)
		// 登录token
		tokenKey = GetTokenKey(user.App, authKey)
		// 身份绑定
		bindKey = GetBindKey(user.App, user.Id)
		// 有效时长
		duration = time.Second * gconv.Duration(config.Expires)
	)

	token := &Token{
		ExpireAt:     now.Unix() + config.Expires,
		RefreshAt:    now.Unix(),
		RefreshCount: 0,
	}

	if err = cache.Instance().Set(ctx, tokenKey, token, duration); err != nil {
		return "", 0, err
	}

	if err = cache.Instance().Set(ctx, bindKey, tokenKey, duration); err != nil {
		return "", 0, err
	}
	return header, config.Expires, nil
}

// Logout 注销登录
func Logout(r *ghttp.Request) (err error) {
	var (
		ctx    = r.Context()
		header = GetAuthorization(r)
	)

	if header == "" {
		err = errorLogin
		return
	}

	claims, err := parseToken(ctx, header)
	if err != nil {
		g.Log().Debugf(ctx, "logout parseToken err:%+v", err)
		err = errorLogin
		return
	}

	var (
		// 认证key
		authKey = GetAuthKey(header)
		// 登录token
		tokenKey = GetTokenKey(contexts.GetModule(ctx), authKey)
		// 身份绑定
		bindKey = GetBindKey(contexts.GetModule(ctx), claims.Id)
	)

	// 删除token
	if _, err = cache.Instance().Remove(ctx, tokenKey); err != nil {
		return
	}

	if !config.MultiLogin {
		if _, err = cache.Instance().Remove(ctx, bindKey); err != nil {
			return
		}
	}
	return
}

// RevokeAll 撤销指定用户此前签发的全部登录令牌。
func RevokeAll(ctx context.Context, appName string, userId int64) error {
	// 纳秒时间戳用于区分紧邻发生的“撤销旧令牌”和“重新登录签发新令牌”。
	return cache.Instance().Set(ctx, GetRevokedAtKey(appName, userId), gtime.TimestampNano(), 0)
}

// ParseLoginUser 解析登录用户信息
func ParseLoginUser(r *ghttp.Request) (user *model.Identity, err error) {
	var (
		ctx    = r.Context()
		header = GetAuthorization(r)
	)

	if header == "" {
		err = errorLogin
		return
	}

	claims, err := parseToken(ctx, header)
	if err != nil {
		g.Log().Debugf(ctx, "parseToken err:%+v", err)
		err = errorLogin
		return
	}

	// 修改密码、重置密码或注销账号后，拒绝此前签发的所有令牌。
	revokedAt, revokeErr := cache.Instance().Get(ctx, GetRevokedAtKey(claims.App, claims.Id))
	if revokeErr != nil {
		g.Log().Debugf(ctx, "get revokedAt err:%+v", revokeErr)
		err = errorLogin
		return nil, err
	}
	if !revokedAt.IsEmpty() && (claims.TokenIssuedAt <= 0 || claims.TokenIssuedAt <= revokedAt.Int64()) {
		g.Log().Debugf(ctx, "token has been revoked userId:%d", claims.Id)
		err = errorLogin
		return nil, err
	}

	var (
		// 认证key
		authKey = GetAuthKey(header)
		// 登录token
		tokenKey = GetTokenKey(claims.App, authKey)
		// 身份绑定
		bindKey = GetBindKey(claims.App, claims.Id)
	)

	// 检查token是否存在
	tk, err := cache.Instance().Get(ctx, tokenKey)
	if err != nil {
		g.Log().Debugf(ctx, "get tokenKey err:%+v", err)
		err = errorLogin
		return
	}

	if tk.IsEmpty() {
		g.Log().Debug(ctx, "token isEmpty")
		err = errorLogin
		return
	}

	var token *Token
	if err = tk.Scan(&token); err != nil {
		g.Log().Debugf(ctx, "token scan err:%+v", err)
		err = errorLogin
		return
	}

	if token == nil {
		g.Log().Debugf(ctx, "token = nil")
		err = errorLogin
		return
	}

	now := gtime.Now()
	if token.ExpireAt < now.Unix() {
		g.Log().Debugf(ctx, "token expired.")
		err = errorLogin
		return
	}

	// 是否允许多端登录
	if !config.MultiLogin {
		origin, err := cache.Instance().Get(ctx, bindKey)
		if err != nil {
			g.Log().Debugf(ctx, "bindKey get err:%+v", err)
			err = errorLogin
			return nil, err
		}

		if origin == nil || origin.IsEmpty() {
			g.Log().Debug(ctx, "bindKey isEmpty")
			err = errorLogin
			return nil, err
		}

		if tokenKey != origin.String() {
			g.Log().Debugf(ctx, "bindKey offsite login tokenKey:%v, origin:%v", tokenKey, origin.String())
			err = errorMultiLogin
			return nil, err
		}
	}

	// 达到刷新间隔后同步续期，确保本次响应可返回准确的绝对过期时间。
	if config.AutoRefresh &&
		(config.MaxRefreshTimes == -1 || token.RefreshCount < config.MaxRefreshTimes) &&
		token.RefreshAt+config.RefreshInterval <= now.Unix() {
		token.ExpireAt = now.Unix() + config.Expires
		token.RefreshAt = now.Unix()
		token.RefreshCount++

		duration := time.Second * gconv.Duration(config.Expires)
		if err = cache.Instance().Set(ctx, tokenKey, token, duration); err != nil {
			return nil, err
		}
		if err = cache.Instance().Set(ctx, bindKey, tokenKey, duration); err != nil {
			return nil, err
		}
	}

	r.Response.Header().Set("X-Token-Expire-At", gconv.String(token.ExpireAt))
	user = claims.Identity
	return
}

// parseToken 解析jwt令牌
func parseToken(ctx context.Context, header string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(header, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		g.Log().Debugf(ctx, "parseToken err:%+v", err)
		return nil, err
	}

	if !token.Valid {
		return nil, errorLogin
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errorLogin
	}
	return claims, nil
}

// GetAuthorization 获取authorization
func GetAuthorization(r *ghttp.Request) string {
	// 默认从请求头获取 Authorization
	var authorization = r.Header.Get("Authorization")

	// 如果 Authorization 不存在，尝试获取 userToken (支持 Headers)
	if authorization == "" {
		authorization = r.Header.Get("userToken")
	}

	// 如果请求头中都不存在，则尝试从 GET 参数中获取 authorization 或 userToken
	if authorization == "" {
		authorization = r.Get("authorization").String()
		if authorization == "" {
			authorization = r.Get("userToken").String()
		}
	}

	return gstr.Replace(authorization, "Bearer ", "")
}

// GetAuthKey 认证key
func GetAuthKey(token string) string {
	return gmd5.MustEncryptString("hotgo" + token)
}

// GetTokenKey 令牌缓存key
func GetTokenKey(appName, authKey string) string {
	return fmt.Sprintf("%v:%v:%v", consts.CacheToken, appName, authKey)
}

// GetBindKey 令牌身份绑定key
func GetBindKey(appName string, userId int64) string {
	return fmt.Sprintf("%v:%v:%v", consts.CacheTokenBind, appName, userId)
}

// GetRevokedAtKey 获取用户全部令牌撤销时间缓存key。
func GetRevokedAtKey(appName string, userId int64) string {
	return fmt.Sprintf("%v:%v:%v", consts.CacheTokenRevokedAt, appName, userId)
}
