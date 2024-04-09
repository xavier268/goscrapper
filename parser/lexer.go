package parser

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// myLexer maintains context needed to parse the data into a Node.
type myLexer struct {
	name   string    // name of the lexer source, user defined
	data   []byte    // entire data to be lexed
	pos    int       // next position to process
	ew     io.Writer // where shall lexer errors be written to ?
	root   Node      // root node for the parsed tree
	params []string  // list of declared input parameters
}

// Minimal Lexer interface.
type Lexer = yyLexer

// Construct new myLexer.
// Use errorWriter to capture detailled error messages.
// If nil, errors will appear on stdout.
func NewLexer(name string, data []byte, errorWriter io.Writer) Lexer {
	return &myLexer{
		name: name,
		data: data,
		pos:  0,
		ew:   errorWriter,
	}
}

// Error implements yyLexer.
func (m *myLexer) Error(e string) {
	if m.ew == nil {
		m.ew = os.Stdout
	}
	windows := 100
	// write to std error writer lx.ew
	fmt.Fprintf(m.ew, "\n%s ********* Error in %s :***************%s\n\n", ColRED, m.name, RESET)
	bef := max(0, m.pos-windows)
	bfs := strings.Split(string(m.data[bef:m.pos]), "\n")
	if len(bfs) > 1 {
		bfs = bfs[1:] // remove potentially incomplete first line
	}

	after := min(len(m.data), m.pos+windows)
	afs := strings.Split(string(m.data[m.pos:after]), "\n")

	// print until the line containing error
	last := 0
	for _, li := range bfs {
		fmt.Fprintf(m.ew, "\n%s%s", ColYELLOW, li)
		last = len(li)
	}
	fmt.Fprintf(m.ew, "%s%s\n", ColYELLOW, afs[0])
	// last points to error position
	last = max(0, last-2)
	fmt.Fprintf(m.ew, " %s%s^ %s %s\n", ColRED, strings.Repeat(" ", last), e, RESET)
	// print rest of context
	fmt.Fprintf(m.ew, "%s%s\n", ColYELLOW, strings.Join(afs[1:], "\n"))
	fmt.Fprintln(m.ew, RESET)
}

// utility to format error message sent to myLexer.Error()
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
		lval.tok.v = string(m.data[m.pos+1 : m.pos+loc[1]-1])   // remove external quotes
		lval.tok.v = strings.Replace(lval.tok.v, `""`, `"`, -1) // replace all doubled quotes escaped inside.
		lval.tok.v = fmt.Sprintf(`%q`, lval.tok.v)              // store as a quoted string
		lval.tok.t = "string"
		lval.tok.c = STRING
		m.pos += loc[1]
		return STRING
	}

	// read strings between ' '
	// You can escape inside ' by adding one more.
	if loc := regexp.MustCompile(`(?s)^'(''+|[^'])*'`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.tok.v = string(m.data[m.pos+1 : m.pos+loc[1]-1])   // remove external quotes
		lval.tok.v = strings.Replace(lval.tok.v, `''`, `'`, -1) // replace all doubled quotes escaped inside.
		lval.tok.v = fmt.Sprintf(`%q`, lval.tok.v)              // stored as a back-quoted string
		lval.tok.t = "string"
		lval.tok.c = STRING
		m.pos += loc[1]
		return STRING
	}

	// read positive integer number.
	if loc := regexp.MustCompile(`^[0-9]+`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.tok.v = string(m.data[m.pos : m.pos+loc[1]])
		lval.tok.t = "int"
		lval.tok.c = NUMBER
		m.pos += loc[1]
		return NUMBER
	}

	// read symbols or operators
	err = m.tryAllOperators(lval)
	if err == nil {
		return lval.tok.c // operator found
	}

	// keywords
	err = m.tryAllKeywords(lval)
	if err == nil {
		return lval.tok.c // keyword found
	}

	if loc := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*`).FindIndex(m.data[m.pos:]); len(loc) == 2 {
		lval.tok.v = string(m.data[m.pos : m.pos+loc[1]])
		lval.tok.t = "IDENTIFIER"
		lval.tok.c = IDENTIFIER
		m.pos += loc[1]
		return IDENTIFIER
	}

	m.errorf("unrecognized token : <%s>", m.data[m.pos:min(len(m.data), m.pos+20)]) // display error
	return 0
}

// try all keywords, returning their code if found.
func (m *myLexer) tryAllKeywords(lval *yySymType) error {

	// special care taken to ensure INPUT is never recognized as IN
	// by testing longer keywords first !
	size := 0
	for _, t := range yyToknames {
		size = max(size, len(t))
	}

	for s := size; s > 0; s-- {
		// test by decreasing token length
		for i, k := range yyToknames {
			// only look at s-sized token
			if len(k) == s && m.try(k) { // Keywords are always upperCase
				// fmt.Printf("Returning %d for %q\n", i+yyPrivate-1, k)
				lval.tok.t = k
				lval.tok.c = i + yyPrivate - 1
				lval.tok.v = k
				return nil
			}
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
		// {"int", INTTYPE},
		// {"bool", BOOLTYPE},
		// {"string", STRINGTYPE},
		// {"bin", BINTYPE},
		//{"nil", NIL},

		// multi bytes
		{"<=", LTE},
		{">=", GTE},
		{"==", EQ},
		{"!=", NEQ},
		{"++", PLUSPLUS},
		{"&&", AND},
		{"||", OR},
		{"..", DOTDOT},
		// {"::", NAMESPACESEPARATOR},
		// {"!~", REGEXNOTMATCH},
		// {"=~", REGEXMATCH},

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
		{"!", BANG},
	}
	for _, k := range opeTable {
		// fmt.Printf("Trying %d %q\n", i, k)
		if m.try(k.ope) {
			lval.tok.t = TokenAsString(k.code)
			lval.tok.c = k.code
			lval.tok.v = k.ope
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

// Prints a token defined by its constant value as a string.
func TokenAsString(t int) string {
	idx := t - yyPrivate + 1 // yyPrivate points to error
	if idx < len(yyToknames) && idx >= 0 {
		return yyToknames[idx]
	} else {
		return fmt.Sprintf("TOK-%d", t)
	}
}
