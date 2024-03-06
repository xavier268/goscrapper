package generator

import (
	"encoding/json"
	"path/filepath"
)

// Pretty print an object, using its json format
func Pretty(st interface{}) string {
	if st == nil {
		return "nil"
	}
	bb, err := json.MarshalIndent(st, "", "     ")
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
