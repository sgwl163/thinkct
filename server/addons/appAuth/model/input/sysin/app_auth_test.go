package sysin

import (
	"context"
	ac "hotgo/addons/appAuth/consts"
	"hotgo/internal/library/dict"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"reflect"
	"testing"
)

func TestAppAuthEditFilterOptions(t *testing.T) {
	for _, tt := range []struct {
		name    string
		mode    string
		domain  string
		ip      string
		status  uint
		wantErr bool
	}{
		{name: "单域名不需要IP", mode: "0", domain: "0", status: 1},
		{name: "双重验证指定IP", mode: "1", domain: "1", ip: "192.0.2.1", status: 2},
		{name: "双重验证缺少IP", mode: "1", domain: "0", status: 1, wantErr: true},
		{name: "非法验证方式", mode: "9", domain: "0", status: 1, wantErr: true},
		{name: "非法域名范围", mode: "0", domain: "9", status: 1, wantErr: true},
		{name: "非法状态", mode: "0", domain: "0", status: 0, wantErr: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			in := &AppAuthEditInp{AppAuth: entity.AppAuth{
				UserId: 1, ProductId: 1, Name: "测试授权", Qq: "1234567",
				Phone: "13800138000", VersionId: "1.0.0", AuthMode: tt.mode,
				AuthDomain: tt.domain, Ip: tt.ip, Status: tt.status,
			}}
			if err := in.Filter(context.Background()); (err != nil) != tt.wantErr {
				t.Fatalf("Filter() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestAppAuthDictionaryValuesRemainStrings(t *testing.T) {
	for _, name := range []string{"AppAuthModeOptions", "AppAuthDomainOptions"} {
		opts := dict.GetEnumsOptions(name)
		if len(opts) != 2 {
			t.Fatalf("%s 未注册完整选项", name)
		}
		for _, option := range opts {
			if _, ok := option.Value.(string); !ok {
				t.Fatalf("%s 的值必须为字符串，实际为 %T", name, option.Value)
			}
		}
	}
	for _, options := range [][]*model.Option{ac.AppAuthModeOptions, ac.AppAuthDomainOptions} {
		if options[0].ListClass == options[1].ListClass {
			t.Fatal("同组授权策略必须使用不同颜色的标签")
		}
	}
	if ac.AppAuthModeOptions[0].ListClass == ac.AppAuthDomainOptions[0].ListClass {
		t.Fatal("验证方式与域名范围应采用不同的标签色系")
	}
	option := dict.GetOption(ac.AppAuthModeOptions, ac.AppAuthModeDomainIP)
	if option == nil || option.Extra.(map[string]bool)["requireIp"] != true {
		t.Fatal("双重验证选项必须声明需要IP")
	}
}

func TestAppAuthUpdateFieldsExcludeStatistics(t *testing.T) {
	fields := reflect.TypeOf(AppAuthUpdateFields{})
	for _, name := range []string{"Views", "CreatedAt", "UpdatedAt"} {
		if _, ok := fields.FieldByName(name); ok {
			t.Errorf("编辑授权不应覆盖系统维护字段 %s", name)
		}
	}
}
