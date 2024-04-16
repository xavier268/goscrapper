package parser

import (
	"reflect"
	"testing"
)

func TestParseLitteral(t *testing.T) {

	data := []struct {
		s string
		a any
	}{
		{"true", true},
		{"false", false},
		{"NIL", nil},
		{"nil", nil},
		{"1", 1},
		{"-1", -1},
		{`'abc'`, "abc"},
		{`"abc"`, "abc"},
		{`"ab""c"`, "ab\"c"},
		{`'ab''c'`, "ab'c"},
		{"1//", 1},
		{"1/*iuy*/", 1},

		{"[]", []any{}},
		{"[-5]", []any{-5}},
		{"[1,'deux', true, nil, -3]", []any{1, "deux", true, nil, -3}},

		{"{}", map[string]any{}},
		{"{aa:-2}", map[string]any{"aa": -2}},
		{"{aa:-2, bb:nil, cc:[1,true, 'toto']}", map[string]any{"aa": -2, "bb": nil, "cc": []any{1, true, "toto"}}},
	}

	for _, d := range data {
		l, err := ParseLitteral(d.s)
		if err != nil {
			t.Fatalf("parseLitteral(%s)  : %v", d.s, err)
		}
		if !reflect.DeepEqual(l, d.a) {
			t.Errorf("parseLitteral(%s) = %v, but wanted %#v", d.s, l, d.a)
		}
	}
}
