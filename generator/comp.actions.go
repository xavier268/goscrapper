package generator

import (
	"fmt"
	"io"
)

// Generate all actions referenced by given state.
func (c *Compiler) generateActions(f io.Writer, statename string) error {

	fmt.Fprintln(f)

	cst, ok := c.conf.States[statename] // cst is the ConfigState
	if !ok {
		return nil // no action for this state
	}
	for i, ca := range cst.Actions { // loop over array of configActions, i is the action index, ca is configurationAction
		fmt.Fprintf(f, "\n// Action n°%d\n", i+1)
		for an, ap := range ca { // an is action name, ap is the action parameters map

			if err := c.generateAction(f, an, ap); err != nil {
				return err
			}
		}
	}
	return nil
}

// Generate inline code for given action with given parameters.
func (c *Compiler) generateAction(f io.Writer, actionname string, params ConfigParameters) error {
	fmt.Fprintf(f, "// %s %#v\n", actionname, params)
	return nil
}
