package site

import (
	"encoding/json"
	"reflect"
	"testing"

	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sitein"
)

func newNavMenuRow(id, pid int64, name string) *entity.WebNavMenu {
	return &entity.WebNavMenu{Id: id, Pid: pid, Name: name, Url: "/" + name, Status: 1}
}

func marshalNavMenu(t *testing.T, v interface{}) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("json marshal failed: %v", err)
	}
	return string(b)
}

// collectNavMenuIds 按前序遍历收集节点ID，用于断言顺序与可见范围。
func collectNavMenuIds(items []*sitein.NavMenuItem) []int64 {
	ids := make([]int64, 0)
	for _, item := range items {
		ids = append(ids, item.Id)
		ids = append(ids, collectNavMenuIds(item.Children)...)
	}
	return ids
}

func TestBuildNavMenuTreeEmptyMarshalsAsArray(t *testing.T) {
	cases := map[string][]*entity.WebNavMenu{
		"nil":   nil,
		"empty": {},
	}
	for name, rows := range cases {
		tree := buildNavMenuTree(rows)
		if tree == nil {
			t.Fatalf("%s: tree should be an empty slice, got nil", name)
		}
		if got := marshalNavMenu(t, tree); got != "[]" {
			t.Fatalf("%s: want [], got %s", name, got)
		}
	}

	got := marshalNavMenu(t, sitein.IndexModel{SiteConfig: map[string]interface{}{}, NavMenu: buildNavMenuTree(nil)})
	if want := `{"siteConfig":{},"navMenu":[]}`; got != want {
		t.Fatalf("index model json mismatch\nwant %s\ngot  %s", want, got)
	}
}

func TestBuildNavMenuTreeFourLevels(t *testing.T) {
	root := newNavMenuRow(1, 0, "products")
	root.Image = "//cdn.example.com/products.jpg"
	root.ImageAlt = "products image"
	root.Level, root.Sort, root.Tree = 1, 1, ""
	rows := []*entity.WebNavMenu{
		root,
		newNavMenuRow(2, 1, "control"),
		newNavMenuRow(3, 2, "power"),
		newNavMenuRow(4, 3, "ups"),
		newNavMenuRow(5, 0, "news"),
	}

	want := `[` +
		`{"id":1,"name":"products","url":"/products","image":"//cdn.example.com/products.jpg","imageAlt":"products image","children":[` +
		`{"id":2,"name":"control","url":"/control","image":"","imageAlt":"","children":[` +
		`{"id":3,"name":"power","url":"/power","image":"","imageAlt":"","children":[` +
		`{"id":4,"name":"ups","url":"/ups","image":"","imageAlt":"","children":[]}` +
		`]}]}]},` +
		`{"id":5,"name":"news","url":"/news","image":"","imageAlt":"","children":[]}` +
		`]`
	if got := marshalNavMenu(t, buildNavMenuTree(rows)); got != want {
		t.Fatalf("four level tree json mismatch\nwant %s\ngot  %s", want, got)
	}
}

func TestBuildNavMenuTreeKeepsSiblingInputOrder(t *testing.T) {
	// 查询已按 sort asc, id asc 排好序；建树必须保留输入顺序，且子节点可先于父节点出现。
	rows := []*entity.WebNavMenu{
		newNavMenuRow(10, 0, "a"),
		newNavMenuRow(22, 3, "b-child-2"),
		newNavMenuRow(3, 0, "b"),
		newNavMenuRow(21, 3, "b-child-1"),
		newNavMenuRow(7, 0, "c"),
		newNavMenuRow(12, 10, "a-child-2"),
		newNavMenuRow(11, 10, "a-child-1"),
	}

	tree := buildNavMenuTree(rows)
	if got, want := collectNavMenuIds(tree), []int64{10, 12, 11, 3, 22, 21, 7}; !reflect.DeepEqual(got, want) {
		t.Fatalf("sibling order mismatch: want %v, got %v", want, got)
	}
}

func TestBuildNavMenuTreeHidesUnreachableNodes(t *testing.T) {
	rows := []*entity.WebNavMenu{
		newNavMenuRow(1, 0, "root"),
		newNavMenuRow(2, 1, "visible-child"),
		// 孤儿：父节点不存在
		newNavMenuRow(3, 99, "orphan"),
		newNavMenuRow(4, 3, "orphan-child"),
		// 禁用父节点 5 已被查询条件过滤，其子树不能提升为顶级
		newNavMenuRow(6, 5, "disabled-parent-child"),
		newNavMenuRow(7, 6, "disabled-parent-grandchild"),
		// 闭环与自引用
		newNavMenuRow(8, 9, "cycle-a"),
		newNavMenuRow(9, 8, "cycle-b"),
		newNavMenuRow(10, 10, "self-parent"),
	}

	tree := buildNavMenuTree(rows)
	if got, want := collectNavMenuIds(tree), []int64{1, 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("visible nodes mismatch: want %v, got %v", want, got)
	}
	want := `[{"id":1,"name":"root","url":"/root","image":"","imageAlt":"","children":[` +
		`{"id":2,"name":"visible-child","url":"/visible-child","image":"","imageAlt":"","children":[]}]}]`
	if got := marshalNavMenu(t, tree); got != want {
		t.Fatalf("unreachable nodes leaked\nwant %s\ngot  %s", want, got)
	}
}
