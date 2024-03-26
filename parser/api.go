package parser

import (
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
// No package name is written.
func Parse(out io.Writer, in io.Reader) error {
	return parse(out, in, "noname")
}

// parse everything and write to out.
// NB : package name should already have been written to out ...
func parse(out io.Writer, in io.Reader, name string) error {

	var err error

	lex := &myLexer{
		name:      name,
		data:      []byte{},
		pos:       0,
		w:         out,
		lines:     []string{},
		inparams:  []string{},
		outparams: []string{},
		vars:      map[string]string{},
		imports:   map[string]bool{},
	}
	lex.data, err = io.ReadAll(in)
	if err != nil {
		return err
	}

	if yyParse(lex) != 0 {
		return fmt.Errorf("parsing error")
	}
	return nil
}

// Parse all files corresponding to the provided glob pattern.
// The package name is derived from the directory name of the outDir directory.
// The generated file names are derived from the input file names
func ParseGlob(outDir, glob string) error {

	files, err := filepath.Glob(glob)
	if err != nil {
		return err
	}
	return ParseFiles(outDir, files...)
}

// Parse all inFiles and generate a scrapper in the provided directory.
// The package name is derived from the directory name.
// The generated file names are derived from the input file names
func ParseFiles(outDir string, inFiles ...string) error {

	outDir = MustAbs(outDir)
	packName := Normalize(filepath.Base(outDir))
	os.MkdirAll(outDir, 0755)
	for _, inFile := range inFiles {
		fmt.Println("Parsing", inFile)
		inFile = MustAbs(inFile)
		base := Normalize(inFile)
		out, err := os.Create(filepath.Join(outDir, base+".go"))
		if err != nil {
			panic(err)
		}
		PrintHeader(out, "package "+packName, "// Generated from "+inFile)
		defer out.Close()
		in, err := os.Open(inFile)
		if err != nil {
			panic(err)
		}
		defer in.Close()
		if err := parse(out, in, base); err != nil {
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

	s = filepath.Base(s)                       // remove path if any
	s = strings.TrimSuffix(s, filepath.Ext(s)) // remove extension if any

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
	fmt.Fprintln(w, "// Autogenerated file. DO NOT EDIT.")
	fmt.Fprintf(w, "// Version: %s\n", goscrapper.VERSION)
	fmt.Fprintf(w, "// Date: %s\n", goscrapper.BUILDDATE)
	fmt.Fprintf(w, "// Built : %s\n", goscrapper.GITHASH)
	fmt.Fprintf(w, "// %s\n", goscrapper.COPYRIGHT)
	fmt.Fprintln(w)

	for _, s := range ss {
		fmt.Fprintln(w, s)
	}
	fmt.Fprintln(w)
}

// Genarate a unique id as a string
func UID() string {
	uid += 1
	return fmt.Sprintf("%06x", uid)
}

var uid = 1
