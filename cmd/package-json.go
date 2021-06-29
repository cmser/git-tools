package cmd

import "github.com/spf13/cobra"

var packageJson = &cobra.Command{
	Use: "package-json",
	Aliases: []string{ "p" },
	Run: func(cmd *cobra.Command, args []string) {

	},
}
