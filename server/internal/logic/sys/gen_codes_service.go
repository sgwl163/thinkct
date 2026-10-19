package sys

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"golang.org/x/tools/imports"

	"hotgo/internal/library/hggen"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
)

var genCodesServiceMu sync.Mutex

// generateBuildServices 先在隔离目录生成和格式化，再发布，避免原地截断被 Windows 映射的文件。
func (s *sSysGenCodes) generateBuildServices(ctx context.Context, record *entity.SysGenCodes, template model.GenerateAppCrudTemplate, clear bool) error {
	genCodesServiceMu.Lock()
	defer genCodesServiceMu.Unlock()

	cfg := hggen.GetServiceConfig()
	if template.IsAddon {
		cfg.SrcFolder = filepath.Join("addons", record.AddonName, "logic")
		cfg.DstFolder = filepath.Join("addons", record.AddonName, "service")
	}
	source := gfile.Abs(cfg.SrcFolder)
	destination := gfile.Abs(cfg.DstFolder)
	if cfg.ImportPrefix == "" {
		content, err := os.ReadFile("go.mod")
		if err != nil {
			return err
		}
		for _, line := range strings.Split(string(content), "\n") {
			fields := strings.Fields(line)
			if len(fields) == 2 && fields[0] == "module" {
				relative, err := filepath.Rel(gfile.Pwd(), source)
				if err != nil {
					return err
				}
				cfg.ImportPrefix = fields[1] + "/" + filepath.ToSlash(relative)
				break
			}
		}
		if cfg.ImportPrefix == "" {
			return gerror.New("读取项目模块名称失败")
		}
	}

	stage, err := os.MkdirTemp("", "hotgo-service-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(stage)
	// 保留 logic/service 目录名称；生成器据此计算 Go 包名。
	cfg.SrcFolder = filepath.Join(stage, "source", filepath.Base(source))
	cfg.DstFolder = filepath.Join(stage, "output", filepath.Base(destination))
	cfg.Clear = clear
	cfg.WatchFile = ""
	if err = gfile.CopyDir(source, cfg.SrcFolder); err != nil {
		return err
	}
	if gfile.Exists(destination) {
		if err = gfile.CopyDir(destination, cfg.DstFolder); err != nil {
			return err
		}
	} else if err = os.MkdirAll(cfg.DstFolder, 0755); err != nil {
		return err
	}
	if err = hggen.ServiceWithCfg(ctx, cfg); err != nil {
		return err
	}

	entries, err := os.ReadDir(cfg.DstFolder)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".go" {
			continue
		}
		path := filepath.Join(cfg.DstFolder, entry.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if !isGeneratedService(content) {
			continue
		}
		// 基础生成器仅记录格式化错误；这里显式检查，禁止发布未格式化的半成品。
		formatted, err := imports.Process(filepath.Join(destination, entry.Name()), content, nil)
		if err != nil {
			return gerror.Wrapf(err, "格式化服务接口失败:%s", entry.Name())
		}
		if err = publishGeneratedFile(filepath.Join(destination, entry.Name()), formatted); err != nil {
			return err
		}
	}
	logicName := filepath.Base(source) + ".go"
	logicContent, err := os.ReadFile(filepath.Join(cfg.SrcFolder, logicName))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil && isGeneratedService(logicContent) {
		if err = publishGeneratedFile(filepath.Join(source, logicName), logicContent); err != nil {
			return err
		}
	}
	if clear && gfile.Exists(destination) {
		oldEntries, err := os.ReadDir(destination)
		if err != nil {
			return err
		}
		for _, entry := range oldEntries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != ".go" || gfile.Exists(filepath.Join(cfg.DstFolder, entry.Name())) {
				continue
			}
			path := filepath.Join(destination, entry.Name())
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if isGeneratedService(content) {
				if err = os.Remove(path); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func isGeneratedService(content []byte) bool {
	return bytes.Contains(content, []byte("Code generated")) && bytes.Contains(content, []byte("DO NOT EDIT."))
}

func publishGeneratedFile(path string, content []byte) error {
	old, err := os.ReadFile(path)
	if err == nil && bytes.Equal(old, content) {
		return nil
	}
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	mode := os.FileMode(0644)
	if info, statErr := os.Stat(path); statErr == nil {
		mode = info.Mode().Perm()
	}
	f, err := os.CreateTemp(filepath.Dir(path), ".hotgo-publish-*")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if err = f.Chmod(mode); err != nil {
		_ = f.Close()
		return err
	}
	if _, err = f.Write(content); err != nil {
		_ = f.Close()
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	// 同目录替换，不对旧文件执行 O_TRUNC；等待编译器释放短暂的映射占用。
	deadline := time.Now().Add(5 * time.Second)
	for {
		if err = os.Rename(f.Name(), path); err == nil {
			return nil
		}
		if runtime.GOOS != "windows" ||
			(!errors.Is(err, syscall.Errno(5)) && !errors.Is(err, syscall.Errno(32)) && !errors.Is(err, syscall.Errno(1224))) ||
			time.Now().After(deadline) {
			return gerror.Wrapf(err, "发布生成文件失败，原文件已保留，请在编译完成后重试:%s", path)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
