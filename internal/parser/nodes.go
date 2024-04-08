package parser

// a Node (and the whole parsed tree) should always evaluate to something.
type Node interface {
	eval(*Interpreter) (any, error)
}

var _ Node = &nodeLitteral{}

type nodeLitteral struct {
	value any
}

func (m *myLexer) newNodeLitteral(value tok) nodeLitteral {
	panic("todo - switch on token type to create correct nodeLitteral..")
}

func (n nodeLitteral) eval(i *Interpreter) (any, error) {
	return n.value, nil
}
