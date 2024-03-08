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
