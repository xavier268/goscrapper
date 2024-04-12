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
	// test IF THEN ELSE
	IF true THEN a=1 ELSE a=2 ;
	IF false THEN b=1 ELSE b=2 ;
	IF true THEN (IF true THEN c=3 ELSE c=4; ) ELSE c=5 ;		
	IF true THEN 
		IF true THEN d=4 ELSE d=5  ;		
	IF false THEN e=4 ELSE ( IF true THEN e=5 ELSE e=6 ;);
	IF false THEN f=5 ELSE  
		IF true THEN f=6 ELSE f=7 ;
	RETURN a , b, c , d , e , f; // 1 2 3 4 5
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
