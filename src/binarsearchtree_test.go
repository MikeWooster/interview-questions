package main

import "testing"

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

func TestBinarySearchTree_ListNodes(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Extend([]int{9, 3, 4, 1, 2, 12})

	got := tree.ListNodes()
	want := []int{9, 3, 1, 2, 4, 12}

	if cmpIntSlice(got, want) == false {
		t.Errorf("Add() = %v want %v", got, want)
	}
}
