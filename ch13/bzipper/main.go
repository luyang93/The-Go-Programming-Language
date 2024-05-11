// See page 365.

// Bzipper reads input, bzip2-compresses it, and writes it out.
package main

import (
	"gitlab.com/luyang93/The-Go-Programming-Language/ch13/bzip"
	"io"
	"log"
	"os"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}
