package parser

import (
	"context"
	"fmt"
	"strings"
)

// =============== compiling a request ===========================

// Compile a request into an abstract syntax tree,
// invars contains the names of declared input parameters for the request.
func Compile(name string, content string) (tree Node, invars []string, err error) {
	buff := new(strings.Builder) // error writer
	lx := NewLexer(name, []byte(content), buff)
	yyParse(lx)
	if buff.Len() > 0 {
		err = fmt.Errorf(buff.String())
	}
	return lx.(*myLexer).root, lx.(*myLexer).ParamsList(), err

}

// =============== interpreting a compiled request ===========================

// Interpreter maintains context for running a compiled request.
type Interpreter struct {
	ctx    context.Context
	vars   []map[string]any // stack of frames, containing values for variables.
	invars []string         // named input variables, declared in the request.
}

// Start a new interpreter
func NewInterpreter(ctx context.Context) *Interpreter {
	it := &Interpreter{
		ctx:    ctx,
		vars:   make([]map[string]any, 0, 1),
		invars: make([]string, 0, 4),
	}
	it.pushFrame()

	return it
}

// Set input parameters for the request to interprete.
func (it *Interpreter) With(params map[string]any) *Interpreter {
	for k, v := range params {
		it.vars[0][k] = v
	}
	return it
}

// Evaluate a compiled request.
func (it *Interpreter) Eval(node Node) (any, error) {
	if node == nil {
		return nil, fmt.Errorf("cannot evaluate nil node")
	}
	return node.eval(it)
}
