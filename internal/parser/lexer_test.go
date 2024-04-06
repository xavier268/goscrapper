package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xavier268/mytest"
)

// compiler checks
var _ yyLexer = new(myLexer)

func TestLexer(t *testing.T) {

	dd := ` 
	un deux22 // trois quatre
	// nothing here .. 
 	< <= .. .
	truefalsetotointboolstring
	true false toto int bool string
	1 2cinq  /* do not
	print this */ 0555 .. .
	"a dq string
	across the line "
	'a sq string'
	"dq with escaped """ "
	'sq with escape '' '
	"string containing
	// a comment"
	PAGE @ PLUS RETURN
	int bool bin string
	LEFT RIGHT MIDDLE INPUT CLICK IN FROM SLOW
	`

	buff := new(strings.Builder)

	lx := &myLexer{
		data: []byte(dd),
		pos:  0,
		w:    buff,
	}

	for {
		lval := new(yySymType)
		tok := lx.Lex(lval)

		//                TokenTypeAsString,  string value, number value
		fmt.Fprintf(buff, "%s   %#v\n", TokenAsString(tok), lval.value)
		if tok == 0 {
			break
		}
	}

	mytest.Verify(t, dd+"\n\n"+buff.String(), "lexer_test")

}
