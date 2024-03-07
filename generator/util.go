package generator

import (
	"encoding/json"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// PrettyJson print an object, using its json format
func PrettyJson(st interface{}) string {
	if st == nil {
		return "nil"
	}
	bb, err := json.MarshalIndent(st, "", "     ")
	if err != nil {
		return string(bb) + " **ERROR during pretty pring ** " + err.Error()
	}
	return string(bb)
}

// PretyYaml print an object, using its yaml format
func PrettyYaml(st interface{}) string {
	if st == nil {
		return "nil"
	}
	bb, err := yaml.Marshal(st)
	if err != nil {
		return string(bb) + " **ERROR during pretty pring ** " + err.Error()
	}
	return string(bb)
}

// Get the absolute path for given string.
func MustAbs(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return abs
}

// Normalize string to lower,removing all invaloiid chars.
func Normalize(s string) string {
	s = strings.ToLower(s)
	patt := regexp.MustCompile(`[a-z][a-z0-9]*`)
	s = patt.FindString(s)
	return s
}

// Upper the first letter of a string.
func UpFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}
