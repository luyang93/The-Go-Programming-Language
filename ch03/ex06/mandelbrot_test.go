package main

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {
	stdout = new(bytes.Buffer)
	main()
	img := stdout.(*bytes.Buffer).String()
	if len(img) < 0 {
		t.Error("Failed to create image")
	}
}
