package cmd

import (
	"github.com/GalaIO/P2Pcrawler/config"
	"github.com/GalaIO/P2Pcrawler/dht"
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "p2pcrawler",
	Short: "p2pcrawler is a tool to clawer data from p2p network",
	Long:  `p2pcrawler is a tool to clawer data from p2p network`,
	Run:   rootExectueFunc,
}

func init() {
	// init rootcmd
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", config.FileAbsPath, "config file path")

	// version cmd
	rootCmd.AddCommand(version)
	rootCmd.AddCommand(torrent)
}

func ExecuteBackgroud() {
	if err := rootCmd.Execute(); err != nil {
		misc.PanicSysErr("execute cobra cmd err", err)
	}
}

func rootExectueFunc(cmd *cobra.Command, args []string) {
	// 指定文件
	if len(cfgFile) > 0 {
		viper.SetConfigFile(cfgFile)
		config.ResetConfig()
	}

	if config.PProfConfig().NeedRun {
		misc.StartPProf()
	}

	for _, host := range config.DhtConfig().BootstrapNodes {
		dht.BootStrap(host)
	}
	go dht.Run()
	misc.Wait4Shutdown(5 * time.Second)
}
