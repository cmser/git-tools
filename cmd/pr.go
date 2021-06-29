package cmd

import (
	"fmt"
	"github.com/cmser/git-tools/pkg/utils"
	"github.com/spf13/cobra"
)

var pr = &cobra.Command{
	Use: "pr",
}

var validatePrName = &cobra.Command{
	Use: "validate",
	RunE: func(cmd *cobra.Command, args []string) error {
		prName := cmd.Flag("name").Value.String()
		fmt.Printf("Validating PR name '%s'", prName)
		return utils.ValidatePrName(prName)
	},
}

func init() {
	validatePrName.Flags().String("name", "", "--name foo bar +semver:patch|major|minor")
	pr.AddCommand(validatePrName)
}
