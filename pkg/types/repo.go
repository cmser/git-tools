package types

import (
	"context"
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"os"
	"sort"
)

type Repo struct {
	Repository    *git.Repository
	GraphQlClient *githubv4.Client
}

type Incrementation string
type IncrementingFunction func(version *semver.Version) semver.Version


const (
	Major Incrementation = "major"
	Minor Incrementation = "minor"
	Patch Incrementation = "patch"
)

var incrementationMap = map[Incrementation] IncrementingFunction {
	Major: func(version *semver.Version) semver.Version {
		return version.IncMajor()
	},
	Minor: func(version *semver.Version) semver.Version {
		return version.IncMinor()
	},
	Patch: func(version *semver.Version) semver.Version {
		return version.IncPatch()
	},
}

func (r *Repo) CreateReleaseBranch(inc Incrementation) error {
	next, err := r.IncrementVersion(inc)
	if err != nil {
		return err
	}
	headRef, err := r.Repository.Head()
	if err != nil {
		return err
	}
	return r.pushReleaseBranch(next, headRef)
}

func (r Repo) FetchLatestTag() (*semver.Version, error) {
	tags, err := r.Repository.Tags()
	if err != nil {
		return nil, err
	}
	var versions []*semver.Version
	err = tags.ForEach(func(reference *plumbing.Reference) error {
		if version, err := semver.NewVersion(reference.Name().Short()); err == nil {
			versions = append(versions, version)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(versions) == 0 {
		return semver.NewVersion("0.0.0")
	}
	sort.Sort(semver.Collection(versions))
	return versions[len(versions) - 1], nil
}

func (r Repo) IncrementVersion(inc Incrementation) (semver.Version, error) {
	current, err := r.FetchLatestTag()
	if err != nil {
		return semver.Version{}, err
	}
	version := incrementationMap[inc](current)
	return version, err
}

func (r *Repo) pushReleaseBranch(next semver.Version, ref *plumbing.Reference) error {
	err := r.pushRef(
		fmt.Sprintf("refs/heads/release/v%s", next.String()),
		ref.Hash().String(),
	)
	if err != nil {
		return err
	}
	return r.pushRef(
		fmt.Sprintf("refs/tags/v%s", next.String()),
		ref.Hash().String(),
	)
}

func (r *Repo) pushRef(name string, hash string) error  {
	var m struct {
		CreateRef struct {
			Ref struct {
				Name githubv4.String
			}
		} `graphql:"createRef(input:$input)"`
	}
	input := githubv4.CreateRefInput{
		RepositoryID:     githubv4.ID(os.Getenv("INPUT_REPO-ID")),
		Name:             githubv4.String(name),
		Oid:              githubv4.GitObjectID(hash),
	}
	return r.GraphQlClient.Mutate(context.Background(), &m, input, nil)
}

func InitializeRepo(dir string, errHandler func(msg interface{})) *Repo {
	repo, err := git.PlainOpen(dir)
	errHandler(err)
	return &Repo{
		Repository: repo,
		GraphQlClient: createClient(),
	}
}

func createClient() *githubv4.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	return githubv4.NewClient(httpClient)
}
