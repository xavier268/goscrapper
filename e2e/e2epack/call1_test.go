// autogenerated package. DO NOT EDIT.
package e2epack

import(
    "fmt"
    "time"
    "testing"
    "context"
    "github.com/xavier268/goscrapper/internal/parser"
)

// Basic test template for call1
// copied verbatim to e2epack from e2e

func TestCall1Sync(t *testing.T) {

	r, err := Do_call1(context.Background(), Input_call1{
		a: 3,
		b: 7,
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("Result call1 : %s\n", parser.Pretty(r))
	if len(r) == 0 {
		t.Fatal("unexepected empty result")
	}
	if r[0].c != 10 {
		t.Fatal("invalid result")
	}
}

func TestCall1Async(t *testing.T) {

	ch := make(chan Output_call1_async, 20)
	go func() { // wait 5 sec and close channel
		time.Sleep(5 * time.Second)
		close(ch)
	}()

	err := DoAsync_call1_async(context.Background(), ch, Input_call1_async{
		a: 3,
		b: 7,
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for out := range ch {
		fmt.Printf("Result call1_async : %s\n", parser.Pretty(out))
		if out.c != 10 {
			t.Fatal("invalid result")
		}
	}

}