package params

import (
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestPack(t *testing.T) {
	s := struct {
		Name string `http:"n"`
		Age  int    `http:"a"`
	}{"Arugula", 35}
	u, err := Pack(&s)
	if err != nil {
		t.Errorf("Pack(%#v): %s", s, err)
	}
	want := "a=35&n=Arugula"
	got := u.RawQuery
	if got != want {
		t.Errorf("Pack(%#v): got %q, want %q", s, got, want)
	}
}

func TestUnpack(t *testing.T) {
	type unpacked struct {
		Post int `http:"p" check:"lt100"`
	}
	tcs := []struct {
		req       *http.Request
		want      unpacked
		errSubstr string
	}{
		{
			&http.Request{Form: url.Values{
				"p": []string{"80"},
			}},
			unpacked{
				80,
			},
			"",
		},
		{
			&http.Request{Form: url.Values{
				"p": []string{"120"},
			}},
			unpacked{
				0,
			},
			">= 100",
		},
	}
	for _, tc := range tcs {
		var got unpacked
		err := Unpack(tc.req, &got)
		if err != nil {
			if !strings.Contains(err.Error(), tc.errSubstr) {
				t.Errorf("Unpack(%v), error %q doesn't contain %q", tc.req, err, tc.errSubstr)
			}
		} else if tc.errSubstr != "" {
			t.Errorf("Unpack(%v): %s", tc.req, err)
		}
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("Unpack(%v), got %v, want %v", tc.req, got, tc.want)
		}
	}
}
