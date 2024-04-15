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
	// Select specific elements in a page

	page = PAGE "http://www.wikipedia.fr" ;
	SELECT "div" AS loop FROM page LIMIT 5;    
		PRINT "**** looping ...*****" , NL, "Captured text : ",  TEXT loop ;
    	RETURN LAST "done" ;
	`
	buff := new(strings.Builder)

	buff.WriteString(data)

	root, err := Compile(t.Name(), data)
	fmt.Fprintf(buff, "\nCOMPILING :\nRoot : %#v\n", root)
	if err != nil {
		fmt.Fprintf(buff, "\nCompilation error :%s%v%s\n", ColRED, err, AnsiRESET)
		t.Fatal(err)
	}
	it := NewInterpreter(context.Background())
	res, err := it.Eval(root)
	fmt.Fprintf(buff, "\nEXECUTING :\nResult :%#v\nResult :%s%s%s\n", res, ColGREEN, rt.MustSerialize(res), AnsiRESET)
	if err != nil {
		fmt.Fprintf(buff, "\nExecution error :%s%v%s\n", ColRED, err, AnsiRESET)
	}
	it.DumpVars(buff, "--- Dumping vars "+t.Name())

	// for visual control
	fmt.Println(buff.String())

	// for non regression control
	// mytest.Verify(t, buff.String(), t.Name())
	time.Sleep(5 * time.Second)
}
