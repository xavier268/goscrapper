package rt

import "fmt"

// Hook to capture runtime browser related errors. Change as needed.
var Errorf func(format string, args ...any) = defaultErrorf

// default error reporting
func defaultErrorf(format string, args ...any) {

	fmt.Println("Runtime : ")
	fmt.Printf(format, args...)
	fmt.Println()
}
