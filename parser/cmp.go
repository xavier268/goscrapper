package parser

import (
	"fmt"
	"reflect"
)

// CompareAny takes two interface{} values and tries to compare them.
// It returns -1 if v1 < v2, 0 if v1 == v2, and 1 if v1 > v2.
// If the types are different or not comparable, it returns an error.
func CompareAny(v1, v2 any) (int, error) {

	// handle broad cases of equality
	if reflect.DeepEqual(v1, v2) {
		return 0, nil
	}

	// handle nils
	if v1 == nil && v2 == nil {
		return 0, nil
	}
	if v1 == nil || v2 == nil {
		return 0, fmt.Errorf("nil can only compare to it self")
	}

	// handle inequality
	rv1, rv2 := reflect.ValueOf(v1), reflect.ValueOf(v2)

	// Check if the types are the same
	if rv1.Type() != rv2.Type() {
		return 0, fmt.Errorf("cannot compare types %s and %s", rv1.Type(), rv2.Type())
	}

	// Ensure the values are valid
	if !rv1.IsValid() || !rv2.IsValid() {
		return 0, fmt.Errorf("invalid value")
	}

	// Actual comparison based on types
	switch rv1.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i1, i2 := rv1.Int(), rv2.Int()
		switch {
		case i1 < i2:
			return -1, nil
		case i1 > i2:
			return 1, nil
		default:
			return 0, nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		u1, u2 := rv1.Uint(), rv2.Uint()
		switch {
		case u1 < u2:
			return -1, nil
		case u1 > u2:
			return 1, nil
		default:
			return 0, nil
		}
	case reflect.Float32, reflect.Float64:
		f1, f2 := rv1.Float(), rv2.Float()
		switch {
		case f1 < f2:
			return -1, nil
		case f1 > f2:
			return 1, nil
		default:
			return 0, nil
		}
	case reflect.String:
		s1, s2 := rv1.String(), rv2.String()
		switch {
		case s1 < s2:
			return -1, nil
		case s1 > s2:
			return 1, nil
		default:
			return 0, nil
		}
	default:
		return 0, fmt.Errorf("cannot compare between type %s", rv1.Type())
	}
}
