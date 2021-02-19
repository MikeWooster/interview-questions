package main

import "testing"

func TestGetFirstNonRepeatedChar(t *testing.T) {
	cases := []struct {
		a    string
		want string
	}{
		{"abc", "a"},
		{"aaa", ""},
		{"bac", "b"},
		{"zzy", "y"},
		{"asdfulkj", "a"},
		{"with white space", "s"},
		{"sspeacwhit", "p"},
	}

	for _, tc := range cases {
		got := GetFirstNonRepeatedChar(tc.a)
		if got != tc.want {
			t.Errorf("GetFirstNonRepeatedChar(%v) = %v, want %v", tc.a, got, tc.want)
		}
	}
}
