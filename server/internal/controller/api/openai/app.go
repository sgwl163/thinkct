package openai

import (
	"context"

	"hotgo/api/api/openai/v1"
	"hotgo/internal/model/input/openaiin"
	"hotgo/internal/service"
)

var (
	App = cApp{}
)

type cApp struct{}

func (c *cApp) List(ctx context.Context, req *v1.AppListReq) (res *v1.AppListRes, err error) {
	list, totalCount, err := service.OpenaiApp().List(ctx, &req.AppListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*openaiin.AppListModel{}
	}

	res = new(v1.AppListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *cApp) CategoryList(ctx context.Context, req *v1.AppCategoryListReq) (res *v1.AppCategoryListRes, err error) {
	list, totalCount, err := service.OpenaiAppCategory().List(ctx, &req.OpenaiAppCategoryListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*openaiin.OpenaiAppCategoryListModel{}
	}

	res = new(v1.AppCategoryListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
