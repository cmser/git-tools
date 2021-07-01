package utils

import (
	"errors"
	"github.com/cmser/git-tools/pkg/types"
	"regexp"
	"strings"
)

func ToIncrementation(name string) (*types.Incrementation, error)  {
	r := regexp.MustCompile(`\+semver:(major|minor|patch)$`)
	if !r.MatchString(name) {
		return nil, errors.New("unable to match an incrementation")
	}
	path := strings.Split(name, ":")
	inc := types.Incrementation(path[len(path) - 1])
	return &inc, nil
}
