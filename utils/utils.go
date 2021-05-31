package utils

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/pcrandall/travelDist/utils/stripansi"
	"github.com/pkg/errors"
)

var (
	regexNumber = regexp.MustCompile(`^\d+$`) // regexNumber is a regex that matches a string that looks like an integer
)

// GenerateUUID returns a uuid v4 in string
func GenerateUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "generating uuid")
	}

	return u.String(), nil
}

// IsNumber checks if the given string is in the form of a number
func IsNumber(s string) bool {
	if s == "" {
		return false
	}
	return regexNumber.MatchString(s)
}

// clean ansi codes and space from string
func StripString(s string) string {
	if s == "" {
		return ""
	}
	// return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
	return stripansi.Strip(strings.TrimSpace(s))
}
