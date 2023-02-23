// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item, domain string) []string, worklist []string) {
	seen := make(map[string]bool)
	for _, w := range worklist {
		url_, err := url.Parse(w)
		if err != nil {
			continue
		}
		domain := url_.Host

		subworklist := make([]string, 1)
		subworklist[0] = w

		for len(subworklist) > 0 {
			items := subworklist
			subworklist = nil
			for _, item := range items {
				if !seen[item] {
					seen[item] = true
					subworklist = append(subworklist, f(item, domain)...)
				}
			}
		}
	}
}

func crawl(url, domain string) []string {
	fmt.Println(url)
	err := savePage(url, domain)
	if err != nil {
		log.Printf("Can't save URL \"%s\": %s", url, err)
	}
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func savePage(rawurl, domain string) error {
	url_, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Errorf("bad url_: %s", err)
	}

	if domain != url_.Host {
		return nil
	}

	dir := url_.Host
	var filename string
	if filepath.Ext(url_.Path) == "" {
		dir = filepath.Join(dir, url_.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(url_.Path))
		filename = url_.Path
	}
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}
	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)

	// Check for delayed write errors, as mentioned at the end of section 5.8.
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}
	return err
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
