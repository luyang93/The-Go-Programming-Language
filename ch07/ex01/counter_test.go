package main

import "testing"

func TestByteCounter_Write(t *testing.T) {
	var tcs = []struct {
		inputs  []byte
		expects int
	}{
		{[]byte(""), 0},
		{[]byte("one"), 3},
		{[]byte("one\n  two"), 9},
		{[]byte("one \ntwo\n three \nfour"), 21},
		{[]byte("one \ntwo\n\nfour\n"), 15},
	}

	for _, tc := range tcs {
		var c ByteCounter
		c.Write(tc.inputs)
		if int(c) != tc.expects {
			t.Errorf("Words of %v, Expects %d, Actual: %d", tc.inputs, tc.expects, c)
		}
	}
}

func TestWordCounter_Write(t *testing.T) {
	var tcs = []struct {
		inputs  []byte
		expects int
	}{
		{[]byte(""), 0},
		{[]byte("one"), 1},
		{[]byte("one two\nthree four"), 4},
		{[]byte("one two \nthree four"), 4},
	}

	for _, tc := range tcs {
		var c WordCounter
		c.Write(tc.inputs)
		if int(c) != tc.expects {
			t.Errorf("Words of %v, Expects %d, Actual: %d", tc.inputs, tc.expects, c)
		}
	}
}

func TestLineCounter_Write(t *testing.T) {
	var tcs = []struct {
		inputs  []byte
		expects int
	}{
		{[]byte(""), 0},
		{[]byte("one"), 1},
		{[]byte("one\n  two"), 2},
		{[]byte("one \ntwo\n three \nfour"), 4},
		{[]byte("one \ntwo\n\nfour\n"), 4},
	}

	for _, tc := range tcs {
		var c LineCounter
		c.Write(tc.inputs)
		if int(c) != tc.expects {
			t.Errorf("Lines of %v, Expects %d, Actual: %d", tc.inputs, tc.expects, c)
		}
	}
}
