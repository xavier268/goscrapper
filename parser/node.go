package parser

import (
	"fmt"
)

var _ Node = tok{}
var _ Node = Nodes{}

// Node for the abstract syntax tree.
type Node interface {
	eval(*Interpreter) (any, error)
}

var _ Node = tok{}

// eval implements Node for token.
// Should never be called. Always eval to nil.
func (t tok) eval(*Interpreter) (any, error) {
	return nil, fmt.Errorf("tok.eval should never be called")
}

// Node with a body (eg : loop statement)
type NodeWithBody interface {
	Node
	appendBody(nodes ...Node) NodeWithBody // return a copy of the node, with selected nodes added to body
}

// ===== Nodes =====

// A list of Node is also a Node.
// There are evaluated sequentially.
// Evaluation stops on error or context cancelation.
type Nodes []Node

func (ns Nodes) eval(i *Interpreter) (any, error) {
	var ret = make([]any, 0, len(ns))
	for _, v := range ns {
		if v == nil {
			continue
		}
		if i.ctx.Err() != nil {
			return nil, i.ctx.Err()
		}
		r, err := v.eval(i)
		if err != nil {
			return nil, err
		}
		ret = append(ret, r)
	}
	return ret, nil
}

var _ Node = Nodes{}

// ===  top level program node ===

type nodeProgram struct {
	req    Node
	invars []string // list of externally provided variables, also called parameters.
}

var _ Node = nodeProgram{}

func (n nodeProgram) eval(it *Interpreter) (any, error) {
	if n.req == nil {
		return nil, fmt.Errorf("cannot evaluate a nil request")
	}

	// evaluate program content
	_, err := n.req.eval(it)

	// If async mode ...
	if it.ch != nil {
		if it.last {
			// send last result to channel now, nothing could be sent before ...
			select {
			case it.ch <- it.results:
				return nil, nil
			case <-it.ctx.Done():
				return nil, it.ctx.Err()
			}
		} else {
			return nil, err // in async mode without last mode, results were already sent
		}
	} else {
		return it.results, err // send aggregated results or just last result
	}
}
