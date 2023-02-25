package intset

import (
	"testing"
)

func TestIntSet_Add(t *testing.T) {
	var tcs = []struct {
		words   []int
		expects []uint64
	}{
		{[]int{}, []uint64{}},
		{[]int{1, 2, 101}, []uint64{6, 137438953472}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 1024}, []uint64{510, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for _, tc := range tcs {
		var s IntSet
		for _, w := range tc.words {
			s.Add(w)
		}

		for i := 0; i < len(tc.expects); i++ {
			if s.words[i] != tc.expects[i] {
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
		IntSet{words: []uint64{1, 2, 3}},
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
		IntSet{words: []uint64{1, 2, 3, 4, 5, 6, 7, 8}},
		IntSet{words: []uint64{1, 2, 1, 2, 1, 2, 1, 2}},
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
