package main

func (l *LinkedList) Reverse() {
	// prev/next will be initialized as nil
	var prev, next *Node

	current := l.Head

	for current != nil {
		// Store the next node for the next iteration.
		next = current.Next

		// Move the pointer to point back at the previous node.
		current.Next = prev

		// Move next and current one node along.
		prev = current
		current = next
	}
	l.Head = prev
}
