// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(stdout, "resp.Status %s\n", resp.Status)

		bytes, err := io.Copy(stdout, resp.Body)
		resp.Body.Close()
		if err != nil || bytes == 0 {
			fmt.Fprintf(os.Stderr, "fetch: coping %v. copied %d byte(s)\n", err, bytes)
			os.Exit(1)
		}
	}
}
