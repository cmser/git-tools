package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"sort"
)

var genTag = &cobra.Command{
	Use: "gen-tag",
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cmd.PersistentFlags().GetString("dir")
		cobra.CheckErr(err)
		repo, err := git.PlainOpen(dir)
		cobra.CheckErr(err)
		tagrefs, err := repo.Tags()
		cobra.CheckErr(err)
		lastCommit, err := repo.Head()
		cobra.CheckErr(err)
		commit, err := repo.CommitObject(lastCommit.Hash())
		cobra.CheckErr(err)
		fmt.Println(commit.Message)
		tagList := []string{};
		repo.Push(&git.PushOptions{})
		err = tagrefs.ForEach(func(t *plumbing.Reference) error {
			fmt.Println(t.Name().IsRemote())
			fmt.Println(t.Name().IsTag())
			fmt.Println(semver.IsValid(t.Name().Short()))
			if t.Name().IsTag() && semver.IsValid(t.Name().Short()) {
				tagList = append(tagList, t.Name().Short())
			}
			return nil
		})
		sort.SliceStable(tagList, func(i, j int) bool {
			return semver.Compare(tagList[i], tagList[j]) > 0
		})
		fmt.Println(tagList)
	},
}
