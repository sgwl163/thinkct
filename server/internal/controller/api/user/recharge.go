package user

import (
	"context"
	v1 "hotgo/api/api/user/v1"
	"hotgo/internal/service"
)

func (c *ControllerV1) RechargeOptions(ctx context.Context, req *v1.RechargeOptionsReq) (*v1.RechargeOptionsRes, error) {
	return service.ApiUser().RechargeOptions(ctx, req)
}

func (c *ControllerV1) RechargeCreate(ctx context.Context, req *v1.RechargeCreateReq) (*v1.RechargeCreateRes, error) {
	return service.ApiUser().RechargeCreate(ctx, req)
}

func (c *ControllerV1) RechargeView(ctx context.Context, req *v1.RechargeViewReq) (*v1.RechargeViewRes, error) {
	return service.ApiUser().RechargeView(ctx, req)
}

func (c *ControllerV1) RechargeList(ctx context.Context, req *v1.RechargeListReq) (*v1.RechargeListRes, error) {
	return service.ApiUser().RechargeList(ctx, req)
}
