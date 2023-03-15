package memo_test

import (
	"testing"

	memo "gitlab.com/luyang93/The-Go-Programming-Language/ch09/memo3"
	"gitlab.com/luyang93/The-Go-Programming-Language/ch09/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func TestSequential(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
