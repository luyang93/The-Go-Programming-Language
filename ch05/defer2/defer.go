// Defer2 demonstrates a deferred call to runtime.Stack during a panic.
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

/*
goroutine 1 [running]:
main.printStack()
        The-Go-Programming-Language/ch05/defer2/defer.go:17 +0x39
panic({0x48cba0, 0x524fd0})
        .goenv/versions/1.19.2/src/runtime/panic.go:884 +0x212
main.f(0x4b9298?)
        The-Go-Programming-Language/ch05/defer2/defer.go:22 +0x113
main.f(0x1)
        The-Go-Programming-Language/ch05/defer2/defer.go:24 +0xf5
main.f(0x2)
        The-Go-Programming-Language/ch05/defer2/defer.go:24 +0xf5
main.f(0x3)
        The-Go-Programming-Language/ch05/defer2/defer.go:24 +0xf5
main.main()
        The-Go-Programming-Language/ch05/defer2/defer.go:12 +0x45
panic: runtime error: integer divide by zero
*/
