package cmd

import (
	"fmt"
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var torrent = &cobra.Command{
	Use:   "torrent",
	Short: "torrent tools",
	Long:  `tools for torrent, ie.parseinfo、findpeers...`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		if misc.IsDir(path) {
			dirs, err := ioutil.ReadDir(path)
			misc.PanicSysErrNonNil(err, "open torrent dir err")
			for _, item := range dirs {
				if !item.IsDir() && strings.HasSuffix(item.Name(), ".torrent") {
					printTorrentInfo(item.Name(), filepath.Join(path, item.Name()))
				}
			}
			return
		}
		printTorrentInfo(path, path)
	},
}

func printTorrentInfo(filename, path string) {
	file, err := os.Open(path)
	misc.PanicSysErrNonNil(err, "open torrent file err", path)
	bytes, err := ioutil.ReadAll(file)
	misc.PanicSysErrNonNil(err, "read torrent file err", path)
	dict, err := misc.DecodeDict(misc.Bytes2Str(bytes))
	misc.PanicSysErrNonNil(err, "DecodeDict torrent file err", path)
	fmt.Printf("%s||%s||%d||%d\r\n", dict.GetStringOrDefault("name", "name err"), filename,
		dict.GetIntegerOrDefault("length", 0), dict.GetIntegerOrDefault("piece length", 0))
}
