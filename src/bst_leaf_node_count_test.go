package main

import "testing"

func TestBinarySearchTree_CountLeafNodes(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Extend([]int{15, 9, 23, 3, 12, 17, 28, 1, 8, 4})
	/* Tree looks like:
	          15
	        /    \
	       9      23
	      / \     / \
	     3  12   17 28
	    / \
	   1   8
	      /
	     4
	*/

	got := tree.CountLeafNodes()
	want := 5
	if got != want {
		t.Errorf("CountLeafNodes() = %v want %v", got, want)
	}
}
