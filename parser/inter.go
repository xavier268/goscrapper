package parser

import (
	"context"
	"fmt"
	"strings"

	"github.com/xavier268/goscrapper/rt"
)

// =============== interpreting a compiled request ===========================

// Interpreter maintains context for running a compiled request.
type Interpreter struct {
	ctx     context.Context
	vars    []map[string]any // stack of frames, containing values for variables.
	ch      chan<- any       // channel to send results in async mode - nil, means synch mode.
	results []any            // aggregated results to be sent at the end in synch mode. Nil in async mode.
	last    bool             // last mode, only care about last result
	unique  *rt.Unique       // uniqueness filter, when distinct filters are expected
	err     error            // last error returned by a node.
}

// Start a new interpreter in default setting.
func NewInterpreter(ctx context.Context) *Interpreter {
	it := &Interpreter{
		ctx:     ctx,
		vars:    make([]map[string]any, 0, 1),
		ch:      nil,
		results: make([]any, 0, 5),
		last:    false,
		unique:  new(rt.Unique),
		err:     nil,
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

// Set asynchroneous mode. Results will be sent to channel for each loop.
// If channel is nil, sets to synchroneous mode.
func (it *Interpreter) SetAsyncMode(ch chan<- any) *Interpreter {
	if ch == nil {
		return it.SetSyncMode()
	}
	it.ch = ch
	it.results = nil
	return it
}

// Set synchroneous mode. Results will be aggregated and sent at the end.
// This is the default mode.
func (it *Interpreter) SetSyncMode() *Interpreter {
	it.ch = nil
	it.results = make([]any, 0, 5)
	return it
}

// Evaluate a compiled request.
func (it *Interpreter) Eval(node Node) (any, error) {
	if node == nil {
		return nil, fmt.Errorf("cannot evaluate nil as a compiled request")
	}

	// check expected input vars have been set
	missing := []string{}
	if np, ok := node.(nodeProgram); ok {
		for _, v := range np.invars {
			if _, ok := it.vars[0][v]; !ok {
				missing = append(missing, v)
			}
		}
		if len(missing) > 0 {
			return nil, fmt.Errorf("missing named input parameters should be provided at execution time : %s", strings.Join(missing, ", "))
		}

		// do actual evaluation
		return node.eval(it)

	} else { // not a nodeProgram
		return nil, fmt.Errorf("not a valid compiled request, cannot evaluate node of type %T", node)
	}
}
