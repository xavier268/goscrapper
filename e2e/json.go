package e2e

import "encoding/json"

// Return an indented json string representation or an human readeable error string.
// Never fails. Used for tests and debugging.
func prettyJson(a any) string {
	bb, err := json.MarshalIndent(a, "", "   ")
	if err != nil {
		return err.Error()
	} else {
		return string(bb)
	}
}
