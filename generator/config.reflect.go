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

// Get the parameter list of a named action :  name, type, kind, value
func (c ConfigAction) getActionParameterList(actname string) (
	params []struct {
		n string
		t string
		k string
		v interface{}
	},
	err error) {

	b := reflect.ValueOf(c)
	if b.Kind() != reflect.Struct {
		err = fmt.Errorf("cannot pull parameter lists : Configuration is not a struct")
		return nil, err
	}
	a := b.FieldByName(actname)
	// if a.IsZero() {
	// 	err = fmt.Errorf("cannot pull parameter lists : Action %s not found", actname)
	// 	return nil, err
	// }

	params = make([]struct {
		n string
		t string
		k string
		v interface{}
	}, a.Type().NumField())

	for i := 0; i < a.Type().NumField(); i++ {
		params[i].n = a.Type().Field(i).Name
		params[i].t = a.Field(i).Type().String()
		params[i].k = a.Field(i).Kind().String()
		params[i].v = a.Field(i).Interface()
	}

	return params, err
}

// Return the function name for an given action name.
func getActionFunctionName(actname string) string {
	return "doAction" + UpFirst(Normalize(actname))
}

// Return the declaration string for the named action function.
// Ex: "func doActionClick( Selector string, Left bool) error"
func (ac ConfigAction) getActionFunctionDeclaration(actname string) (string, error) {

	pp, err := ac.getActionParameterList(actname)
	if err != nil {
		return "", err
	}

	sb := new(strings.Builder)
	fmt.Fprintf(sb, "func (j *Job)%s( ", getActionFunctionName(actname))
	for i, p := range pp {
		if i > 0 {
			fmt.Fprintf(sb, ", ")
		}
		fmt.Fprintf(sb, "%s %s", p.n, p.t)
	}
	sb.WriteString(") error")

	return sb.String(), err
}
