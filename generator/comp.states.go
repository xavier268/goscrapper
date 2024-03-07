package generator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// generate all states and create file
func (c Compiler) generateStates() error {

	f, err := os.Create(filepath.Join(c.PackageDir, "states.go"))
	if err != nil {
		return fmt.Errorf("failed to create states.go: %v", err)
	}
	defer f.Close()

	err = c.writeHeader(f)
	if err != nil {
		return fmt.Errorf("failed to write header to states.go: %v", err)
	}

	fmt.Fprintln(f, "// A State is just any function that can be applied to *Scrapper")
	fmt.Fprintln(f, "type State func(*Scrapper) error")
	fmt.Fprintln(f)

	for sn := range c.conf.States {
		err = c.generateState(f, sn)
		if err != nil {
			return fmt.Errorf("failed to write state %s: %v", sn, err)
		}
	}
	return nil
}

// generate one state function for the named state
func (c *Compiler) generateState(f io.Writer, sname string) error {

	fmt.Fprintf(f, "\n/***** state : %s *******\n%s\n **********************/\n", sname, PrettyJson(c.conf.States[sname]))
	S := StateName(sname) // Normalized name

	fmt.Fprintf(f, "func  %s(s *Scrapper) error {\n", S)
	fmt.Fprintln(f, `panic("Not implemented")`)
	fmt.Fprintln(f, "}")

	return nil
}
