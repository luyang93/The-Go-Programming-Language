package main

import (
	"fmt"
	"io"
	"log"
	"os"

	uncp "gitlab.com/luyang93/The-Go-Programming-Language/ch10/ex02/uncompressor"
	_ "gitlab.com/luyang93/The-Go-Programming-Language/ch10/ex02/uncompressor/tar"
	_ "gitlab.com/luyang93/The-Go-Programming-Language/ch10/ex02/uncompressor/zip"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: extract FILE ...")
	}
	exitCode := 0
	for _, filename := range os.Args[1:] {
		err := printArchive(filename)
		if err != nil {
			log.Print(err)
			exitCode = 2
		}
	}
	os.Exit(exitCode)
}

func printArchive(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := uncp.Open(f)
	if err != nil {
		return fmt.Errorf("open archive reader: %s", err)
	}
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		return fmt.Errorf("printing: %s", err)
	}
	return nil
}
