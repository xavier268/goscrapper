package parser

import "testing"

func TestCmp(t *testing.T) {
	data := []struct {
		a   any
		b   any
		res int
		ok  bool
	}{
		{nil, nil, 0, true},
		{1, 2, -1, true},
		{"zzz", "aaa", 1, true},
		{nil, "deux", 0, false},
		{nil, 2, 0, false},
		{[]string{"zzz"}, []string{"zzz"}, 0, true},
		{[]string{"zzz"}, []string{"aaa"}, 0, false},
	}

	for _, d := range data {

		i, err := CompareAny(d.a, d.b)

		// direct test
		if err != nil && d.ok {
			t.Fatalf("Unexpected error compare %v %v => %d %v", d.a, d.b, i, err)
		}
		if err == nil && !d.ok {
			t.Fatalf("Compare %v %v => %d should fail,but did not", d.a, d.b, i)
		}
		if i != d.res {
			t.Fatalf("Unexpected compare result %v %v => %d", d.a, d.b, i)
		}

		// reverse order, and repeat
		d.a, d.b, d.res = d.b, d.a, -d.res
		i, err = CompareAny(d.a, d.b)

		if i != d.res {
			t.Fatalf("Unexpected compare result %v %v => %d", d.a, d.b, i)
		}
		if err != nil && d.ok {
			t.Fatalf("Unexpected error compare %v %v => %d %v", d.a, d.b, i, err)
		}
		if err == nil && !d.ok {
			t.Fatalf("Compare %v %v => %d should fail,but did not", d.a, d.b, i)
		}

	}
}
