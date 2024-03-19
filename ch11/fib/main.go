package fib

import (
	"sync"
	"time"
)

// Import fmt to use Println to show output
func FibRecursive(n int) uint64 {
	// Function Creation to accept integer till which the Fibonacci series runs
	if n == 0 {
		return 0
		// Base case for the recursive call
	} else if n == 1 {
		return 1
		// Base case for the recursive call
	} else {
		return FibRecursive(n-1) + FibRecursive(n-2)
		// Recursive call for finding the Fibonacci number
	}
}

var f [100010]uint64

func Fib(n int) uint64 {
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

var ch = make(chan int)

func Block() {
	// WaitGroup is used to make the function wait for G1 and G2 to finish
	wg := sync.WaitGroup{}
	wg.Add(2)
	go G1(&wg)
	go G2(&wg)
	wg.Wait()
}
func G1(wg *sync.WaitGroup) {
	// Write to the channel
	ch <- 100
	wg.Done()
}
func G2(wg *sync.WaitGroup) {
	// Sleep for 1 second
	time.Sleep(time.Second)
	// Read from the channel
	<-ch
	wg.Done()
}

func Mutx(s int) int {
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(1000 * 1000)
	// Launching 1,000,000 goroutines
	for i := 0; i < 1000*1000; i++ {
		go func(i int) {
			// Locking the mutex
			mu.Lock()
			// Unlocking the mutex
			defer mu.Unlock()
			defer wg.Done()
			s++
		}(i)
	}
	wg.Wait()
	return s
}
