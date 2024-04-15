package parser

import "fmt"

// === IF THEN ELSE ====

type nodeIf struct {
	cond Node
	t    Node // then
	e    Node // else
}

var _ Node = nodeIf{}

func (n nodeIf) eval(it *Interpreter) (any, error) {

	// evaluate condition
	if n.cond == nil {
		return nil, fmt.Errorf("cannot evaluate a nil condition")
	}
	v, err := n.cond.eval(it)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, fmt.Errorf("expected a boolean condition, but got nil")
	}
	// verify v is a bool
	vv, ok := v.(bool)
	if !ok {
		return nil, fmt.Errorf("expected a boolean condition, but got a %T", v)
	}

	if vv {
		// evaluate then, if not nil
		if n.t != nil {
			_, err := n.t.eval(it)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
	} else {
		// evaluate else, if not nil
		if n.e != nil {
			_, err := n.e.eval(it)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
	}
	return nil, nil // do noting if did not evaluate either then or else
}
