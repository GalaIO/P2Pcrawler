package main

import (
	"github.com/GalaIO/P2Pcrawler/cmd"
	"github.com/GalaIO/P2Pcrawler/config"
	"github.com/GalaIO/P2Pcrawler/misc"
)

func main() {
	misc.SetLevel(misc.TransLogLevel(config.LoggerConfig().Level))
	cmd.ExecuteBackgroud()
}
