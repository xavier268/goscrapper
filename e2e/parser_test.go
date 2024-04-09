package e2e

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/xavier268/goscrapper/parser"
)

func TestParserFull(t *testing.T) {

	data := []struct {
		req    string
		params map[string]any
	}{
		{
			req: `
			// basic arithm & var playing
		// z = 5; // this should fail later, when z will be (re)declared as input variable
		x = 3 + @z;
		// z = 5; // this would fail now, since z is already an input variable
		PRINT 1 ;
		PRINT RAW 1;
		PRINT 2, x , @z ; // note that z should remain the input variable
		RETURN ;
		`,
			params: map[string]any{"z": 66},
		},
		{
			req: `
			// arrays
			x = [1,2];
			y = [x,x];
			PRINT x, y;
			PRINT RAW x,y;
			RETURN ;
			`,
			params: nil,
		},
		{
			req: `
			// objects 
			x = {a:1, b:2};
			y = {c:x,d:x};
			PRINT x, y;
			PRINT RAW x, y;
			RETURN ;
			`,
			params: nil,
		},
	}

	for i, req := range data {
		rs := fmt.Sprintf("%s-req#%d", t.Name(), i)
		fmt.Println("--- Testing", rs)
		fmt.Println(req.req)
		fmt.Println("--- Compiling", rs)
		root, ins, err := parser.Compile(rs, req.req)
		fmt.Printf("compiled reqst : %#v\nparams : %v\nerr : %v\n", root, ins, err)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("--- Executing", rs, "with", req.params)
		it := parser.NewInterpreter(context.Background()).With(req.params)
		res, err := it.Eval(root)
		fmt.Printf("result : %#v\nerr :%v\n", res, err)
		if err != nil {
			t.Fatal(err)
		}
		it.DumpVars(os.Stdout, "--- Dumping vars "+rs)
	}

}

func TestParserCompile(t *testing.T) {

	data := `
			// objects - BUGGY !
			x = {a:1, b:2};
			PRINT x;
			RETURN ;
			`
	root, ins, err := parser.Compile(t.Name(), data)
	fmt.Printf("compiled reqst : %#v\nparams : %v\nerr : %v\n", root, ins, err)
	if err != nil {
		t.Fatal(err)
	}
}
