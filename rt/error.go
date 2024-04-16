package rt

import "fmt"

// Set these hooks to capture runtime errors orlog messages.
// If not set, will report to stdout.
var Errorf func(format string, args ...any)
var Logf func(format string, args ...any)

func init() {
	Errorf = // set default error reporting function
		func(format string, args ...any) {
			fmt.Println("Runtime error : ")
			fmt.Printf(format, args...)
			fmt.Println()
		}
	Logf = func(format string, args ...any) {
		fmt.Printf("Runtime log : ")
		fmt.Printf(format, args...)
		fmt.Println()
	}
}
