package main

import "testing"

func TestReverseLinkedList(t *testing.T) {
	l := LinkedList{}

	l.Extend([]int{11, 34, 62, 10, 96, 17, 12, 83, 23, 51, 49})

	l.Reverse()

	want := LinkedList{}
	want.Extend([]int{49, 51, 23, 83, 12, 17, 96, 10, 62, 34, 11})
	if l.Equals(want) == false {
		t.Errorf("Reverse(%v) = %v want %v", l.Values(), l.Values(), want.Values())
	}
}
