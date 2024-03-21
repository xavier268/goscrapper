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
	select "input[name=btnK]"	 // <== Keywords are alse recognized in full lowercase (no mixed case !!)
	CLICK "input[name=btnK]"
	RETURN "input[name=btnK]"
	`

	err := parser.Parse(strings.NewReader(input), os.Stdout)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
