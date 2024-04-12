package parser

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/xavier268/goscrapper/rt"
)

func TestParserVisual(t *testing.T) {

	data := `
	// test ASSERT and FAIL
	// ASSERT 1 == 2; 	// should fail
	// ASSERT 30 ; 	// should fail
	ASSERT 1 == 1; 	// should not fail
	// FAIL "user message";	// should fail
	// FAIL ;			// should fail
	RETURN ;
			`
	buff := new(strings.Builder)

	buff.WriteString(data)

	root, err := Compile(t.Name(), data)
	fmt.Fprintf(buff, "\nCOMPILING :\nRoot : %#v\nRoot : %s\nCompilation error : %s%v%s\n", root, rt.Pretty(root), ColRED, err, AnsiRESET)
	if err != nil {
		t.Fatal(err)
	}
	it := NewInterpreter(context.Background())
	res, err := it.Eval(root)
	fmt.Fprintf(buff, "\nEXECUTING :\nResult :%#v\nResult :%s%s%s\nExecution error :%s%v%s\n", res, ColGREEN, rt.Pretty(res), AnsiRESET, ColRED, err, AnsiRESET)
	it.DumpVars(buff, "--- Dumping vars "+t.Name())

	// for visual control
	fmt.Println(buff.String())

	// for non regression control
	// mytest.Verify(t, buff.String(), t.Name())
	time.Sleep(5 * time.Second)
}
