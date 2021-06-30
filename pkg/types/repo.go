package types

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type Repo struct {
	Repository *git.Repository
}

func (r *Repo) CreateReleaseBranch(errHandler func(msg interface{})) error {
	headRef, err := r.Repository.Head()
	errHandler(err)
	ref := plumbing.NewHashReference("refs/heads/release/toto", headRef.Hash())
	err = r.Repository.Storer.SetReference(ref)
	errHandler(err)
	return r.Repository.Push(&git.PushOptions{
		RemoteName: "origin",
		RefSpecs:   []config.RefSpec{ config.RefSpec(fmt.Sprintf("%s:%s", ref.Strings()[0], ref.Strings()[0])) },
	})
}

func (r Repo) FetchTags(errHandler func(msg interface{})) error  {
	tags, err := r.Repository.Tags()
	errHandler(err)
	tags.ForEach(func(reference *plumbing.Reference) error {
		fmt.Println(reference.Strings())
		return nil
	})
	return nil
}

func InitializeRepo(dir string, errHandler func(msg interface{})) *Repo {
	repo, err := git.PlainOpen(dir)
	errHandler(err)
	return &Repo{
		Repository: repo,
	}
}
