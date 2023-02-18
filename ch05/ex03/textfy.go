package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr
var stdin io.Reader = os.Stdin

func main() {
	doc, err := html.Parse(stdin)
	if err != nil {
		fmt.Fprintf(stderr, "textfy: %v\n", err)
	}
	for _, text := range textfy(nil, doc) {
		fmt.Fprintln(stdout, text)
	}
}

func textfy(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}

	if n.Type == html.TextNode {
		if n.Parent.Data != "script" && n.Parent.Data != "style" {
			for _, line := range strings.Split(n.Data, "\n") {
				if len(line) != 0 {
					texts = append(texts, line)
				}
			}
		}
	}
	texts = textfy(texts, n.FirstChild)
	return textfy(texts, n.NextSibling)
}
