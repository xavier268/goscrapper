package parser

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/xavier268/goscrapper"
)

// ======================================================
// This file contains the glue for the public parser api
// ======================================================

// Parse from in and write production on out.
func Parse(out io.Writer, in io.Reader) error {

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

// Parse all inFiles and generate a scrapper in the provided directory.
// The package name is derived from the directory name.
// The generated file names are derived from the input file names
func ParseFiles(outDir string, inFiles ...string) error {

	outDir = MustAbs(outDir)
	packName := Normalize(filepath.Base(outDir))
	os.MkdirAll(outDir, 0755)
	for _, inFile := range inFiles {
		base := Normalize(filepath.Base(inFile))
		out, err := os.Create(filepath.Join(outDir, base+".go"))
		if err != nil {
			panic(err)
		}
		PrintHeader(out, "package "+packName, "// From "+inFile)
		defer out.Close()
		in, err := os.Open(inFile)
		if err != nil {
			panic(err)
		}
		defer in.Close()
		if err := Parse(out, in); err != nil {
			return err
		}
		out.Close()
		in.Close()
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

// Convert filename to absolute file name.
func MustAbs(filename string) string {
	fn, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}
	return fn
}

// Normalize a string to a valid, lowercase, identifier without any special chars.
// return "invalid" if invalid.
func Normalize(s string) string {

	// remove non valid chars
	re1 := regexp.MustCompile(`[^0-9a-z]`)
	s = strings.ToLower(s)
	s = re1.ReplaceAllString(s, "")

	// verify that the string is a valid identifier.
	re2 := regexp.MustCompile(`^[a-z][a-z0-9]*$`)
	if re2.MatchString(s) {
		return s
	} else {
		return "invalid"
	}
}

func PrintHeader(w io.Writer, ss ...string) {
	fmt.Println("// Autogenerated file. DO NOT EDIT.")
	fmt.Printf("// Version: %s\n", goscrapper.VERSION)
	fmt.Printf("// Date: %s\n", goscrapper.BUILDDATE)
	fmt.Printf("// Built : %s\n", goscrapper.GITHASH)
	fmt.Printf("// %s\n", goscrapper.COPYRIGHT)
	fmt.Println()

	for _, s := range ss {
		fmt.Println(s)
	}
	fmt.Println()
}
