package parser

import (
	"fmt"
	"strings"
)

// ===========================================================================================
// When generating code, global state is kept within myLexer.
// It can be reached from code in the grammar file using lex.(*myLexer).doSomething() ...
// ===========================================================================================

// finalize code generation, and write it to the output file from the stored code lines.
// ensure generation are called in the right order.
func (m *myLexer) finalize() {

	m.wImports()

	m.wInputParams()
	m.wOutputParams()

	m.wCommentedSource()

	m.wFuncDeclaration()

	// write function body
	for _, l := range m.lines {
		fmt.Fprintf(m.w, "%s\n", l)
	}

	// return function results, finish function definition.
	fmt.Fprintln(m.w, "return _res, _err")
	fmt.Fprintln(m.w, "}")

}

// write import code
func (m *myLexer) wImports() {

	fmt.Fprintln(m.w, "import (")
	// set import defaults
	m.imports["github.com/xavier268/goscrapper/rt"] = true
	// add imports added during parsing
	for k := range m.imports {
		fmt.Fprintf(m.w, "\t%q\n", k)
	}
	fmt.Fprintln(m.w, ")")
	fmt.Fprintln(m.w)
}

// define name and type of inputParams
func (m *myLexer) wInputParams() {
	fmt.Fprintln(m.w)
	fmt.Fprintf(m.w, "type %s struct {\n", "Input_"+m.name)
	for _, l := range m.inparams {
		ty := m.vars[l] // golang type of the variable
		fmt.Fprintf(m.w, "\t%s %s\n", l, ty)
	}
	fmt.Fprintln(m.w, "}")
	fmt.Fprintln(m.w)
}

// define name and type of outputParams
func (m *myLexer) wOutputParams() {
	fmt.Fprintln(m.w)
	fmt.Fprintf(m.w, "type %s struct {\n", "Output_"+m.name)
	for _, l := range m.outparams {
		ty, ok := m.vars[l] // verify type was set ...
		if !ok || ty == "" {
			m.errorf("cannot return %s because its value (and its type) has not been set", l)
		}
		fmt.Fprintf(m.w, "\t%s %s\n", l, ty)
	}
	fmt.Fprintln(m.w, "}")
	fmt.Fprintln(m.w)
}

// writes the data source used to create function, as a comment.
func (m *myLexer) wCommentedSource() {
	fmt.Fprintln(m.w)
	dd := strings.Split(string(m.data), "\n")
	for _, d := range dd {
		fmt.Fprintf(m.w, "// %s\n", d)
	}
}

// writes function declaration, with input/output types.
func (m *myLexer) wFuncDeclaration() {
	fmt.Fprintf(m.w,
		"func Do_%s(_in Input_%s) (_res []Output_%s, _err error) {\n",
		m.name, m.name, m.name)
}
