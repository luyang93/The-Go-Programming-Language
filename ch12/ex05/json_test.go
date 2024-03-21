package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
//	$ go test -v gopl.io/ch12/sexpr
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
	fmt.Println(string(data))

	// Decode it
	var movie Movie
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}

func TestBool(t *testing.T) {
	var tcs = []struct {
		b    bool
		want string
	}{
		{true, "true"},
		{false, "false"},
	}
	for _, tc := range tcs {
		actual, err := Marshal(tc.b)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}
		if string(actual) != tc.want {
			t.Errorf("Result = %s, Expected %v", actual, tc.want)
		}
	}
}

func TestFloat32(t *testing.T) {
	tcs := []struct {
		v    float32
		want string
	}{
		{3.2e9, "3.2e+09"},
		{1.0, "1"},
		{0, "0"},
	}

	for _, tc := range tcs {
		data, err := Marshal(tc.v)
		if err != nil {
			t.Errorf("Marshal(%v): %s", tc.v, err)
		}
		if string(data) != tc.want {
			t.Errorf("Marshal(%v) got %s, wanted %s", tc.v, data, tc.want)
		}
	}
}

func TestFloat64(t *testing.T) {
	tcs := []struct {
		v    float64
		want string
	}{
		{3.2e9, "3.2e+09"},
		{1.0, "1"},
		{0, "0"},
	}
	for _, tc := range tcs {
		data, err := Marshal(tc.v)
		if err != nil {
			t.Errorf("Marshal(%v): %s", tc.v, err)
		}
		if string(data) != tc.want {
			t.Errorf("Marshal(%v) got %s, wanted %s", tc.v, data, tc.want)
		}
	}
}
