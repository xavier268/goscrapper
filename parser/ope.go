package parser

import (
	"fmt"
	"path/filepath"
	"reflect"
	"time"

	"github.com/xavier268/goscrapper"
	"github.com/xavier268/goscrapper/rt"
)

// ==== No argument operators ====
type nodeOpe0 tok

var _ Node = nodeOpe0{}

// eval implements Node.
func (n nodeOpe0) eval(it *Interpreter) (rv any, err error) {
	switch n.c {
	case NOW:
		return time.Now(), nil
	case VERSION:
		return goscrapper.VERSION, nil
	case FILE_SEPARATOR:
		return string(filepath.Separator), nil
	default:
		return nil, fmt.Errorf("unknown zero-ary operator %s", TokenAsString(n.c))
	}
}

// ==== Unary operators ====

type nodeOpe1 struct {
	operator int  // code for operation
	right    Node // right operand
}

func (m *myLexer) newNodeOpe1(op Node, right Node) nodeOpe1 {
	if ope1, ok := op.(tok); ok {
		return nodeOpe1{
			operator: ope1.c,
			right:    right,
		}
	}
	m.errorf("error trying to contruct unary operator node %#v %#v", op, right)
	return nodeOpe1{}
}

var _ Node = nodeOpe1{}

// eval implements Node.
func (n nodeOpe1) eval(it *Interpreter) (rv any, err error) {

	// evaluate argument
	if n.right != nil {
		rv, err = n.right.eval(it)
		if err != nil {
			return nil, err
		}
	}

	// apply operator
	switch n.operator {
	case PLUS:
		switch v := rv.(type) {
		case int:
			return v, nil
		case nil:
			return nil, nil
		default:
			return nil, fmt.Errorf("cannot apply unary %s to %T", TokenAsString(n.operator), rv)
		}
	case MINUS:
		switch v := rv.(type) {
		case int:
			return -v, nil
		case nil:
			return nil, nil
		default:
			return nil, fmt.Errorf("cannot apply unary %s to %T", TokenAsString(n.operator), rv)
		}
	case PLUSPLUS:
		switch v := rv.(type) {
		case int:
			return v + 1, nil
		default:
			return nil, fmt.Errorf("cannot apply unary %s to %T", TokenAsString(n.operator), rv)
		}
	case MINUSMINUS:
		switch v := rv.(type) {
		case int:
			return v - 1, nil
		default:
			return nil, fmt.Errorf("cannot apply unary %s to %T", TokenAsString(n.operator), rv)
		}
	case NOT:
		switch v := rv.(type) {
		case bool:
			return !v, nil
		default:
			return nil, fmt.Errorf("cannot apply unary %s to %T", TokenAsString(n.operator), rv)
		}
	case LEN:
		if rv == nil {
			return 0, nil
		}
		if v, ok := rv.(string); ok {
			return len(v), nil
		}
		if reflect.TypeOf(rv).Kind() == reflect.Slice {
			return reflect.ValueOf(rv).Len(), nil
		}
		if reflect.TypeOf(rv).Kind() == reflect.Map {
			return reflect.ValueOf(rv).Len(), nil
		}
		return nil, fmt.Errorf("cannot apply unary %s to %T", TokenAsString(n.operator), rv)

	case ABS:
		switch v := rv.(type) {
		case int:
			if v < 0 {
				return -v, nil
			}
			return v, nil
		default:
			return nil, fmt.Errorf("cannot apply unary %s to %T", TokenAsString(n.operator), rv)
		}
	case PAGE:
		switch url := rv.(type) {
		case string:
			return rt.GetPage(it.ctx, url), nil
		default:
			return nil, fmt.Errorf("cannot apply unary %s to %T, expected string", TokenAsString(n.operator), rv)
		}
	case TEXT:
		switch elem := rv.(type) {
		case *rt.Element:
			return rt.GetElemText(elem), nil
		case *rt.Page:
			return rt.GetPageText(elem), nil
		default:
			return nil, fmt.Errorf("cannot apply unary %s to %T, expected *rt.Element", TokenAsString(n.operator), rv)
		}
	default:
		return nil, fmt.Errorf("unknown unary operator: %d", n.operator)
	}
}

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

	default:
		return nil, fmt.Errorf("unkown binary operator %s", TokenAsString(n.operator))
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
