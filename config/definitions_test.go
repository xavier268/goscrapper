package config

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseVersion(t *testing.T) {

	tdata := []struct {
		v string
		M int
		m int
		p int
	}{
		{"1.2.3", 1, 2, 3},
		{"v1.2.3", 1, 2, 3},
		{"v1.2.3.beta", 1, 2, 3},
		{"v1.2.3beta", 1, 2, 3},
		{"1.2.", 1, 2, 0},
		{"1.2", 1, 2, 0},
		{"v1.2", 1, 2, 0},
		{"v1.2beta", 1, 2, 0},
		{"1.2", 1, 2, 0},
		{"1.", 1, 0, 0},
		{"1", 1, 0, 0},
		{"v1", 1, 0, 0},
		{"v1beta", 1, 0, 0},
		{"v1.beta", 1, 0, 0},
		{"v1.2.3.4", 1, 2, 3},
	}

	for _, v := range tdata {
		M, m, p := parseVersion(v.v)
		if !(M == v.M && m == v.m && p == v.p) {
			t.Errorf("parseVersion(%v) = %v,%v,%v, want %v,%v,%v", v.v, M, m, p, v.M, v.m, v.p)
		}
	}
}

func TestParseConfiguration(t *testing.T) {

	testfile := "test.yaml"

	c, err := ParseDefinitions(testfile)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("\nCaptured the following definitions from file %s\n%s\n%s\n%s\n",
		testfile,
		strings.Repeat("=", 40),
		pretty(c),
		strings.Repeat("=", 40))
}
