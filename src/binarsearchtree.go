package main

import (
	"fmt"
	"strings"
)

// BinarySearchNode represents the node in a binary tree.
type BinarySearchNode struct {
	Value  int
	Left   *BinarySearchNode
	Right  *BinarySearchNode
	Parent *BinarySearchNode
}

// HasChildren returns whether the node has any children.
func (n *BinarySearchNode) HasChildren() bool {
	return n.Left != nil || n.Right != nil
}

// Print retuns a string representation of a node.
// mainly for testing purposes.
func (n *BinarySearchNode) Print() string {
	var lv, rv, pv string
	if n.Left == nil {
		lv = "<nil>"
	} else {
		lv = fmt.Sprintf("%v", n.Left.Value)
	}
	if n.Right == nil {
		rv = "<nil>"
	} else {
		rv = fmt.Sprintf("%v", n.Right.Value)
	}
	if n.Parent == nil {
		pv = "<nil>"
	} else {
		pv = fmt.Sprintf("%v", n.Parent.Value)
	}
	return fmt.Sprintf("Node: %v, Parent: %v, LeftChild: %v, RightChild: %v", n.Value, pv, lv, rv)
}

// BinarySearchTree represents the tree itself.
type BinarySearchTree struct {
	Root *BinarySearchNode
}

// Add adds a new value to the tree
func (t *BinarySearchTree) Add(v int) {
	node := &BinarySearchNode{Value: v}

	if t.Root == nil {
		// Empty tree - the node goes here and we can stop.
		t.Root = node
		return
	}

	current := t.Root
	for {
		if v < current.Value && current.Left == nil {
			node.Parent = current
			current.Left = node
			return
		}
		if v > current.Value && current.Right == nil {
			node.Parent = current
			current.Right = node
			return
		}

		if v < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
}

// Extend adds a list of nodes into a tree one after the other
func (t *BinarySearchTree) Extend(values []int) {
	for _, v := range values {
		t.Add(v)
	}
}

// ListNodes returns the integer values in the tree in
// the trees order.
func (t *BinarySearchTree) ListNodes() []int {
	var nodeVals []int
	if t.Root == nil {
		// Empty tree - return here
		return nodeVals
	}

	var nodes []*BinarySearchNode

	nodes = append(nodes, t.Root)

	// Non-recursively get the values out of the tree.
	for len(nodes) > 0 {
		// Take the last item from the slice.
		n := pop(&nodes)
		nodeVals = append(nodeVals, n.Value)

		if n.Right != nil {
			nodes = append(nodes, n.Right)
		}

		if n.Left != nil {
			nodes = append(nodes, n.Left)
		}
	}

	return nodeVals
}

// Print prints out the tree
func (t *BinarySearchTree) Print() string {

	if t.Root == nil {
		// Empty tree - return here
		return ""
	}

	var nodes []*BinarySearchNode

	nodes = append(nodes, t.Root)

	var nodeVals []string

	// Non-recursively get the values out of the tree.
	for len(nodes) > 0 {
		// Take the last item from the slice.
		n := pop(&nodes)
		nodeVals = append(nodeVals, n.Print())

		if n.Right != nil {
			nodes = append(nodes, n.Right)
		}

		if n.Left != nil {
			nodes = append(nodes, n.Left)
		}
	}

	return strings.Join(nodeVals, "\n")
}

// Min returns the smallest value in the tree.
func (t *BinarySearchTree) Min() (int, error) {
	node := t.minNode(t.Root)
	if node == nil {
		return 0, fmt.Errorf("tree is empty")
	}
	return node.Value, nil
}

func (t *BinarySearchTree) minNode(node *BinarySearchNode) *BinarySearchNode {
	if node == nil {
		return node
	}

	for {
		if node.Left != nil {
			node = node.Left
		} else {
			// No left node means this is the largest node
			return node
		}
	}
}

// Max returns the greatest value in the tree.
func (t *BinarySearchTree) Max() (int, error) {
	node := t.maxNode(t.Root)
	if node == nil {
		return 0, fmt.Errorf("tree is empty")
	}
	return node.Value, nil
}

func (t *BinarySearchTree) maxNode(node *BinarySearchNode) *BinarySearchNode {
	if node == nil {
		return node
	}

	for {
		if node.Right != nil {
			node = node.Right
		} else {
			// No left node means this is the largest node
			return node
		}
	}
}

// Search finds a the value of element in the tree, and reports whether it
// is present or not.
func (t *BinarySearchTree) Search(v int) bool {
	finder := t.getNode(v)
	return finder != nil
}

// Delete removes a value from the tree if it is in the tree.
func (t *BinarySearchTree) Delete(v int) {
	node := t.getNode(v)
	if node == nil {
		// Node not in tree, return at this point.
		return
	}
	t.deleteNode(node)
}

func (t *BinarySearchTree) deleteNode(node *BinarySearchNode) {

	var promoted *BinarySearchNode

	if node.Left == nil && node.Right == nil {
		promoted = nil

	} else if node.Right == nil {
		// The node has one child on the left. promote this child.
		promoted = node.Left
		promoted.Parent = node.Parent

	} else if node.Left == nil {
		// The node has one child on the right. promote this child.
		promoted = node.Right
		promoted.Parent = node.Parent

	} else {
		// Find the largest node from this nodes left subtree
		// and move to the deleted nodes location
		promoted = t.maxNode(node.Left)

		// Delete this node, once this has been removed, we can
		// move the removed node to the desired position.
		t.deleteNode(promoted)

		// Set the children of the new child node to the children of the
		// node we are deleting.
		promoted.Right = node.Right
		promoted.Left = node.Left
		// Make sure the children of the newly promoted node
		// have the correct parents set.
		promoted.Right.Parent = promoted
		promoted.Left.Parent = promoted
		promoted.Parent = node.Parent
	}

	if node.Parent == nil {
		// This is the root node, move the child to this node.
		t.Root = promoted
	} else if node.Parent.Left == node {
		node.Parent.Left = promoted
	} else {
		node.Parent.Right = promoted
	}

}

// getNode finds the node in the tree and returns it, or nil
// if it can't be found.
func (t *BinarySearchTree) getNode(v int) *BinarySearchNode {
	node := t.Root

	for node != nil {
		if v == node.Value {
			return node
		}

		if v < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return nil
}

func pop(nodes *[]*BinarySearchNode) *BinarySearchNode {
	size := len(*nodes)

	if size == 0 {
		return nil
	}

	n := (*nodes)[size-1]
	*nodes = (*nodes)[:size-1]
	return n
}
