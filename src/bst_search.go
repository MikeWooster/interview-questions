package main

// Search finds a the value of element in the tree, and reports whether it
// is present or not.
func (t *BinarySearchTree) Search(v int) bool {
	finder := t.getNode(v)
	return finder != nil
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
