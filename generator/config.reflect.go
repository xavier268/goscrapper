package generator

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

// Use reflection to access configuration structure.

// Extract all fields from a struct instance.
// Fields are extracted exactly as there are in the struct, no lowercasing.
func extractStructFields(d interface{}) (names []string) {
	if d == nil {
		return nil
	}
	sf := reflect.VisibleFields(reflect.TypeOf(d))
	for _, f := range sf {
		names = append(names, f.Name)
	}
	return names
}

// Returns a description of the structure of a valid Configuration :
func HelpConfiguration() string {

	sb := new(strings.Builder)
	var bb []byte
	cf := Configuration{
		Schema:  1,
		Run:     "startstate1",
		AppName: "myapp1",
		Define: map[string]string{
			"definition1": "value1",
			"definition2": "value2",
			"...":         "...",
		},
		Buses: map[string]ConfigBus{
			"busname1": {},
			"busname2": {},
			"...":      {},
		},
		States: map[string]ConfigState{
			"statename1": {},
			"statename2": {},
			"...":        {},
		},
	}
	fmt.Fprintln(sb, "\n=== The main Configuration object has the following structure :\n ")
	//bb, _ = json.MarshalIndent(cf, "", "  ")
	bb, _ = yaml.Marshal(cf)
	sb.Write(bb)

	fmt.Fprintln(sb, "\n==== Valid actions are the following, with the following parameters :\n ")
	//bb, _ = json.MarshalIndent(cf, "", "  ")
	bb, _ = yaml.Marshal(ConfigAction{})
	sb.Write(bb)

	return sb.String()

}
