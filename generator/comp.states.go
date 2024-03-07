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
func (c *Compiler) generateState(f io.Writer, sname string) (err error) {

	fmt.Fprintf(f, "\n/***** state : %s *******\n%s\n **********************/\n", sname, PrettyJson(c.conf.States[sname]))

	fmt.Fprintf(f, "func  %s(j *job) {\n", StateName(sname))
	fmt.Fprintln(f, "defer j.sc.wg.Done()")
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
