package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xavier268/mytest"
)

func TestLexer(t *testing.T) {
	data := `
// test for lexer
/* block 
comment */
BOOL
NUMBER
STRING 
IDENTIFIER

ASSIGN SEMICOLON CLICK INPUT IN 
PRINT RAW SLOW LEFT RIGHT MIDDLE 
RETURN COMMA FOR
SELECT AS FROM
WHERE LIMIT
LPAREN RPAREN
LBRACKET RBRACKET
LBRACE RBRACE
DOT LEN 
PLUS MINUS PLUSPLUS MINUSMINUS MULTI DIV MOD ABS
NOT AND OR XOR
EQ NEQ LT LTE GT GTE
CONTAINS
FIND PATH WITH JOIN PAGE
COLON TEXT ATTR OF
DISTINCT 
AT /* @ */
DOTDOT /* .. */
QUESTION
()[]{}
123 c56 5de
2.56

++ + - -- . .. ...  =  <= >= < > 
= != @ !

&& || ;:%

// testing simple strings
"un"
"  un  "
'deux'
' deux '
"trois & quatre"
'cinq et six'

// testing complex strings
"double quote, 
multi-line,
string with """ escaped double quotes and 'single quotes'"

'single quote
with escaped single ''' quotes and "double" quotes'

EndOfTest
	`
	buff := new(strings.Builder)
	buff.WriteString(data + "\n")

	lx := NewLexer(t.Name(), []byte(data), buff)
	for {
		lval := new(yySymType)
		c := lx.Lex(lval)
		fmt.Fprintf(buff, "%#v\n", lval)
		if lval.tok.c == 0 || c == 0 {
			break
		}

	}
	mytest.Verify(t, buff.String(), t.Name())
}
