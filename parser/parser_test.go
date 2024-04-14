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
	// Printing and formating
	a = [ 1, "deux", false, nil ];
	b = { z : 1, b : "deux", c : false, d : nil };
	c = { a: a, b : b , c:[ a,b]};
	PRINT ;		
		// empty line
	PRINT a;	
		// [1 deux false <nil>]
	PRINT a,b;	
		// [1 deux false <nil>]map[b:deux c:false d:<nil> z:1]
	PRINT a,b,c;
		// [1 deux false <nil>]map[b:deux c:false d:<nil> z:1]map[a:[1 deux false <nil>] b:map[b:deux c:false d:<nil> z:1] c:[[1 deux false <nil>] map[b:deux c:false d:<nil> z:1]]]
	
	PRINT "---";
	
	PRINT "Raw format : ",  RAW c ;	
		// Raw format : map[string]interface {}{"a":[]interface {}{1, "deux", false, interface {}(nil)}, "b":map[string]interface {}{"b":"deux", "c":false, "d":interface {}(nil), "z":1}, "c":[]interface {}{[]interface {}{1, "deux", false, interface {}(nil)}, map[string]interface {}{"b":"deux", "c":false, "d":interface {}(nil), "z":1}}}
	PRINT "Json format : ", JSON c ;
		// Json format : {"a":[1,"deux",false,null],"b":{"b":"deux","c":false,"d":null,"z":1},"c":[[1,"deux",false,null],{"b":"deux","c":false,"d":null,"z":1}]}
	PRINT "Go format : ",GO c; 		
		// Go format : map[a:[1 deux false <nil>] b:map[b:deux c:false d:<nil> z:1] c:[[1 deux false <nil>] map[b:deux c:false d:<nil> z:1]]]
	// PRINT "GSC format : ", GSC c ;

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
