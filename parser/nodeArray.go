package parser

import "fmt"

// ===== NODE ARRAY ACCESS ====

type nodeArrayAccess struct {
	a Node // array
	i Node // index
}

var _ Node = nodeArrayAccess{}

// eval implements Node.
func (n nodeArrayAccess) eval(it *Interpreter) (any, error) {
	if n.a == nil || n.i == nil {
		return nil, fmt.Errorf("cannot evaluate a nil array or index")
	}
	// evaluate array
	a, err := n.a.eval(it)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, nil
	}
	// verify a is an array
	aa, ok := a.([]any)
	if !ok {
		return nil, fmt.Errorf("expected an array, but got a %T", a)
	}
	// evaluate index
	i, err := n.i.eval(it)
	if err != nil {
		return nil, err
	}
	if i == nil {
		return nil, nil
	}
	// verify i is an int
	ii, ok := i.(int)
	if !ok {
		return nil, fmt.Errorf("expected an int index, but got a %T", i)
	}
	// verify bounds
	if ii < 0 || ii >= len(aa) {
		return nil, fmt.Errorf("index %d is out of bounds for array of length %d", ii, len(aa))
	}
	// Access array content
	return aa[ii], nil
}
