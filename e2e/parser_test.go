// contains end-to-end parser tests
package e2e

import (
	"fmt"
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
	SELECT "input[name=btnK]"	 
	CLICK "input[name=btnK]"
	RETURN "input[name=btnK]"
	`

	err := parser.Parse(os.Stdout, strings.NewReader(input))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
