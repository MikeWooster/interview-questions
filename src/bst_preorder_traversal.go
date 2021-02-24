package main

// PreorderTraversalRecursive uses a recursive algorithm to
// traverse the tree in `preorder`, i.e. root, left, right.
func (t *BinarySearchTree) PreorderTraversalRecursive() []int {
	// Setup the result array
	var r []int

	n := t.Root
	if n == nil {
		// Empty tree.
		return r
	}

	return t.preorderTraversalRecursive(n)
}

func (t *BinarySearchTree) preorderTraversalRecursive(n *BinarySearchNode) []int {
	var r []int
	r = append(r, n.Value)

	if n.Left != nil {
		for _, v := range t.preorderTraversalRecursive(n.Left) {
			r = append(r, v)
		}
	}
	if n.Right != nil {
		for _, v := range t.preorderTraversalRecursive(n.Right) {
			r = append(r, v)
		}

	}
	return r
}

// PreorderTraversal traverse the tree in `preorder`, i.e. root, left, right.
func (t *BinarySearchTree) PreorderTraversal() []int {
	// Setup the result array
	var r []int

	// Create the stack starting with the root node
	s := nodeStack{}
	if t.Root != nil {
		s.Append(t.Root)
	}

	for s.Len() > 0 {
		n := s.Pop()

		// Add root to the result.
		r = append(r, n.Value)

		// As we are working with LIFO, to visit the left before the right,
		// the left must added visited last.

		if n.Right != nil {
			s.Append(n.Right)
		}

		if n.Left != nil {
			s.Append(n.Left)
		}

	}

	return r
}
