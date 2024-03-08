package generator

import (
	"fmt"
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

	fmt.Fprintln(f)
	fmt.Fprintf(f, "func (j *Job) Run(state State) error {")
	fmt.Fprintln(f, `
	j.state = state
	for {
	select{
	case <- Done :  	// external close request
		j.sc.Close()
		return nil

	case <- j.sc.ctx.Done() : // internal close request
		j.sc.Close()
		return nil

	default: 
		switch j.state {

		


	
	
	`)

	fmt.Fprintln(f, `
		} // switch
	} // select
	} // for
} // Run
	`)
	return nil
}
