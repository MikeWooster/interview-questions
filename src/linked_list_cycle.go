package main

func CheckLinkedListCycle(l LinkedList) bool {
	// Keep a record of the visited nodes.
	// using an empty struct type as we don't need
	// any value stored and structs take up 0 space.
	visited := make(map[*Node]struct{})

	node := l.Head

	// While node isn't nil
	for node != nil {
		if _, ok := visited[node]; ok {
			// This node is in the map, so we have already seen this.
			// Therefore there is a cycle in the linked list.
			return true
		}
		visited[node] = struct{}{}

		node = node.Next
	}

	// No cycles detected.
	return false
}

func FindLinkedListCycle(l LinkedList) *Node {
	visited := make(map[*Node]struct{})

	node := l.Head

	// While node isn't nil
	for node != nil {

		if _, ok := visited[node.Next]; ok {
			// If the next node is in the map, then this node is the start of a cycle.
			return node
		}
		visited[node] = struct{}{}

		node = node.Next
	}

	// No cycles detected.
	return nil
}
