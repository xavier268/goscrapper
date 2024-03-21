package parser

import (
	"bytes"
	"fmt"
	"io"
)

// ======================================================
// This file contains the glue for the public parser api
// ======================================================

// Parse from in and write production on out.
func Parse(in io.Reader, out io.Writer) error {

	var err error
	var bb = new(bytes.Buffer) // collect out (errors) from lex.Error()

	lex := &myLexer{
		data: []byte{},
		pos:  0,
		w:    bb,
	}
	lex.data, err = io.ReadAll(in)
	if err != nil {
		return err
	}

	if yyParse(lex) != 0 {
		return fmt.Errorf(bb.String())
	}
	return nil
}

// Prints a token defined by its constant value as a string.
func TokenAsString(t int) string {
	idx := t - yyPrivate + 1 // yyPrivate points to error
	if idx < len(yyToknames) && idx >= 0 {
		return yyToknames[idx]
	} else {
		return fmt.Sprintf("TOK-%d", t)
	}
}
