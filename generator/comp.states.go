package generator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// generate enums for all states
func (c *Compiler) generateStates() error {

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

	// Print template for actions functions
	fmt.Fprintln(f, "/* The following action functions should be defined elsewhere")
	generateActionFunctionTemplates(f)
	fmt.Fprintln(f, "*/")

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
		switch j.state {`)

	for sn := range c.conf.States {
		fmt.Fprintf(f, `
			case %s :
			
	`, StateName(sn))

		err := c.generateStateCase(f, sn)
		if err != nil {
			return err
		}
	}

	fmt.Fprintln(f, `
			} // switch
		} // select
	} // for
} // Run
	`)
	return nil
}

func (c *Compiler) generateStateCase(f io.Writer, stateKey string) error {
	fmt.Fprintf(f, "		// generating state case for %s\n", stateKey)
	for _, confAct := range c.conf.States[stateKey].Actions {
		actName, err := confAct.configActionVerify()
		if err != nil {
			fmt.Fprintln(f, err)
			return err
		}
		err = c.generateAction(f, actName, confAct)
		if err != nil {
			fmt.Fprintln(f, err)
			return err
		}
	}
	return nil

}
