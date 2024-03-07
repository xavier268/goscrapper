package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// crete the Bus definition go file.
// Buses are publicly available.
func (c *Compiler) generateBuses() (err error) {

	f, err := os.Create(filepath.Join(c.PackageDir, "buses.go"))
	if err != nil {
		return fmt.Errorf("failed to create buses.go: %v", err)
	}
	defer f.Close()

	err = c.writeHeader(f)
	if err != nil {
		return fmt.Errorf("failed to write header to buses.go: %v", err)
	}
	fmt.Fprintln(f, "\nvar (")

	fmt.Fprintln(f, "	// Predefined build-in buses")
	fmt.Fprintln(f, "	Done = make( chan interface{} ) 	// will be closed upon termination request from scrapper")
	fmt.Fprintln(f, "	Messages = make( chan string , 30) 	// messages sent. Make sure you read them to avoid blocking !")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "	// Buses from configuration files")

	for b, bd := range c.conf.Buses {
		fmt.Fprintf(f, "	%s = make( chan string, %d )\n", BusName(b), bd.Limit)
	}
	fmt.Fprintln(f, ")")
	return nil
}
