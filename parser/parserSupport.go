package parser

import (
	"fmt"
	"strings"
)

// ===========================================================================================
// When generating code, global state is kept within myLexer.
// It can be reached from code in the grammar file using lex.(*myLexer).doSomething() ...
// ===========================================================================================

// add lines to the code to be generated.
func (m *myLexer) addLines(lines ...string) {
	m.lines = append(m.lines, lines...)
}

// finalize code generation, and write it to the output file.
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

	// write type of inputParams
	fmt.Fprintln(m.w)
	fmt.Fprintf(m.w, "type %s struct {\n", "Input_"+m.name)
	for _, l := range m.inparams {
		fmt.Fprintf(m.w, "\t%s\n", l)
	}
	fmt.Fprintln(m.w, "}")
	fmt.Fprintln(m.w)

	// wrtite type of output parameters
	fmt.Fprintln(m.w)
	fmt.Fprintf(m.w, "type %s struct {\n", "Output_"+m.name)
	for _, l := range m.outparams {
		fmt.Fprintf(m.w, "\t%s\n", l)
	}
	fmt.Fprintln(m.w, "}")
	fmt.Fprintln(m.w)

	// write function header, with parameters
	fmt.Fprintf(m.w, "\nfunc Do_%s(_in Input_%s) (%s, _err error) {\n", m.name, m.name, "_out []Output_"+m.name)

	// write function body
	for _, l := range m.lines {
		fmt.Fprintf(m.w, "%s\n", l)
	}

	fmt.Fprintln(m.w, "return _out, _err\n}")

	// save the source function at the bottom of the file
	m.printCommentedSource()
}

// define an input parameter.
func (m *myLexer) setParam(name string, typ string) {
	m.inparams = append(m.inparams, name+" "+typ)
}

// print the data source used to create function, as a comment.
func (m *myLexer) printCommentedSource() {
	fmt.Fprintln(m.w)
	dd := strings.Split(string(m.data), "\n")
	for _, d := range dd {
		fmt.Fprintf(m.w, "// %s\n", d)
	}
	fmt.Fprintln(m.w)
}
