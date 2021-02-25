package main

import "testing"

func TestBinarySearchTree_Search(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Extend([]int{9, 3, 4, 1, 2, 12})

	cases := []struct {
		find     int
		expected bool
	}{
		{find: 9, expected: true},
		{find: 10, expected: false},
		{find: 2, expected: true},
		{find: 5, expected: false},
	}

	for _, c := range cases {
		got := tree.Search(c.find)
		if got != c.expected {
			t.Errorf("Search(%v) = %v, want %v", c.find, got, c.expected)
		}
	}
}
