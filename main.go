package main

import "github.com/GalaIO/P2Pcrawler/dht"

func main() {
	dht.BootStrap("87.98.162.88:6881")
	dht.Run()
}
