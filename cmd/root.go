package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command {
	Use: "semver",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	curDir, _ := os.Getwd()
	genTag.PersistentFlags().String("dir", curDir, "--dir")
	RootCmd.AddCommand(genTag)
}
