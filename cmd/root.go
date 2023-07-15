package cmd

import (
	"context"
	"fmt"
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/consts"
	frplog "github.com/fatedier/frp/pkg/util/log"
	"github.com/routehead/client/pkg/confs"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "routehead",
	Short: "Routehead is a service that allows private servers go public",
	Long: "A simple-to-use proxy server that can forward your local servers" +
		"to the internet. Go to https://routehead.com to know more.",
	Run: func(cmd *cobra.Command, args []string) {
		frplog.InitLog("file", "frp.logs", "info", 0, false)
		log.Println("Starting connection to server")
		cfg := config.GetDefaultClientConf()
		confs.CommonConf(&cfg)
		pxyConfigs := make(map[string]config.ProxyConf)
		confs.CreateHTTPConf(pxyConfigs)

		svr, err := client.NewService(cfg, pxyConfigs, nil, "")
		if err != nil {
			return
		}

		if pxyConfigs["test.http"].GetBaseConfig().ProxyType == consts.HTTPProxy {
			p := pxyConfigs["test.http"].(*config.HTTPProxyConf)
			log.Printf("Connected to the server. Visit http://%s.app.routehead.com to know more\n", p.SubDomain)
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
