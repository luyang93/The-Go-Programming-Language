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
			if s.words[i] != uint(tc.expects[i]) {
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

func TestIntSet_Len(t *testing.T) {
	var tcs = []struct {
		words   []int
		expects int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 101}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 1024}, 9},
	}

	for _, tc := range tcs {
		var s IntSet
		for _, w := range tc.words {
			s.Add(w)
		}
		if s.Len() != tc.expects {
			t.Errorf("Len of %v, Actual: %d, Expects %d", tc.words, s.Len(), tc.expects)
		}
	}
}

func TestIntSet_Remove(t *testing.T) {
	var tcs = []struct {
		words   []int
		remove  int
		expects string
	}{
		{[]int{}, 1, "{}"},
		{[]int{1, 2, 3}, 2, "{1 3}"},
	}

	for _, tc := range tcs {
		var s IntSet
		for _, w := range tc.words {
			s.Add(w)
		}

		s.Remove(tc.remove)

		if s.String() != tc.expects {
			t.Errorf("Remove %d from %v, Actual = %s, Expects %s", tc.remove, tc.words, s.String(), tc.expects)
		}
	}
}

func TestIntSet_Clear(t *testing.T) {
	var tcs = []IntSet{
		IntSet{},
		IntSet{words: []uint{1, 2, 3}},
	}

	for _, tc := range tcs {
		tc.Clear()
		if tc.Len() != 0 {
			t.Errorf("Clear failed. IntSet: %q", tc)
		}
	}
}

func TestIntSet_Copy(t *testing.T) {
	var tcs = []IntSet{
		IntSet{},
		IntSet{words: []uint{1, 2, 3, 4, 5, 6, 7, 8}},
		IntSet{words: []uint{1, 2, 1, 2, 1, 2, 1, 2}},
	}

	for _, tc := range tcs {
		actual := tc.Copy()
		for i := 0; i < len(tc.words); i++ {
			if actual.words[i] != tc.words[i] {
				t.Errorf("Copy: %v, Actual: %v", tc, actual)
			}
		}
	}
}

func TestIntSet_AddAll(t *testing.T) {
	var tcs = []struct {
		ints    []int
		values  []int
		expects string
	}{
		{[]int{}, []int{1, 2, 3}, "{1 2 3}"},
		{[]int{1, 2, 3}, []int{}, "{1 2 3}"},
		{[]int{1, 2, 3}, []int{2}, "{1 2 3}"},
		{[]int{1, 2, 3, 4, 6, 7, 8}, []int{1024}, "{1 2 3 4 6 7 8 1024}"},
	}

	for _, tc := range tcs {
		s := &IntSet{}
		for _, i := range tc.ints {
			s.Add(i)
		}

		s.AddAll(tc.values...)

		if s.String() != tc.expects {
			t.Errorf("IntSet: %v AddAll %v, Expects: %s, Actual: %s", tc.ints, tc.values, tc.expects, s.String())
		}
	}
}

func TestIntSet_IntersectWith(t *testing.T) {
	var tcs = []struct {
		s       []int
		t       []int
		expects string
	}{
		{[]int{1}, []int{2, 3, 4}, "{}"},
		{[]int{1, 2, 3}, []int{2, 3, 4}, "{2 3}"},
		{[]int{1, 2, 3}, []int{2}, "{2}"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 2}, "{2 8}"},
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

		s.IntersectWith(o)

		if s.String() != tc.expects {
			t.Errorf("%v IntersectWith %v, Expects: %s, Actual: %s", tc.s, tc.t, tc.expects, s.String())
		}
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	var tcs = []struct {
		s       []int
		t       []int
		expects string
	}{
		{[]int{1}, []int{2, 3, 4}, "{1}"},
		{[]int{1, 2, 3}, []int{2, 3, 4}, "{1}"},
		{[]int{1, 2, 3}, []int{1, 2, 3}, "{}"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 2}, "{1 3 4 5 6 7}"},
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

		s.DifferenceWith(o)

		if s.String() != tc.expects {
			t.Errorf("%v DifferenceWith %v, Expects: %s, Actual: %s", tc.s, tc.t, tc.expects, s.String())
		}
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	var tcs = []struct {
		s       []int
		t       []int
		expects string
	}{
		{[]int{1}, []int{2, 3, 4}, "{1 2 3 4}"},
		{[]int{1, 2, 3}, []int{1, 2, 3}, "{}"},
		{[]int{1, 2, 3}, []int{2}, "{1 3}"},
		{[]int{1, 2, 3}, []int{2, 3, 4}, "{1 4}"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 2, 9, 10}, "{1 3 4 5 6 7 9 10}"},
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

		s.SymmetricDifference(o)

		if s.String() != tc.expects {
			t.Errorf("%v SymmetricDifference with %v, Expects: %s, Actual: %s", tc.s, tc.t, tc.expects, s.String())
		}
	}
}

func TestIntSet_Elems(t *testing.T) {
	var tcs = [][]int{
		[]int{},
		[]int{1},
		[]int{1, 2, 3},
	}

	for _, tc := range tcs {
		s := new(IntSet)
		for _, i := range tc {
			s.Add(i)
		}

		for i, e := range s.Elems() {
			if e != tc[i] {
				t.Errorf("Expects: %v, Actual: %v", tc, e)
			}
		}
	}
}
