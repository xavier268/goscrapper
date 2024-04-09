package parser

import "fmt"

// Unary operators

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
		default:
			return nil, fmt.Errorf("cannot apply unary %s to %T", TokenAsString(n.operator), rv)
		}
	case MINUS:
		switch v := rv.(type) {
		case int:
			return -v, nil
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
		switch v := rv.(type) {
		case string:
			return len(v), nil
		case []any:
			return len(v), nil
		case map[string]any:
			return len(v), nil
		default:
			return nil, fmt.Errorf("cannot apply unary %s to %T", TokenAsString(n.operator), rv)
		}
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
	default:
		return nil, fmt.Errorf("unknown unary operator: %d", n.operator)
	}
}
