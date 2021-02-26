package main

import "testing"

func TestStringIsNumber(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{s: "asdflkwerclksadf", want: false},
		{s: "asdf234asdf", want: false},
		{s: "092384", want: true},
	}

	for _, c := range cases {
		got := StringIsNumber(c.s)
		if got != c.want {
			t.Errorf("StringIsNumber(%v) = %v want %v", c.s, got, c.want)
		}
	}
}
