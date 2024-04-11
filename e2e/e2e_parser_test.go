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
		}, {
			req: `
			// combined data structures
			x = {a:1, b:2};
			y = ["a a", "bb", true, 66, x];
			z = {x:x, y:y};
			PRINT x, y,z;
			RETURN ;
			`,
			params: nil,
		}, {
			req: `
			// unary tests
			a=1;
			b = "deux";
			c = true;
			d = false;
			e = [1,2,3];
			f = {a:1, b:2};
			g = {a:e, b:f};
			h = [g,g,g];
			z1 = 2*4+6-5;
			b1 = (b == "deux") OR false;
			PRINT a, +a,-a,++a,--a;
			PRINT z1 == 9; // should be true
			RETURN ;
			`,
			params: nil,
		}, {
			req: `
			// returning values
			a=1;
			b = "deux";
			c = true;
			d = [1,2,3];
			e = {a:1, b:2};
			f = {a:e, b:d};
			g =  [a,f];
			RETURN a,b,c,d,e,f,g,5*6+4,"fin du test";
			`,
			params: nil,
		}, {
			req: `
			a=1;	
			RETURN {first: a,second:a+3} ; 
			`,
			params: nil,
		},
	}

	for i, req := range data {
		rs := fmt.Sprintf("%s-req#%d", t.Name(), i)
		fmt.Println("--- Testing", rs)
		fmt.Println(req.req)
		fmt.Println("--- Compiling", rs)
		root, err := parser.Compile(rs, req.req)
		fmt.Printf("COMPILED : \n%#v\nCompilation error : %s%v%s\n", root, parser.ColRED, err, parser.RESET)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("--- Executing", rs, "with", req.params)
		it := parser.NewInterpreter(context.Background()).With(req.params)
		res, err := it.Eval(root)
		fmt.Printf("EXECUTION RESULT : %s%s%s\nExecution error :%v\n", parser.ColGREEN, parser.PrettyJson(res), parser.RESET, err)
		if err != nil {
			t.Fatal(err)
		}
		it.DumpVars(os.Stdout, "--- Dumping vars "+rs)
	}

}
