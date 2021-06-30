package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command {
	Use: "git-tools",
}

func init() {
	RootCmd.AddCommand(pr)
	RootCmd.AddCommand(release)
}
