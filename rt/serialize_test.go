package rt

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-rod/rod"
)

func TestSerialize(t *testing.T) {

	dataOk := []struct {
		d any
		s string
	}{
		{d: nil, s: "nil"},
		{d: true, s: "true"},
		{d: false, s: "false"},
		{d: 1, s: "1"},
		{d: `something like this`, s: `"something like this"`},
		{d: `something "" like this`, s: `"something """ like this"`},
		{d: `something '' like this`, s: `"something '' like this"`},
		{d: true, s: "true"},

		{d: []any{1, 2, 3}, s: "[1, 2, 3]"},
		{d: map[string]any{"un": 1, "deux": "something like this", "trois": true}, s: `{deux: "something like this", trois: true, un: 1}`}, // sorting !

		{d: time.Time{}, s: "time{0001-01-01T00:00:00Z}"},

		{d: new(rod.Page), s: "page{--}"},
		{d: new(rod.Element), s: "element{}"},

		{d: Hash{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, s: "hash{0102030405060708090a0b0c0d0e0f10}"},
	}

	dataNotOk := []any{
		[]int{1, 2, 3},          // should be []any, not []int
		[]any{1, rod.Page{}, 3}, // no rod.Page, but *rod.Page
		map[int]int{1: 2, 3: 4},
		map[string]any{"invalid_key": 1, "b": false},
	}

	for _, dd := range dataOk {
		fmt.Println(dd.d)
		s, err := Serialize(dd.d)
		if err != nil {
			t.Fatalf("Unexpected error serializing %#v : %v", dd.d, err)
		}
		if s != dd.s {
			t.Fatalf("Unexpected result serializing %#v : got <%s> - want <%s>", dd.d, s, dd.s)
		}
	}

	fmt.Println()

	for _, dd := range dataNotOk {
		fmt.Println(dd)
		_, err := Serialize(dd)
		if err == nil {
			t.Fatalf("Expected error serializing %v did not happen", dd)
		}
	}
}
