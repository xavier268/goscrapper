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
		RETURN x ; // 69
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
			RETURN x , y;
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
			RETURN x, y,z; // {a:1, b:2}, ["a a", "bb", true, 66, {a:1, b:2}] , {x:{a:1, b:2}, y:["a a", "bb", true, 66, {a:1, b:2}]}
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
		}, {
			req: `
			a=1;	
			FOR b FROM 4 TO 10 STEP 3;
				FOR c FROM -1 TO -3 STEP -1;
					PRINT b,c;
					RETURN {first: a,second:b, third :c} ; 
					`,
			params: nil,
		}, {
			req: `
			// test map access
			a = {un:1, deux:"dd", trois:true};
			RETURN a.un, a.deux, a.trois, a.quatre; // a.quatre should give nil/null with no error
					`,
			params: nil,
		}, {
			req: `
			// test map access on non map
			b = [1,5];
			// c = b.trois; // should fail at execution if uncommented
			RETURN ;
					`,
			params: nil,
		}, {
			req: `
			// test array access
			a = "xxx";
			b = [1+2,5,a+a];
			RETURN b[0], b[1], b[1+1] ;
					`,
			params: nil,
		}, {
			req: `// test empty array and objects
			a = [];
			b = {};
			// c = a[0]; // should fail "out of bound" at execution if uncommented
			d = b.two ; // should not fail, justr null/nil
			RETURN a, b, b.six; // should not fail
					`,
			params: nil,
		}, {
			req: `
			// test PAGE & TEXT
			wk = PAGE "http://www.wikipedia.fr";
			empty = PAGE "";
			t = TEXT wk;
			RETURN t ;
					`,
			params: nil,
		}, {
			req: `
			// test VERSION , SLOW & NOW
			t = NOW;
			v = "Version is " + VERSION;
			s = "File separator is " + FILE_SEPARATOR;
			SLOW;
			RETURN t != NOW , t == NOW , t == t ; // true, false, true
					`,
			params: nil,
		}, {
			req: `
			// test IF THEN ELSE
			IF true THEN a=1 ELSE a=2 ;
			IF false THEN b=1 ELSE b=2 ;
			IF true THEN (IF true THEN c=3 ELSE c=4; ) ELSE c=5 ;		
			IF true THEN 
				IF true THEN d=4 ELSE d=5  ;		
			IF false THEN e=4 ELSE ( IF true THEN e=5 ELSE e=6 ;);
			IF false THEN f=5 ELSE  
				IF true THEN f=6 ELSE f=7 ;
			RETURN a , b, c , d , e , f; // 1 2 3 4 5 6
					`,
			params: nil,
		}, {
			req: `// Reserved
			RETURN;
					`,
			params: nil,
		}, {
			req: `// Reserved
			RETURN;
					`,
			params: nil,
		}, {
			req: `// Reserved
			RETURN;
					`,
			params: nil,
		}, {
			req: `// Reserved
			RETURN;
					`,
			params: nil,
		},
	}

	for i, req := range data {
		rs := fmt.Sprintf("%s-req#%d", t.Name(), i)
		fmt.Println("--- Testing", rs)
		fmt.Println(req.req)
		fmt.Println("\n--- Compiling", rs)
		root, err := parser.Compile(rs, req.req)
		fmt.Printf("COMPILED : \n%#v\n\nCompilation error : %s%v%s\n", root, parser.ColRED, err, parser.AnsiRESET)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("\n--- Executing", rs, "with", req.params)
		it := parser.NewInterpreter(context.Background()).With(req.params)
		res, err := it.Eval(root)
		fmt.Printf("EXECUTION RESULT : %s%s%s\n\nExecution error :%s%v%s\n",
			parser.ColGREEN, parser.PrettyJson(res), parser.AnsiRESET, parser.ColRED, err, parser.AnsiRESET)
		if err != nil {
			t.Fatal(err)
		}
		it.DumpVars(os.Stdout, "\n--- Dumping vars "+rs)
	}

}
