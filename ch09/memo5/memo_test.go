package memo_test

import (
	"testing"

	memo "gitlab.com/luyang93/The-Go-Programming-Language/ch09/memo5"
	"gitlab.com/luyang93/The-Go-Programming-Language/ch09/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}
