package sys

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"

	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/addons"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/service"
)

// addonWithMenus 将默认菜单初始化接入原有安装事务，不改变插件接口和开发者钩子。
type addonWithMenus struct {
	addons.Module
	ctx context.Context
}

func (m *addonWithMenus) Ctx() context.Context { return m.ctx }

func (m *addonWithMenus) Install(ctx context.Context) error {
	if err := m.Module.Install(ctx); err != nil {
		return err
	}
	return syncAddonMenus(ctx, m.GetSkeleton())
}

func (m *addonWithMenus) Upgrade(ctx context.Context) error {
	if err := m.Module.Upgrade(ctx); err != nil {
		return err
	}
	return syncAddonMenus(ctx, m.GetSkeleton())
}

// registeredAddon 只使用服务端注册信息，避免信任列表请求携带的标签、路径等字段。
func registeredAddon(name string) (addons.Module, error) {
	for _, skeleton := range addons.GetSkeletons() {
		if skeleton.Name == name {
			return skeleton.GetModule(), nil
		}
	}
	return nil, gerror.New("插件不存在，请确认插件已注册并重启服务")
}

func addonMenuName(name string) string { return consts.AddonsTag + name }

// findAddonMenuRoot 优先使用固定编码，其次复用只属于该插件的人工一级目录。
// 不移动既有菜单，也不把其他插件或主模块的目录误认为本插件目录。
func findAddonMenuRoot(menus []*entity.AdminMenu, name string) (*entity.AdminMenu, error) {
	byID := make(map[int64]*entity.AdminMenu, len(menus))
	for _, menu := range menus {
		byID[menu.Id] = menu
		if menu.Name == addonMenuName(name) {
			if menu.Pid != 0 || menu.Type != 1 || menu.Component != "LAYOUT" {
				return nil, gerror.New("插件一级菜单编码已被其他菜单占用，请先在菜单管理中核对")
			}
		}
	}
	for _, menu := range menus {
		if menu.Name == addonMenuName(name) {
			return menu, nil
		}
	}
	top := func(menu *entity.AdminMenu) *entity.AdminMenu {
		seen := map[int64]bool{}
		for menu != nil && menu.Pid > 0 {
			if seen[menu.Id] {
				return nil
			}
			seen[menu.Id] = true
			menu = byID[menu.Pid]
		}
		return menu
	}
	prefix := "/addons/" + name + "/"
	// 旧生成器也会为单个业务模块建立 LAYOUT 一级目录，不能将它当作插件总目录。
	moduleRoots := make(map[int64]bool)
	for _, menu := range menus {
		parent := byID[menu.Pid]
		if parent != nil && menu.Type == 2 && menu.Name == parent.Name+"Index" && menu.Component == prefix+parent.Name+"/index" {
			moduleRoots[parent.Id] = true
		}
	}
	var root *entity.AdminMenu
	for _, menu := range menus {
		if !strings.HasPrefix(menu.Component, prefix) {
			continue
		}
		candidate := top(menu)
		if candidate == nil || candidate.Type != 1 || candidate.Component != "LAYOUT" || moduleRoots[candidate.Id] {
			continue
		}
		if root != nil && root.Id != candidate.Id {
			return nil, nil
		}
		root = candidate
	}
	if root != nil {
		for _, menu := range menus {
			if menu.Type != 2 {
				continue
			}
			ancestor := top(menu)
			if ancestor != nil && ancestor.Id == root.Id && !strings.HasPrefix(menu.Component, prefix) {
				return nil, nil
			}
		}
	}
	return root, nil
}

type addonMenuPage struct {
	module, component string
	config            bool
}

// addonMenuPages 只识别脚手架配置页和标准模块入口，不把弹窗和子组件当成菜单。
// 源码未随二进制部署时返回空页面列表，仍可复用数据库中已有菜单。
func addonMenuPages(ctx context.Context, name string) ([]addonMenuPage, error) {
	config, err := service.SysConfig().GetLoadGenerate(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "读取插件页面生成配置失败")
	}
	if config == nil || config.Addon == nil || config.Addon.WebViewsPath == "" {
		return nil, nil
	}
	base := strings.ReplaceAll(config.Addon.WebViewsPath, "{$name}", name)
	base, err = filepath.Abs(base)
	if err != nil {
		return nil, gerror.Wrap(err, "解析插件页面目录失败")
	}
	info, err := os.Lstat(base)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, gerror.Wrap(err, "读取插件页面目录失败")
	}
	if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return nil, gerror.New("插件页面目录必须为普通目录")
	}
	normalized := filepath.ToSlash(base)
	marker := "/web/src/views/"
	index := strings.LastIndex(normalized, marker)
	if index < 0 {
		return nil, gerror.New("插件页面目录必须位于 web/src/views 下")
	}
	componentBase := "/" + strings.Trim(normalized[index+len(marker):], "/")
	entries, err := os.ReadDir(base)
	if err != nil {
		return nil, gerror.Wrap(err, "读取插件页面文件失败")
	}
	var pages []addonMenuPage
	regularFile := func(path string) (bool, error) {
		info, err := os.Lstat(path)
		if os.IsNotExist(err) {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		return info.Mode().IsRegular(), nil
	}
	for _, entry := range entries {
		if entry.Type()&os.ModeSymlink != 0 || !entry.IsDir() {
			continue
		}
		module := entry.Name()
		if module == "components" || module == "component" || strings.HasPrefix(module, ".") {
			continue
		}
		filename := "index.vue"
		isConfig := module == "config"
		if isConfig {
			filename = "system.vue"
		}
		ok, err := regularFile(filepath.Join(base, module, filename))
		if err != nil {
			return nil, gerror.Wrap(err, "检查插件页面入口失败")
		}
		if !ok {
			continue
		}
		pages = append(pages, addonMenuPage{module: module, component: componentBase + "/" + module + "/" + strings.TrimSuffix(filename, ".vue"), config: isConfig})
	}
	return pages, nil
}

func syncAddonMenus(ctx context.Context, skeleton *addons.Skeleton) error {
	pages, err := addonMenuPages(ctx, skeleton.Name)
	if err != nil {
		return err
	}
	var menus []*entity.AdminMenu
	if err = dao.AdminMenu.Ctx(ctx).Scan(&menus); err != nil {
		return gerror.Wrap(err, "查询插件菜单失败")
	}
	root, err := findAddonMenuRoot(menus, skeleton.Name)
	if err != nil {
		return err
	}
	if root == nil {
		root = &entity.AdminMenu{Name: addonMenuName(skeleton.Name), Title: skeleton.Label, Path: "/addons/" + skeleton.Name, Type: 1, Component: "LAYOUT", AlwaysShow: 1, Sort: 99, Status: consts.StatusEnabled, Icon: consts.AddonsGroupIconMap[skeleton.Group]}
		if root.Title == "" {
			root.Title = skeleton.Name
		}
		if err = insertAddonMenu(ctx, root, menus); err != nil {
			return err
		}
		menus = append(menus, root)
	}
	var records []*entity.SysGenCodes
	if len(pages) > 0 {
		if err = dao.SysGenCodes.Ctx(ctx).Where(dao.SysGenCodes.Columns().AddonName, skeleton.Name).Scan(&records); err != nil {
			return gerror.Wrap(err, "读取插件代码生成配置失败")
		}
	}
	for _, page := range pages {
		var menu *entity.AdminMenu
		for _, existing := range menus {
			if existing.Type == 2 && existing.Component == page.component {
				menu = existing
				break
			}
		}
		if menu == nil {
			menu = &entity.AdminMenu{Pid: root.Id, Name: addonMenuName(skeleton.Name) + "_" + page.module, Title: skeleton.Label + "-" + page.module, Path: page.module, Type: 2, Component: page.component, Sort: 99, Status: consts.StatusEnabled}
			if page.config {
				menu.Title = skeleton.Label + "配置"
				menu.Permissions = "/" + skeleton.Name + "/config/get"
			} else {
				for _, record := range records {
					if gstr.LcFirst(record.VarName) == page.module && (record.GenType == consts.GenCodesTypeCurd || record.GenType == consts.GenCodesTypeTree) {
						// 与生成器使用相同的根菜单编码，后续删除生成代码能清理对应页面。
						menu.Name = gstr.LcFirst(record.VarName)
						menu.Title = record.TableComment
						menu.Permissions = "/" + skeleton.Name + "/" + page.module + "/list"
						break
					}
				}
			}
			if err = insertAddonMenu(ctx, menu, menus); err != nil {
				return err
			}
			menus = append(menus, menu)
		}
		if page.config {
			permission := "/" + skeleton.Name + "/config/update"
			exists := false
			for _, existing := range menus {
				if existing.Pid == menu.Id && gstr.InArray(strings.Split(existing.Permissions, ","), permission) {
					exists = true
					break
				}
			}
			if !exists {
				button := &entity.AdminMenu{Pid: menu.Id, Name: addonMenuName(skeleton.Name) + "_config_update", Title: skeleton.Label + "配置保存", Type: 3, Permissions: permission, Sort: 10, Status: consts.StatusEnabled}
				if err = insertAddonMenu(ctx, button, menus); err != nil {
					return err
				}
				menus = append(menus, button)
			}
		}
	}
	return nil
}

// insertAddonMenu 复用菜单输入校验和关系树维护；只插入缺失项，不覆盖人工菜单。
func insertAddonMenu(ctx context.Context, menu *entity.AdminMenu, existing []*entity.AdminMenu) error {
	in := &adminin.MenuEditInp{AdminMenu: *menu}
	if err := in.Filter(ctx); err != nil {
		return err
	}
	for _, other := range existing {
		if other.Name == menu.Name || other.Title == menu.Title || (menu.Type != 3 && other.Pid == menu.Pid && other.Path == menu.Path) {
			return gerror.Newf("插件菜单存在名称、编码或路由冲突，请在菜单管理中核对：%s", menu.Title)
		}
	}
	var err error
	menu.Pid, menu.Level, menu.Tree, err = hgorm.AutoUpdateTree(ctx, &dao.AdminMenu, 0, menu.Pid)
	if err != nil {
		return gerror.Wrap(err, "生成插件菜单关系树失败")
	}
	menu.Id, err = dao.AdminMenu.Ctx(ctx).Data(menu).OmitEmptyData().InsertAndGetId()
	if err != nil {
		return gerror.Wrap(err, "创建插件菜单失败")
	}
	return nil
}

// bindAddonMenuParent 同时用于保存、预览和构建；保持这三个入口的父菜单一致。
// 只补全默认根目录，不覆盖已明确选择的上级，也不自动打开被取消的菜单生成选项。
func bindAddonMenuParent(ctx context.Context, record *entity.SysGenCodes) error {
	if record.GenType != consts.GenCodesTypeCurd && record.GenType != consts.GenCodesTypeTree {
		return nil
	}
	if record.Options == nil || !gstr.InArray(record.Options.Get("autoOps").Strings(), "genMenuPermissions") || record.Options.Get("menu.pid").Int64() > 0 {
		return nil
	}
	config, err := service.SysConfig().GetLoadGenerate(ctx)
	if err != nil {
		return err
	}
	if config == nil || record.GenTemplate < 0 || record.GenTemplate >= len(config.Application.Crud.Templates) || config.Application.Crud.Templates[record.GenTemplate] == nil {
		return gerror.New("生成模板不存在")
	}
	if !config.Application.Crud.Templates[record.GenTemplate].IsAddon {
		return nil
	}
	if record.AddonName == "" {
		return gerror.New("插件模板必须选择一个有效的插件")
	}
	var menus []*entity.AdminMenu
	if err = dao.AdminMenu.Ctx(ctx).Scan(&menus); err != nil {
		return gerror.Wrap(err, "查询插件上级菜单失败")
	}
	root, err := findAddonMenuRoot(menus, record.AddonName)
	if err != nil {
		return err
	}
	if root == nil {
		return gerror.New("插件一级菜单不存在，请先安装插件；已安装的插件可点击更新配置补齐菜单")
	}
	return record.Options.Set("menu.pid", root.Id)
}
