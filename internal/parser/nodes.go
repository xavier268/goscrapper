package parser

import (
	"fmt"
	"strconv"
)

// a Node (and the whole parsed tree) should always evaluate to something.
type Node interface {
	eval(*Interpreter) (any, error)
}

// A list of Node is also a Node per se.
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

var _ Node = &nodeLitteral{}
var _ Node = Nodes{}

// ===== nodeLitteral =====

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

type nodePrint struct {
	nodes Nodes
}

var _ Node = &nodePrint{}

// eval implements Node.
func (n nodePrint) eval(*Interpreter) (any, error) {
	for _, n := range n.nodes {
		fmt.Printf("PRINT : %#v", n)
	}
	fmt.Println()
	return nil, nil
}
