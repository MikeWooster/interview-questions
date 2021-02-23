package main

import (
	"testing"
)

func TestBinarySearchTree_Add(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Add(9)
	tree.Add(4)
	tree.Add(1)
	tree.Add(3)
	tree.Add(12)

	if tree.Root.Value != 9 {
		t.Errorf("Add(%v) = %v", 9, tree.Root)
	}
	if tree.Root.Left.Value != 4 {
		t.Errorf("Add(%v) = %v", 4, tree.Root.Left)
	}
	if tree.Root.Left.Left.Value != 1 {
		t.Errorf("Add(%v) = %v", 1, tree.Root.Left.Left)
	}
	if tree.Root.Left.Left.Right.Value != 3 {
		t.Errorf("Add(%v) = %v", 3, tree.Root.Left.Left.Right)
	}
	if tree.Root.Right.Value != 12 {
		t.Errorf("Add(%v) = %v", 3, tree.Root.Right)
	}
}

func TestBinarySearchTree_Extend(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Extend([]int{9, 4, 1, 3, 12})

	if tree.Root.Value != 9 {
		t.Errorf("Extend(%v) = %v", 9, tree.Root)
	}
	if tree.Root.Left.Value != 4 {
		t.Errorf("Extend(%v) = %v", 4, tree.Root.Left)
	}
	if tree.Root.Left.Left.Value != 1 {
		t.Errorf("Extend(%v) = %v", 1, tree.Root.Left.Left)
	}
	if tree.Root.Left.Left.Right.Value != 3 {
		t.Errorf("Extend(%v) = %v", 3, tree.Root.Left.Left.Right)
	}
	if tree.Root.Right.Value != 12 {
		t.Errorf("Extend(%v) = %v", 3, tree.Root.Right)
	}
}

func TestBinarySearchTree_ListNodes(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Extend([]int{9, 3, 4, 1, 2, 12})

	got := tree.ListNodes()
	want := []int{9, 3, 1, 2, 4, 12}

	if cmpIntSlice(got, want) == false {
		t.Errorf("ListNodes() = %v want %v", got, want)
	}
}

func TestBinarySearchTree_Min(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Extend([]int{9, 3, 4, 1, 2, 12})

	got, _ := tree.Min()
	want := 1

	if got != want {
		t.Errorf("Min() = %v want %v", got, want)
	}
}

func TestBinarySearchTree_Max(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Extend([]int{9, 3, 4, 1, 2, 12})

	got, _ := tree.Max()
	want := 12

	if got != want {
		t.Errorf("Max() = %v want %v", got, want)
	}
}

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

func TestBinarySearchTree_Delete(t *testing.T) {
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

	tree.Delete(28)
	/* Tree now looks like:
	          15
	        /    \
	       9      23
	      / \     /
	     3  12   17
	    / \
	   1   8
	      /
	     4
	*/

	if tree.Root.Right.Right != nil {
		t.Errorf("Delete(%v) did not delete expected element from tree: %v", 28, tree.Print())
	}

	tree.Delete(8)
	/* Tree now looks like:
	          15
	        /    \
	       9      23
	      / \     /
	     3  12   17
	    / \
	   1   4
	*/

	if tree.Root.Left.Left.Right.Value != 4 {
		t.Errorf("Delete(%v) did not delete expected element from tree: %v", 8, tree.Print())
	}

	tree.Delete(9)
	/* Tree now looks like:
	          15
	        /    \
	       4      23
	      / \     /
	     3  12   17
	    /
	   1
	*/

	if tree.Root.Left.Value != 4 || tree.Root.Left.Left.Value != 3 || tree.Root.Left.Right.Value != 12 || tree.Root.Left.Left.Left.Value != 1 {
		t.Errorf("Delete(%v) did not delete expected element from tree: %v", 9, tree.Print())
	}

	tree.Delete(4)
	/* Tree now looks like:
	    15
	  /    \
	 3       23
	/ \     /
	1  12  17
	*/

	if tree.Root.Left.Value != 3 || tree.Root.Left.Left.Value != 1 || tree.Root.Left.Right.Value != 12 {
		t.Errorf("Delete(%v) did not delete expected element from tree: %v", 4, tree.Print())
	}

	tree.Delete(15)
	/* Tree now looks like:
	    12
	   /  \
	  3    23
	 /     /
	1     17
	*/

	if tree.Root.Value != 12 || tree.Root.Left.Right != nil {
		t.Errorf("Delete(%v) did not delete expected element from tree: %v", 15, tree.Print())
	}
}
