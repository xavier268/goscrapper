package parser

import (
	"regexp"
	"slices"
)

// type for all the functions available while parsing, taht appy to a list of values, retuning a value.
// these functions are NOT vaialble to the compoiled code, once parsing is complete.
// These are NOT part of the Runtime.
type myValueFunction func(lx *myLexer, values []value) (value, error)

// map of registered functions, addresed by their names.
// names should be valid idenifiers and NOT keywords.
// use Register to set them. Do not set directly.
var myFunctions map[string]myValueFunction = make(map[string]myValueFunction, 10)

func init() {
	registerValueFunction("test", testFunction)
}

// Register a function to make it available to parser
func registerValueFunction(name string, f myValueFunction) {
	// enforce function names to be valid identifiers
	patt := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]*$")
	if !patt.MatchString(name) {
		panic("cannot register function : invalid function name: " + name)
	}
	// enforce name is not a keyword
	if slices.Contains(yyToknames[:], name) {
		panic("cannot register function : name is a keyword: " + name)
	}
	myFunctions[name] = f
}

func (m *myLexer) callFunction(name string, values []value) value {
	f, ok := myFunctions[name]
	if !ok {
		m.errorf("function not registered: %s", name)
	}
	v, err := f(m, values)
	if err != nil {
		m.errorf("error calling function %s : %s", name, err.Error())
	}
	return v
}

// ======================================================================

func testFunction(m *myLexer, values []value) (value, error) {
	return value{v: "test is ok !", t: "string"}, nil
}
