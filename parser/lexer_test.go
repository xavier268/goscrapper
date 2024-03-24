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
 	< <= .. .
	truefalsetotointboolstring
	true false toto int bool string
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
		fmt.Printf("%s   %#v\n", TokenAsString(tok), lval.value)
		if tok == 0 {
			break
		}
	}

	// Output:
	// IDENTIFIER   parser.value{v:"un", t:"IDENTIFIER", c:57415}
	// IDENTIFIER   parser.value{v:"deux22", t:"IDENTIFIER", c:57415}
	// LT   parser.value{v:"<", t:"LT", c:57354}
	// LTE   parser.value{v:"<=", t:"LTE", c:57352}
	// DOTDOT   parser.value{v:"..", t:"DOTDOT", c:57371}
	// DOT   parser.value{v:".", t:"DOT", c:57360}
	// BOOL   parser.value{v:"true", t:"BOOL", c:57413}
	// BOOL   parser.value{v:"false", t:"BOOL", c:57413}
	// IDENTIFIER   parser.value{v:"totointboolstring", t:"IDENTIFIER", c:57415}
	// BOOL   parser.value{v:"true", t:"BOOL", c:57413}
	// BOOL   parser.value{v:"false", t:"BOOL", c:57413}
	// IDENTIFIER   parser.value{v:"toto", t:"IDENTIFIER", c:57415}
	// INTTYPE   parser.value{v:"int", t:"INTTYPE", c:57378}
	// BOOLTYPE   parser.value{v:"bool", t:"BOOLTYPE", c:57379}
	// STRINGTYPE   parser.value{v:"string", t:"STRINGTYPE", c:57380}
	// NUMBER   parser.value{v:"1", t:"int", c:57418}
	// NUMBER   parser.value{v:"2", t:"int", c:57418}
	// IDENTIFIER   parser.value{v:"cinq", t:"IDENTIFIER", c:57415}
	// NUMBER   parser.value{v:"0555", t:"int", c:57418}
	// DOTDOT   parser.value{v:"..", t:"DOTDOT", c:57371}
	// DOT   parser.value{v:".", t:"DOT", c:57360}
	// STRING   parser.value{v:"\"a dq string\\n\\taccross the line \"", t:"string", c:57417}
	// STRING   parser.value{v:"\"a sq string\"", t:"string", c:57417}
	// STRING   parser.value{v:"\"dq with escaped \\\"\\\" \"", t:"string", c:57417}
	// STRING   parser.value{v:"\"sq with escape ' \"", t:"string", c:57417}
	// STRING   parser.value{v:"\"string containing\\n\\t// a comment\"", t:"string", c:57417}
	// TOK-0   parser.value{v:"", t:"", c:0}

}
