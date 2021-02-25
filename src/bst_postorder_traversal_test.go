package main

import "testing"

func TestBinarySearchTree_PostorderTraversalRecursive(t *testing.T) {
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

	got := tree.PostorderTraversalRecursive()
	want := []int{1, 4, 8, 3, 12, 9, 17, 28, 23, 15}
	if cmpIntSlice(got, want) == false {
		t.Errorf("PostorderTraversalRecursive() = %v want %v", got, want)
	}
}

func TestBinarySearchTree_PostorderTraversal(t *testing.T) {
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

	got := tree.PostorderTraversal()
	want := []int{1, 4, 8, 3, 12, 9, 17, 28, 23, 15}
	if cmpIntSlice(got, want) == false {
		t.Errorf("PostorderTraversal() = %v want %v", got, want)
	}
}
