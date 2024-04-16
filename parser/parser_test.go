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
	// variables and their scopes
	// inp named parameter must be set before executing the request, or a runtime error will happen :
	//          NewInterpreter(context.Background()).With(map[string]any{"inp": 33})
	
	ii = 3 ;                // ERROR : ii is declared and assigned to,
	// wrong = @ ii ;       // ii being alrerady a known variable, cannot be declared as an input parameter
	$a = 100 + @ inp ;      // a is forced to global scope ; inp is registerd as input parameter
	// inp = 23 ;           // ERROR : assigning to an input parameter
	// b = a + inp ;        // ERROR : inp was never declared as a "normal" variable
	b = a + @inp ;          // OK : inp was declared as an input parameter
	FOR i FROM 1 TO 5 ;     // i declared as local (loop) variable
		e = b + i + @inp ;  // b and i are read from local scope, inp is read from input parameter, e is assigned to local scope
		RETURN e ;	`
	buff := new(strings.Builder)

	buff.WriteString(data)

	root, err := Compile(t.Name(), data)
	fmt.Fprintf(buff, "\nCOMPILING :\nRoot : %#v\n", root)
	if err != nil {
		fmt.Fprintf(buff, "\nCompilation error :%s%v%s\n", ColRED, err, AnsiRESET)
		t.Fatal(err)
	}
	it := NewInterpreter(context.Background()).With(map[string]any{"inp": 33})
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
