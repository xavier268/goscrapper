package generator

import (
	"fmt"
	"reflect"
)

// List of defined ActionKeys.
// Not lowercase, because the corresponding fields are public.
var ActionKeys []string = extractFields(ConfigAction{})

// ConfigAction defines action to conduct.
// It incorporates the syntax for each possible action.
// There should be only ONE and ONLY ONE field set.
type ConfigAction struct {
	// === Only one of these will be not nil. Check will happen at compile time.
	// It is define as a pointer to a struct to facilitate detection of zero values.
	Click *struct { // click on an element
		Selector string // element selector, relative to job scope.
		Left     bool   // left button (right is default)
	}
	Scope *struct { // set the scope
		Selector string // element selector
	}
}

// Extract all fields from a struct instance.
// Fields are extracted exactly as there are in the struct, no lowercasing.
func extractFields(d interface{}) (names []string) {
	if d == nil {
		return nil
	}
	sf := reflect.VisibleFields(reflect.TypeOf(d))
	for _, f := range sf {
		names = append(names, f.Name)
	}
	return names
}

// Verify ONE and only ONE Action is set, and return the corresponding action.
func configActionVerify(a ConfigAction) (actionName string, err error) {
	count := 0
	for _, f := range ActionKeys {
		if !reflect.ValueOf(a).FieldByName(f).IsNil() {
			count++
			actionName = f
		}
	}
	if count == 1 {
		return actionName, nil
	}
	return "", fmt.Errorf("exactly ONE valid action can be set, but found %d", count)
}
