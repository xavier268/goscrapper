package rt

import "fmt"

func Serialize(a any) (string, error) {
	panic("not implemented")
}

// SafePrintf is like Sprintf, but return formatting errors without panic
func SafeSprintf(format string, a ...any) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error formatting with %s : %v", format, r)
		}
	}()
	result = fmt.Sprintf(format, a...)
	return result, err
}
