package config

import "encoding/json"

// pretty print an object, using its json format
func pretty(st interface{}) string {
	if st == nil {
		return "nil"
	}
	bb, err := json.MarshalIndent(st, "", "     ")
	if err != nil {
		return string(bb) + " **ERROR** " + err.Error()
	}
	return string(bb)
}
