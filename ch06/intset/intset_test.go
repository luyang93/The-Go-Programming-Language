package intset

import (
	"testing"
)

func TestIntSet_Add(t *testing.T) {
	var tcs = []struct {
		words   []int
		expects []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 101}, []int{6, 137438953472}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 1024}, []int{510, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for _, tc := range tcs {
		var s IntSet
		for _, w := range tc.words {
			s.Add(w)
		}

		for i := 0; i < len(tc.expects); i++ {
			if s.words[i] != uint64(tc.expects[i]) {
				t.Errorf("IntSet Add, Expect: %v, Actual: %v", tc.expects, s)
			}
		}
	}
}

func TestIntSet_UnionWith(t *testing.T) {
	var tcs = []struct {
		s       []int
		t       []int
		expects string
	}{
		{[]int{}, []int{1}, "{1}"},
		{[]int{1}, []int{2}, "{1 2}"},
		{[]int{1, 2}, []int{2}, "{1 2}"},
		{[]int{2}, []int{1, 2}, "{1 2}"},
		{[]int{1, 2}, []int{3, 4}, "{1 2 3 4}"},
	}

	for _, tc := range tcs {
		s := &IntSet{}
		o := &IntSet{}
		for _, i := range tc.s {
			s.Add(i)
		}
		for _, i := range tc.t {
			o.Add(i)
		}

		s.UnionWith(o)

		if s.String() != tc.expects {
			t.Errorf("IntSet Add, Expect: %v, Actual: %v", tc.expects, s.String())
		}
	}
}
func TestIntSet_Has(t *testing.T) {
	var tcs = []struct {
		words   []int
		has     int
		expects bool
	}{
		{[]int{}, 1, false},
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 4, false},
	}

	for _, tc := range tcs {
		var s IntSet
		for _, w := range tc.words {
			s.Add(w)
		}

		if s.Has(tc.has) != tc.expects {
			t.Errorf("IntSet Has, Expect: %v, Actual: %v", tc.expects, s.Has(tc.has))
		}
	}
}

func TestIntSet_String(t *testing.T) {
	var tcs = []struct {
		words   []int
		expects string
	}{
		{[]int{}, "{}"},
		{[]int{1, 2, 101}, "{1 2 101}"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 1024}, "{1 2 3 4 5 6 7 8 1024}"},
	}

	for _, tc := range tcs {
		var s IntSet
		for _, w := range tc.words {
			s.Add(w)
		}

		if s.String() != tc.expects {
			t.Errorf("IntSet IntSet, Expect: %v, Actual: %v", tc.expects, s.String())
		}
	}
}
