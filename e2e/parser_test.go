// contains end-to-end parser tests
package e2e

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/xavier268/goscrapper/parser"
	"github.com/xavier268/mytest"
)

// generate compiled files for visual examination.
func TestParserVisual(t *testing.T) {

	// generate the compiled files
	err := parser.ParseGlob("e2epack", "*.sc")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	err = parser.ParseGlobAsync("e2epack", "*.sc")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	t.Skip() // uncomment to monitor any changes in compiled files.

	// verify changes in compiled files
	ff, err := filepath.Glob(filepath.Join("e2epack", "*.go"))
	if err != nil {
		panic(err)
	}
	for _, f := range ff {
		mytest.VerifyFile(t, f)
	}

}

func TestCall1Sync(t *testing.T) {

	err := parser.ParseGlob(".", "call1.sc")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	r, err := Do_call1(context.Background(), Input_call1{
		a: 3,
		b: 7,
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("Result call1 : %#v\n", r)
	if len(r) != 1 {
		t.Fatal("unexpected length")
	}
	if r[0].c != 10 {
		t.Fatal("invalid result")
	}
}

func TestCall1Async(t *testing.T) {

	err := parser.ParseGlobAsync(".", "call1.sc")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	ch := make(chan Output_call1_async, 5)
	defer close(ch)

	err = DoAsync_call1_async(context.Background(), ch, Input_call1_async{
		a: 3,
		b: 7,
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	out, open := <-ch
	if !open {
		t.Fatal("channel was closed unexpectedly")
	}
	fmt.Printf("Result call1_async : %#v\n", out)
	if out.c != 10 {
		t.Fatal("invalid result")
	}
}
