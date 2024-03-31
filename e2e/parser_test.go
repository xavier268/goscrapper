// contains end-to-end parser tests
package e2e

import (
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
