package parser

import (
	"fmt"
	"slices"
	"strconv"
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

// ===== Nodes =====

// A list of Node is also a Node.
type Nodes []Node

func (ns Nodes) eval(i *Interpreter) (any, error) {
	var ret = make([]any, 0, len(ns))
	for _, v := range ns {
		if v == nil {
			continue
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

// ======== PRINT statement ========

var _ Node = &nodePrint{}

type nodePrint struct {
	nodes Nodes
	raw   bool // if raw, will print using %#v, otherwise, use %v
}

// eval implements Node.
func (n nodePrint) eval(it *Interpreter) (any, error) {
	var err error
	fmt.Print("PRINT: ")
	for _, nn := range n.nodes {
		var nne any
		if nn != nil {

			nne, err = nn.eval(it)
		}
		if err != nil {
			return nil, err
		}
		if n.raw {
			fmt.Printf("%#v ", nne)
		} else {
			fmt.Printf("%v ", nne)
		}
	}
	fmt.Println()
	return nil, nil
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

// newInputVar creates a new nodeVariable node to GET the variable contanet later,
// either registering variable as an input var if input is set to true, or as a normal variable.
func (m *myLexer) newNodeVariable(tok tok, input bool) nodeVariable {
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
		// register normal var access
		if m.params[tok.v] {
			m.errorf("variable %s is already an input variable", tok.v)
			return nodeVariable{id: tok.v}
		}
		m.vars[tok.v] = true
		return nodeVariable{id: tok.v}
	}
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
	return n.req.eval(it)
}
