package rt

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

// See Serialize. Panic on error.
func MustSerialize(a any) string {
	s, err := Serialize(a)
	if err != nil {
		panic(err)
	}
	return s
}

// Serialize a value to string, using the GSC syntax.
// Only values produced by GSC itself can be serialized.
// Maps keys are sorted to guarantee determistic serialization.
// Non recognized values will trigger an error.
func Serialize(a any) (string, error) {
	sb := new(strings.Builder)
	err := serialize(sb, a)
	// fmt.Println("DEBUG SERIALIZE :", sb.String())
	return sb.String(), err
}

// actual serialization for known types
func serialize(sb *strings.Builder, a any) error {
	switch v := a.(type) {
	case nil:
		fmt.Fprint(sb, "nil")
		return nil
	case int:
		fmt.Fprintf(sb, "%d", v)
		return nil
	case bool:
		if v {
			fmt.Fprint(sb, "true")
		} else {
			fmt.Fprint(sb, "false")
		}
		return nil
	case string:
		// escape one or more double quotes by adding one
		esc := regexp.MustCompile(`"+`)
		ss := esc.ReplaceAllStringFunc(v, func(match string) string {
			return strings.Repeat(`"`, len(match)+1)
		})
		fmt.Fprintf(sb, "\"%s\"", ss)
		return nil
	case *int:
		return serialize(sb, *v)
	case *bool:
		return serialize(sb, *v)
	case *string:
		return serialize(sb, *v)
	case []any:
		fmt.Fprint(sb, "[")
		for i, e := range v {
			if i > 0 {
				fmt.Fprint(sb, ", ")
			}
			if err := serialize(sb, e); err != nil {
				return err
			}
		}
		fmt.Fprint(sb, "]")
		return nil
	case map[string]any:
		// check and sort keys
		keys := make([]string, 0, len(v))
		for k := range v {
			if regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]*$").MatchString(k) {
				keys = append(keys, k)
			} else {
				return fmt.Errorf("cannot serialize map key %s (invalid key syntax)", k)
			}
		}
		sort.Strings(keys)
		first := true
		// serialize keys and values
		fmt.Fprint(sb, "{")
		for _, k := range keys {
			if !first {
				fmt.Fprint(sb, ", ")
			}
			first = false
			fmt.Fprintf(sb, "%s: ", k)
			if err := serialize(sb, v[k]); err != nil {
				return err
			}
		}
		fmt.Fprint(sb, "}")
		return nil
	case time.Time:
		fmt.Fprintf(sb, "time{%s}", v.Format(time.RFC3339))
		return nil
	case *Page:
		fmt.Fprintf(sb, "page{%s-%s-%s}", v.TargetID, v.FrameID, v.SessionID)
		return nil
	case *Element:
		if v == nil || v.Object == nil {
			fmt.Fprint(sb, "element{}")
		} else {
			xp, err := v.GetXPath(true) // optimized path will use id if available
			if err != nil {
				return err
			}
			fmt.Fprintf(sb, "element{%s}", xp)
		}
		return nil
	case Hash:
		fmt.Fprintf(sb, "hash{%x}", v)
		return nil
	default:
		fmt.Fprintf(sb, "???%#v???", a) // default to golang raw value.
		return fmt.Errorf("cannot serialize value of type %T", a)
	}
}

// SafePrintf is like Sprintf, but return runtime formatting errors without panic
func SafeSprintf(format string, a ...any) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error formatting with %s : %v", format, r)
		}
	}()
	result = fmt.Sprintf(format, a...)
	return result, err
}

// ==== Hashing ====

// Hash type (an array, not a slice)
type Hash [md5.Size]byte

// Compute the Hash of any GSC serializable value.
// Error if a cannot be GSC-serializable.
func HashGSC(a any) (Hash, error) {
	sb := new(strings.Builder)
	err := serialize(sb, a)
	if err != nil {
		return [md5.Size]byte{}, err
	}
	return md5.Sum([]byte(sb.String())), nil
}

// === Unique filter ===

// Uniqueness filter, using a GSC Hash.
// Use u := new(Unique) to initialize it.
// Typically for RETURN DISTINCT.
// For rod.Element, you may prefer using element.Equal(element) to comprare actual DOM elements in browser.
type Unique struct {
	values map[Hash]bool
}

var (
	ErrNotUnique = fmt.Errorf("value is not unique")
	ErrNilFilter = fmt.Errorf("cannot use a nil Unique filter")
)

// Add a value to the uniqueness filter.
// Return error if value was already there, or cannot be serialized.
func (u *Unique) Add(a any) error {
	if u == nil {
		return ErrNilFilter
	}
	if u.values == nil {
		u.values = make(map[Hash]bool, 5)
	}
	h, err := HashGSC(a)
	if err != nil {
		return err
	}
	if u.values[h] {
		return ErrNotUnique
	}
	u.values[h] = true
	return nil
}
