package parser

import (
	"fmt"
	"os"
)

// compiler checks
var _ yyLexer = new(myLexer)

func Example_lexer() {

	dd := ` 
	un deux22 // trois quatre
	// nothing here .. 

	1 2cinq  /* do not
	print this */ 0555 .. .
	"a dq string
	accross the line "
	'a sq string'
	"dq with escaped """ "
	'sq with escape '' '
	"string containing
	// a comment"

	`

	lx := &myLexer{
		data: []byte(dd),
		pos:  0,
		w:    os.Stdout,
	}

	for {
		lval := new(yySymType)
		tok := lx.Lex(lval)

		//                TokenTypeAsString,  string value, number value
		fmt.Printf("%s   %s\n", TokenAsString(tok), lval.value)
		if tok == 0 {
			break
		}
	}

	// Output:
	// IDENTIFIER   un
	// IDENTIFIER   deux22
	// NUMBER   1
	// NUMBER   2
	// IDENTIFIER   cinq
	// NUMBER   0555
	// DOTDOT   ..
	// DOT   .
	// STRING   "a dq string\n\taccross the line "
	// STRING   "a sq string"
	// STRING   "dq with escaped \"\" "
	// STRING   "sq with escape ' "
	// STRING   "string containing\n\t// a comment"
	// TOK-0
}
