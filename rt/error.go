package rt

import "fmt"

// hook to handle errors emitted at runtime
// can be set to any user designed error handler.
var Errorf = func(format string, args ...interface{}) {
	fmt.Println("**************")
	fmt.Println("Runtime Error:")
	fmt.Printf(format, args...)
}
