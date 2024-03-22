package parser

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type myLexer struct {
	data []byte    // entire data to be lexed
	pos  int       // next position to process
	w    io.Writer // where are error messages written to ?
}

// lexed symbol
type yySymType struct {
	// potential values
	int    int
	string string
	bool   bool
	// not sure what this yys field is used for ?
	yys int
}

// Error implements yyLexer.
func (m *myLexer) Error(e string) {
	bef := max(0, m.pos-20)
	after := min(len(m.data), m.pos+20)
	fmt.Fprint(m.w, string(m.data[bef:m.pos]))
	fmt.Fprintf(m.w, "%s<<<%s>>>%s", ColRED, e, RESET)
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
		lval.string = string(m.data[m.pos+1 : m.pos+loc[1]-1])    // remove external quotes
		lval.string = strings.Replace(lval.string, `""`, `"`, -1) // replace all doubled quotes escaped inside.
		m.pos += loc[1]
		return STRING
	}

	// read strings between ' '
	// You can escape inside ' by adding one more.
	if loc := regexp.MustCompile(`(?s)^'(''+|[^'])*'`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.string = string(m.data[m.pos+1 : m.pos+loc[1]-1])    // remove external quotes
		lval.string = strings.Replace(lval.string, `''`, `'`, -1) // replace all doubled quotes escaped inside.
		m.pos += loc[1]
		return STRING
	}

	// read integer number.
	if loc := regexp.MustCompile(`^[+-]?[0-9]+`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.int, err = strconv.Atoi(string(m.data[m.pos : m.pos+loc[1]])) // convert to int
		if err == nil {
			m.pos += loc[1]
			return NUMBER
		}
		// continue ...
	}

	// read boolean. Use true or false, in lowercase.
	if loc := regexp.MustCompile(`^true|false`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.bool = (m.data[m.pos] == 't')
		m.pos += loc[1] // skip
		return BOOL
	}

	// read symbols
	// start with the multichar operators before single chars.
	switch {

	// start with mutibyte
	case m.try("<="):
		return LTE
	case m.try(">="):
		return GTE
	case m.try("=="):
		return EQ
	case m.try("!="):
		return NEQ
	case m.try("++"):
		return PLUSPLUS
	case m.try("--"):
		return MINUSMINUS
	case m.try("&&"), m.try("AND"):
		return AND
	case m.try("||"), m.try("OR"):
		return OR
	case m.try(".."):
		return DOTDOT
	case m.try("::"):
		return NAMESPACESEPARATOR
	case m.try("!~"):
		return REGEXNOTMATCH
	case m.try("=~"):
		return REGEXMATCH

		// single bytes
	case m.try(":"):
		return COLON
	case m.try(";"):
		return SEMICOLON
	case m.try("("):
		return LPAREN
	case m.try(")"):
		return RPAREN
	case m.try("{"):
		return LBRACE
	case m.try("}"):
		return RBRACE
	case m.try("["):
		return LBRACKET
	case m.try("]"):
		return RBRACKET
	case m.try(","):
		return COMMA
	case m.try("."):
		return DOT

	case m.try("<"):
		return LT
	case m.try(">"):
		return GT

	case m.try("*"):
		return MULTI
	case m.try("/"):
		return DIV
	case m.try("%"):
		return MOD
	case m.try("+"):
		return PLUS
	case m.try("-"):
		return MINUS
	case m.try("="):
		return ASSIGN
	case m.try("?"):
		return QUESTION

	case m.try("@"):
		return PARAM
	}

	// keywords
	key, err := m.tryAllKeywords()
	if err == nil {
		return key // keyword found
	}

	if loc := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.string = string(m.data[m.pos : m.pos+loc[1]])
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
