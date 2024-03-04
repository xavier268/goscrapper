package config

import (
	"fmt"
	"path/filepath"
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

	tf1 := filepath.Join("testfiles", "test1.yml")
	tf2 := filepath.Join("testfiles", "test2.yml")

	c, err := ParseDefinitions(tf1, tf2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("\nCaptured configuration :\n%s\n", Pretty(c))
}
