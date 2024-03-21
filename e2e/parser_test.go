// contains end-to-end parser tests
package e2e

import (
	"os"
	"strings"
	"testing"

	"github.com/xavier268/goscrapper/parser"
)

func TestParserVisual(t *testing.T) {

	input := `
	HEADLESS
	DOCUMENT "http://www.google.fr"
	SELECT "input[name=q]"
	TYPE "test"
	SELECT "input[name=btnK]"
	TYPE "test"
	CLICK "input[name=btnK]"
	RETURN "input[name=btnK]"
	`
	err := parser.Parse(strings.NewReader(input), os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}
