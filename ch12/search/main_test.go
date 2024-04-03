package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	var tcs = []struct {
		Url        string
		Body       string
		StatusCode int
	}{
		{
			"http://localhost:12345/search",
			"{Labels:[] MaxResults:10 Exact:false}",
			200,
		},
		{
			"http://localhost:12345/search?l=golang&l=programming",
			"{Labels:[golang programming] MaxResults:10 Exact:false}",
			200,
		},
		{
			"http://localhost:12345/search?l=golang&l=programming&max=100",
			"{Labels:[golang programming] MaxResults:100 Exact:false}",
			200,
		},
		{
			"http://localhost:12345/search?x=true&l=golang&l=programming",
			"{Labels:[golang programming] MaxResults:10 Exact:true}",
			200,
		},
		{
			"http://localhost:12345/search?q=hello&x=123",
			"x: strconv.ParseBool: parsing \"123\": invalid syntax",
			400,
		},
		{
			"http://localhost:12345/search?q=hello&max=lots",
			"max: strconv.ParseInt: parsing \"lots\": invalid syntax",
			400,
		},
	}
	for _, tc := range tcs {
		resp, err := http.Get(tc.Url)
		if err != nil {
			t.Error(err)
		}

		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		if !strings.Contains(string(b), tc.Body) {
			t.Errorf("Body is not equal. Expect = %q, Get %q", tc.Body, string(b))
		}

		if resp.StatusCode != tc.StatusCode {
			t.Errorf("StatusCode is not equal. Expect = %q, Get %q", tc.StatusCode, resp.StatusCode)
		}
	}
}
