package confs

import "github.com/fatedier/frp/pkg/config"

func CommonConf(cfg *config.ClientCommonConf) {
	cfg.NatHoleSTUNServer = ""
	cfg.ServerAddr = "app.routehead.com"
	cfg.ServerPort = 7000
	cfg.TLSEnable = true
	cfg.Protocol = "tcp"
}
