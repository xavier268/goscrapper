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
		lval.value = string(m.data[m.pos+1 : m.pos+loc[1]-1])   // remove external quotes
		lval.value = strings.Replace(lval.value, `""`, `"`, -1) // replace all doubled quotes escaped inside.
		lval.value = fmt.Sprintf(`%q`, lval.value)              // store as a quoted string
		m.pos += loc[1]
		return STRING
	}

	// read strings between ' '
	// You can escape inside ' by adding one more.
	if loc := regexp.MustCompile(`(?s)^'(''+|[^'])*'`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.value = string(m.data[m.pos+1 : m.pos+loc[1]-1])   // remove external quotes
		lval.value = strings.Replace(lval.value, `''`, `'`, -1) // replace all doubled quotes escaped inside.
		lval.value = fmt.Sprintf(`%q`, lval.value)              // stored as a back-quoted string
		m.pos += loc[1]
		return STRING
	}

	// read integer number.
	if loc := regexp.MustCompile(`^[+-]?[0-9]+`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.value = string(m.data[m.pos : m.pos+loc[1]])
		m.pos += loc[1]
		return NUMBER
	}

	// read boolean. Use true or false, in lowercase.
	if loc := regexp.MustCompile(`^true|false`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.value = string(m.data[m.pos : m.pos+loc[1]])
		m.pos += loc[1] // skip
		return BOOL
	}

	// read symbols
	// start with the multichar operators before single chars.
	switch {

	// start with mutibyte
	case m.try("<="):
		lval.value = "<="
		return LTE
	case m.try(">="):
		lval.value = ">="
		return GTE
	case m.try("=="):
		lval.value = "=="
		return EQ
	case m.try("!="):
		lval.value = "!="
		return NEQ
	case m.try("++"):
		lval.value = "++"
		return PLUSPLUS
	case m.try("--"):
		lval.value = "--"
		return MINUSMINUS
	case m.try("&&"), m.try("AND"):
		lval.value = "&&"
		return AND
	case m.try("||"), m.try("OR"):
		lval.value = "||"
		return OR
	case m.try(".."):
		lval.value = ".."
		return DOTDOT
	case m.try("::"):
		lval.value = "::"
		return NAMESPACESEPARATOR
	case m.try("!~"):
		lval.value = "!~"
		return REGEXNOTMATCH
	case m.try("=~"):
		lval.value = "=~"
		return REGEXMATCH

		// single bytes
	case m.try(":"):
		lval.value = ":"
		return COLON
	case m.try(";"):
		lval.value = ";"
		return SEMICOLON
	case m.try("("):
		lval.value = "("
		return LPAREN
	case m.try(")"):
		lval.value = ")"
		return RPAREN
	case m.try("{"):
		lval.value = "{"
		return LBRACE
	case m.try("}"):
		lval.value = "}"
		return RBRACE
	case m.try("["):
		lval.value = "["
		return LBRACKET
	case m.try("]"):
		lval.value = "]"
		return RBRACKET
	case m.try(","):
		lval.value = ","
		return COMMA
	case m.try("."):
		lval.value = "."
		return DOT

	case m.try("<"):
		lval.value = "<"
		return LT
	case m.try(">"):
		lval.value = ">"
		return GT

	case m.try("*"):
		lval.value = "*"
		return MULTI
	case m.try("/"):
		lval.value = "/"
		return DIV
	case m.try("%"):
		lval.value = "%"
		return MOD
	case m.try("+"):
		lval.value = "+"
		return PLUS
	case m.try("-"):
		lval.value = "-"
		return MINUS
	case m.try("="):
		lval.value = "="
		return ASSIGN
	case m.try("?"):
		lval.value = "?"
		return QUESTION
	case m.try("@"):
		lval.value = "@"
		return AT
	}

	// keywords
	key, err := m.tryAllKeywords()
	if err == nil {
		return key // keyword found
	}

	if loc := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.value = string(m.data[m.pos : m.pos+loc[1]])
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
