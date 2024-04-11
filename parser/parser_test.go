package parser

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestParserVisual(t *testing.T) {

	data := `
	// test array access
	b = [1+2,5];
	c = b[0];
	RETURN ;
			`
	buff := new(strings.Builder)

	buff.WriteString(data)

	root, err := Compile(t.Name(), data)
	fmt.Fprintf(buff, "\nCOMPILING : %#v\nCompilation error : %s%v%s\n", root, ColRED, err, AnsiRESET)
	if err != nil {
		t.Fatal(err)
	}
	it := NewInterpreter(context.Background())
	res, err := it.Eval(root)
	fmt.Fprintf(buff, "\nEXECUTING : %s%s%s\nExecution error :%s%v%s\n", ColGREEN, PrettyJson(res), AnsiRESET, ColRED, err, AnsiRESET)
	it.DumpVars(buff, "--- Dumping vars "+t.Name())

	// for visual control
	fmt.Println(buff.String())

	// for non regression control
	// mytest.Verify(t, buff.String(), t.Name())
}
