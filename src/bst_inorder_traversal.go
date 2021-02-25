package main

// InorderTraversalRecursive traverses the binary tree recursively inorder
// i.e. left, root, right
func (t *BinarySearchTree) InorderTraversalRecursive() []int {
	return t.inorderTraversalRecursive(t.Root)
}

func (t *BinarySearchTree) inorderTraversalRecursive(n *BinarySearchNode) []int {
	var r []int

	if n == nil {
		return r
	}

	if n.Left != nil {
		for _, v := range t.inorderTraversalRecursive(n.Left) {
			r = append(r, v)
		}
	}

	r = append(r, n.Value)

	if n.Right != nil {
		for _, v := range t.inorderTraversalRecursive(n.Right) {
			r = append(r, v)
		}
	}

	return r
}

// InorderTraversal traverses the binary tree inorder
// i.e. left, root, right
func (t *BinarySearchTree) InorderTraversal() []int {
	var r []int

	if t.Root == nil {
		return r
	}

	// Create the stack starting with the root node
	s := nodeStack{}

	current := t.Root

	for {

		if current != nil {
			s.Append(current)
			current = current.Left

		} else if s.Len() > 0 {
			current = s.Pop()
			r = append(r, current.Value)
			current = current.Right

		} else {
			break

		}
	}
	return r
}
