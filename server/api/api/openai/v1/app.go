package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/openaiin"
)

type AppListReq struct {
	g.Meta `path:"/openaiApp/list" method:"get" tags:"OpenAI应用" summary:"获取应用列表"`
	openaiin.AppListInp
}

type AppListRes struct {
	form.PageRes
	List []*openaiin.AppListModel `json:"list" dc:"数据列表"`
}

type AppCategoryListReq struct {
	g.Meta `path:"/openaiAppCategory/list" method:"get" tags:"OpenAI应用分类" summary:"获取应用分类列表"`
	openaiin.OpenaiAppCategoryListInp
}

type AppCategoryListRes struct {
	form.PageRes
	List []*openaiin.OpenaiAppCategoryListModel `json:"list" dc:"数据列表"`
}
