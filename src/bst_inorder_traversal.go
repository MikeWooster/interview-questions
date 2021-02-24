package main

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
