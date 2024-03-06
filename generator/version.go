package generator

import (
	"regexp"
	"strconv"
)

const (
	VERSION   = "0.1.0"
	COPYRIGHT = "(c) Xavier Gandillot 2024"
	SCHEMA    = "1" // required schema in configuration files

	// Debugging levels
	LEVEL_SILENT  = 0 // No output
	LEVEL_INFO    = 1 // only critical output
	LEVEL_VERBOSE = 2 // Verbose output for end user
	LEVEL_DEBUG   = 3 // Debugging data
)

// build time variables - will be subsituted at build time.
var (
	GITHASH   string = "n/a"
	BUILDDATE string = "n/a"
)

// runtime variables (from flags or test configuration)
var (
	// debugging level during generation process
	DEBUG_LEVEL int = LEVEL_DEBUG
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
