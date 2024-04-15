package parser

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/xavier268/goscrapper/rt"
)

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
	case RAW:
		return fmt.Sprintf("%#v", rv), nil
	case GO:
		return fmt.Sprintf("%v", rv), nil
	case JSON:
		bb, err := json.Marshal(rv)
		if err != nil {
			return nil, err
		}
		return string(bb), nil
	case GSC:
		str, err := rt.Serialize(rv)
		if err != nil {
			return nil, err
		}
		return str, nil
	default:
		return nil, fmt.Errorf("unknown unary operator: %d", n.operator)
	}
}
