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
()[]{}
123 c56 5de
2.56
nil

++ + - -- . .. ...  =  <= >= < > $
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
string with 2 "" escaped double quotes and 'single quotes'"
"double quote, 
multi-line,
string with 3 """ escaped double quotes and 'single quotes'"
"double quote, 
multi-line,
string with 4 """" escaped double quotes and 'single quotes'"

'single quote
with 2 single '' quotes and "double" quotes'
'single quote
with 3 single ''' quotes and "double" quotes'
'single quote
with 4 single '''' quotes and "double" quotes'


BOOL
NUMBER
STRING 
IDENTIFIER

ASSIGN SEMICOLON CLICK INPUT IN 
SLOW LEFT RIGHT MIDDLE 
RETURN COMMA FOR
SELECT AS FROM TO STEP
WHERE LIMIT
LPAREN RPAREN
LBRACKET RBRACKET
LBRACE RBRACE
DOT LEN 
PLUS MINUS PLUSPLUS MINUSMINUS MULTI DIV MOD ABS
NOT AND OR XOR NAND
EQ NEQ LT LTE GT GTE
CONTAINS
FIND PATH WITH JOIN PAGE
COLON TEXT ATTR OF
DISTINCT 
AT /* @ */
DOTDOT /* .. */
QUESTION /*?*/

NOW VERSION FILE_SEPARATOR

IF THEN ELSE

ASSERT FAIL

PRINT FORMAT RAW GO JSON GSC NL

DOLLAR NIL
LAST

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
