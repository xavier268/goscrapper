// contains end-to-end parser tests
package e2e

import (
	"fmt"
	"testing"

	"github.com/xavier268/goscrapper/parser"
)

func TestParserVisual(t *testing.T) {
	err := parser.ParseGlob("e2epack", "*.sc")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
