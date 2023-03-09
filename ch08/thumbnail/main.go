//go:build ignore
// +build ignore

// The thumbnail command produces thumbnails of JPEG files
// whose names are provided on each line of the standard input.
//
// The "+build ignore" tag (see p.295) excludes this file from the
// thumbnail package, but it can be compiled as a command and run like
// this:
//
// Run with:
//
//	$ go run $GOPATH/src/gitlab.com/luyang93/The-Go-Programming-Language/ch08/thumbnail/main.go
//	foo.jpeg
//	^D
package main
