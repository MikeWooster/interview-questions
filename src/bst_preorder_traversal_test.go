package main

import "testing"

func TestBinarySearchTree_PreorderTraversalRecursive(t *testing.T) {
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

	got := tree.PreorderTraversalRecursive()
	want := []int{15, 9, 3, 1, 8, 4, 12, 23, 17, 28}
	if cmpIntSlice(got, want) == false {
		t.Errorf("PreorderTraversalRecursive() = %v want %v", got, want)
	}
}

func TestBinarySearchTree_PreorderTraversal(t *testing.T) {
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

	got := tree.PreorderTraversal()
	want := []int{15, 9, 3, 1, 8, 4, 12, 23, 17, 28}
	if cmpIntSlice(got, want) == false {
		t.Errorf("PreorderTraversal() = %v want %v", got, want)
	}
}
