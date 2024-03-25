package parser

import (
	"fmt"
	"strings"
)

// start a loop, saving the output and adding the opening {
func (m *myLexer) forNameInExpression(name string, expr value) {

	// inc loop counter
	m.loops += 1

	if !strings.HasPrefix(expr.t, "[]]") {
		m.errorf("FOR ... IN ... needs an array to loop, but %s is not an array but a %s", expr.v, expr.t)
	}
	m.vars[name] = expr.v[2:] // set type of loop variable.
	li := fmt.Sprintf("for %s := range %s {", name, expr.v)
	m.addLines(li)
	m.incOut()

}

// finish a loop, saving the output and adding the closing }
func (m *myLexer) finishLoop() {
	m.saveOut()
	m.addLines("}")
}
