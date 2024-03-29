package parser

import (
	"fmt"
	"sort"
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
	// clean lines, and write function body
	m.cleanOut()
	for _, l := range m.lines {
		if l != "" {
			fmt.Fprintf(m.w, "%s\n", l)
		}
	}
	// print final return statement
	// there is laways an empty preallocated element in the slice, so remove it.
	fmt.Fprintln(m.w, "return _out[:len(_out) -1], _err")
	fmt.Fprintln(m.w, "}")
}

// write import code
func (m *myLexer) wImports() {
	// write nothing if no imports were added during parsing.
	if len(m.imports) == 0 {
		return
	}

	fmt.Fprintln(m.w, "import (")
	// sort imports for reproductibility
	keys := make([]string, 0, len(m.imports))
	for k := range m.imports {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// add imports added during parsing
	for _, k := range keys {
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
// NB : the _err returned, will be a RUNTIME error, not a parse time error !
func (m *myLexer) wFuncDeclaration() {
	fmt.Fprintf(m.w,
		"func Do_%s(_in Input_%s) (_out []Output_%s, _err error) {\n",
		m.name, m.name, m.name)

	// write lateDecl lines, sorted.
	ld := make([]string, 0, len(m.lateDecl)) // lateDecl lines, sorted.
	for l := range m.lateDecl {
		ld = append(ld, l)
	}
	sort.Strings(ld)
	for _, l := range ld {
		fmt.Fprintf(m.w, "%s\n", l)
	}
}
