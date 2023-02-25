// Defer1 demonstrates a deferred call being invoked during a panic.
package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

/*
//!+stdout
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
//!-stdout

//!+stderr
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.f(0x4b91b8?)
        The-Go-Programming-Language/ch05/defer1/defer.go:11 +0x113
main.f(0x1)
        The-Go-Programming-Language/ch05/defer1/defer.go:13 +0xf5
main.f(0x2)
        The-Go-Programming-Language/ch05/defer1/defer.go:13 +0xf5
main.f(0x3)
        The-Go-Programming-Language/ch05/defer1/defer.go:13 +0xf5
main.main()
        The-Go-Programming-Language/ch05/defer1/defer.go:7 +0x1e
//!-stderr
*/
