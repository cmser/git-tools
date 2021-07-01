package cmd

import (
	"fmt"
	"github.com/cmser/git-tools/pkg/types"
	"github.com/cmser/git-tools/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var release = &cobra.Command{
	Use: "release",
}

var createRelease = &cobra.Command{
	Use: "create",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := cmd.Flag("name").Value.String()
		if inc, err := utils.ToIncrementation(name); err == nil {
			repo := types.InitializeRepo(cmd.Flag("dir").Value.String(), cobra.CheckErr)
			return repo.CreateReleaseBranch(*inc)
		}
		return nil
	},
}

func init() {
	currDir, _ := os.Getwd()
	release.PersistentFlags().String("dir", currDir, fmt.Sprintf("--dir %s", currDir))
	release.PersistentFlags().String("name", "", "--name foo bar +semver:patch|major|minor")
	release.AddCommand(createRelease)
}
