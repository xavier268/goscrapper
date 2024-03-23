package parser

import (
	"fmt"
	"slices"
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

	// write type of output parameters
	fmt.Fprintln(m.w)
	fmt.Fprintf(m.w, "type %s struct {\n", "Output_"+m.name)
	for _, l := range m.outparams {
		fmt.Fprintf(m.w, "\t%s\n", l)
	}
	fmt.Fprintln(m.w, "}")
	fmt.Fprintln(m.w)

	// write function header, with parameters
	fmt.Fprintf(m.w, "\nfunc Do_%s(_in Input_%s) (_res []Output_%s, _err error) {\n", m.name, m.name, m.name)

	// initialize _res
	m.nextRes()

	// write function body
	for _, l := range m.lines {
		fmt.Fprintf(m.w, "%s\n", l)
	}

	fmt.Fprintln(m.w, "return _res, _err\n}")

	// save the source function at the bottom of the file
	m.printCommentedSource()
}

// define an input parameter.
func (m *myLexer) setParam(name string, typ string) {
	m.inparams = append(m.inparams, name+" "+typ)
}

func (m *myLexer) errorf(format string, args ...interface{}) {
	m.Error(fmt.Sprintf(format, args...))
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

// prepare res to accept a new set of output variables,
// by writing the code to do that.
func (m *myLexer) nextRes() {
	fmt.Fprintf(m.w, "_res = append(_res, Output_%s{})\n", m.name)
}

// set the value of a variable in the current set of output variables.
func (m *myLexer) setVar(name string, value string) {
	m.addLines(fmt.Sprintf("\t\t_res[len(_res)-1].%s = %s\n", name, value))
}

func (m *myLexer) getVar(name string) string {
	if slices.Contains(m.inparams, name) {
		return name
	}
	if slices.Contains(m.outparams, name) {
		return name
	}
	m.errorf("variable %s is not yet defined", name)
	return name
}
