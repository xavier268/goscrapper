package parser

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/xavier268/mytest"
)

func TestParserLab(t *testing.T) {

	data := `
	a=100;
	b = "deux";
	b2 = "un " + 'deux';
	c = true;
	d = false;
	e = [1,2,3];
	f = {a:1, b:2*3};
	g = {a:e, b:f};
	h = [g,g,g];
	z1 = 2*4+6-5;
	b1 = (b == "deux") OR false;
	PRINT a, +a,-a,++a,--a;
	PRINT z1 == 9; // should be true
	RETURN ;
			`
	buff := new(strings.Builder)

	buff.WriteString(data)

	root, ins, err := Compile(t.Name(), data)
	fmt.Fprintf(buff, "compiled reqst : %#v\nparams : %v\nerr : %v\n", root, ins, err)
	if err != nil {
		t.Fatal(err)
	}
	it := NewInterpreter(context.Background())
	res, err := it.Eval(root)
	fmt.Fprintf(buff, "result : %#v\nerr :%v\n", res, err)
	it.DumpVars(buff, "--- Dumping vars "+t.Name())

	mytest.Verify(t, buff.String(), t.Name())
}
