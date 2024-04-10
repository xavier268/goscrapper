package parser

import "encoding/json"

func PrettyJson(a any) string {
	bb, err := json.MarshalIndent(a, "", "   ")
	if err != nil {
		return err.Error()
	} else {
		return string(bb)
	}
}
