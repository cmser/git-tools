package cmd

import (
	"fmt"
	"github.com/cmser/git-tools/pkg/types"
	"github.com/spf13/cobra"
	"os"
)

var release = &cobra.Command{
	Use: "release",
}

var createRelease = &cobra.Command{
	Use: "create",
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := types.InitializeRepo(cmd.Flag("dir").Value.String(), cobra.CheckErr)
		//return repo.CreateReleaseBranch(cobra.CheckErr)
		return repo.FetchTags(cobra.CheckErr)
	},
}

func init() {
	currDir, _ := os.Getwd()
	release.PersistentFlags().String("dir", currDir, fmt.Sprintf("--dir %s", currDir))
	release.AddCommand(createRelease)
}
