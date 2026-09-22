// Package userin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package userin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserUpdateFields 修改会员管理字段过滤
type UserUpdateFields struct {
	GroupId    int         `json:"groupId"    dc:"分组ID"`
	Username   string      `json:"username"   dc:"用户名"`
	Nickname   string      `json:"nickname"   dc:"昵称"`
	Email      string      `json:"email"      dc:"邮箱"`
	Mobile     string      `json:"mobile"     dc:"手机"`
	Avatar     string      `json:"avatar"     dc:"头像"`
	Gender     int         `json:"gender"     dc:"性别:0=未知,1=男,2=女"`
	Birthday   *gtime.Time `json:"birthday"   dc:"生日"`
	Money      int         `json:"money"      dc:"余额"`
	Score      int         `json:"score"      dc:"积分"`
	JoinTime   int64       `json:"joinTime"   dc:"加入时间"`
	Motto      string      `json:"motto"      dc:"签名"`
	Status     string      `json:"status"     dc:"状态:enable=启用,disable=禁用"`
	CreateTime int64       `json:"createTime" dc:"创建时间"`
	UpdateTime int64       `json:"updateTime" dc:"更新时间"`
	Password   string      `json:"password"   dc:"密码"`
	Salt       string      `json:"salt"       dc:"密码盐"`
}

// UserInsertFields 新增会员管理字段过滤
type UserInsertFields struct {
	GroupId    int         `json:"groupId"    dc:"分组ID"`
	Username   string      `json:"username"   dc:"用户名"`
	Nickname   string      `json:"nickname"   dc:"昵称"`
	Email      string      `json:"email"      dc:"邮箱"`
	Mobile     string      `json:"mobile"     dc:"手机"`
	Avatar     string      `json:"avatar"     dc:"头像"`
	Gender     int         `json:"gender"     dc:"性别:0=未知,1=男,2=女"`
	Birthday   *gtime.Time `json:"birthday"   dc:"生日"`
	Money      int         `json:"money"      dc:"余额"`
	Score      int         `json:"score"      dc:"积分"`
	JoinTime   int64       `json:"joinTime"   dc:"加入时间"`
	Motto      string      `json:"motto"      dc:"签名"`
	Status     string      `json:"status"     dc:"状态:enable=启用,disable=禁用"`
	CreateTime int64       `json:"createTime" dc:"创建时间"`
	UpdateTime int64       `json:"updateTime" dc:"更新时间"`
	Password   string      `json:"password"   dc:"密码"`
	Salt       string      `json:"salt"       dc:"密码盐"`
}

// UserEditInp 修改/新增会员管理
type UserEditInp struct {
	entity.User
}

func (in *UserEditInp) Filter(ctx context.Context) (err error) {
	// 验证分组ID
	if err := g.Validator().Rules("required").Data(in.GroupId).Messages("分组ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证用户名
	if err := g.Validator().Rules("required").Data(in.Username).Messages("用户名不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证邮箱
	if in.Email != "" {
		if err := g.Validator().Rules("email").Data(in.Email).Messages("邮箱不是邮箱地址").Run(ctx); err != nil {
			return err.Current()
		}
	}
	// 验证手机
	if in.Mobile != "" {
		if err := g.Validator().Rules("phone").Data(in.Mobile).Messages("手机不是手机号码").Run(ctx); err != nil {
			return err.Current()
		}
	}
	// 验证余额
	if err := g.Validator().Rules("min:0").Data(in.Money).Messages("余额不能小于0").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证积分
	if err := g.Validator().Rules("min:0").Data(in.Score).Messages("积分不能小于0").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态:enable=启用,disable=禁用
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态:enable=启用,disable=禁用不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type UserEditModel struct{}

// UserDeleteInp 删除会员管理
type UserDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *UserDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type UserDeleteModel struct{}

// UserViewInp 获取指定会员管理信息
type UserViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *UserViewInp) Filter(ctx context.Context) (err error) {
	return
}

type UserViewModel struct {
	entity.User
}

// UserListInp 获取会员管理列表
type UserListInp struct {
	form.PageReq
	Id int `json:"id" dc:"ID"`
}

func (in *UserListInp) Filter(ctx context.Context) (err error) {
	return
}

type UserListModel struct {
	Id            int         `json:"id"            dc:"ID"`
	GroupId       int         `json:"groupId"       dc:"分组ID"`
	Username      string      `json:"username"      dc:"用户名"`
	Nickname      string      `json:"nickname"      dc:"昵称"`
	Email         string      `json:"email"         dc:"邮箱"`
	Mobile        string      `json:"mobile"        dc:"手机"`
	Avatar        string      `json:"avatar"        dc:"头像"`
	Gender        int         `json:"gender"        dc:"性别:0=未知,1=男,2=女"`
	Birthday      *gtime.Time `json:"birthday"      dc:"生日"`
	Money         int         `json:"money"         dc:"余额"`
	Score         int         `json:"score"         dc:"积分"`
	LastLoginTime int64       `json:"lastLoginTime" dc:"上次登录时间"`
	LastLoginIp   string      `json:"lastLoginIp"   dc:"上次登录IP"`
	LoginFailure  int         `json:"loginFailure"  dc:"登录失败次数"`
	JoinIp        string      `json:"joinIp"        dc:"加入IP"`
	JoinTime      int64       `json:"joinTime"      dc:"加入时间"`
	Status        string      `json:"status"        dc:"状态:enable=启用,disable=禁用"`
	CreateTime    int64       `json:"createTime"    dc:"创建时间"`
}

// UserExportModel 导出会员管理
type UserExportModel struct {
	Id            int         `json:"id"            dc:"ID"`
	GroupId       int         `json:"groupId"       dc:"分组ID"`
	Username      string      `json:"username"      dc:"用户名"`
	Nickname      string      `json:"nickname"      dc:"昵称"`
	Email         string      `json:"email"         dc:"邮箱"`
	Mobile        string      `json:"mobile"        dc:"手机"`
	Avatar        string      `json:"avatar"        dc:"头像"`
	Gender        int         `json:"gender"        dc:"性别:0=未知,1=男,2=女"`
	Birthday      *gtime.Time `json:"birthday"      dc:"生日"`
	Money         int         `json:"money"         dc:"余额"`
	Score         int         `json:"score"         dc:"积分"`
	LastLoginTime int64       `json:"lastLoginTime" dc:"上次登录时间"`
	LastLoginIp   string      `json:"lastLoginIp"   dc:"上次登录IP"`
	LoginFailure  int         `json:"loginFailure"  dc:"登录失败次数"`
	JoinIp        string      `json:"joinIp"        dc:"加入IP"`
	JoinTime      int64       `json:"joinTime"      dc:"加入时间"`
	Motto         string      `json:"motto"         dc:"签名"`
	Status        string      `json:"status"        dc:"状态:enable=启用,disable=禁用"`
	CreateTime    int64       `json:"createTime"    dc:"创建时间"`
}

// UserStatusInp 更新会员管理状态
type UserStatusInp struct {
	Id     int    `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status string `json:"status" dc:"状态"`
}

func (in *UserStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status == "" {
		err = gerror.New("状态不能为空")
		return
	}

	if in.Status != "enable" && in.Status != "disable" {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type UserStatusModel struct{}

// UserSwitchInp 更新会员管理开关状态
type UserSwitchInp struct {
	Key   string `json:"key" v:"required#开关参数错误" dc:"键"`
	Value string `json:"value" dc:"值"`
	Id    int    `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *UserSwitchInp) Filter(ctx context.Context) (err error) {
	return
}

type UserSwitchModel struct{}
