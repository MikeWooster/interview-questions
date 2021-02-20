package main

import (
	"errors"
	"fmt"
)

// Node stores the Value of the linked list as an integer.
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents a singly linked list
// containing integer values.
type LinkedList struct {
	Head *Node
	Tail *Node
	len  int
}

// GetLen returns the current length of the linked list.
func (l *LinkedList) GetLen() int { return l.len }

// Append adds a single new integer to the linked list.
func (l *LinkedList) Append(v int) {
	// Create a new node with no Next defined.
	node := &Node{Value: v}

	if l.GetLen() == 0 {
		// The list is empty, we need to start the Head.
		l.Head = node
	} else {
		// Get the last element and set it's Next Value
		prev := l.Tail
		prev.Next = node
	}
	// This new node is now the new last position in the list.
	l.Tail = node

	// We can increment the length of the list now.
	l.len++
}

// Extend adds a x number of integers to the linked list.
func (l *LinkedList) Extend(vals []int) {
	for _, v := range vals {
		l.Append(v)
	}
}

// AtIndex returns the integer stored at position `i` in the linked list.
func (l *LinkedList) AtIndex(i int) (int, error) {
	if i >= l.GetLen()-1 {
		return 0, errors.New("index out of range")
	}
	node := l.Head
	for x := 0; x < i; x++ {
		node = node.Next
		if node == nil {
			return 0, fmt.Errorf("unexpected error accessing node at position %v", i)
		}
	}

	return node.Value, nil
}

// Equals compares the value and order of 2 linked lists
func (l *LinkedList) Equals(other LinkedList) bool {
	if l.GetLen() != other.GetLen() {
		return false
	}

	thisNode := l.Head
	otherNode := other.Head

	for thisNode != nil {
		if thisNode.Value != otherNode.Value {
			return false
		}
		thisNode = thisNode.Next
		otherNode = otherNode.Next
	}
	return true
}

// Values returns the list of all values in the linked list
func (l *LinkedList) Values() []int {
	var values []int

	node := l.Head
	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}

	return values
}
