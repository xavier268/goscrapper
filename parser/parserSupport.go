package parser

import (
	"fmt"
	"strings"
)

// When generating code, global state is kept within myLexer.
// It can be reached from code in the grammar file using lex.(*myLexer).doSomething() ...

func (m *myLexer) addLines(lines ...string) {
	m.lines = append(m.lines, lines...)
}

func (m *myLexer) finalize() {

	// imports
	fmt.Fprintln(m.w, "import (")
	// set import defaults
	m.imports["github.com/xavier268/goscrapper/rt"] = true
	// add imports added during parsing
	for k := range m.imports {
		fmt.Fprintf(m.w, "\t%q\n", k)
	}
	fmt.Fprintln(m.w, ")")
	fmt.Fprintln(m.w)
	fmt.Fprintln(m.w)
	fmt.Fprintln(m.w)

	// write function header, with parameters
	fmt.Fprintf(m.w, "\nfunc Do%s(%s) (%s) {\n", m.name, strings.Join(m.inparams, ","), strings.Join(m.outparams, ", "))

	// write function body
	for _, l := range m.lines {
		fmt.Fprintf(m.w, "%s\n", l)
	}
	fmt.Fprintln(m.w, "}")
	// possibly, save the source function at the bottom of the file
}

func (m *myLexer) setParam(name string, typ string) {
	m.inparams = append(m.inparams, name+" "+typ)
}
