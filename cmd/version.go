package cmd

import (
	"fmt"
	"github.com/GalaIO/P2Pcrawler/config"
	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions. This is p2pcrlawer's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("p2pcrlawer version %s\r\n", config.VersionNum)
	},
}
