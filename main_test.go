package main

import (
	"testing"
)

func CmpStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGetDuplicateChars(t *testing.T) {
	cases := []struct {
		input string
		want  []string
	}{
		{input: "thisisatest", want: []string{"t", "i", "s"}},
		{input: "abcabcabcd", want: []string{"a", "b", "c"}},
		{input: "a bb ccc dddd", want: []string{"b", "c", "d"}},
	}

	for _, tc := range cases {
		got := GetDuplicateChars(tc.input)
		if CmpStringSlice(got, tc.want) == false {
			t.Errorf("GetDuplicateChars(%v) = %v, want %v", tc.input, got, tc.want)
		}
	}

}
