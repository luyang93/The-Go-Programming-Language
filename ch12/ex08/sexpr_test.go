package sexpr

import (
	"bytes"
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
//	$ go test -v ch12/sexpr
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

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}

func TestFloat(t *testing.T) {
	var tcs = []struct {
		num32 float32
		num64 float64
	}{
		{12.3, 3.21},
		{0.0, 10000},
	}
	for _, tc := range tcs {
		actual32, err := Marshal(tc.num32)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}
		actual64, err := Marshal(tc.num64)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}

		if string(actual32) != fmt.Sprintf("%4.4f", tc.num32) {
			t.Errorf("Result = %v, Expected %v", actual32, tc.num32)
		}
		if string(actual64) != fmt.Sprintf("%4.4f", tc.num64) {
			t.Errorf("Result = %v, Expected %v", actual64, tc.num64)
		}
	}
}

func TestComplex(t *testing.T) {
	var tcs = []struct {
		num64  complex64
		num128 complex128
	}{
		{12.0 - 3i, 3.2 + 1i},
		{0.0 - 0i, 10000 + 0i},
	}
	for _, tc := range tcs {
		actual64, err := Marshal(tc.num64)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}
		actual128, err := Marshal(tc.num128)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}

		if string(actual64) != fmt.Sprintf("#C(%4.4f %4.4f)", real(tc.num64), imag(tc.num64)) {
			t.Errorf("Result = %s, Expected %v", actual64, tc.num64)
		}
		if string(actual128) != fmt.Sprintf("#C(%4.4f %4.4f)", real(tc.num128), imag(tc.num128)) {
			t.Errorf("Result = %s, Expected %v", actual128, tc.num128)
		}
	}
}

func TestBool(t *testing.T) {
	var tcs = []struct {
		b    bool
		want string
	}{
		{true, "t"},
		{false, "nil"},
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

func TestStreamingDecode(t *testing.T) {
	type Book struct {
		Title, Author string
	}
	book := Book{"Point Counterpoint", "Aldous Huxley"}
	data, err := Marshal(book)
	if err != nil {
		t.Errorf("setting up test: %s", err)
		return
	}
	data = bytes.Repeat(data, 2)
	t.Logf("%s", data)

	dec := NewDecoder(bytes.NewReader(data))
	books := make([]Book, 0)
	for dec.More() {
		var b Book
		err := dec.Decode(&b)
		if err != nil {
			t.Errorf("Error after decoding %s: %s", books, err)
			return
		}
		books = append(books, b)
	}
	want := []Book{book, book}
	if !reflect.DeepEqual(want, books) {
		t.Errorf("Got %s, want %s", books, want)
	}
}
