package generator

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflection(t *testing.T) {

	// Extract all public fields from struct
	sf := reflect.VisibleFields(reflect.TypeOf(ConfigAction{}))
	for _, f := range sf {
		fmt.Println(f.Name)
	}

}

func TestFunctionDescription(_ *testing.T) {

	d, e := (ConfigAction{}).getActionFunctionDeclaration("Scope")
	if e != nil {
		panic(e)
	}
	fmt.Println(d)

}
