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
	// RETURN LAST or DISTINCT ...
	$a = 10;
	FOR i FROM 1 TO 10 ;
		$a = $a + i;
		// RETURN  $a%4;				// [[3], [1], [0], [0], [1], [3], [2], [2], [3], [1]]
		// RETURN DISTINCT $a%4;		// [[3], [1], [0], [2]]
		RETURN LAST a%4;				// [[1]]
		`
	buff := new(strings.Builder)

	buff.WriteString(data)

	root, err := Compile(t.Name(), data)
	fmt.Fprintf(buff, "\nCOMPILING :\nRoot : %#v\nCompilation error : %s%v%s\n", root, ColRED, err, AnsiRESET)
	if err != nil {
		t.Fatal(err)
	}
	it := NewInterpreter(context.Background())
	res, err := it.Eval(root)
	fmt.Fprintf(buff, "\nEXECUTING :\nResult :%#v\nResult :%s%s%s\nExecution error :%s%v%s\n", res, ColGREEN, rt.MustSerialize(res), AnsiRESET, ColRED, err, AnsiRESET)
	it.DumpVars(buff, "--- Dumping vars "+t.Name())

	// for visual control
	fmt.Println(buff.String())

	// for non regression control
	// mytest.Verify(t, buff.String(), t.Name())
	time.Sleep(5 * time.Second)
}
