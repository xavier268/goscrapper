package parser

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type myLexer struct {
	name string // name of the lexer source
	data []byte // entire data to be lexed
	pos  int    // next position to process
	// code building support
	w         io.Writer         // where are error messages written to ?
	lines     []string          // list of lines in the code, pending writing.
	inparams  []string          // contains the names of the input parameters.
	outparams []string          // contains the the name of the output parameters. Type should not change between scopes !
	vars      map[string]string // associate a type to a given var. All var defined  with same name in different scopes should have same types.
	imports   map[string]bool   // set of imports required, written at the end during finalize
	lateDecl  map[string]bool   // set of local variables that will be declared once before function starts. Ex : _page.

}

// add lines to the code to be generated.
func (m *myLexer) addLines(lines ...string) {
	m.lines = append(m.lines, lines...)
}

// Error implements yyLexer.
func (m *myLexer) Error(e string) {
	fmt.Fprintf(m.w, "\n%s ********* Error in %s :***************%s\n\n", ColRED, m.name, RESET)
	bef := max(0, m.pos-80)
	after := min(len(m.data), m.pos+80)
	fmt.Fprint(m.w, string(m.data[bef:m.pos]))
	fmt.Fprintf(m.w, " %s <<<<<<<<<<<<<<<<< %s %s\n", ColRED, e, RESET)
	fmt.Fprintln(m.w, string(m.data[m.pos:after]))
}

// utility to format error
func (m *myLexer) errorf(format string, args ...interface{}) {
	m.Error(fmt.Sprintf(format, args...))
}

// Lex implements yyLexer.
func (m *myLexer) Lex(lval *yySymType) int {

	var err error

startLoop:

	// check if we reached the end
	if m.pos >= len(m.data) {
		return 0
	}

	// skip white spaces and terminators
	if loc := regexp.MustCompile(`^[\s]+`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		m.pos += loc[1] // skip
		goto startLoop
	}

	// skip multiline comments
	if loc := regexp.MustCompile(`(?s)^\/\*.*\*\/+`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		m.pos += loc[1] // skip
		goto startLoop
	}

	// skip single line comments
	if loc := regexp.MustCompile(`^\/\/.*`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		m.pos += loc[1] // skip
		goto startLoop
	}

	// read strings between " "
	// You can escape inside " by adding one more.
	if loc := regexp.MustCompile(`(?s)^"(""+|[^"])*"`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.value.v = string(m.data[m.pos+1 : m.pos+loc[1]-1])     // remove external quotes
		lval.value.v = strings.Replace(lval.value.v, `""`, `"`, -1) // replace all doubled quotes escaped inside.
		lval.value.v = fmt.Sprintf(`%q`, lval.value.v)              // store as a quoted string
		lval.value.t = "string"
		lval.value.c = STRING
		m.pos += loc[1]
		return STRING
	}

	// read strings between ' '
	// You can escape inside ' by adding one more.
	if loc := regexp.MustCompile(`(?s)^'(''+|[^'])*'`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.value.v = string(m.data[m.pos+1 : m.pos+loc[1]-1])     // remove external quotes
		lval.value.v = strings.Replace(lval.value.v, `''`, `'`, -1) // replace all doubled quotes escaped inside.
		lval.value.v = fmt.Sprintf(`%q`, lval.value.v)              // stored as a back-quoted string
		lval.value.t = "string"
		lval.value.c = STRING
		m.pos += loc[1]
		return STRING
	}

	// read integer number.
	if loc := regexp.MustCompile(`^[+-]?[0-9]+`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.value.v = string(m.data[m.pos : m.pos+loc[1]])
		lval.value.t = "int"
		lval.value.c = NUMBER
		m.pos += loc[1]
		return NUMBER
	}

	// read symbols or operators
	err = m.tryAllOperators(lval)
	if err == nil {
		return lval.value.c // operator found
	}

	// keywords
	err = m.tryAllKeywords(lval)
	if err == nil {
		return lval.value.c // keyword found
	}

	if loc := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.value.v = string(m.data[m.pos : m.pos+loc[1]])
		lval.value.t = "IDENTIFIER"
		lval.value.c = IDENTIFIER
		m.pos += loc[1]
		return IDENTIFIER
	}

	m.errorf("unrecognized token : <%s>", m.data[m.pos:min(len(m.data), m.pos+20)]) // display error
	return 0
}

// try all keywords, returning their code if found.
func (m *myLexer) tryAllKeywords(lval *yySymType) error {

	for i, k := range yyToknames {
		// fmt.Printf("Trying %d %q\n", i, k)
		if m.try(k) { // Keywords are always upperCase
			// fmt.Printf("Returning %d for %q\n", i+yyPrivate-1, k)
			lval.value.t = k
			lval.value.c = i + yyPrivate - 1
			lval.value.v = k
			return nil
		}
	}

	return fmt.Errorf("no keyword found")
}

// try all operators, using an internal table
// update lval if found.
func (m *myLexer) tryAllOperators(lval *yySymType) error {
	opeTable := []struct {
		ope  string
		code int
	}{
		// special values
		{"true", BOOL},
		{"false", BOOL},
		{"int", INTTYPE},
		{"bool", BOOLTYPE},
		{"string", STRINGTYPE},
		{"bin", BINTYPE},
		//{"nil", NIL},

		// multi bytes
		{"<=", LTE},
		{">=", GTE},
		{"==", EQ},
		{"!=", NEQ},
		{"++", PLUSPLUS},
		{"--", MINUSMINUS},
		{"&&", AND},
		{"||", OR},
		{"..", DOTDOT},
		{"::", NAMESPACESEPARATOR},
		{"!~", REGEXNOTMATCH},
		{"=~", REGEXMATCH},
		// single bytes
		{":", COLON},
		{";", SEMICOLON},
		{"(", LPAREN},
		{")", RPAREN},
		{"{", LBRACE},
		{"}", RBRACE},
		{"[", LBRACKET},
		{"]", RBRACKET},
		{",", COMMA},
		{".", DOT},
		{"<", LT},
		{">", GT},
		{"*", MULTI},
		{"/", DIV},
		{"%", MOD},
		{"+", PLUS},
		{"-", MINUS},
		{"=", ASSIGN},
		{"?", QUESTION},
		{"@", AT},
	}
	for _, k := range opeTable {
		// fmt.Printf("Trying %d %q\n", i, k)
		if m.try(k.ope) {
			lval.value.t = TokenAsString(k.code)
			lval.value.c = k.code
			lval.value.v = k.ope
			return nil
		}
	}
	return fmt.Errorf("no operator found")
}

// Try a keyword, if success, update the lexer position and return true.
// Otherwise, return false.
// Keywords are case sensitive.
func (m *myLexer) try(what string) bool {
	if bytes.HasPrefix(m.data[m.pos:], []byte(what)) {
		m.pos += len(what)
		return true
	} else {
		return false
	}
}
