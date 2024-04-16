package examples

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/xavier268/goscrapper/parser"
)

func TestAllExamples(t *testing.T) {

	files, err := filepath.Glob("./*.gsc")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		t.Logf("Running example %s", file)
		// skip examples requiring input
		if strings.Contains(file, ".input.") {
			t.Logf("%sSkipping example %s%s", parser.ColYELLOW, file, parser.AnsiRESET)
			continue
		}
		_, err := parser.EvalFile(file)
		if err != nil {
			t.Fatal(parser.ColRED, err, parser.AnsiRESET)
		}
		t.Logf("Example %s passed", file)
	}

}
