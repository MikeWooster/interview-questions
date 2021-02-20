package main

import "testing"

func TestCheckLinkedListCycle(t *testing.T) {
	l := LinkedList{}

	l.Extend([]int{11, 34, 62, 10, 96, 17, 12, 83, 23, 51, 49})
	// Set the last nodes next value equal to the first.
	l.Tail.Next = l.Head
	l.Tail = l.Head

	got := CheckLinkedListCycle(l)

	want := true
	if got != want {
		t.Errorf("CheckLinkedListCycle(%v) = %v want %v", l, got, want)
	}
}

func TestCheckLinkedListCycleNoCycle(t *testing.T) {
	l := LinkedList{}

	l.Extend([]int{11, 34, 62, 10, 96, 17, 12, 83, 23, 51, 49})

	got := CheckLinkedListCycle(l)

	want := false
	if got != want {
		t.Errorf("CheckLinkedListCycle(%v) = %v want %v", l, got, want)
	}
}

func TestCheckLinkedListCycleEmpty(t *testing.T) {
	l := LinkedList{}

	got := CheckLinkedListCycle(l)

	want := false
	if got != want {
		t.Errorf("CheckLinkedListCycle(%v) = %v want %v", l, got, want)
	}
}

func TestFindLinkedListCycle(t *testing.T) {
	l := LinkedList{}

	l.Extend([]int{11, 34, 62, 10, 96, 17, 12, 83, 23, 51, 49})
	// Set the last nodes next value equal to the first.
	start := l.Tail
	l.Tail.Next = l.Head
	l.Tail = l.Head

	got := FindLinkedListCycle(l)

	want := start
	if got != want {
		t.Errorf("FindLinkedListCycle(%v) = %v want %v", l, got, want)
	}
}

func TestFindLinkedListCycleNoCycle(t *testing.T) {
	l := LinkedList{}

	l.Extend([]int{11, 34, 62, 10, 96, 17, 12, 83, 23, 51, 49})

	got := FindLinkedListCycle(l)

	if got != nil {
		t.Errorf("FindLinkedListCycle(%v) = %v want %v", l, got, nil)
	}
}
