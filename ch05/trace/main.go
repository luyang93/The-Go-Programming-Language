// The trace program uses defer to add entry/exit diagnostics to a function.
package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}

/*
$ go build gitlab.com/luyang93/The-Go-Programming-Language/ch5/trace
$ ./trace
2023/02/25 01:17:26 enter bigSlowOperation
2023/02/25 01:17:36 exit bigSlowOperation (10.009250011s)
*/
