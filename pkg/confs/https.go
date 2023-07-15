package confs

import (
	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/consts"
)

func CreateHTTPSConf(pxyConfigs map[string]config.ProxyConf) {
	httpConf := &config.HTTPSProxyConf{}
	httpConf.ProxyName = "test.http"
	httpConf.ProxyType = consts.HTTPProxy
	httpConf.LocalIP = "127.0.0.1"
	httpConf.LocalPort = 8080
	httpConf.SubDomain = "test"
	pxyConfigs[httpConf.ProxyName] = httpConf
}
