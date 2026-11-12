package sys

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofrs/flock"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"

	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hggen"
	"hotgo/internal/library/hggen/views"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	dbutil "hotgo/utility/db"
)

// generationLock 使用项目独立的进程锁，防止并发生成及热重载后的进程同时写共享目录。
func generationLock(id int64) *flock.Flock {
	root := gfile.Pwd()
	key := sha256.Sum256([]byte(root))
	return flock.New(filepath.Join(os.TempDir(), fmt.Sprintf("hotgo-generate-%x-%d.lock", key[:12], id)))
}

// runBuildStage 记录关键阶段和耗时；不记录生成内容、SQL 或数据库配置。
func runBuildStage(ctx context.Context, id int64, stage string, execute func(context.Context) error) error {
	started := time.Now()
	g.Log().Info(ctx, "代码生成阶段开始", g.Map{"id": id, "stage": stage})
	succeeded := false
	defer func() {
		g.Log().Info(ctx, "代码生成阶段结束", g.Map{
			"id": id, "stage": stage, "elapsedMs": time.Since(started).Milliseconds(), "succeeded": succeeded,
		})
	}()
	err := execute(ctx)
	succeeded = err == nil
	return err
}

func (s *sSysGenCodes) build(ctx context.Context, in *sysin.GenCodesBuildInp, execute func(context.Context, *sysin.GenCodesBuildInp) error) (err error) {
	lock := generationLock(0)
	locked, err := lock.TryLock()
	if err != nil {
		return gerror.Wrap(err, "获取代码生成锁失败")
	}
	if !locked {
		return gerror.New("已有代码生成任务正在执行，请稍后重试")
	}
	defer lock.Close()

	saved, err := s.Edit(ctx, &sysin.GenCodesEditInp{SysGenCodes: in.SysGenCodes})
	if err != nil {
		return err
	}
	// 新建后必须使用真实 ID；不能继续用请求中的 0 更新状态。
	in.SysGenCodes = saved.SysGenCodes
	recordLock := generationLock(in.Id)
	locked, err = recordLock.TryLock()
	if err != nil {
		return gerror.Wrap(err, "获取生成记录锁失败")
	}
	if !locked {
		return gerror.New("该配置正在生成，请稍后重试")
	}
	defer recordLock.Close()

	// 客户端超时或离开页面不能取消最后的生成结果更新。
	buildCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Minute)
	defer cancel()
	if err = runBuildStage(buildCtx, in.Id, "保存生成中状态", func(ctx context.Context) error {
		return s.Status(ctx, &sysin.GenCodesStatusInp{Id: in.Id, Status: consts.GenCodesStatusRunning})
	}); err != nil {
		return gerror.Wrap(err, "保存生成中状态失败")
	}
	var buildErr error
	if panicErr := g.Try(buildCtx, func(ctx context.Context) {
		buildErr = runBuildStage(ctx, in.Id, "执行代码生成", func(ctx context.Context) error { return execute(ctx, in) })
	}); panicErr != nil {
		buildErr = panicErr
	}
	err = buildErr
	status := consts.GenCodesStatusOk
	if err != nil {
		status = consts.GenCodesStatusFail
	}
	statusCtx, statusCancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Second)
	defer statusCancel()
	if statusErr := runBuildStage(statusCtx, in.Id, "保存生成结果", func(ctx context.Context) error {
		return s.Status(ctx, &sysin.GenCodesStatusInp{Id: in.Id, Status: status})
	}); statusErr != nil {
		if err != nil {
			return gerror.Newf("生成未完成，且结果状态保存失败，请刷新核对：%s；%s", gerror.Current(err), gerror.Current(statusErr))
		}
		return gerror.Wrap(statusErr, "代码已生成，但结果状态保存失败，请刷新核对")
	}
	if err != nil {
		g.Log().Warningf(statusCtx, "生成代码失败 id:%d err:%+v", in.Id, err)
		return gerror.Newf("生成未完成，可能已写入部分文件：%s", gerror.Current(err))
	}
	return nil
}

// executeBuild 沿用现有预览和生成事件，MySQL 菜单使用无连接副作用的事务导入。
func (s *sSysGenCodes) executeBuild(ctx context.Context, in *sysin.GenCodesBuildInp) error {
	if in.GenType != consts.GenCodesTypeCurd && in.GenType != consts.GenCodesTypeTree {
		return hggen.Build(ctx, in)
	}
	if err := bindAddonMenuParent(ctx, &in.SysGenCodes); err != nil {
		return err
	}
	config, err := s.loadBuildConfig(ctx, in)
	if err != nil {
		return err
	}
	buildInput := &views.CurdBuildInput{
		PreviewIn: &views.CurdPreviewInput{
			In:        &sysin.GenCodesPreviewInp{SysGenCodes: in.SysGenCodes},
			DaoConfig: hggen.GetDaoConfig(in.DbName), Config: config,
		},
		BeforeEvent: views.CurdBuildEvent{"runDao": func(ctx context.Context) error {
			return runBuildStage(ctx, in.Id, "生成 DAO", hggen.Dao)
		}},
		AfterEvent: views.CurdBuildEvent{"runService": func(ctx context.Context) error {
			return runBuildStage(ctx, in.Id, "生成服务接口", func(ctx context.Context) error {
				return s.generateBuildServices(ctx, &in.SysGenCodes, *config.Application.Crud.Templates[in.GenTemplate], false)
			})
		}},
	}
	if g.DB().GetConfig().Type != consts.DBMysql {
		return views.Curd.DoBuild(ctx, buildInput)
	}
	var preview *sysin.GenCodesPreviewModel
	if err = runBuildStage(ctx, in.Id, "生成预览", func(ctx context.Context) error {
		var previewErr error
		preview, previewErr = views.Curd.DoPreview(ctx, buildInput.PreviewIn)
		return previewErr
	}); err != nil {
		return err
	}
	before := func(ctx context.Context) error {
		if gstr.InArray(in.Options.Get("autoOps").Strings(), "runDao") {
			return buildInput.BeforeEvent["runDao"](ctx)
		}
		return nil
	}
	after := func(ctx context.Context) error {
		if gstr.InArray(in.Options.Get("autoOps").Strings(), "runService") {
			return buildInput.AfterEvent["runService"](ctx)
		}
		return nil
	}
	return s.publishBuildPreview(ctx, in.Id, preview, before, after)
}

// publishBuildPreview 保持 DAO、代码、service、菜单的先后顺序，不调用会改变连接会话的旧导入器。
func (s *sSysGenCodes) publishBuildPreview(ctx context.Context, id int64, preview *sysin.GenCodesPreviewModel, before, after func(context.Context) error) error {
	if err := before(ctx); err != nil {
		return err
	}
	if err := runBuildStage(ctx, id, "发布代码文件", func(ctx context.Context) error {
		for name, file := range preview.Views {
			if name == "source.sql" || !shouldWriteBuildFile(file) {
				continue
			}
			if err := publishGeneratedFile(file.Path, []byte(strings.TrimSpace(file.Content))); err != nil {
				return gerror.Wrapf(err, "写入生成文件失败:%s", file.Path)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	if err := after(ctx); err != nil {
		return err
	}
	file := preview.Views["source.sql"]
	if !shouldWriteBuildFile(file) {
		return nil
	}
	return runBuildStage(ctx, id, "发布并导入菜单", func(ctx context.Context) error {
		if err := publishGeneratedFile(file.Path, []byte(strings.TrimSpace(file.Content))); err != nil {
			return gerror.Wrap(err, "写入生成菜单 SQL 失败")
		}
		if err := dbutil.ImportMysqlMenuSQL(ctx, g.DB(), file.Content); err != nil {
			_ = gfile.RemoveFile(file.Path)
			return gerror.Wrap(err, "导入生成菜单 SQL 失败")
		}
		return nil
	})
}

func shouldWriteBuildFile(file *sysin.GenFile) bool {
	return file != nil && (file.Meth == consts.GenCodesBuildMethCreate || file.Meth == consts.GenCodesBuildMethCover)
}

// inspectBuild 显式区分执行结果和文件存在情况，旧的异常记录也可安全清理残留文件。
func (s *sSysGenCodes) inspectBuild(ctx context.Context, row *sysin.GenCodesViewModel) error {
	row.GenCodesBuildInfo = sysin.GenCodesBuildInfo{}
	defer func() {
		row.BuildStatusLabel = consts.GenCodesStatusNameMap[row.Status]
		if row.Status == consts.GenCodesStatusRunning && !row.Building {
			row.BuildStatusLabel = "结果待确认"
		}
		if row.BuildFiles > 0 {
			if label, ok := consts.GenCodesExistingFilesStatusNameMap[row.Status]; ok {
				row.BuildStatusLabel = label
			}
		}
	}()
	lock := generationLock(row.Id)
	locked, err := lock.TryLock()
	if err != nil {
		return gerror.Wrap(err, "读取生成任务状态失败")
	}
	if !locked {
		row.Status = consts.GenCodesStatusRunning
		row.Building = true
		row.BuildMessage = "代码正在生成，请等待完成后再操作"
		return nil
	}
	defer lock.Close()
	// 首次查询可能发生在生成开始前；持锁后重读状态，避免返回过期的未生成或失败结果。
	var latest entity.SysGenCodes
	if err = dao.SysGenCodes.Ctx(ctx).Fields(dao.SysGenCodes.Columns().Status, dao.SysGenCodes.Columns().UpdatedAt).
		Where(dao.SysGenCodes.Columns().Id, row.Id).Scan(&latest); err != nil {
		return err
	}
	row.Status, row.UpdatedAt = latest.Status, latest.UpdatedAt
	if row.Status == consts.GenCodesStatusRunning {
		// 锁只能判断当前任务是否仍被观察到，不能证明上次执行失败。
		// 热重启和并发查询期间保持查询只读，终态仅由生成任务写入。
		row.BuildMessage = "当前未检测到生成任务，但上次结果尚未确认；请核对服务日志和已有文件，确认中断后可重新提交"
	}
	row.BuildSucceeded = row.Status == consts.GenCodesStatusOk
	if row.GenType != consts.GenCodesTypeCurd && row.GenType != consts.GenCodesTypeTree {
		return nil
	}
	paths, _, err := s.getBuildFiles(ctx, &row.SysGenCodes)
	if err != nil {
		row.BuildMessage = gerror.Current(err).Error()
		return nil
	}
	for _, path := range paths {
		if gfile.Exists(path) {
			row.BuildFiles++
		}
	}
	if row.BuildMessage == "" {
		switch row.Status {
		case consts.GenCodesStatusWait:
			if row.BuildFiles > 0 {
				row.BuildMessage = "检测到已有生成文件，上次完成状态未确认；可提交生成补全，默认保留已存在文件"
			}
		case consts.GenCodesStatusFail:
			row.BuildMessage = "上次生成未完整完成，可能已有部分文件；请查看生成错误并重试"
		case consts.GenCodesStatusOk:
			if row.BuildFiles < len(paths) {
				row.BuildMessage = "生成成功记录对应的部分文件已缺失，请检查或重新生成"
			}
		}
	}
	return nil
}

func (s *sSysGenCodes) loadBuildConfig(ctx context.Context, in *sysin.GenCodesBuildInp) (*model.GenerateConfig, error) {
	config, err := service.SysConfig().GetLoadGenerate(ctx)
	if err != nil {
		return nil, err
	}
	if config == nil || in.GenTemplate < 0 || in.GenTemplate >= len(config.Application.Crud.Templates) || config.Application.Crud.Templates[in.GenTemplate] == nil {
		return nil, gerror.New("生成模板不存在")
	}
	// 保持实体名规范与预览一致。
	in.VarName = gstr.UcFirst(in.VarName)
	return config, nil
}
