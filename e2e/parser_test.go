package e2e

import (
	"context"
	"fmt"
	"testing"

	"github.com/xavier268/goscrapper/internal/parser"
)

func TestParser(t *testing.T) {

	reqs := []string{
		`
		PRINT 1 ;
		RETURN 2 ;
		`,
	}

	for i, rq := range reqs {
		rs := fmt.Sprintf("%s-req#%d", t.Name(), i)
		fmt.Println("--- Compiling", rs)
		root, ins, err := parser.Compile(rs, rq)
		fmt.Printf("compiled reqst : %#v\nparams : %v\nerr : %v\n", root, ins, err)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("--- Executing", rs)
		it := parser.NewInterpreter(context.Background())
		res, err := it.Eval(root)
		fmt.Println(res, err)
		fmt.Println()
		if err != nil {
			t.Fatal(err)
		}
	}

}
