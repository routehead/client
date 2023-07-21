package confs

import (
	"fmt"

	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/consts"
)

func CreateHTTPConf(pxyConfigs map[string]config.ProxyConf, subdomain string, port int) {
	file, err := LoadConfigFile()
	if err != nil {
		return
	}

	httpConf := &config.HTTPProxyConf{}
	httpConf.ProxyName = fmt.Sprintf("%s.http", file.Username)
	httpConf.ProxyType = consts.HTTPProxy
	httpConf.LocalIP = "127.0.0.1"
	httpConf.LocalPort = port
	httpConf.SubDomain = subdomain
	pxyConfigs[httpConf.ProxyName] = httpConf
}
