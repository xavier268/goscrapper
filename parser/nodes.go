package parser

import (
	"fmt"
	"strconv"
)

// Node for the abstract syntax tree.
type Node interface {
	eval(*Interpreter) (any, error)
}

// A list of Node is also a Node.
type Nodes []Node

func (ns Nodes) eval(i *Interpreter) (any, error) {
	var ret = make([]any, len(ns))
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
}

// eval implements Node.
func (n nodePrint) eval(it *Interpreter) (any, error) {
	var err error
	for _, nn := range n.nodes {
		var nne any
		if nn != nil {
			nne, err = nn.eval(it)
		}
		if err != nil {
			return nil, err
		}
		fmt.Printf("PRINT : %#v\n", nne)
	}
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
func (n nodeAssign) eval(it *Interpreter) (any, error) {
	value, err := n.node.eval(it)
	if err != nil {
		return nil, err
	}
	err = it.assignVar(n.id, value)
	return nil, err
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
