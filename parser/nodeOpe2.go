package parser

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/xavier268/goscrapper/rt"
)

// ==== Binary operators ====

type nodeOpe2 struct {
	operator int  // code for operation
	left     Node // left operand
	right    Node // right operand
}

var _ Node = nodeOpe2{}

func (m *myLexer) newNodeOpe2(left Node, op Node, right Node) nodeOpe2 {
	if ope2, ok := op.(tok); ok {
		return nodeOpe2{
			operator: ope2.c,
			left:     left,
			right:    right,
		}
	}
	m.errorf("error trying to contruct binary operator node %#v %#v %#v", left, op, right)
	return nodeOpe2{}
}

// eval implements Node.
func (n nodeOpe2) eval(it *Interpreter) (any, error) {
	var left, right any
	var err error

	// evaluate both arguments
	if n.left != nil {
		left, err = n.left.eval(it)
		if err != nil {
			return nil, err
		}
	}
	if n.right != nil {
		right, err = n.right.eval(it)
		if err != nil {
			return nil, err
		}
	}

	// apply operator
	switch n.operator {
	case PLUS:
		switch v := left.(type) {
		case int:
			switch w := right.(type) {
			case int:
				return v + w, nil
			default:
				return nil, fmt.Errorf("cannot apply binary %s to %T and %T", TokenAsString(n.operator), v, w)
			}
		case string:
			switch w := right.(type) {
			case string:
				return v + w, nil
			default:
				return nil, fmt.Errorf("cannot apply binary %s to %T and %T", TokenAsString(n.operator), v, w)
			}
		case []any: // add element to array
			return append(v, right), nil
		default:
			return nil, fmt.Errorf("cannot apply binary %s to %T", TokenAsString(n.operator), v)
		}
	case PLUSPLUS:
		switch v := left.(type) {
		case []any:
			switch w := right.(type) {
			case []any:
				return append(v, w...), nil
			default:
				return nil, fmt.Errorf("expected two arrays, cannot apply binary %s to %T and %T", TokenAsString(n.operator), v, w)
			}
		default:
			return nil, fmt.Errorf("cannot apply binary %s to %T", TokenAsString(n.operator), v)
		}
	case MINUS:
		switch v := left.(type) {
		case int:
			switch w := right.(type) {
			case int:
				return v - w, nil
			default:
				return nil, fmt.Errorf("cannot apply binary %s to %T and %T", TokenAsString(n.operator), v, w)
			}
		default:
			return nil, fmt.Errorf("cannot apply binary %s to %T", TokenAsString(n.operator), v)
		}
	case MULTI:
		switch v := left.(type) {
		case int:
			switch w := right.(type) {
			case int:
				return v * w, nil
			default:
				return nil, fmt.Errorf("cannot apply binary %s to %T and %T", TokenAsString(n.operator), v, w)
			}
		default:
			return nil, fmt.Errorf("cannot apply binary %s to %T", TokenAsString(n.operator), v)
		}
	case DIV:
		switch v := left.(type) {
		case int:
			switch w := right.(type) {
			case int:
				return v / w, nil
			default:
				return nil, fmt.Errorf("cannot apply binary %s to %T and %T", TokenAsString(n.operator), v, w)
			}
		default:
			return nil, fmt.Errorf("cannot apply binary %s to %T", TokenAsString(n.operator), v)
		}
	case MOD:
		switch v := left.(type) {
		case int:
			switch w := right.(type) {
			case int:
				return v % w, nil
			default:
				return nil, fmt.Errorf("cannot apply binary %s to %T and %T", TokenAsString(n.operator), v, w)
			}
		default:
			return nil, fmt.Errorf("cannot apply binary %s to %T", TokenAsString(n.operator), v)
		}

	case EQ:
		return reflect.DeepEqual(left, right), nil
	case NEQ:
		return !reflect.DeepEqual(left, right), nil

	case LT:
		i, err := CompareAny(left, right)
		return i < 0, err
	case LTE:
		i, err := CompareAny(left, right)
		return i <= 0, err
	case GT:
		i, err := CompareAny(left, right)
		return i > 0, err
	case GTE:
		i, err := CompareAny(left, right)
		return i >= 0, err

	case ATTR: // elem ATTR att
		if left == nil || right == nil {
			return nil, fmt.Errorf("invalid nil argument to %s operator", TokenAsString(n.operator))
		}
		if v, ok := left.(*rt.Element); ok {
			if w, ok := right.(string); ok {
				return rt.GetAttribute(v, w), nil
			}
		}
		return nil, fmt.Errorf("cannot apply binary %s to %T and %T", TokenAsString(n.operator), left, right)

	case FORMAT: // any FORMAT format
		// check format is a non nil string
		if right == nil {
			return nil, fmt.Errorf("invalid nil format argument to %s operator", TokenAsString(n.operator))
		}
		ft, ok := right.(string)
		if !ok {
			return nil, fmt.Errorf("cannot use a non string format : %#v", right)
		}
		// format
		return rt.SafeSprintf(ft, left)

	// nil CONTAINS nil (but nothing else)
	// string CONTAINS substring
	// array CONTAINS any
	case CONTAINS:
		if left == nil {
			return (right == nil), nil
		}
		if v, ok := left.(string); ok {
			if w, ok := right.(string); ok {
				return strings.Contains(v, w), nil
			}
		}
		if v, ok := left.([]any); ok {
			return slices.Contains(v, right), nil
		}
		return nil, fmt.Errorf("cannot apply binary %s to %T and %T", TokenAsString(n.operator), left, right)

	default:
		return nil, fmt.Errorf("unknown binary operator %s", TokenAsString(n.operator))
	}

}

// === binary bool operator ===

type nodeOpe2Bool struct {
	operator int // code for operation
	left     Node
	right    Node
}

var _ Node = nodeOpe2Bool{}

func (m *myLexer) newNodeOpe2Bool(left, op, right Node) nodeOpe2Bool {
	if ope2, ok := op.(tok); ok {
		return nodeOpe2Bool{
			operator: ope2.c,
			left:     left,
			right:    right,
		}
	}
	m.errorf("error trying to contruct binary bool operator node %#v %#v %#v", left, op, right)
	return nodeOpe2Bool{}
}

// eval implements Node.
// DECISION : do not use lazy evaluation of second argument, since it would generate hard to diagnose
// run time errors if types are mismatched.
func (n nodeOpe2Bool) eval(it *Interpreter) (any, error) {

	var left, right any
	var err error
	var lb, rb, okl, okr bool

	// evaluate both arguments
	if n.left != nil {
		left, err = n.left.eval(it)
		if err != nil {
			return nil, err
		}
	}
	if n.right != nil {
		right, err = n.right.eval(it)
		if err != nil {
			return nil, err
		}
	}
	if left == nil || right == nil {
		return nil, fmt.Errorf("invalid nil argument to boolean operator")
	}
	lb, okl = left.(bool)
	rb, okr = right.(bool)
	if !okl || !okr {
		return nil, fmt.Errorf("invalid non boolean arguments to boolean operator : %T and %T", left, right)
	}
	switch n.operator {
	case AND:
		return lb && rb, nil
	case OR:
		return lb || rb, nil
	case XOR:
		return lb != rb, nil
	case NAND:
		return !(lb && rb), nil
	default:
		return nil, fmt.Errorf("unknown boolean binary operator %s", TokenAsString(n.operator))
	}
}
