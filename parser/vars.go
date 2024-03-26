package parser

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

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
// 		//********* increment _out ***********
//		 _out = append(_out, Output_xxx{})
//		// ************************************
//
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
//		// ********* save  _out **********
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
//				// Last statement contained a RETURN, that generates the save code ...
//				// ********* save  _out **********
//      		// loop over the output variables, and set the known variables with the same name.
//				// set c because c is both a known variable and is requested by return.
//				_out[len(_out) - 1]. c = c
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
func (m *myLexer) declInputParam(name string, typ string) {

	// fmt.Println("DEBUG : calling declInputParam with ", name, typ)

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
	li := fmt.Sprintf("var %s %s = _in.%s ; _ = %s", name, typ, name, name)
	m.addLines(li)
}

// declare list of identifiers as output to be returned.
func (m *myLexer) declOutputParams(names []string) {
	for _, name := range names {
		// check
		if typ, ok := m.vars[name]; !ok || typ == "" {
			m.errorf("variable %s cannot be returned because it was never declared", name)
		}
		for _, oo := range m.outparams {
			if oo == name {
				m.errorf("variable %s duplicated in output parameters", name)
			}
		}
		// register output name
		m.outparams = append(m.outparams, name)

		// clean non relevant _out affectations.
		m.cleanOut()

	}
}

// Set a variable (local or global).
// The type is typically derived from the expression generating the value.
func (m *myLexer) vSetVar(name string, v value) {

	// checks
	for _, k := range m.inparams {
		if k == name {
			m.errorf("you cannot allocate a value to the input parameter %s", name)
		}
	}
	if tt, ok := m.vars[name]; ok && tt != v.t {
		m.errorf("the variable %s already exists under type %s, but trying to set a value of type %s", name, m.vars[name], v.t)
	}

	// register variable
	m.vars[name] = v.t

	// generate code
	li := fmt.Sprintf("var %s %s= %s;_=%s", name, v.t, v.v, name)
	m.addLines(li)

	// fmt.Printf("DEBUG : vars = %#v\n", m.vars)
}

// Get the value and type of a named variable.
func (m *myLexer) vGetVar(name string) (v value) {
	// checks
	tt, ok := m.vars[name]
	if !ok || tt == "" {
		m.errorf("variable %s is not defined", name)
	}
	// return value
	return value{
		v: fmt.Sprintf(" %s ", name),
		t: tt,
		c: 0,
	}
}

// increment the result variable _out.
func (m *myLexer) incOut() {
	m.addLines("// call to incOut")
	li := fmt.Sprintf(" _out = append(_out, Output_%s{})", m.name)
	m.addLines(li)
}

// make a snapshot of relevant vars into _out.
// imediately reincrement the output.
func (m *myLexer) saveOut() {
	// TODO - this should run only for known variables AND for variables that are already set by the code.
	// the known variables at this stage are available from m.vars.
	// but the out variables are not yet available.
	// This may required to create lines allocating all knwonwn variables as if they were all out params, and at the end, revisiting and commenting all lines atht are not part of out ?
	m.addLines("//call to saveOut")
	// sort vars to make output deterministic for easier testing
	vv := make([]string, len(m.vars))
	for v := range m.vars {
		vv = append(vv, v)
	}
	sort.Strings(vv)
	for _, v := range vv {
		li := fmt.Sprintf("//_out[len(_out)-1].%s=%s", v, v) // relevant lines will be uncommented later.
		m.addLines(li)
	}
	m.incOut()
}

// uncomment lines that correspond to valid out params.
// This should be called only after all return params are defined.
func (m *myLexer) cleanOut() {
	patt := regexp.MustCompile(`^//_out\[len\(_out\)-1\]\.([A-Za-z][A-Za-z0-9]*)=([A-Za-z][A-Za-z0-9]*)$`) // same as IDENTIFIER PATTERN
	for i, li := range m.lines {
		if ss := patt.FindStringSubmatch(li); len(ss) == 3 {
			suppress := true
			for _, v := range m.outparams {
				// fmt.Printf("DEBUG : cleanOut : %s and %s -> %#v\n", v, li, patt.FindStringSubmatch(li))
				if ss[1] == ss[2] && ss[1] == v {
					m.lines[i] = li[2:] // uncomment line ...
					suppress = false
				}
			}
			if suppress {
				// m.lines[i] += "// *SUPPRESSED !"
				m.lines[i] = ""
			}
		}
	}
}

// get arr[idx]
func (m *myLexer) vGetElementOf(arr value, idx value) value {

	if idx.t != "int" {
		m.errorf("array index should be a number, but it is a %s", idx.t)
	}
	if !strings.HasPrefix(arr.t, "[]") {
		m.errorf("expecting an array but got a %s", arr.t)
	}

	return value{
		v: fmt.Sprintf("(%s)[%s]", arr.v, idx.v),
		t: arr.t[2:],
		c: 0,
	}
}

// create an arry, checking all types are the same (or nil)
func (m *myLexer) vMakeArray(li []value) value {
	if len(li) == 0 { // that should never happen ...
		m.errorf("you cannot define a litteral array wit no elements")
	}

	t0 := li[0].t
	v0 := "[]" + t0 + "{"

	for i, v := range li {
		if i != 0 {
			v0 += ","
		}
		if v.t != t0 {
			m.errorf("array element type %s differs from first element type %s", v.t, t0)
		}
		v0 += v.v
	}

	return value{
		v: v0 + "}",
		t: "[]" + t0,
		c: 0,
	}

}
