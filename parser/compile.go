package parser

import (
	"context"
	"fmt"
	"io"
	"os"
	"sort"
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

// Extracts the (sorted) list of parameters a compiled requests expects.
func GetParamsList(compiledRequest Node) []string {

	switch cr := compiledRequest.(type) {
	case nil:
		return []string{}
	case nodeProgram:
		ret := cr.invars
		sort.Strings(ret)
		return ret
	default:
		return []string{}
	}
}

// =============== compile and evaluate at once ===========================

// Compile a request, create a defalt interpreter, and evaluate the request.
func Eval(name string, content string) (result any, err error) {
	compiledReq, err := Compile(name, content)
	if err != nil {
		return nil, err
	}
	it := NewInterpreter(context.Background())
	return it.Eval(compiledReq)
}

// Same a Eval, but read request from fileName.
func EvalFile(fname string) (result any, err error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return Eval(fname, string(content))
}

// Same a Eval, but with parameters map
func EvalWithParams(name string, content string, params map[string]any) (result any, err error) {
	compiledReq, err := Compile(name, content)
	if err != nil {
		return nil, err
	}
	it := NewInterpreter(context.Background()).With(params)
	return it.Eval(compiledReq)
}

// Eval from file with parameter map
func EvalFileWithParams(fname string, params map[string]any) (result any, err error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return EvalWithParams(fname, string(content), params)
}
