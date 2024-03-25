package parser

import (
	"fmt"
	"strings"
)

// start a loop, saving the output and adding the opening {
func (m *myLexer) forNameInExpression(name string, expr value) {

	// inc loop counter
	m.loops += 1

	if !strings.HasPrefix(expr.t, "[]") {
		m.errorf("FOR name IN expression requires expression to be an array, but %s is not an array; it is a %s", expr.v, expr.t)
	}

	for _, k := range m.inparams {
		if k == name {
			m.errorf("you cannot allocate the loop variable to an input parameter %s", name)
		}
	}
	if _, ok := m.vars[name]; ok {
		m.errorf("the variable %s already exists and cannot be used as a loop variable", name)
	}

	m.vars[name] = expr.t[2:] // set type of loop variable.

	li := fmt.Sprintf("for _, %s := range %s { \n _ = %s", name, expr.v, name)
	m.addLines(li)

}

func (m *myLexer) selectExpression(expr value) {

	// fmt.Println("DEBUG : calling selectExpression with", expr.v)

	// increment loop counter
	m.loops += 1

	if expr.t != "string" {
		m.errorf("FOR expression requires expression to be a string, but %s is not a string but a %s", expr.v, expr.t)
	}

	li := "for true { // this will be a loop with expr ..." // to do ...
	m.addLines(li)
}
