package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command {
	Use: "git-tools",
}

func init() {
	curDir, _ := os.Getwd()
	genTag.PersistentFlags().String("dir", curDir, "--dir")
	RootCmd.AddCommand(genTag)
	RootCmd.AddCommand(packageJson)
	RootCmd.AddCommand(pr)
}
