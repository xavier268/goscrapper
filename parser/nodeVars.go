package parser

// ======= ASSIGN statement =========

var _ Node = nodeAssign{}

type nodeAssign struct {
	id     string
	node   Node
	global bool
}

func (m *myLexer) newNodeAssign(tok tok, node Node, global bool) nodeAssign {
	if !isValidId(tok.v) || tok.c != IDENTIFIER {
		m.errorf("variable %s is not a valid input variable", tok.v)
	}
	if m.params[tok.v] {
		m.errorf("you may not assign to input variable %s", tok.v)
	}
	// register variable as a normal variable.
	m.vars[tok.v] = true
	return nodeAssign{id: tok.v, node: node, global: global}
}

// eval nodeAssign
func (n nodeAssign) eval(it *Interpreter) (value any, err error) {
	if n.node != nil {
		value, err = n.node.eval(it)
	}
	if err != nil {
		return nil, err
	}

	err = it.assignVar(n.id, value, n.global)

	return value, err
}

// ========= NODE VARIABLE =============

var _ Node = nodeVariable{}

type nodeVariable struct {
	id     string
	global bool
}

// eval implements Node.
func (n nodeVariable) eval(it *Interpreter) (any, error) {
	return it.getVar(n.id, n.global)
}

// newInputVar creates a new nodeVariable node to GET the variable content later,
// either registering variable as an input var if input is set to true, or as a normal variable.
// if exists is set, verify the variable was already declared (not for input)
// global force a global variable creation
func (m *myLexer) newNodeVariable(tok tok, input bool, exists bool, global bool) nodeVariable {
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
				return nodeVariable{id: tok.v, global: global}
			}
			m.vars[tok.v] = true
		} else {
			// variable is supposed to be already declared
			if !m.vars[tok.v] {
				m.errorf("variable %s is not yet declared", tok.v)
			}
			return nodeVariable{id: tok.v, global: global}
		}
	}
	return nodeVariable{id: tok.v, global: global}
}
