// compile and execute GSC requests
package parser

//go:generate go install golang.org/x/tools/cmd/goyacc@latest
//go:generate goyacc -o parser.go grammar.y

func init() {
	// set debugging
	yyDebug = 1
	yyErrorVerbose = true
}
