package parser

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestParserLab(t *testing.T) {

	data := `
	a=1;	
	RETURN {first: a,second:5+a} ; 
			`
	buff := new(strings.Builder)

	buff.WriteString(data)

	root, err := Compile(t.Name(), data)
	fmt.Fprintf(buff, "\nCOMPILING : %#v\nCompilation error : %s%v%s\n", root, ColRED, err, RESET)
	if err != nil {
		t.Fatal(err)
	}
	it := NewInterpreter(context.Background())
	res, err := it.Eval(root)
	fmt.Fprintf(buff, "\nEXECUTING : %s%s%s\nExecution error :%s%v%s\n", ColGREEN, PrettyJson(res), RESET, ColRED, err, RESET)
	it.DumpVars(buff, "--- Dumping vars "+t.Name())

	// for visual control
	fmt.Println(buff.String())

	// for non regression control
	// mytest.Verify(t, buff.String(), t.Name())
}
