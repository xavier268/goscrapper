package parser

import (
	"fmt"

	"github.com/xavier268/goscrapper/rt"
)

// ==== SELECT NODE ===

type nodeSelect struct {
	css     Node   // css expression
	page    Node   // page or element expression
	body    Nodes  // body of the loop
	loopVar string // local loop variable or tok{}
	limit   Node   // integer limit >0, or else, no limit
	where   Nodes  // list of conditions
	xpath   bool   // use xpath instead of css
}

var _ NodeWithBody = nodeSelect{}

// appendBody implements NodeWithBody.
func (n nodeSelect) appendBody(nodes ...Node) NodeWithBody {
	return nodeSelect{
		css:     n.css,
		page:    n.page,
		body:    append(n.body, nodes...),
		loopVar: n.loopVar,
		limit:   n.limit,
		where:   n.where,
		xpath:   n.xpath,
	}
}

// merge select options
func (n nodeSelect) mergeOptions(n2 Node) nodeSelect {
	switch ns2 := n2.(type) {
	case nil:
		return n
	case nodeSelect:
		if ns2.limit != nil && n.limit != nil && n.limit != n.limit {
			lx.errorf("you may not specify a limit twice")
			return n
		}
		res := nodeSelect{
			where: append(n.where, ns2.where...),
		}
		if n.limit == nil {
			res.limit = ns2.limit
		}
		n.xpath = n.xpath || ns2.xpath
		return res
	default:
		lx.errorf("cannot merge select options of type %T", n2)
		return n
	}
}

// eval implements NodeWithBody.
func (n nodeSelect) eval(it *Interpreter) (any, error) {

	// check context
	if it.ctx.Err() != nil {
		return nil, it.ctx.Err()
	}

	// prepare a new loop frame
	it.pushFrame()
	defer it.popFrame()

	// evaluate and verify page or element expression
	var page rt.Elementer
	if n.page == nil {
		return nil, fmt.Errorf("cannot select from a nil page or element expression")
	}
	p, err := n.page.eval(it)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, fmt.Errorf("expected a page or element expression, but got nil")
	}
	switch pp := p.(type) {
	case *rt.Page:
		page = pp
	case *rt.Element:
		page = pp
	default:
		return nil, fmt.Errorf("expected a page or element expression, but got a %T", p)
	}

	// evaluate and verify css
	if n.css == nil {
		return nil, fmt.Errorf("cannot evaluate a nil css/xpath expression")
	}
	cs, err := n.css.eval(it)
	if err != nil {
		return nil, err
	}
	if cs == nil {
		return nil, fmt.Errorf("expected a string css/xpath expression, but got nil")
	}
	if _, ok := cs.(string); !ok {
		return nil, fmt.Errorf("expected a string css/xpath expression, but got a %T", cs)
	}
	css := cs.(string)
	if css == "" {
		return nil, fmt.Errorf("expected an expression evaluating to a non-empty string css/xpath expression, but got an empty string")
	}

	// evaluate and verify limit
	limit := -1         // means no limit
	if n.limit != nil { // if a limit is given, it should be valid
		l, err := n.limit.eval(it)
		if err != nil {
			return nil, err
		}
		if l == nil {
			return nil, fmt.Errorf("expected an integer limit, but expression evaluated to nil")
		}
		ll, ok := l.(int)
		if !ok {
			return nil, fmt.Errorf("expected an integer limit, but got a %T", l)
		}
		limit = ll
	}

	// prepare stack frame
	it.pushFrame()
	defer it.popFrame()

	// start main loop
	iter := rt.NewSelectAllIterator[rt.Elementer](it.ctx, page, css, limit, n.xpath)
loop:
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {

		// update loop variable, if one was defined
		if n.loopVar != "" {
			err := it.assignVar(n.loopVar, next, false)
			if err != nil {
				return nil, err
			}
		}
		// evaluate where conditions, if any
		if n.where != nil {
			for _, p := range n.where {
				if it.ctx.Err() != nil {
					return nil, it.ctx.Err()
				}
				v, err := p.eval(it)
				if err != nil {
					return nil, err
				}
				if v == nil {
					return nil, fmt.Errorf("expected a boolean condition, but got nil")
				}
				vv, ok := v.(bool)
				if !ok {
					return nil, fmt.Errorf("expected a boolean condition, but got a %T", v)
				}
				if !vv { // condition does not match
					continue loop
				}
			}
		}

		// evaluate loop body
		for _, p := range n.body {
			if it.ctx.Err() != nil {
				return nil, it.ctx.Err()
			}
			_, err := p.eval(it)
			if err != nil {
				return nil, err
			}
		}

		// update frame
		if err := it.resetFrame(); err != nil {
			return nil, err
		}
	}
	return nil, nil // done !
}

// Create a nodeSelect without body.
// If xpath flag is set, use xpath instead of css (default).
func (lx *myLexer) newNodeSelect(loopVar Node, css Node, page Node, options Node, xpath bool) nodeSelect {

	// Setup loop var
	ret := nodeSelect{}
	if loopVar != nil { // loopVariable can possibly be nil, represented as empty string internally in that case
		//check loopVariable, and register it
		lv, ok := loopVar.(tok)
		if !ok || lv.c != IDENTIFIER || !isValidId(lv.v) {
			lx.errorf("loop variable %s is not a valid identifier", loopVar)
		}
		if lx.vars[lv.v] {
			lx.errorf("loop variable %s is already declared as a variable or parameter", lv.v)
		}
		lx.vars[lv.v] = true
		ret.loopVar = lv.v
	}

	// setup and verify css
	if css == nil {
		lx.errorf("css expression is required and cannot be nil")
	}
	ret.css = css

	// setup and verify page
	if page == nil {
		lx.errorf("page or element expression is required and cannot be nil")
	}
	ret.page = page
	if options != nil {
		if opt, ok := options.(nodeSelect); ok {
			ret.where = opt.where
			ret.limit = opt.limit
		} else {
			lx.errorf("expected a nodeSelect options, but got a %T", options)
		}
	}
	ret.xpath = ret.xpath || xpath
	return ret
}
