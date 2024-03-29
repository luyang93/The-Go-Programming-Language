package sexpr

import (
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

		if string(actual32) != fmt.Sprintf("%g", tc.num32) {
			t.Errorf("Result = %v, Expected %v", actual32, tc.num32)
		}
		if string(actual64) != fmt.Sprintf("%g", tc.num64) {
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

		if string(actual64) != fmt.Sprintf("#C(%g %g)", real(tc.num64), imag(tc.num64)) {
			t.Errorf("Result = %s, Expected %v", actual64, tc.num64)
		}
		if string(actual128) != fmt.Sprintf("#C(%g %g)", real(tc.num128), imag(tc.num128)) {
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
func TestMarshal(t *testing.T) {
	type Interface interface{}
	type Record struct {
		B    bool
		F32  float32
		F64  float64
		C64  complex64
		C128 complex128
		I    Interface
	}
	tcs := []struct {
		r    Record
		want string
	}{
		{
			Record{true, 2.5, 0, 1 + 2i, 2 + 3i, Interface(5)},
			`((B t) (F32 2.5) (F64 0) (C64 #C(1 2)) (C128 #C(2 3)) (I ("sexpr.Interface" 5)))`,
		},
		{
			Record{false, 0, 1.5, 0, 1i, Interface(0)},
			`((B nil) (F32 0) (F64 1.5) (C64 #C(0 0)) (C128 #C(0 1)) (I ("sexpr.Interface" 0)))`,
		},
	}
	for _, tc := range tcs {
		data, err := Marshal(tc.r)
		s := string(data)
		if err != nil {
			t.Errorf("Marshal(%s): %s", s, err)
		}
		if s != tc.want {
			t.Errorf("Marshal(%#v) got %s, wanted %s", tc.r, s, tc.want)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	type Interface interface{}
	type Record struct {
		B    bool
		F32  float32
		F64  float64
		C64  complex64
		C128 complex128
		I    Interface
	}
	Interfaces["sexpr.Interface"] = reflect.TypeOf(int(0))
	tcs := []struct {
		s    string
		want Record
	}{
		{
			`((B t) (F32 2.5) (F64 0) (I ("sexpr.Interface" 5)))`,
			Record{true, 2.5, 0, 0, 0, Interface(5)},
		},
		{
			`((B nil) (F32 0) (F64 1.5) (I ("sexpr.Interface" 0)))`,
			Record{false, 0, 1.5, 0, 0, Interface(0)},
		},
	}
	for _, tc := range tcs {
		var r Record
		err := Unmarshal([]byte(tc.s), &r)
		if err != nil {
			t.Errorf("Unmarshal(%q): %s", tc.s, err)
		}
		if !reflect.DeepEqual(r, tc.want) {
			t.Errorf("Unmarshal(%q) got %#v, wanted %#v", tc.s, r, tc.want)
		}
	}
}
