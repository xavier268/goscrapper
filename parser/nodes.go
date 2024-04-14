package parser

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/xavier268/goscrapper/rt"
)

// Node for the abstract syntax tree.
type Node interface {
	eval(*Interpreter) (any, error)
}

var _ Node = tok{}

// eval implements Node for token.
// Should never be called. Always eval to nil.
func (t tok) eval(*Interpreter) (any, error) {
	return nil, fmt.Errorf("tok.eval should never be called")
}

// Node with a body (eg : loop statement)
type NodeWithBody interface {
	Node
	appendBody(nodes ...Node) NodeWithBody // return a copy of the node, with selected nodes added to body
}

// ===== Nodes =====

// A list of Node is also a Node.
type Nodes []Node

func (ns Nodes) eval(i *Interpreter) (any, error) {
	var ret = make([]any, 0, len(ns))
	for _, v := range ns {
		if v == nil {
			continue
		}
		if i.ctx.Err() != nil {
			return nil, i.ctx.Err()
		}
		r, err := v.eval(i)
		if err != nil {
			return nil, err
		}
		ret = append(ret, r)
	}
	return ret, nil
}

var _ Node = Nodes{}

// ===== NodeMap =====

// A map of string to Node is also a Node.
type NodeMap map[string]Node

var _ Node = NodeMap{}

// Eval returns a map of string to values from each member evaluation.
// Keys are left unchanged.
func (n NodeMap) eval(it *Interpreter) (any, error) {
	res := make(map[string]any, len(n))
	for k, v := range n {
		r, err := v.eval(it)
		if err != nil {
			return nil, err
		}
		res[k] = r
	}
	return res, nil
}

// add a keyvalue node to a NodeMap.
// if NodeSet is nil, a new NodeSet is created.
// if both are nil, create an empty node set.
func (m *myLexer) newNodeMap(set Node, kv Node) NodeMap {

	res := make(map[string]Node, 4) // result will copied be here, to remain immutable.

	if set == nil {
		if kv == nil {
			return res
		}
		kk, okk := kv.(nodeKeyValue)
		if !okk {
			m.errorf("cannot add keyValue %#v to %#v", kv, set)
			return res
		}
		res[kk.key] = kk.value
		return res
	}

	// now, set != nil
	ss, oks := set.(NodeMap)
	if !oks {
		m.errorf("cannot add keyValue %#v to %#v, not a NodeSet", kv, set)
		return res
	}
	// copy set
	for k, v := range ss {
		res[k] = v
	}
	// try to update ?
	if kv == nil {
		return res
	}
	kk, okk := kv.(nodeKeyValue)
	if !okk {
		m.errorf("cannot add %#v to %#v, not a keyValue", kv, ss)
		return nil
	}
	// update and return copy
	res[kk.key] = kk.value
	return res
}

// ===== nodeLitteral =====
var _ Node = &nodeLitteral{}

type nodeLitteral struct {
	value any
}

func (m *myLexer) newNodeLitteral(tok tok) nodeLitteral {
	switch tok.c {
	case BOOL:
		var v bool = false
		switch tok.v {
		case "true":
			v = true
		case "false":
			v = false
		default:
			m.errorf("cannot convert %s to bool litteral", tok.v)
		}
		return nodeLitteral{value: v}
	case STRING:
		return nodeLitteral{value: tok.v}
	case NUMBER:
		i, err := strconv.Atoi(tok.v)
		if err != nil {
			m.errorf("cannot convertr number %s to int", tok.v)
		}
		return nodeLitteral{value: i}
	default:
		m.errorf("could not recognize token %#v for node litteral", tok)
	}
	return nodeLitteral{value: fmt.Errorf("invalid litteral")}
}

func (n nodeLitteral) eval(i *Interpreter) (any, error) {
	return n.value, nil
}

// ======= ASSIGN statement =========

var _ Node = &nodeAssign{}

type nodeAssign struct {
	id   string
	node Node
}

func (m *myLexer) newNodeAssign(tok tok, node Node) nodeAssign {
	if !isValidId(tok.v) || tok.c != IDENTIFIER {
		m.errorf("variable %s is not a valid input variable", tok.v)
	}
	if m.params[tok.v] {
		m.errorf("you may not assign to input variable %s", tok.v)
	}
	// register variable as a normal variable.
	m.vars[tok.v] = true
	return nodeAssign{id: tok.v, node: node}
}

// eval nodeAssign
func (n nodeAssign) eval(it *Interpreter) (value any, err error) {
	if n.node != nil {
		value, err = n.node.eval(it)
	}
	if err != nil {
		return nil, err
	}
	err = it.assignVar(n.id, value)
	return value, err
}

// ========= nodeVariable =============
var _ Node = &nodeVariable{}

type nodeVariable struct {
	id string
}

// eval implements Node.
func (n nodeVariable) eval(it *Interpreter) (any, error) {
	return it.getVar(n.id)
}

// newInputVar creates a new nodeVariable node to GET the variable content later,
// either registering variable as an input var if input is set to true, or as a normal variable.
// if exists is set, verify the variable was already declared (not for input)
func (m *myLexer) newNodeVariable(tok tok, input bool, exists bool) nodeVariable {
	if !isValidId(tok.v) || tok.c != IDENTIFIER {
		m.errorf("variable %s is not a valid variable", tok.v)
	}
	if input {
		// register as input var
		if m.vars[tok.v] {
			m.errorf("variable %s is already a normal variable", tok.v)
			return nodeVariable{id: tok.v}
		}
		m.params[tok.v] = true
		return nodeVariable{id: tok.v}
	} else {
		if !exists {
			// register normal var access
			if m.params[tok.v] {
				m.errorf("variable %s is already an input variable", tok.v)
				return nodeVariable{id: tok.v}
			}
			m.vars[tok.v] = true
		} else {
			// variable is supposed to be already declared
			if !m.vars[tok.v] {
				m.errorf("variable %s is not yet declared", tok.v)
			}
			return nodeVariable{id: tok.v}
		}
	}
	return nodeVariable{id: tok.v}
}

// ===== nodeKey ========

type nodeKey struct {
	key string
}

var _ Node = nodeKey{}

func (m *myLexer) newNodeKey(tok tok) nodeKey {
	if !isValidId(tok.v) || tok.c != IDENTIFIER {
		m.errorf("%s is not a valid key", tok.v)
	}
	return nodeKey{key: tok.v}
}

// eval implements Node.
func (n nodeKey) eval(*Interpreter) (any, error) {
	return n.key, nil
}

// ===== nodeKeyValue ========

type nodeKeyValue struct {
	key   string
	value Node
}

var _ Node = nodeKeyValue{}

// eval implements Node. Evaluate only its value. Key is unchanged.
func (n nodeKeyValue) eval(it *Interpreter) (any, error) {
	if n.value == nil {
		return nil, nil
	}
	return n.value.eval(it)
}

func (m *myLexer) newNodeKeyValue(key Node, node Node) nodeKeyValue {
	if kk, ok := key.(nodeKey); ok {
		return nodeKeyValue{key: kk.key, value: node}
	} else {
		m.errorf("key %v is not a valid nodeKey node", key)
	}
	return nodeKeyValue{}
}

// ===  top level program node ===

type nodeProgram struct {
	req    Node
	invars []string // list of externally provided variables, also called parameters.
}

var _ Node = nodeProgram{}

func (n nodeProgram) eval(it *Interpreter) (any, error) {
	if n.req == nil {
		return nil, fmt.Errorf("cannot evaluate a nil request")
	}
	// check that all params provided to the interpreter are knwown to the request
	for _, v := range it.invars {
		if !slices.Contains(n.invars, v) {
			return nil, fmt.Errorf("provided parameter %s in not known to this request", v)
		}
	}
	// evaluate program content
	_, err := n.req.eval(it)
	if it.ch != nil {
		return nil, err // in async mode, results were already sent
	} else {
		return it.results, err // in sync mode, results were aggregated
	}
}

// ==== RETURN NODE ====

type nodeReturn struct {
	what Nodes
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
	if it.ch != nil {
		// try to send res to channel while monitoring context cancelation
		select {
		case it.ch <- res:
			return nil, nil
		case <-it.ctx.Done():
			return nil, it.ctx.Err()
		}
	} else {
		// aggregate result
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
		err := it.assignVar(n.loopVar, i)
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
		err := it.assignVar(n.loopVar, ae)
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

// === NODE MAP ACCESS ===

type nodeMapAccess struct {
	m Node
	k Node
}

var _ Node = nodeMapAccess{}

// eval implements Node.
func (n nodeMapAccess) eval(it *Interpreter) (any, error) {
	if n.m == nil || n.k == nil {
		return nil, fmt.Errorf("cannot evaluate a nil map or key")
	}
	// evaluate map
	m, err := n.m.eval(it)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, nil
	}
	// verify m is a map
	mm, ok := m.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("expected a map, but got a %T", m)
	}
	// evaluate key
	k, err := n.k.eval(it)
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, nil
	}
	// verify k is a string
	kk, ok := k.(string)
	if !ok {
		return nil, fmt.Errorf("expected a string a key, but got a %T", k)
	}
	// Access map content
	return mm[kk], nil
}

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

// ==== NODE SLOW ====

type nodeSlow struct {
	m Node // duration in milliseconds
}

var _ Node = nodeSlow{}

// eval Node
func (n nodeSlow) eval(it *Interpreter) (any, error) {
	// default millis duration
	dur := rt.SLOW_DELAY

	// evaluate duration
	if n.m != nil {
		d, err := n.m.eval(it)
		if err != nil {
			return nil, err
		}
		if d != nil { // if d is nil, use default dur ...
			// verify d is an int
			dd, ok := d.(int)
			if !ok {
				return nil, fmt.Errorf("expected an int millis duration, but got a %T", d)
			}
			if dd > 0 {
				dur = time.Duration(dd) * time.Millisecond // only set for valid, strictly positive values.
			}
		}
	}
	// sleep until duration has expired or context is cancelled
	timer := time.NewTimer(dur)
	select {
	case <-it.ctx.Done():
		if !timer.Stop() {
			<-timer.C // drain the channel if the timer had already expired.
			return nil, it.ctx.Err()
		}
		return nil, it.ctx.Err()
	case <-timer.C:
		return nil, nil
	}
}

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

// === NODE FAIL ===

type nodeFail struct {
	what Node
}

var _ Node = nodeFail{}

func (n nodeFail) eval(it *Interpreter) (any, error) {
	if n.what != nil {
		v, err := n.what.eval(it)
		if err != nil {
			return nil, err
		}
		if v != nil {
			if vv, ok := v.(string); ok {
				return nil, fmt.Errorf("fail message : %s", vv)
			}
		}
	}
	return nil, fmt.Errorf("fail requested")
}

// === NODE ASSERT ===
type nodeAssert struct {
	cond Node
}

var _ Node = nodeAssert{}

func (n nodeAssert) eval(it *Interpreter) (any, error) {
	// evaluate condition
	if n.cond != nil {

		v, err := n.cond.eval(it)
		if err != nil {
			return nil, err
		}
		if v != nil {

			// verify v is a bool
			vv, ok := v.(bool)
			if ok && vv {
				return nil, nil // assertion success
			}
		}
	}
	return nil, fmt.Errorf(fmt.Sprintf("assertion %#v failed", n.cond))
}

// ==== PRINT NODE ===

type nodePrint struct {
	nodes Nodes
}

var _ Node = nodePrint{}

func (n nodePrint) eval(it *Interpreter) (any, error) {
	if n.nodes != nil {
		// evaluate and print
		for _, p := range n.nodes {
			v, err := p.eval(it)
			if err != nil {
				return nil, err
			}
			fmt.Print(v)
		}
	}
	fmt.Println()
	return nil, nil
}
