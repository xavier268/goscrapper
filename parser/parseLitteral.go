package parser

import (
	"fmt"
	"os"
	"strconv"
)

// ParseLitteral parse a gsc litteral with a small subset of the gsc syntax.
// Only accepts nil, bool, ints, strings, [] and {}.
func ParseLitteral(content string) (any, error) {
	lx := NewLexer("litteral parser", []byte(content), os.Stdout)
	return parseLitteral(lx.(*myLexer))
}

// parseLitteral parse a gsc litteral with a small subset of the gsc syntax.
func parseLitteral(lx *myLexer) (any, error) {

	lval := new(yySymType)
	c := lx.Lex(lval)

	switch c {
	case 0:
		return nil, fmt.Errorf("unexpected EOF while parsing a litteral")
	case NIL:
		return nil, nil
	case NUMBER:
		i, err := strconv.Atoi(lval.tok.v)
		if err != nil {
			return nil, err
		}
		return i, nil
	case MINUS:
		i, err := parseLitteral(lx)
		if err != nil {
			return nil, err
		}
		if i, ok := i.(int); ok {
			return -i, nil
		}
		return nil, fmt.Errorf("unexpected value %#v while parsing a negative number", i)
	case STRING:
		return lval.tok.v, nil
	case IDENTIFIER:
		return lval.tok.v, nil // for map keys
	case BOOL:
		return (lval.tok.v == "true"), nil
	case LBRACE:
		return parseMap(lx)
	case LBRACKET:
		return parseArray(lx)
	default:
		return nil, fmt.Errorf("unexpected token %s while parsing a gsc litteral", lval.tok.v)
	}
}

// the opening bracket was already parsed ...
func parseArray(lx *myLexer) ([]any, error) {

	arr := make([]any, 0)
	lval := new(yySymType)

	var c int

	// peek for closing bracket ?
	if lx.data[lx.pos] == ']' {
		lx.pos++
		return arr, nil
	}

loop:

	// parse array element
	el, err := parseLitteral(lx)
	if err != nil {
		return nil, err
	}
	arr = append(arr, el)

	// read next token, expect ] or ,
	c = lx.Lex(lval)
	if c == 0 {
		return nil, fmt.Errorf("unexpected EOF while parsing an array")
	}
	if c == RBRACKET {
		return arr, nil
	}
	if c == COMMA {
		goto loop
	}
	return nil, fmt.Errorf("unexpected token %s while parsing an array", lval.tok.v)
}

// the opening brace was already parsed ...
func parseMap(lx *myLexer) (any, error) {
	mm := make(map[string]any)
	lval := new(yySymType)
	c := 0

	// peek for closing brace
	if lx.data[lx.pos] == '}' {
		lx.pos++
		return mm, nil
	}

loop:
	// parse a map key
	c = lx.Lex(lval)
	if c != IDENTIFIER {
		return nil, fmt.Errorf("unexpected token %s while parsing a map key", lval.tok.v)
	}
	key := lval.tok.v
	if key == "" {
		return nil, fmt.Errorf("unexpected empty map key")
	}
	c = lx.Lex(lval)
	if c != COLON {
		return nil, fmt.Errorf("unexpected token %s while expecting a colon after map key", lval.tok.v)
	}
	val, err := parseLitteral(lx)
	if err != nil {
		return nil, err
	}
	mm[key] = val

	// read next token - expect } or ,
	c = lx.Lex(lval)
	if c == 0 {
		return nil, fmt.Errorf("unexpected EOF while parsing a map")
	}
	if c == RBRACE {
		return mm, nil
	}
	if c == COMMA {
		goto loop
	}

	return nil, fmt.Errorf("unexpected token %s while parsing a map", lval.tok.v)

}
