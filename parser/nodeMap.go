package parser

import "fmt"

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

// ===== NODE KEY ========

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

// ===== NODE KEY VALUE ========

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
