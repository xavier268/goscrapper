package parser

import "fmt"

// compiler checks
var _ yyLexer = new(MyLexer)

func ExampleNewLexer() {

	dd := ` 
	un deux22 // trois quatre
	// nothing here .. 

	1 2cinq -- /* do not
	print this */ 0555 .. .
	"a dq string
	accross the line "
	'a sq string'
	"dq with escaped """ "
	'sq with escape '' '
	"string containing
	// a comment"


	`

	lx := NewLexer([]byte(dd))

	for {
		lval := new(yySymType)
		tok := lx.Lex(lval)
		if tok == 0 {
			break
		}
		//                TokenTypeAsString,  string value, number value
		fmt.Printf("%s   %q  %d\n", TokenAsString(tok), lval.string, lval.int)
	}

	// Output:
	// IDENTIFIER   "un"  0
	// IDENTIFIER   "deux22"  0
	// NUMBER   ""  1
	// NUMBER   ""  2
	// IDENTIFIER   "cinq"  0
	// MINUSMINUS   ""  0
	// NUMBER   ""  555
	// DOTDOT   ""  0
	// DOT   ""  0
	// STRING   "a dq string\n\taccross the line "  0
	// STRING   "a sq string"  0
	// STRING   "dq with escaped \"\" "  0
	// STRING   "sq with escape ' "  0
	// STRING   "string containing\n\t// a comment"  0

}
