package main

import (
	"github.com/GalaIO/P2Pcrawler/dht"
	"github.com/GalaIO/P2Pcrawler/misc"
)

func main() {
	misc.SetLevel(misc.INFO)
	dht.BootStrap("87.98.162.88:6881")
	dht.Run()
}
