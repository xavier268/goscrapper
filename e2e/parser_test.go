// contains end-to-end parser tests
package e2e

import (
	"fmt"
	"io"
	"os"
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

	// look for changes in compiled files
	ff, err := filepath.Glob(filepath.Join("e2epack", "*.go"))
	fmt.Println(ff)
	if err != nil {
		panic(err)
	}
	for _, f := range ff {
		fmt.Printf("Verifying %s\n", f)
		r, err := os.Open(f)
		if err != nil {
			panic(err)
		}
		bb, err := io.ReadAll(r)
		if err != nil {
			panic(err)
		}
		r.Close()
		mytest.Verify(t, string(bb), f)
	}

}
