package main

import (
	"errors"
)

func GetLinkedListMiddleElement(l LinkedList) (int, error) {
	middle := l.GetLen() / 2
	return l.AtIndex(middle)
}

// GetLinkedListMiddleElementNoLen performs the same operation,
// but assumes the linked list doesnt store the length.
func GetLinkedListMiddleElementNoLen(l LinkedList) (int, error) {
	node := l.Head
	if node == nil {
		return 0, errors.New("index out of range")
	}

	half_node := node
	count := 0

	// Move the main pointer all the way to the end of the list
	// whilst every other node, move the other pointer along.
	// This means that when you get to the end of the list using
	// the first pointer, the other pointer will be in the middle.
	for {
		node = node.Next
		if node == nil {
			break
		}

		if count%2 == 0 {
			half_node = half_node.Next
		}

		count++
	}

	return half_node.Value, nil
}
