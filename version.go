// top level package, not much at this level.
package goscrapper

import (
	"regexp"
	"strconv"
)

const (
	VERSION   = "0.4.12"
	COPYRIGHT = "(c) Xavier Gandillot 2024"
)

// build time variables - will be substituted at build time.
var (
	GITHASH   string = "n/a"
	BUILDDATE string = "n/a"
)

// Parse the version into its components.
// Shorter version string are accepted, corresponding values default to 0.
// Non numbers also default to 0.
// There can be a string prefix or suffix.
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
