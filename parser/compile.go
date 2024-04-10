package parser

import (
	"fmt"
	"strings"
)

// =============== compiling a request ===========================

// Compile a request.
// Once compiled, a  request can be executed multiple times in different Interpreter contexts.
func Compile(name string, content string) (compiledReq Node, err error) {
	buff := new(strings.Builder) // error writer
	lx := NewLexer(name, []byte(content), buff)
	yyParse(lx)
	if buff.Len() > 0 {
		err = fmt.Errorf(buff.String())
	}
	return lx.(*myLexer).root, err
}
