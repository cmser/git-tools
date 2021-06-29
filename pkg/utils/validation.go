package utils

import (
	"errors"
	"fmt"
	"regexp"
)

const pattern = `\+semver:(major|minor|patch)$`

func ValidatePrName(name string) error {
	if match, _ := regexp.MatchString(pattern, name); match {
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to match valid PR name against pattern %s", pattern))
}
