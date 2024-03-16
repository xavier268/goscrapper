package generator

import (
	"fmt"
	"reflect"
)

// List of defined ActionKeys.
// Not lowercase, because the corresponding fields are public.
var ActionKeys []string = extractStructFields(ConfigAction{})

// ConfigAction defines action to conduct.
// It incorporates the syntax for each possible action, like a UNION.
// There should be only ONE and ONLY ONE first level field set
// and the second levels fields should not all be the ZeroValue.
type ConfigAction struct {
	// === Only one of these will be not nil. Check will happen at compile time.
	// It is impossible to select an Action without setting at least one of its fields to a non zero value.

	Load struct {
		Url   string // url to load
		Blank bool   // load Blank page if url is empty
	}
	Quit struct {
		Job bool // only quit current job
		All bool // quit all jobs and stop application
	}
	Scope struct { // set the scope
		Selector string // element selector
	}
	Click struct { // click on an element
		Selector   string   // required element selector, relative to job scope.
		Left       bool     // left button (right is default)
		Experiment struct { // experiment
			ex   int
			peri float64
			ment []string
		}
	}
}

// Verify ONE and only ONE Action is set per ConfigACtion, and return the corresponding action.
func (a ConfigAction) configActionVerify() (actionName string, err error) {
	count := 0
	for _, f := range ActionKeys {
		if !reflect.ValueOf(a).FieldByName(f).IsZero() {
			count++
			actionName = f
		}
	}
	if count == 1 {
		if DEBUG >= LEVEL_DEBUG {
			fmt.Printf("action verification for %s succeeded\n", actionName)
		}
		return actionName, nil
	}
	return "", fmt.Errorf("exactly ONE valid action should be set, but found %d", count)
}
