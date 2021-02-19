package main

import "testing"

func TestCheckStringAnagrams(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"abc", "cab", true},
		{"aaa", "bbb", false},
		{"abc", "bac", true},
		{"zzz", "zzy", false},
		{"asdfulkj", "fajksudl", true},
		{"with spaces", "sspeacwhit", true},
		{"sspeacwhit", "with spaces", true},
	}

	for _, tc := range cases {
		got := CheckStringAnagrams(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("CheckStringAnagrams(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}
