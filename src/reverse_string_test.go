package main

import "testing"

func TestReverseStringWithRecursion(t *testing.T) {
	cases := []struct {
		a    string
		want string
	}{
		{"", ""},
		{"abc", "cba"},
		{"aaa", "aaa"},
		{"bac", "cab"},
		{"zzy", "yzz"},
		{"asdfulkj", "jklufdsa"},
		{"with white space", "ecaps etihw htiw"},
		{"sspeacwhit", "tihwcaepss"},
	}

	for _, tc := range cases {
		got := ReverseStringWithRecursion(tc.a)
		if got != tc.want {
			t.Errorf("ReverseStringWithRecursion(%v) = %v, want %v", tc.a, got, tc.want)
		}
	}
}

func TestReverseString(t *testing.T) {
	cases := []struct {
		a    string
		want string
	}{
		{"", ""},
		{"abc", "cba"},
		{"aaa", "aaa"},
		{"bac", "cab"},
		{"zzy", "yzz"},
		{"asdfulkj", "jklufdsa"},
		{"with white space", "ecaps etihw htiw"},
		{"sspeacwhit", "tihwcaepss"},
	}

	for _, tc := range cases {
		got := ReverseString(tc.a)
		if got != tc.want {
			t.Errorf("ReverseString(%v) = %v, want %v", tc.a, got, tc.want)
		}
	}
}
