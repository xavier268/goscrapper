package generator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// generate enums for all states
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

	fmt.Fprintln(f)
	fmt.Fprintln(f, "type State int")
	fmt.Fprintln(f, "const (")
	first := true
	for sn := range c.conf.States {
		if first {
			fmt.Fprintf(f, "%s State = iota\n", StateName(sn))
			first = false
		} else {
			fmt.Fprintf(f, "%s\n", StateName(sn))
		}
	}
	fmt.Fprintln(f, ")")

	return nil
}

// generate one state function for the named state
func (c *Compiler) generateState(f io.Writer, sname string) (err error) {

	fmt.Fprintf(f, "\n/***** state : %s *******\n%s\n **********************/\n", sname, PrettyJson(c.conf.States[sname]))

	fmt.Fprintf(f, "func  %s(j *Job) {\n", StateName(sname))
	err = c.generateActions(f, sname)
	if err != nil {
		return fmt.Errorf("failed to generate actions for state %s: %v", sname, err)
	}
	fmt.Fprintln(f, `panic("Not implemented")`)
	fmt.Fprintln(f, "}")
	fmt.Fprintln(f, "\n// ensure no compiler warning for stets that are never called")
	fmt.Fprintf(f, "var _ = %s\n\n", StateName(sname))

	return nil
}
