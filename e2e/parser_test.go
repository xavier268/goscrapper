// contains end-to-end parser tests
package e2e

import (
	"fmt"
	"testing"

	"github.com/xavier268/goscrapper/parser"
)

func TestParserVisual(t *testing.T) {

	err := parser.ParseFiles("e2epack", "test1.sc")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
