package sys

import (
	"hotgo/internal/model/entity"
	"testing"
)

func TestAppAuthProductOptionUsesName(t *testing.T) {
	for _, tt := range []struct {
		name    string
		product entity.AppProducts
		label   string
	}{
		{name: "优先使用name而不是title", product: entity.AppProducts{Id: 7, Name: "product-name", Title: "产品标题"}, label: "product-name"},
		{name: "空名称保留产品ID", product: entity.AppProducts{Id: 8, Title: "产品标题"}, label: "产品 #8"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			option := appAuthProductOption(&tt.product)
			if option.Label != tt.label || option.Value != tt.product.Id {
				t.Fatalf("产品选项 = %+v, want label %q and value %d", option, tt.label, tt.product.Id)
			}
		})
	}
}

func TestAppAuthUserOptionOnlyContainsIdentity(t *testing.T) {
	for _, tt := range []struct {
		name     string
		username string
		nickname string
		label    string
	}{
		{name: "用户名和昵称", username: "demo", nickname: "演示用户", label: "demo（演示用户） #12"},
		{name: "仅用户名", username: "demo", label: "demo #12"},
		{name: "相同昵称不重复", username: "demo", nickname: "demo", label: "demo #12"},
		{name: "仅昵称", nickname: "演示用户", label: "演示用户 #12"},
		{name: "空名称保留ID", label: "用户 #12"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			user := &entity.User{Id: 12, Username: tt.username, Nickname: tt.nickname, Password: "secret", Salt: "private", Mobile: "13800138000"}
			option := appAuthUserOption(user)
			if option.Label != tt.label || option.Value != user.Id || option.ListClass != "primary" {
				t.Fatalf("用户选项 = %+v, want label %q, numeric ID and primary tag", option, tt.label)
			}
			if option.Extra != nil {
				t.Fatal("用户选项不应附带会员敏感资料")
			}
		})
	}
}
