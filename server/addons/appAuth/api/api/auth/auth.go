package auth

import (
	"hotgo/addons/appAuth/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// QueryReq 查询域名授权登记，公开接口不执行授权码或IP双重校验。
type QueryReq struct {
	g.Meta `path:"/auth/query" method:"get" tags:"应用授权系统" summary:"查询域名授权登记状态"`
	sysin.AuthQueryInp
}

type QueryRes struct {
	*sysin.AuthQueryModel
}

type ConfigReq struct {
	g.Meta `path:"/auth/config" method:"get" tags:"应用授权系统" summary:"获取公开产品授权配置"`
	sysin.AuthConfigInp
}

type ConfigRes struct {
	*sysin.AuthConfigModel
}
