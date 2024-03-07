package generator

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestParseConfiguration(t *testing.T) {

	tf1 := filepath.Join("..", "testfiles", "test1.yml")
	tf2 := filepath.Join("..", "testfiles", "test2.yml")

	c := NewCompiler()
	err := c.Parse(tf1, tf2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("\nCaptured configuration :\n%s\n", PrettyJson(c.conf))
}
