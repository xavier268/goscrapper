package parser

import (
	"context"
	"fmt"
	"strings"
)

// =============== compiling a request ===========================

// a Node (and the whole parsed tree) should always evaluate to something.
type Node interface {
	eval(*Interpreter) (any, error)
}

func Compile(name string, content string) (rootTree Node, invars []string, err error) {
	buff := new(strings.Builder) // error writer
	lx := NewLexer(name, []byte(content), buff)
	yyParse(lx)
	if buff.Len() > 0 {
		err = fmt.Errorf(buff.String())
	}
	return lx.(*myLexer).root, lx.(*myLexer).params, err

}

// =============== interpreting a compiled request ===========================

// Interpreter maintains context for running a compiled request.
type Interpreter struct {
	ctx    context.Context
	vars   []map[string]any // stack of frames, containing values for variables.
	invars []string         // named input variables, declared in the request.
}

// start a new interpreter, passing input variables.
// it is ok to pass more or less input variables as declared in the request.
// extra will be ignored, others will be nil.
func NewInterpreter(ctx context.Context, params map[string]any) *Interpreter {
	it := &Interpreter{
		ctx:    ctx,
		vars:   make([]map[string]any, len(params)),
		invars: make([]string, len(params)),
	}
	it.pushFrame()
	// initialize input in global frame
	for k, v := range params {
		it.invars = append(it.invars, k)
		it.vars[0][k] = v
	}
	return it
}

func (it *Interpreter) Eval(node Node) (any, error) {
	if node == nil {
		return nil, fmt.Errorf("cannot evaluate nil node")
	}
	return node.eval(it)
}
