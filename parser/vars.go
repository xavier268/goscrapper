package parser

import "fmt"

// this files implements functions to manipulate variable, and input/output parameters.

// Typical generated code structure for a function with input parameters:
//
// func Do_xxx ( _in Input_XXX) (_out []Output_XXX, _err error) {
//
//		// initialize values from input params
// 		var a typa = _in.a
// 		var b typb = _in.b
//		.../...
//
// 		// intitialize global result structure
// 		_out = []Output_XXX{}
//
//		// verbatim lines generated during parsing, usin :=.
//		// this applies for both output and non output variables.
//		// input vars cannot be set.
//		c := a + b // c can only be allocated once per scope. Its type is captured from a and b type.
//
//		../... generated statement lines
//		tata := "a tata value" // tata is added to the knwon variables, with its type.
//
//
//		// ********* save and increment _out **********
//      // loop over the output variables, and set the known variables with the same name.
//		// set c because c is both a known variable and is requested by return.
//		// _out[len(_out) - 1]. c = c
//		// ********************************************
//
//		for titi :=  ... {
//			// always increment at beginning of any loop :
//			// ********* increment _out ***********
//			 _out = append(_out, Output_xxx{})
//			// ************************************
//
//
//
//			// do stuff ...
//			c := "a new value" // c is given a new value, shadocing previous c. Compiler checks that the type has not changed if a 'c' var was already known.
//
//			// setting tutu(local) from tata(global)
//			tutu := tata // compiler verify tata type is known, set tt to known and type of tata
//
//			// embedded loops as last stement.
//			for zzz := ... {
//				// always increment at beginning of any loop :
//				// ********* increment _out ***********
//				 _out = append(_out, Output_xxx{})
//				// ************************************
//
//
//				// Last statement contained a RETUN, that generates the save code ...
//				// ********* save  _out **********
//      		// loop over the output variables, and set the known variables with the same name.
//				// set c because c is both a known variable and is requested by return.
//				 _out[len(_out) - 1]. c = c
//			 	_out[len(_out) - 1]. tata = tata // based on global value
//			 	_out[len(_out) - 1]. tutu = tutu // based on local value
//			 	_out = append(_out, Output_xxx{})
//				// ********************************************
//
//					} // for zz
// 		// nothing serious should happen here - the FOR expression must always be the last statement of the main function or of the outer loop.
// 		} // for titi
//
//		// return result
//		return _out, nil
// }

// Declare a new input parameter.
func (m *myLexer) wDeclInputParam(name string, typ string) {

	// cheks ...
	if typ == "" {
		m.errorf("the type for the input parameter %s should be specified", name)
	}
	for _, k := range m.inparams {
		if name == k {
			m.errorf("the input parameter %s was already defined", name)
		}
	}

	// register in lexer status
	m.inparams = append(m.inparams, name)
	m.vars[name] = typ

	// declare a golang variable with same name and type in current (gloabl ) scope
	li := fmt.Sprintf("var %s %s = _in.%s", name, typ, name)
	m.addLines(li)
}

// Set a variable (local or global).
// The type is typically derived from the expression generating the value.
func (m *myLexer) setVar(name string, val string, typ string) {

	// checks
	for _, k := range m.inparams {
		if k == name {
			m.errorf("you cannot allocate a value to the input parameter %s", name)
		}
	}
	if tt, ok := m.vars[name]; ok && tt != typ {
		m.errorf("the variable %s already exists under a different type %s but the value is of type %s", name, m.vars[typ], typ)
	}

	// register variable
	m.vars[name] = typ

	// generate code
	li := fmt.Sprintf("%s := %s", name, val)
	m.addLines(li)
}

// Get the value and type of a named variable.
func (m *myLexer) getVar(name string) (val string, typ string) {
	// checks
	tt, ok := m.vars[name]
	if !ok || tt == "" {
		m.errorf("variable %s was never defined", name)
	}
	// return value & type
	return fmt.Sprintf("( %s )", name), tt
}
