package main

// BinarySearchNode represents the node in a binary tree.
type BinarySearchNode struct {
	Value int
	Left  *BinarySearchNode
	Right *BinarySearchNode
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
			current.Left = node
			return
		}
		if v > current.Value && current.Right == nil {
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

func pop(nodes *[]*BinarySearchNode) *BinarySearchNode {
	size := len(*nodes)

	if size == 0 {
		return nil
	}

	n := (*nodes)[size-1]
	*nodes = (*nodes)[:size-1]
	return n
}
