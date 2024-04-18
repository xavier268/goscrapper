package parser

import (
	"fmt"
	"io"
	"regexp"
	"slices"
	"sort"
)

// Write all known variables to the given writer,
// including input parameters. Sorted for test stability.
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
			fmt.Fprintf(w, "\t%s = %#v\n", k, v)
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

// reset just pop and push a new frame.
func (it *Interpreter) resetFrame() error {
	if err := it.popFrame(); err != nil {
		return err
	}
	it.pushFrame()
	return nil
}

// Assign (and declare if needed) a var in the current stack frame.
// Local value will shadow the more global value.
// Multiple reassignements are ok.
// Assigning to a var declared and assigned an input value by NewInterpreter is illegal, and should be prevented at compile time.
// If the bool flag is set, assign to the root scope (global scope).
func (it *Interpreter) assignVar(varName string, value any, global bool) error {
	if !isValidId(varName) {
		return fmt.Errorf("invalid var identifier: %s", varName)
	}
	if global {
		it.vars[0][varName] = value
	} else {
		it.vars[len(it.vars)-1][varName] = value
	}
	return nil
}

// Retrieve the value for the var.
// Local values shadow the more global values, even if assigned to nil.
// Works for both internal or input vars.
// If global flag, will only fetch the global scope value.
func (it *Interpreter) getVar(varName string, global bool) (value any, err error) {
	if global {
		// only serach global scope
		if v, ok := it.vars[0][varName]; ok {
			return v, nil
		}
		return nil, fmt.Errorf("no such global variable : %s", varName)
	} else {
		// look through local scopes
		for i := len(it.vars) - 1; i >= 0; i-- {
			if v, ok := it.vars[i][varName]; ok {
				return v, nil
			}
		}
		return nil, fmt.Errorf("no such variable : %s", varName)
	}
}

// verify valid identifier - exclude all known tokens.
// case sensitive.
func isValidId(varName string) bool {
	patt := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*$`)
	return patt.MatchString(varName) && !slices.Contains(yyToknames[:], varName)
}
