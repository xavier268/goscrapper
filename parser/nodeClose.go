package parser

import (
	"fmt"

	"github.com/xavier268/goscrapper/rt"
)

type nodeClose struct {
	node Node
}

var _ Node = nodeClose{}

// Close if expression can be closed.
// Ignore nil. Else, fail.
func (n nodeClose) eval(it *Interpreter) (any, error) {
	if n.node == nil {
		return nil, nil
	}
	v, err := n.node.eval(it)
	if err != nil {
		return nil, err
	}

	switch v := v.(type) {
	case nil:
		return nil, nil
	case *rt.Page:
		v.Close()
	default:
		fmt.Printf("DEBUG : unexpected CLOSE type %T", v)
	}
	return nil, nil
}
