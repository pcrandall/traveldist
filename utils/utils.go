package utils

import (
	"regexp"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// GenerateUUID returns a uuid v4 in string
func GenerateUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "generating uuid")
	}

	return u.String(), nil
}

// regexNumber is a regex that matches a string that looks like an integer
var regexNumber = regexp.MustCompile(`^\d+$`)

// IsNumber checks if the given string is in the form of a number
func IsNumber(s string) bool {
	if s == "" {
		return false
	}

	return regexNumber.MatchString(s)
}
