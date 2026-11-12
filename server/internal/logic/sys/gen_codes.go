// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"database/sql"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/casbin"
	"hotgo/internal/library/hggen"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/tree"
	"path/filepath"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

type sSysGenCodes struct{}

func NewSysGenCodes() *sSysGenCodes {
	return &sSysGenCodes{}
}

func init() {
	service.RegisterSysGenCodes(NewSysGenCodes())
}

// Delete 删除未生成的配置。已开始生成的记录必须通过 DeleteBuild 清理，防止遗留文件或菜单。
func (s *sSysGenCodes) Delete(ctx context.Context, in *sysin.GenCodesDeleteInp) (err error) {
	lock := generationLock(0)
	locked, err := lock.TryLock()
	if err != nil {
		return gerror.Wrap(err, "获取代码生成锁失败")
	}
	if !locked {
		return gerror.New("代码正在生成，请完成后再删除")
	}
	defer lock.Close()
	var records []entity.SysGenCodes
	if err = dao.SysGenCodes.Ctx(ctx).Where(dao.SysGenCodes.Columns().Id, in.Id).Scan(&records); err != nil {
		return gerror.Wrap(err, "获取生成配置失败")
	}
	for _, record := range records {
		row := &sysin.GenCodesViewModel{SysGenCodes: record}
		if err = s.inspectBuild(ctx, row); err != nil {
			return err
		}
		if row.Building || row.Status != consts.GenCodesStatusWait || row.BuildFiles > 0 {
			return gerror.New("所选配置已有生成记录或代码文件，请使用删除生成代码，避免遗留文件和菜单")
		}
	}
	_, err = dao.SysGenCodes.Ctx(ctx).
		Where(dao.SysGenCodes.Columns().Id, in.Id).
		Where(dao.SysGenCodes.Columns().Status, consts.GenCodesStatusWait).
		Delete()
	if err != nil {
		return gerror.Wrap(err, "删除生成配置失败")
	}
	return
}

// DeleteBuild 删除已生成代码、菜单SQL和已导入的菜单。
// DAO 和 service 文件由多个生成记录共享，只会在删除 logic 文件后重新生成 service 内容。
func (s *sSysGenCodes) DeleteBuild(ctx context.Context, in *sysin.GenCodesDeleteBuildInp) (err error) {
	lock := generationLock(0)
	locked, err := lock.TryLock()
	if err != nil {
		return gerror.Wrap(err, "获取代码生成锁失败")
	}
	if !locked {
		return gerror.New("代码正在生成，请完成后再删除")
	}
	defer lock.Close()
	var record entity.SysGenCodes
	if err = dao.SysGenCodes.Ctx(ctx).Where(dao.SysGenCodes.Columns().Id, in.Id).Scan(&record); err != nil {
		return gerror.Wrap(err, "获取生成配置失败")
	}
	if record.Id == 0 {
		return gerror.New("生成配置不存在")
	}
	row := &sysin.GenCodesViewModel{SysGenCodes: record}
	if err = s.inspectBuild(ctx, row); err != nil {
		return err
	}
	record = row.SysGenCodes
	if row.Building {
		return gerror.New("代码正在生成，请完成后再删除")
	}
	if record.Status != consts.GenCodesStatusOk && record.Status != consts.GenCodesStatusFail && row.BuildFiles == 0 {
		return gerror.New("该配置尚未生成代码")
	}
	if record.GenType != consts.GenCodesTypeCurd && record.GenType != consts.GenCodesTypeTree {
		return gerror.New("该生成类型暂不支持删除生成代码")
	}

	files, template, err := s.getBuildFiles(ctx, &record)
	if err != nil {
		return err
	}
	for _, path := range files {
		if !gfile.Exists(path) {
			continue
		}
		if err = gfile.RemoveFile(path); err != nil {
			return gerror.Wrapf(err, "删除生成文件失败:%s", path)
		}
	}

	if record.Options != nil && gstr.InArray(record.Options.Get("autoOps").Strings(), "genMenuPermissions") {
		if err = s.deleteBuildMenus(ctx, gstr.LcFirst(record.VarName)); err != nil {
			return err
		}
	}

	if err = s.refreshBuildServices(ctx, &record, template); err != nil {
		return gerror.Wrap(err, "更新业务服务失败")
	}

	if err = s.Status(ctx, &sysin.GenCodesStatusInp{Id: record.Id, Status: consts.GenCodesStatusWait}); err != nil {
		return gerror.Wrap(err, "更新生成状态失败")
	}
	return
}

func (s *sSysGenCodes) refreshBuildServices(ctx context.Context, record *entity.SysGenCodes, template model.GenerateAppCrudTemplate) error {
	return s.generateBuildServices(ctx, record, template, true)
}

func (s *sSysGenCodes) getBuildFiles(ctx context.Context, record *entity.SysGenCodes) (paths []string, template model.GenerateAppCrudTemplate, err error) {
	config, err := service.SysConfig().GetLoadGenerate(ctx)
	if err != nil {
		return nil, template, gerror.Wrap(err, "加载生成配置失败")
	}
	if config == nil || record.GenTemplate < 0 || record.GenTemplate >= len(config.Application.Crud.Templates) {
		return nil, template, gerror.New("生成模板不存在")
	}
	if config.Application.Crud.Templates[record.GenTemplate] == nil {
		return nil, template, gerror.New("生成模板不存在")
	}

	template = *config.Application.Crud.Templates[record.GenTemplate]
	if template.IsAddon {
		if record.AddonName == "" {
			return nil, template, gerror.New("插件名称不能为空")
		}
		pathsToReplace := []*string{
			&template.ApiPath,
			&template.InputPath,
			&template.ControllerPath,
			&template.LogicPath,
			&template.RouterPath,
			&template.SqlPath,
			&template.WebApiPath,
			&template.WebViewsPath,
		}
		for _, path := range pathsToReplace {
			*path = gstr.Replace(*path, "{$name}", record.AddonName)
		}
	}

	varName := gstr.UcFirst(record.VarName)
	var columnOps, autoOps []string
	if record.Options != nil {
		columnOps = record.Options.Get("columnOps").Strings()
		autoOps = record.Options.Get("autoOps").Strings()
	}
	paths = []string{
		filepath.Join(template.ApiPath, gstr.ToLower(varName), gstr.ToLower(varName)+".go"),
		filepath.Join(template.InputPath, convert.CamelCaseToUnderline(varName)+".go"),
		filepath.Join(template.ControllerPath, convert.CamelCaseToUnderline(varName)+".go"),
		filepath.Join(template.LogicPath, convert.CamelCaseToUnderline(varName)+".go"),
		filepath.Join(template.RouterPath, convert.CamelCaseToUnderline(varName)+".go"),
		filepath.Join(template.WebApiPath, gstr.LcFirst(varName), "index.ts"),
		filepath.Join(template.WebViewsPath, gstr.LcFirst(varName), "model.ts"),
		filepath.Join(template.WebViewsPath, gstr.LcFirst(varName), "index.vue"),
	}
	if gstr.InArray(columnOps, "edit") {
		paths = append(paths, filepath.Join(template.WebViewsPath, gstr.LcFirst(varName), "edit.vue"))
	}
	if gstr.InArray(columnOps, "view") {
		paths = append(paths, filepath.Join(template.WebViewsPath, gstr.LcFirst(varName), "view.vue"))
	}
	if gstr.InArray(autoOps, "genMenuPermissions") {
		paths = append(paths, filepath.Join(template.SqlPath, convert.CamelCaseToUnderline(varName)+"_menu.sql"))
	}
	// filepath.Abs 无需目录已存在，避免 RealPath 空值将文件错误指向工作目录。
	for index, path := range paths {
		if paths[index], err = filepath.Abs(path); err != nil {
			return nil, template, gerror.Wrap(err, "解析生成文件路径失败")
		}
	}
	return
}

func (s *sSysGenCodes) deleteBuildMenus(ctx context.Context, name string) (err error) {
	var root entity.AdminMenu
	if err = dao.AdminMenu.Ctx(ctx).Where(dao.AdminMenu.Columns().Name, name).Scan(&root); err != nil {
		// 未导入过菜单或上次清理已删除菜单时，允许继续清理文件和刷新服务。
		if gerror.Is(err, sql.ErrNoRows) {
			return nil
		}
		return gerror.Wrap(err, "查询生成菜单失败")
	}
	if root.Id == 0 {
		return nil
	}

	var menus []*entity.AdminMenu
	if err = dao.AdminMenu.Ctx(ctx).Scan(&menus); err != nil {
		return gerror.Wrap(err, "查询菜单树失败")
	}

	ids := []int64{root.Id}
	for _, menu := range menus {
		for _, ancestorID := range tree.GetIds(menu.Tree) {
			if ancestorID == root.Id {
				ids = append(ids, menu.Id)
				break
			}
		}
	}
	if _, err = dao.AdminMenu.Ctx(ctx).WhereIn(dao.AdminMenu.Columns().Id, convert.UniqueSlice(ids)).Delete(); err != nil {
		return gerror.Wrap(err, "删除生成菜单失败")
	}
	return casbin.Refresh(ctx)
}

// Edit 修改/新增
func (s *sSysGenCodes) Edit(ctx context.Context, in *sysin.GenCodesEditInp) (res *sysin.GenCodesEditModel, err error) {
	if err = bindAddonMenuParent(ctx, &in.SysGenCodes); err != nil {
		return nil, err
	}
	if in.GenType == 0 {
		err = gerror.New("生成类型不能为空")
		return
	}
	if in.VarName == "" {
		err = gerror.New("实体名称不能为空")
		return
	}

	if in.GenType == consts.GenCodesTypeCurd {
		var temp *model.GenerateAppCrudTemplate
		cfg := fmt.Sprintf("hggen.application.crud.templates.%v", in.GenTemplate)
		if err = g.Cfg().MustGet(ctx, cfg).Scan(&temp); err != nil {
			return
		}

		if temp == nil {
			err = gerror.Newf("选择的模板不存在:%v", cfg)
			return
		}

		if temp.IsAddon && in.AddonName == "" {
			err = gerror.New("插件模板必须选择一个有效的插件")
			return
		}
	}

	if in.Id > 0 {
		lock := generationLock(in.Id)
		locked, lockErr := lock.TryLock()
		if lockErr != nil {
			return nil, gerror.Wrap(lockErr, "获取生成配置锁失败")
		}
		if !locked {
			return nil, gerror.New("代码正在生成，请完成后再保存配置")
		}
		defer lock.Close()
	}
	// 修改时排除状态、主键和创建时间，防止旧页面覆盖生成结果。
	if in.Id > 0 {
		_, err = dao.SysGenCodes.Ctx(ctx).Fields(sysin.GenCodesUpdateFields{}).
			Where(dao.SysGenCodes.Columns().Id, in.Id).Data(in).Update()
		if err != nil {
			return nil, gerror.Wrap(err, "保存生成配置失败")
		}
		var record entity.SysGenCodes
		if err = dao.SysGenCodes.Ctx(ctx).Where(dao.SysGenCodes.Columns().Id, in.Id).Scan(&record); err != nil {
			return nil, gerror.Wrap(err, "读取生成配置失败")
		}
		if record.Id == 0 {
			return nil, gerror.New("生成配置不存在")
		}
		return &sysin.GenCodesEditModel{SysGenCodes: record}, nil
	}

	// 新增
	if in.Options.IsNil() {
		in.Options = gjson.New(consts.NilJsonToString)
	}

	in.MasterColumns = gjson.New("[]")
	in.Status = consts.GenCodesStatusWait
	id, err := dao.SysGenCodes.Ctx(ctx).Fields(sysin.GenCodesInsertFields{}).Data(in).OmitEmptyData().InsertAndGetId()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	in.Id = id
	return &sysin.GenCodesEditModel{SysGenCodes: in.SysGenCodes}, nil
}

// Status 更新生成状态，事务提交成功后才允许向调用方报告完成。
func (s *sSysGenCodes) Status(ctx context.Context, in *sysin.GenCodesStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if _, ok := consts.GenCodesStatusNameMap[in.Status]; !ok {
		return gerror.New("生成状态不正确")
	}

	// 即使复用了 AUTOCOMMIT=0 的旧连接，也必须显式提交；不能仅凭 UPDATE 成功判断落库。
	return dao.SysGenCodes.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := dao.SysGenCodes.Ctx(ctx).Where(dao.SysGenCodes.Columns().Id, in.Id).
			Data(dao.SysGenCodes.Columns().Status, in.Status).Update()
		if err != nil {
			return gerror.Wrap(err, "保存生成状态失败")
		}
		affected, err := result.RowsAffected()
		if err != nil {
			return gerror.Wrap(err, "读取生成状态更新结果失败")
		}
		if affected == 0 {
			// MySQL 同值更新也可能影响 0 行，不能直接将其判定为记录不存在。
			record, queryErr := dao.SysGenCodes.Ctx(ctx).Fields(dao.SysGenCodes.Columns().Id).
				Where(dao.SysGenCodes.Columns().Id, in.Id).One()
			if queryErr != nil {
				return gerror.Wrap(queryErr, "核对生成配置失败")
			}
			if len(record) == 0 {
				return gerror.New("生成配置不存在，无法保存生成状态")
			}
		}
		return nil
	})
}

// MaxSort 最大排序
func (s *sSysGenCodes) MaxSort(ctx context.Context, in *sysin.GenCodesMaxSortInp) (res *sysin.GenCodesMaxSortModel, err error) {
	if in.Id > 0 {
		if err = dao.SysGenCodes.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	if res == nil {
		res = new(sysin.GenCodesMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取指定字典类型信息
func (s *sSysGenCodes) View(ctx context.Context, in *sysin.GenCodesViewInp) (res *sysin.GenCodesViewModel, err error) {
	err = dao.SysGenCodes.Ctx(ctx).Where(dao.SysGenCodes.Columns().Id, in.Id).Scan(&res)
	if err != nil || res == nil {
		return
	}
	err = s.inspectBuild(ctx, res)
	return
}

// List 获取列表
func (s *sSysGenCodes) List(ctx context.Context, in *sysin.GenCodesListInp) (list []*sysin.GenCodesListModel, totalCount int, err error) {
	mod := dao.SysGenCodes.Ctx(ctx)

	if in.GenType > 0 {
		mod = mod.Where("gen_type", in.GenType)
	}

	if in.VarName != "" {
		mod = mod.Where("var_name", in.VarName)
	}

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil || totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	typeSelect, err := hggen.GenTypeSelect(ctx)
	if err != nil {
		return
	}

	getTpGroup := func(row *sysin.GenCodesListModel) string {
		if row == nil {
			return ""
		}

		genType := int(row.GenType)
		if genType == consts.GenCodesTypeTree {
			genType = consts.GenCodesTypeCurd
		}

		for _, v := range typeSelect {
			if v.Value == genType {
				for index, template := range v.Templates {
					if index == row.GenTemplate {
						return template.Label
					}
				}
			}
		}
		return ""
	}

	if len(list) > 0 {
		for _, v := range list {
			v.GenTemplateGroup = getTpGroup(v)
			view := &sysin.GenCodesViewModel{SysGenCodes: v.SysGenCodes}
			if err = s.inspectBuild(ctx, view); err != nil {
				return
			}
			v.SysGenCodes, v.GenCodesBuildInfo = view.SysGenCodes, view.GenCodesBuildInfo
		}
	}
	return
}

// Selects 选项
func (s *sSysGenCodes) Selects(ctx context.Context, in *sysin.GenCodesSelectsInp) (res *sysin.GenCodesSelectsModel, err error) {
	return hggen.TableSelects(ctx, in)
}

// TableSelect 表选项
func (s *sSysGenCodes) TableSelect(ctx context.Context, in *sysin.GenCodesTableSelectInp) (res []*sysin.GenCodesTableSelectModel, err error) {
	var (
		sql           string
		config        = g.DB(in.Name).GetConfig()
		disableTables = g.Cfg().MustGet(ctx, "hggen.disableTables").Strings()
		lists         []*sysin.GenCodesTableSelectModel
	)

	// 根据数据库类型使用不同的SQL
	if config.Type == consts.DBPgsql {
		// PostgreSQL: 使用pg_catalog查询表和注释
		sql = `
			SELECT 
				c.relname as value,
				COALESCE(obj_description(c.oid), '') as label
			FROM pg_class c
			JOIN pg_namespace n ON c.relnamespace = n.oid
			WHERE n.nspname = 'public' 
				AND c.relkind = 'r'
			ORDER BY c.relname`
	} else {
		// MySQL: 使用information_schema.TABLES
		sql = fmt.Sprintf("SELECT TABLE_NAME as value, TABLE_COMMENT as label FROM information_schema.TABLES WHERE TABLE_SCHEMA = '%s'", config.Name)
	}

	if err = g.DB(in.Name).Ctx(ctx).Raw(sql).Scan(&lists); err != nil {
		return
	}

	patternStr := `addon_(\w+)_`
	repStr := ``

	for _, v := range lists {
		if gstr.InArray(disableTables, v.Value) {
			continue
		}

		newValue := v.Value
		if config.Prefix != "" {
			newValue = gstr.SubStrFromEx(v.Value, config.Prefix)
		}
		if newValue == "" {
			err = gerror.Newf("表名[%v]前缀必须和配置中的前缀设置[%v] 保持一致", v.Value, config.Prefix)
			return
		}

		// 如果是插件模块，则移除掉插件表前缀
		bt, err := gregex.Replace(patternStr, []byte(repStr), []byte(newValue))
		if err != nil {
			err = gerror.Newf("表名[%v] gregex.Replace err:%v", v.Value, err.Error())
			return nil, err
		}

		row := v
		row.DefTableComment = v.Label
		row.DaoName = gstr.CaseCamel(newValue)
		row.DefVarName = gstr.CaseCamel(string(bt))
		row.DefAlias = gstr.CaseCamelLower(newValue)
		row.Name = fmt.Sprintf("%s (%s)", v.Value, v.Label)
		row.Label = row.Name
		res = append(res, row)
	}
	return
}

// ColumnSelect 表字段选项
func (s *sSysGenCodes) ColumnSelect(ctx context.Context, in *sysin.GenCodesColumnSelectInp) (res []*sysin.GenCodesColumnSelectModel, err error) {
	var (
		sql    string
		config = g.DB(in.Name).GetConfig()
	)

	// 根据数据库类型使用不同的SQL
	if config.Type == consts.DBPgsql {
		// PostgreSQL: 使用pg_catalog查询列注释
		sql = `
			SELECT 
				a.attname as value,
				COALESCE(col_description(a.attrelid, a.attnum), '') as label
			FROM pg_attribute a
			JOIN pg_class c ON a.attrelid = c.oid
			JOIN pg_namespace n ON c.relnamespace = n.oid
			WHERE n.nspname = 'public' 
				AND c.relname = '%s'
				AND a.attnum > 0 
				AND NOT a.attisdropped
			ORDER BY a.attnum`
		sql = fmt.Sprintf(sql, in.Table)
	} else {
		// MySQL: 使用information_schema.COLUMNS
		sql = fmt.Sprintf("SELECT COLUMN_NAME as value, COLUMN_COMMENT as label FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'", config.Name, in.Table)
	}

	if err = g.DB(in.Name).Ctx(ctx).Raw(sql).Scan(&res); err != nil {
		return
	}

	if len(res) == 0 {
		err = gerror.Newf("table does not exist:%v", in.Table)
		return
	}

	for k, v := range res {
		res[k].Name = fmt.Sprintf("%s (%s)", v.Value, v.Label)
		res[k].Label = res[k].Name
	}
	return
}

// ColumnList 表字段列表
func (s *sSysGenCodes) ColumnList(ctx context.Context, in *sysin.GenCodesColumnListInp) (res []*sysin.GenCodesColumnListModel, err error) {
	return hggen.TableColumns(ctx, in)
}

// Preview 生成预览
func (s *sSysGenCodes) Preview(ctx context.Context, in *sysin.GenCodesPreviewInp) (res *sysin.GenCodesPreviewModel, err error) {
	if err = bindAddonMenuParent(ctx, &in.SysGenCodes); err != nil {
		return nil, err
	}
	return hggen.Preview(ctx, in)
}

// Build 提交生成
func (s *sSysGenCodes) Build(ctx context.Context, in *sysin.GenCodesBuildInp) (err error) {
	return s.build(ctx, in, s.executeBuild)
}
