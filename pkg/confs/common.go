package confs

import "github.com/fatedier/frp/pkg/config"

type RouteheadServerConfig struct {
	ServerAddr        string
	ServerPort        int
	TLSEnable         bool
	Protocol          string
	NatHoleSTUNServer string
}

func CommonConf(cfg *config.ClientCommonConf, serverConfig *RouteheadServerConfig) {
	cfg.NatHoleSTUNServer = serverConfig.NatHoleSTUNServer
	cfg.ServerAddr = serverConfig.ServerAddr
	cfg.ServerPort = serverConfig.ServerPort
	cfg.TLSEnable = serverConfig.TLSEnable
	cfg.Protocol = serverConfig.Protocol
}

func SetAuthentication(cfg *config.ClientCommonConf, configFile *Config) {
	cfg.AuthenticationMethod = "token"
	cfg.AuthenticateHeartBeats = false
	cfg.AuthenticateNewWorkConns = true
	cfg.Metas["token"] = configFile.Token
}
