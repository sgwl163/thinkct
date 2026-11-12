package model

import (
	"net"
	"strings"

	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/net/idna"
	"golang.org/x/net/publicsuffix"
)

// NormalizeAuthHost 规范化用于授权登记查询的裸域名，不解析URL、不进行DNS请求。
func NormalizeAuthHost(host string) (string, error) {
	host = strings.TrimSuffix(strings.ToLower(strings.TrimSpace(host)), ".")
	ascii, err := idna.Lookup.ToASCII(host)
	if err != nil || len(ascii) > 253 || !strings.Contains(ascii, ".") ||
		net.ParseIP(ascii) != nil || !validate.IsDNSName(ascii) {
		return "", gerror.New("请输入正确的域名，不要包含协议、端口或路径")
	}
	for _, label := range strings.Split(ascii, ".") {
		if len(label) == 0 || len(label) > 63 || strings.Contains(label, "_") ||
			strings.HasPrefix(label, "-") || strings.HasSuffix(label, "-") {
			return "", gerror.New("请输入正确的域名，不要包含协议、端口或路径")
		}
	}
	topLabel := ascii[strings.LastIndexByte(ascii, '.')+1:]
	if strings.Trim(topLabel, "0123456789") == "" {
		return "", gerror.New("请输入正确的域名，不支持IP地址")
	}
	if _, err = publicsuffix.EffectiveTLDPlusOne(ascii); err != nil {
		return "", gerror.New("请输入完整的域名")
	}
	return ascii, nil
}
