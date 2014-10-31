package string

import (
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"Backward", "drawkcaB"},
		{"", ""},
		{"Hello Украина", "аниаркУ olleH"},
	}

	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}

func TestIsPangram(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"We promptly judged antique ivory buckles for the next prize", "pangram"},
		{"We promptly judged antique ivory buckles for the prize", "not pangram"},
	}

	for _, c := range tests {
		got := IsPangram(c.s)
		if got != c.want {
			t.Errorf("IsPangram(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}
