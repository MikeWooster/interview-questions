package main

// PostorderTraversalRecursive traverses the binary tree recursively in post order
// i.e. left, right, root
func (t *BinarySearchTree) PostorderTraversalRecursive() []int {
	return t.postorderTraversalRecursive(t.Root)
}

func (t *BinarySearchTree) postorderTraversalRecursive(n *BinarySearchNode) []int {
	var r []int

	if n == nil {
		return r
	}

	if n.Left != nil {
		for _, v := range t.postorderTraversalRecursive(n.Left) {
			r = append(r, v)
		}
	}

	if n.Right != nil {
		for _, v := range t.postorderTraversalRecursive(n.Right) {
			r = append(r, v)
		}
	}

	r = append(r, n.Value)

	return r
}

// PostorderTraversal traverses the binary tree in post order
// i.e. left, right, root
func (t *BinarySearchTree) PostorderTraversal() []int {
	var r []int

	if t.Root == nil {
		return r
	}

	// Create the stack starting with the root node
	s := nodeStack{}
	s.Append(t.Root)

	tracker := make(map[*BinarySearchNode]struct{})
	added := func(n *BinarySearchNode) bool {
		_, ok := tracker[n]
		return ok
	}
	tmp := nodeStack{}

	for s.Len() > 0 {
		n := s.Pop()

		// If left/right nodes have children, move them onto
		// the temporary stack for later processing.
		// This means we can ensure the order of processing is
		// correct.

		if n.Left != nil && added(n.Left) == false {
			tmp.Append(n.Left)
		}

		if n.Right != nil && added(n.Right) == false {
			tmp.Append(n.Right)
		}

		if tmp.Len() == 0 {
			r = append(r, n.Value)
			tracker[n] = struct{}{}
		} else {
			// Move everything back onto the main stack for further
			// processing starting with the root node which will
			// therefore be processed last.
			s.Append(n)

			for tmp.Len() > 0 {
				tn := tmp.Pop()
				s.Append(tn)
			}
		}

	}
	return r
}
