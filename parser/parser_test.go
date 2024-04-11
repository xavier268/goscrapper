package parser

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestParserVisual(t *testing.T) {

	data := `
	a=1;	
	FOR b FROM 4 TO 10 STEP 3;
		FOR c FROM -1 TO -3 STEP -1;
			PRINT b,c;
			RETURN {first: a,second:b} ; 
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
