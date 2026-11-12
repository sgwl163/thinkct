package api

import (
	"context"
	"hotgo/addons/appAuth/api/api/auth"
	"hotgo/addons/appAuth/service"
)

var Auth = cAuth{}

type cAuth struct{}

func (c *cAuth) Query(ctx context.Context, req *auth.QueryReq) (res *auth.QueryRes, err error) {
	data, err := service.SysAppAuth().Query(ctx, &req.AuthQueryInp)
	if err != nil {
		return
	}
	res = new(auth.QueryRes)
	res.AuthQueryModel = data
	return
}

func (c *cAuth) Config(ctx context.Context, req *auth.ConfigReq) (res *auth.ConfigRes, err error) {
	data, err := service.SysConfig().GetPublic(ctx, &req.AuthConfigInp)
	if err != nil {
		return
	}
	res = new(auth.ConfigRes)
	res.AuthConfigModel = data
	return
}
