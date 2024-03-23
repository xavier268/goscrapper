package parser

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type myLexer struct {
	name string // name of the lexer
	data []byte // entire data to be lexed
	pos  int    // next position to process
	// code building support
	w         io.Writer       // where are error messages written to ?
	lines     []string        // code lines collected so far
	inparams  []string        // contains the defnitions for the name type definition of the function input parameters, ex :  "name string"
	outparams []string        // contains the definitions for the name type definition of the function output parameters, ex :  "name string"
	imports   map[string]bool // set of imports required

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
		m.pos += loc[1]
		return STRING
	}

	// read integer number.
	if loc := regexp.MustCompile(`^[+-]?[0-9]+`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.value.v = string(m.data[m.pos : m.pos+loc[1]])
		lval.value.t = "int"
		m.pos += loc[1]
		return NUMBER
	}

	// read symbols or operators
	key, err := m.tryAllOperators()
	if err == nil {
		return key // operator found
	}

	// keywords
	key, err = m.tryAllKeywords()
	if err == nil {
		return key // keyword found
	}

	if loc := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.value.v = string(m.data[m.pos : m.pos+loc[1]])
		lval.value.t = "" // unknown for the moment
		m.pos += loc[1]
		return IDENTIFIER
	}

	m.Error("unrecognized token") // display error
	return 0
}

// try all keywords, returning their code if found.
func (m *myLexer) tryAllKeywords() (int, error) {

	for i, k := range yyToknames {
		// fmt.Printf("Trying %d %q\n", i, k)
		if m.try(k) { // Keywords are always upperCase
			// fmt.Printf("Returning %d for %q\n", i+yyPrivate-1, k)
			return i + yyPrivate - 1, nil
		}
	}

	return 0, fmt.Errorf("no keyword found")
}

// try all operators, using an internal table
func (m *myLexer) tryAllOperators() (int, error) {
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

		// multi bytes
		{"<=", LTE},
		{">=", GTE},
		{"==", EQ},
		{"!=", NEQ},
		{"++", PLUSPLUS},
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
			return k.code, nil
		}
	}
	return 0, fmt.Errorf("no operator found")
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
