// contains end-to-end parser tests
package e2e

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/xavier268/goscrapper/internal/parser"
	"github.com/xavier268/mytest"
)

// generate compiled files for visual examination.
func TestParserCompile(t *testing.T) {

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

	// t.Skip() // uncomment to monitor any changes in compiled files.

	// verify changes in compiled files
	ff, err := filepath.Glob(filepath.Join("e2epack", "test*.go"))
	if err != nil {
		panic(err)
	}
	for _, f := range ff {
		mytest.VerifyFile(t, f)
	}

	// copy call1 test to e2epack
	target, _ := os.Create(filepath.Join("e2epack", "call1_test.go"))
	defer target.Close()
	source, _ := os.Open("test.go.tpl")
	defer source.Close()
	io.Copy(target, source)

	// make a doc.go file to document the package
	target, _ = os.Create(filepath.Join("e2epack", "doc.go"))
	defer target.Close()
	target.WriteString("//autogenerated end-to-end testing package. DO NOT EDIT.\n// Generated on ")
	target.WriteString(time.Now().String())
	target.WriteString("\npackage e2epack\n")
}
