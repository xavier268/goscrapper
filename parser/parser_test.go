package parser

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestParserLab(t *testing.T) {

	data := `
	a=1;b=2;c=3;d=4;e=5;
	
	RETURN {first: a,second:b} ;
			`
	buff := new(strings.Builder)

	buff.WriteString(data)

	root, err := Compile(t.Name(), data)
	fmt.Fprintf(buff, "compiled reqst : %#v\nerr : %v\n", root, err)
	if err != nil {
		t.Fatal(err)
	}
	it := NewInterpreter(context.Background())
	res, err := it.Eval(root)
	fmt.Fprintf(buff, "interpreted reqst : %s\nerr :%v\n", PrettyJson(res), err)
	it.DumpVars(buff, "--- Dumping vars "+t.Name())

	// for visual control
	fmt.Println(buff.String())

	// for non regression control
	// mytest.Verify(t, buff.String(), t.Name())
}
