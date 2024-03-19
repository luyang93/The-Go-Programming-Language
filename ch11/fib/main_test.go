package fib

import "testing"

func TestFibRecursive(t *testing.T) {
	if FibRecursive(40) != 102334155 {
		t.Error("Incorrect!")
	}
}

func TestFib(t *testing.T) {
	if Fib(40) != 102334155 {
		t.Error("Incorrect!")
	}
}

func BenchmarkFibRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRecursive(40)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(40)
	}
}

func TestBlock(t *testing.T) {
	// Call the block function
	Block()
}

func TestMutx(t *testing.T) {
	result := Mutx(4)
	t.Logf("%d", result)
}
