package main

import (
	"errors"
	"fmt"
)

// Node stores the value of the linked list as an integer.
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a singly linked list
// containing integer values.
type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

// GetLen returns the current length of the linked list.
func (l *LinkedList) GetLen() int { return l.len }

// Append adds a single new integer to the linked list.
func (l *LinkedList) Append(v int) {
	// Create a new node with no next defined.
	node := &Node{value: v}

	if l.GetLen() == 0 {
		// The list is empty, we need to start the head.
		l.head = node
	} else {
		// Get the last element and set it's next value
		prev := l.tail
		prev.next = node
	}
	// This new node is now the new last position in the list.
	l.tail = node

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
	node := l.head
	for x := 0; x < i; x++ {
		node = node.next
		if node == nil {
			return 0, fmt.Errorf("unexpected error accessing node at position %v", i)
		}
	}

	return node.value, nil
}

func GetLinkedListMiddleElement(l LinkedList) (int, error) {
	middle := l.GetLen() / 2
	return l.AtIndex(middle)
}

// GetLinkedListMiddleElementNoLen performs the same operation,
// but assumes the linked list doesnt store the length.
func GetLinkedListMiddleElementNoLen(l LinkedList) (int, error) {
	node := l.head
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
		node = node.next
		if node == nil {
			break
		}

		if count%2 == 0 {
			half_node = half_node.next
		}

		count++
	}

	return half_node.value, nil
}
