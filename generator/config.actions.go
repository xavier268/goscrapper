package generator

import (
	"fmt"
	"reflect"
)

// List of defined ActionKeys.
// Not lowercase, because the corresponding fields are public.
var ActionKeys []string = extractStructFields(ConfigAction{})

// ConfigAction defines action to conduct.
// It incorporates the syntax for each possible action.
// There should be only ONE and ONLY ONE field set.
type ConfigAction struct {
	// === Only one of these will be not nil. Check will happen at compile time.
	Click struct { // click on an element
		Selector string // element selector, relative to job scope.
		Left     bool   // left button (right is default)
	}
	Scope struct { // set the scope
		Selector string // element selector
	}
}

// Verify ONE and only ONE Action is set, and return the corresponding action.
func (a ConfigAction) configActionVerify() (actionName string, err error) {
	count := 0
	for _, f := range ActionKeys {
		if !reflect.ValueOf(a).FieldByName(f).IsZero() {
			count++
			actionName = f
		}
	}
	if count == 1 {
		return actionName, nil
	}
	return "", fmt.Errorf("exactly ONE valid action can be set, but found %d", count)
}
