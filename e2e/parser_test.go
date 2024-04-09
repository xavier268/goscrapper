package e2e

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/xavier268/goscrapper/parser"
)

func TestParser(t *testing.T) {

	reqs := []string{
		`
		// z = 5; // this should fail later, when z will be decalred as input variable
		x = 3 + @z;
		// z = 5; // this would fail, since z is an input variable
		PRINT 1 ;
		PRINT 2, x , z ;
		RETURN ;
		`,
	}

	for i, rq := range reqs {
		rs := fmt.Sprintf("%s-req#%d", t.Name(), i)
		fmt.Println("--- Testing", rs)
		fmt.Println(rq)
		fmt.Println("--- Compiling", rs)
		root, ins, err := parser.Compile(rs, rq)
		fmt.Printf("compiled reqst : %#v\nparams : %v\nerr : %v\n", root, ins, err)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("--- Executing", rs)
		it := parser.NewInterpreter(context.Background())
		res, err := it.Eval(root)
		fmt.Printf("result : %#v\nerr :%v\n", res, err)
		if err != nil {
			t.Fatal(err)
		}
		it.DumpVars(os.Stdout, "--- Dumping vars "+rs)
	}

}
