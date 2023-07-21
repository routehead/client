package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/consts"
	frplog "github.com/fatedier/frp/pkg/util/log"
	"github.com/routehead/client/pkg/confs"
	"github.com/spf13/cobra"
)

var (
	proxyType  string
	port       int
	subdomain  string
	serverAddr string
	serverPort int
)

func init() {
	rootCmd.Flags().StringVarP(&proxyType, "proxyType", "t", "http", "Define the type: http/https")
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to forward")
	rootCmd.Flags().StringVarP(&subdomain, "subdomain", "s", "test", "Subdomain at the server default domain")
	rootCmd.Flags().StringVar(&serverAddr, "serverAddr", "app.routehead.com", "The frps server address")
	rootCmd.Flags().IntVar(&serverPort, "serverPort", 7000, "The frps server port")
}

var rootCmd = &cobra.Command{
	Use:   "routehead",
	Short: "Routehead is a service that allows private servers go public",
	Long: "A simple-to-use proxy server that can forward your local servers" +
		"to the internet. Go to https://routehead.com to know more.",
	Run: func(cmd *cobra.Command, args []string) {
		frplog.InitLog("console", "", "info", 0, false)
		log.Println("Starting connection to server")

		configFile, err := confs.LoadConfigFile()
		if err != nil {
			fmt.Printf("Invalid config %v\n", err)
			return
		}

		cfg := config.GetDefaultClientConf()
		pxyConfigs := make(map[string]config.ProxyConf)

		serverCfg := confs.RouteheadServerConfig{
			ServerAddr:        serverAddr,
			ServerPort:        7000,
			TLSEnable:         true,
			Protocol:          "tcp",
			NatHoleSTUNServer: "",
		}

		confs.CommonConf(&cfg, &serverCfg)
		confs.SetAuthentication(&cfg, configFile)
		confs.CreateHTTPConf(pxyConfigs, subdomain, port)

		svr, err := client.NewService(cfg, pxyConfigs, nil, "")
		if err != nil {
			fmt.Println(err)
			return
		}

		if pxyConfigs[fmt.Sprintf("%s.http", configFile.Username)].GetBaseConfig().ProxyType == consts.HTTPProxy {
			p := pxyConfigs[fmt.Sprintf("%s.http", configFile.Username)].(*config.HTTPProxyConf)
			log.Printf("Connected to the server. Visit http://%s.via.routehead.com to know more\n", p.SubDomain)
		}

		closedDoneCh := make(chan struct{})

		err = svr.Run(context.TODO())
		if err == nil {
			<-closedDoneCh
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
