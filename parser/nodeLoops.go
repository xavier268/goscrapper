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
