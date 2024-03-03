package config

import (
	"regexp"
	"strconv"
)

const (
	VERSION   = "0.0.3"
	COPYRIGHT = "(c) Xavier Gandillot 2024"
)

// build time variables
var (
	GITHASH   string = "n/a"
	BUILDDATE string = "n/a"
)

// runtime variables (from flags or test configuration)
var (
	DEBUG = 0
)

// Parse the version into its components.
// Shorter version string are accepted, corresponding values default to 0.
// Non numbers also default to 0.
// There can ba a string prefix or suffix.
func ParseVersion() (major int, minor int, patch int) {
	return parseVersion(VERSION)
}

func parseVersion(version string) (major int, minor int, patch int) {
	var err error

	patt := regexp.MustCompile(`(\d+)(\.((\d+)(\.(\d+)?)?)?)?`)
	matches := patt.FindStringSubmatch(version)

	if len(matches) > 0 {
		major, err = strconv.Atoi(matches[1])
		if err != nil {
			major = 0
		}
	}

	if len(matches) > 4 {
		minor, err = strconv.Atoi(matches[4])
		if err != nil {
			minor = 0
		}
	}

	if len(matches) > 6 {
		patch, err = strconv.Atoi(matches[6])
		if err != nil {
			patch = 0
		}
	}

	return major, minor, patch
}
