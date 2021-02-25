package main

import "testing"

func TestBinarySearchTree_InorderTraversalRecursive(t *testing.T) {
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

	got := tree.InorderTraversalRecursive()
	want := []int{1, 3, 4, 8, 9, 12, 15, 17, 23, 28}
	if cmpIntSlice(got, want) == false {
		t.Errorf("InorderTraversalRecursive() = %v want %v", got, want)
	}
}

func TestBinarySearchTree_InorderTraversal(t *testing.T) {
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

	got := tree.InorderTraversal()
	want := []int{1, 3, 4, 8, 9, 12, 15, 17, 23, 28}
	if cmpIntSlice(got, want) == false {
		t.Errorf("InorderTraversal() = %v want %v", got, want)
	}
}
