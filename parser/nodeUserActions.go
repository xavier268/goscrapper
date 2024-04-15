package parser

import (
	"fmt"

	"github.com/xavier268/goscrapper/rt"
)

// === INPUT text IN element ===

type nodeInput struct {
	element Node
	text    Node
}

var _ Node = nodeInput{}

// eval implements Node.
func (n nodeInput) eval(*Interpreter) (any, error) {
	txt := ""
	if n.text != nil {
		v, err := n.text.eval(nil)
		if err != nil {
			return nil, err
		}
		if v != nil {
			return nil, fmt.Errorf("expected a string input, but got nil")
		}
		vv, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("expected a string input, but got a %T", v)
		}
		txt = vv
	}
	if n.element == nil {
		return nil, fmt.Errorf("expected an element expression, but got nil")
	}
	e, err := n.element.eval(nil)
	if err != nil {
		return nil, err
	}
	if e == nil {
		return nil, fmt.Errorf("expected an element expression, but got nil")
	}
	ee, ok := e.(*rt.Element)
	if !ok {
		return nil, fmt.Errorf("expected an element expression, but got a %T", e)
	}
	return nil, ee.Input(txt)
}

// === CLICK with OPTIONS ====
type nodeClick struct {
	element Node
	// left is default
	middle bool
	right  bool
	count  Node
}

var _ Node = nodeClick{}

func (*myLexer) mergeNodeClick(n1 Node, n2 Node) nodeClick {
	ret := nodeClick{}
	if n1 != nil {
		if n1c, ok := n1.(nodeClick); ok {
			ret = n1c
		}
	}
	if n2 != nil {
		if n2c, ok := n2.(nodeClick); ok {
			ret.middle = ret.middle || n2c.middle
			ret.right = ret.right || n2c.right
			if n2c.count != nil && ret.count != nil {
				lx.errorf("you may not specify a click count twice")
			}
			if n2c.count != nil {
				ret.count = n2c.count
			}
		}
	}
	return ret
}

func (n nodeClick) eval(it *Interpreter) (any, error) {
	if n.element == nil {
		return nil, fmt.Errorf("expected an element expression, but got nil")
	}
	e, err := n.element.eval(it)
	if err != nil {
		return nil, err
	}
	if e == nil {
		return nil, fmt.Errorf("expected an element expression, but got nil")
	}
	ee, ok := e.(*rt.Element)
	if !ok {
		return nil, fmt.Errorf("expected an element expression, but got a %T", e)
	}

	count := 1 // default
	if n.count != nil {
		v, err := n.count.eval(it)
		if err != nil {
			return nil, err
		}
		if v == nil {
			return nil, fmt.Errorf("expected an integer count, but got nil")
		}
		vv, ok := v.(int)
		if !ok {
			return nil, fmt.Errorf("expected an integer count, but got a %T", v)
		}
		if vv > 0 { // only change default for > 0.
			count = vv
		}
	}

	switch {
	case n.middle:
		rt.Click(ee, "middle", count)
	case n.right:
		rt.Click(ee, "right", count)
	default:
		rt.Click(ee, "left", count)
	}
	return nil, nil
}
