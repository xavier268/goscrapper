package parser

import (
	"fmt"
	"regexp"
)

type yySymType struct {
	start int
	end   int
	text  string
}
type yyLexer interface {
	Lex(lval *yySymType) int
	Error(e string)
}

// Lexer implementation
type myLexer struct {
	data []byte // entire data to be lexed
	pos  int    // next position to process
}

func NewLexer(data []byte) *myLexer {
	return &myLexer{data: data, pos: 0}
}

// Error implements yyLexer.
func (m *myLexer) Error(e string) {
	bef := max(0, m.pos-20)
	after := min(len(m.data), m.pos+20)
	fmt.Print(string(m.data[bef:m.pos]))
	fmt.Printf("%s<<<%s>>>%s", ColRED, e, RESET)
	fmt.Println(string(m.data[m.pos:after]))
}

// Lex implements yyLexer.
func (m *myLexer) Lex(lval *yySymType) int {

startLoop:

	// check if we reached the end
	if m.pos >= len(m.data) {
		return 0
	}

	// skip white spaces and terminators
	if loc := regexp.MustCompile(`^[\s]+`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		m.pos = loc[1] // skip
		goto startLoop
	}

	// skip multiline comments
	if loc := regexp.MustCompile(`^\/\*.*\*\/+`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		m.pos = loc[1] // skip
		goto startLoop
	}

	// read strings
	if loc := regexp.MustCompile(`^"[^"]*"`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.start = m.pos
		lval.end = loc[1]
		m.pos = loc[1]
		lval.end = m.pos
		lval.text = string(m.data[lval.start:lval.end])
		return STRING
	}
	err := "cannot lex beyond this point"
	m.Error(err) // display error
	return 0
}

const (
	STRING = iota + 1
)
