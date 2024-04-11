package rt

import "fmt"

// Set this hook to capture runtime browser related errors.
// If not set, will report to stdout.
var Errorf func(format string, args ...any)

func init() {
	Errorf = // set default error reporting function
		func(format string, args ...any) {
			fmt.Println("Runtime error : ")
			fmt.Printf(format, args...)
			fmt.Println()
		}
}
