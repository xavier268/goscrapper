package parser

import (
	"fmt"
	"io"
	"regexp"
	"slices"
	"sort"
)

// Write all known variables to the given writer.
// Input parameters. Sorted for test stability.
func (it *Interpreter) DumpVars(w io.Writer, title string) {
	fmt.Fprintln(w, title)
	for lev, stack := range it.vars {
		fmt.Fprintf(w, "\nLevel %d :\n", lev)

		// collect sorted keys
		kk := make([]string, 0, len(stack))
		for k := range stack {
			kk = append(kk, k)
		}
		sort.Strings(kk)

		// print stack level dump
		for _, k := range kk {
			v := stack[k]
			if it.isInputVar(k) {
				fmt.Fprintf(w, "\t%s (input param) = %#v \n", k, v)
			} else {
				fmt.Fprintf(w, "\t%s = %#v\n", k, v)
			}
		}
	}
	fmt.Fprintln(w)
}

// push a new stack frame
func (it *Interpreter) pushFrame() {
	it.vars = append(it.vars, make(map[string]any))
}

// pop the stack frame
func (it *Interpreter) popFrame() error {
	if len(it.vars) == 1 {
		return fmt.Errorf("cannot pop root frame")
	}
	it.vars = it.vars[:len(it.vars)-1]
	return nil
}

// Assign (and declare if needed) a var in the current stack frame.
// Local value will shadow the more global value.
// Multiple reassignements are ok.
// Assigning to a var declared and assigned an input value by NewInterpreter is illegal.
func (it *Interpreter) assignVar(varName string, value any) error {
	if !isValidId(varName) {
		return fmt.Errorf("invalid var identifier: %s", varName)
	}
	if it.isInputVar(varName) {
		return fmt.Errorf("cannot reassign to input parameter: %s", varName)
	}
	it.vars[len(it.vars)-1][varName] = value
	return nil
}

// Retrieve the value for the var.
// Local values shadow the more global values, even if assigned to nil.
// Works for both interanl or input vars.
func (it *Interpreter) getVar(varName string) (value any, err error) {
	for i := len(it.vars) - 1; i >= 0; i-- {
		if v, ok := it.vars[i][varName]; ok {
			return v, nil
		}
	}
	return nil, fmt.Errorf("unknown var: %s", varName)
}

// check if name was declared as input var name.
func (it *Interpreter) isInputVar(varName string) bool {
	return slices.Contains(it.invars, varName)
}

// verify valid identifier - exclude all known tokens.
// case sensitive.
func isValidId(varName string) bool {
	patt := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*$`)
	return patt.MatchString(varName) && !slices.Contains(yyToknames[:], varName)
}
