// contains lexer, grammar, and parser generation tools
package parser

//go:generate go install golang.org/x/tools/cmd/goyacc@latest
//go:generate goyacc -o parser.go grammar.y

func init() {
	// set debugging level to max
	yyDebug = 3
}
