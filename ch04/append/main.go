// Append illustrates the behavior of the built-in append function.
package main

import "fmt"

func main() {
	var x, y []int
	var x1, y1 []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("appendInt   %d  cap=%d\t%v\n", i, cap(y), y)
		x = y
		y1 = appendslice(x1, x...)
		x1 = y1
		fmt.Printf("appendSlice %d  cap=%d\t%v\n", i, cap(y1), y1)
	}
	/*
		appendInt   0  cap=1    [0]
		appendSlice 0  cap=1    [0]
		appendInt   1  cap=2    [0 1]
		appendSlice 1  cap=3    [0 0 1]
		appendInt   2  cap=4    [0 1 2]
		appendSlice 2  cap=6    [0 0 1 0 1 2]
		appendInt   3  cap=4    [0 1 2 3]
		appendSlice 3  cap=12   [0 0 1 0 1 2 0 1 2 3]
		appendInt   4  cap=8    [0 1 2 3 4]
		appendSlice 4  cap=20   [0 0 1 0 1 2 0 1 2 3 0 1 2 3 4]
		appendInt   5  cap=8    [0 1 2 3 4 5]
		appendSlice 5  cap=30   [0 0 1 0 1 2 0 1 2 3 0 1 2 3 4 0 1 2 3 4 5]
		appendInt   6  cap=8    [0 1 2 3 4 5 6]
		appendSlice 6  cap=30   [0 0 1 0 1 2 0 1 2 3 0 1 2 3 4 0 1 2 3 4 5 0 1 2 3 4 5 6]
		appendInt   7  cap=8    [0 1 2 3 4 5 6 7]
		appendSlice 7  cap=56   [0 0 1 0 1 2 0 1 2 3 0 1 2 3 4 0 1 2 3 4 5 0 1 2 3 4 5 6 0 1 2 3 4 5 6 7]
		appendInt   8  cap=16   [0 1 2 3 4 5 6 7 8]
		appendSlice 8  cap=56   [0 0 1 0 1 2 0 1 2 3 0 1 2 3 4 0 1 2 3 4 5 0 1 2 3 4 5 6 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 8]
		appendInt   9  cap=16   [0 1 2 3 4 5 6 7 8 9]
		appendSlice 9  cap=56   [0 0 1 0 1 2 0 1 2 3 0 1 2 3 4 0 1 2 3 4 5 0 1 2 3 4 5 6 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 8 0 1 2 3 4 5 6 7 8 9]
	*/
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

func appendslice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to expand the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
