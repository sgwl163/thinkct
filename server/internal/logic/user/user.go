// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package user

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/userin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type sUserUser struct{}

func NewUserUser() *sUserUser {
	return &sUserUser{}
}

func init() {
	service.RegisterUserUser(NewUserUser())
}

// Model 会员管理ORM模型
func (s *sUserUser) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.User.Ctx(ctx), option...)
}

// List 获取会员管理列表
func (s *sUserUser) List(ctx context.Context, in *userin.UserListInp) (list []*userin.UserListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(userin.UserListModel{})

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.User.Columns().Id, in.Id)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.User.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取会员管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出会员管理
func (s *sUserUser) Export(ctx context.Context, in *userin.UserListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(userin.UserExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出会员管理-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []userin.UserExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增会员管理
func (s *sUserUser) Edit(ctx context.Context, in *userin.UserEditInp) (err error) {
	// 验证'Username'唯一
	if err = hgorm.IsUnique(ctx, &dao.User, g.Map{dao.User.Columns().Username: in.Username}, "用户名已存在", in.Id); err != nil {
		return
	}
	if in.Mobile != "" {
		if err = hgorm.IsUnique(ctx, &dao.User, g.Map{dao.User.Columns().Mobile: in.Mobile}, "手机号已存在", in.Id); err != nil {
			return
		}
	}
	if in.Email != "" {
		if err = hgorm.IsUnique(ctx, &dao.User, g.Map{dao.User.Columns().Email: in.Email}, "邮箱已存在", in.Id); err != nil {
			return
		}
	}

	if in.Id <= 0 && in.Password == "" {
		err = gerror.New("新增会员时密码不能为空")
		return
	}

	if in.Password != "" {
		in.Salt = grand.S(6)
		in.Password = gmd5.MustEncryptString(in.Password + in.Salt)
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			var before *entity.User
			if err = s.Model(ctx).WherePri(in.Id).Lock("FOR UPDATE").Scan(&before); err != nil {
				return gerror.Wrap(err, "获取会员信息失败，请稍后重试！")
			}
			if before == nil {
				return gerror.New("会员不存在")
			}

			in.UpdateTime = uint64(gtime.Now().Unix())
			data := gconv.Map(in)
			if in.Password == "" {
				delete(data, dao.User.Columns().Password)
				delete(data, dao.User.Columns().Salt)
			}
			if _, err = s.Model(ctx).
				Fields(userin.UserUpdateFields{}).
				WherePri(in.Id).Data(data).Update(); err != nil {
				return gerror.Wrap(err, "修改会员管理失败，请稍后重试！")
			}
			if err = s.writeCreditsLog(ctx, before, in); err != nil {
				return err
			}
			return
		}

		// 新增
		in.CreateTime = uint64(gtime.Now().Unix())
		in.UpdateTime = uint64(gtime.Now().Unix())
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(userin.UserInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增会员管理失败，请稍后重试！")
		}
		return
	})
}

func (s *sUserUser) writeCreditsLog(ctx context.Context, before *entity.User, in *userin.UserEditInp) error {
	if int64(before.Money) == int64(in.Money) && int64(before.Score) == int64(in.Score) {
		return nil
	}

	now := gtime.Now()
	cols := dao.UserCreditsLog.Columns()
	remark := fmt.Sprintf("后台用户:%d 编辑会员资料", contexts.GetUserId(ctx))
	if int64(before.Money) != int64(in.Money) {
		beforeNum, afterNum := int64(before.Money), int64(in.Money)
		if _, err := dao.UserCreditsLog.Ctx(ctx).Data(g.Map{
			cols.UserId:     before.Id,
			cols.CreditType: consts.UserCreditTypeMoney,
			cols.BeforeNum:  beforeNum,
			cols.Num:        afterNum - beforeNum,
			cols.AfterNum:   afterNum,
			cols.Remark:     remark,
			cols.Status:     consts.StatusEnabled,
			cols.CreatedAt:  now,
			cols.UpdatedAt:  now,
		}).Insert(); err != nil {
			return gerror.Wrap(err, "记录余额变动流水失败，请稍后重试！")
		}
	}
	if int64(before.Score) != int64(in.Score) {
		beforeNum, afterNum := int64(before.Score), int64(in.Score)
		if _, err := dao.UserCreditsLog.Ctx(ctx).Data(g.Map{
			cols.UserId:     before.Id,
			cols.CreditType: consts.UserCreditTypeScore,
			cols.BeforeNum:  beforeNum,
			cols.Num:        afterNum - beforeNum,
			cols.AfterNum:   afterNum,
			cols.Remark:     remark,
			cols.Status:     consts.StatusEnabled,
			cols.CreatedAt:  now,
			cols.UpdatedAt:  now,
		}).Insert(); err != nil {
			return gerror.Wrap(err, "记录积分变动流水失败，请稍后重试！")
		}
	}
	return nil
}

// Delete 删除会员管理
func (s *sUserUser) Delete(ctx context.Context, in *userin.UserDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除会员管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取会员管理指定信息
func (s *sUserUser) View(ctx context.Context, in *userin.UserViewInp) (res *userin.UserViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取会员管理信息，请稍后重试！")
		return
	}
	return
}

// Status 更新会员管理状态
func (s *sUserUser) Status(ctx context.Context, in *userin.UserStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.User.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新会员管理状态失败，请稍后重试！")
		return
	}
	return
}

// Switch 更新会员管理开关
func (s *sUserUser) Switch(ctx context.Context, in *userin.UserSwitchInp) (err error) {
	var fields = []string{
		dao.User.Columns().Status,
		// ...
	}

	if !validate.InSlice(fields, in.Key) {
		err = gerror.New("开关键名不在白名单")
		return
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		in.Key: in.Value,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新会员管理开关失败，请稍后重试！")
		return
	}
	return
}
