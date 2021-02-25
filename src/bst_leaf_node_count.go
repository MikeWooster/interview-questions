package main

// CountLeafNodes returns the number of leaf nodes
// in the tree.  A leaf node is defined as a node
// that has no children.
func (t *BinarySearchTree) CountLeafNodes() int {
	s := nodeStack{}
	s.Append(t.Root)

	count := 0
	for s.Len() > 0 {
		n := s.Pop()
		if n.Right != nil {
			s.Append(n.Right)
		}
		if n.Left != nil {
			s.Append(n.Left)
		}
		if n.HasChildren() == false {
			count++
		}
	}

	return count
}
