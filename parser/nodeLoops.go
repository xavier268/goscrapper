package parser

import (
	"fmt"

	"github.com/xavier268/goscrapper/rt"
)

// ==== RETURN NODE ====

type nodeReturn struct {
	what     Nodes
	distinct bool // only return/send distincts
	last     bool // only return last
}

var _ Node = nodeReturn{}

// evaluating a return will evaluate the return expressionList,
// and either send it to the channel, or aggregate it for the final result.
func (n nodeReturn) eval(it *Interpreter) (any, error) {

	if n.what == nil || len(n.what) == 0 {
		// nothing to send, we're done.
		return nil, nil
	}

	// evaluate return expression list
	res, err := n.what.eval(it)
	if err != nil {
		return nil, err
	}

	// if we only care about last result ...
	if n.last {
		it.last = true          // remember last mode fro program finalization
		it.results = []any{res} // only keep last result
		return nil, nil
	}

	// if we only care about distinct, check if we already sent it
	if n.distinct {
		if it.unique.Add(res) != nil {
			// already seen, ignore ...
			return nil, nil
		}
	}

	// send result to channel, or aggregate it for the final result
	if it.ch != nil {
		// try to send res to channel while monitoring context cancelation
		select {
		case it.ch <- res:
			return nil, nil
		case <-it.ctx.Done():
			return nil, it.ctx.Err()
		}
	} else {
		// aggregate result for end of program return
		it.results = append(it.results, res)
		return nil, it.ctx.Err() // err if ctx cancelled
	}
}

// ====== FOR LOOP NODE ======

type nodeForLoop struct {
	from, to, step Node
	loopVar        string
	body           Nodes
}

var _ Node = nodeForLoop{}
var _ NodeWithBody = nodeForLoop{}

// body will be set later
func (m *myLexer) newNodeForLoop(loopVar Node, from Node, to Node, step Node) nodeForLoop {
	ret := nodeForLoop{}
	if loopVar != nil { // loopVariable can possibly be nil, represented as empty string in nodeForLoop
		//check loopVariable, and register it
		lv, ok := loopVar.(tok)
		if !ok || lv.c != IDENTIFIER || !isValidId(lv.v) {
			m.errorf("loop variable %s is not a valid identifier", loopVar)
		}
		if m.vars[lv.v] {
			m.errorf("loop variable %s is already declared as a variable or parameter", lv.v)
		}
		m.vars[lv.v] = true
		ret.loopVar = lv.v
	}
	ret.from = from
	ret.to = to
	ret.step = step
	ret.body = Nodes{}
	return ret
}

// make a copy of the node with updated body.
func (n nodeForLoop) appendBody(nodes ...Node) NodeWithBody {
	return nodeForLoop{
		from:    n.from,
		to:      n.to,
		step:    n.step,
		loopVar: n.loopVar,
		body:    append(n.body, nodes...),
	}
}

// eval implements Node.
func (n nodeForLoop) eval(it *Interpreter) (any, error) {
	// check context
	if it.ctx.Err() != nil {
		return nil, it.ctx.Err()
	}

	// defaults
	var from, to int
	step := 1

	// set from
	if n.from != nil {
		v, err := n.from.eval(it)
		if err != nil {
			return nil, err
		}
		if v != nil {
			if i, ok := v.(int); ok {
				from = i
			} else {
				return nil, fmt.Errorf("expected an int as a loop limit, but got a %T", v)
			}
		}
	}
	// set to
	if n.to != nil {
		v, err := n.to.eval(it)
		if err != nil {
			return nil, err
		}
		if v != nil {
			if i, ok := v.(int); ok {
				to = i
			} else {
				return nil, fmt.Errorf("expected an int as a loop limit, but got a %T", v)
			}
		}
	}

	// set step
	if n.step != nil {
		v, err := n.step.eval(it)
		if err != nil {
			return nil, err
		}
		if v != nil {
			if i, ok := v.(int); ok {
				step = i
			} else {
				return nil, fmt.Errorf("expected an int as a loop step, but got a %T", v)
			}
		}
	}

	// trigger iteration
	iter := rt.NewForLoopIterator(from, to, step)

	// create a stack frame
	it.pushFrame()
	defer it.popFrame()

	// do the actual looping
	for i, ok := iter.Next(); ok; i, ok = iter.Next() {
		// check context
		if it.ctx.Err() != nil {
			return nil, it.ctx.Err()
		}
		// assign loopVar
		err := it.assignVar(n.loopVar, i, false)
		if err != nil {
			return nil, err
		}
		// run loop
		_, err = n.body.eval(it)
		if err != nil {
			return nil, err
		}
		// reset stack frame and iterate
		err = it.resetFrame()
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// ==== FOR ARRAY LOOP NODE ====

type nodeForArray struct {
	array   Node
	loopVar string
	body    Nodes
}

var _ NodeWithBody = nodeForArray{}
var _ Node = nodeForArray{}

func (m *myLexer) newNodeForArray(loopVar Node, array Node) nodeForArray {
	ret := nodeForArray{}
	if loopVar != nil { // loopVariable can possibly be nil, represented as empty string in nodeForLoop
		//check loopVariable, and register it
		lv, ok := loopVar.(tok)
		if !ok || lv.c != IDENTIFIER || !isValidId(lv.v) {
			m.errorf("loop variable %s is not a valid identifier", loopVar)
		}
		if m.vars[lv.v] {
			m.errorf("loop variable %s is already declared as a variable or parameter", lv.v)
		}
		m.vars[lv.v] = true
		ret.loopVar = lv.v
	}
	ret.array = array
	ret.body = Nodes{}
	return ret
}

func (n nodeForArray) appendBody(nodes ...Node) NodeWithBody {
	return nodeForArray{
		array:   n.array,
		loopVar: n.loopVar,
		body:    append(n.body, nodes...),
	}
}

func (n nodeForArray) eval(it *Interpreter) (any, error) {
	// check context
	if it.ctx.Err() != nil {
		return nil, it.ctx.Err()
	}

	// evaluate array
	a, err := n.array.eval(it)
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

	// prepare a new loop frame
	it.pushFrame()
	defer it.popFrame()

	// iterate over array elements
	for _, ae := range aa {

		// check context
		if it.ctx.Err() != nil {
			return nil, it.ctx.Err()
		}

		// assign loopVar
		err := it.assignVar(n.loopVar, ae, false)
		if err != nil {
			return nil, err
		}
		// run loop
		_, err = n.body.eval(it)
		if err != nil {
			return nil, err
		}
		// reset stack frame and iterate
		err = it.resetFrame()
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// ==== SELECT NODE ===

type nodeSelect struct {
	css     Node   // css expression
	page    Node   // page or element expression
	body    Nodes  // body of the loop
	loopVar string // local loop variable or tok{}
	limit   Node   // integer limit >0, or else, no limit
	where   Nodes  // list of conditions
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
		return nil, fmt.Errorf("cannot evaluate a nil css expression")
	}
	cs, err := n.css.eval(it)
	if err != nil {
		return nil, err
	}
	if cs == nil {
		return nil, fmt.Errorf("expected a string css expression, but got nil")
	}
	if _, ok := cs.(string); !ok {
		return nil, fmt.Errorf("expected a string css expression, but got a %T", cs)
	}
	css := cs.(string)
	if css == "" {
		return nil, fmt.Errorf("expected an expression evaluating to a non-empty string css expression, but got an empty string")
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
	iter := rt.NewSelectAllIterator[rt.Elementer](it.ctx, page, css, limit)
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
func (lx *myLexer) newNodeSelect(loopVar Node, css Node, page Node, options Node) nodeSelect {

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

	// setup ând verify page
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
	return ret
}
