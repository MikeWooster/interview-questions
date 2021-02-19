package main

import "testing"

func TestGetLinkedListMiddleElement(t *testing.T) {
	ll := LinkedList{}

	ll.Extend([]int{11, 34, 62, 10, 96, 17, 12, 83, 23, 51, 49})

	got, err := GetLinkedListMiddleElement(ll)
	if err != nil {
		t.Fatalf("GetLinkedListMiddleElement(%v) returned unexpected error %v", ll, err)
	}
	want := 17
	if got != want {
		t.Errorf("GetLinkedListMiddleElement(%v) = %v want %v", ll, got, want)
	}
}

func TestGetLinkedListMiddleElementNoLen(t *testing.T) {
	ll := LinkedList{}

	ll.Extend([]int{11, 34, 62, 10, 96, 17, 12, 83, 23, 51, 49})

	got, err := GetLinkedListMiddleElementNoLen(ll)
	if err != nil {
		t.Fatalf("GetLinkedListMiddleElement(%v) returned unexpected error %v", ll, err)
	}
	want := 17
	if got != want {
		t.Errorf("GetLinkedListMiddleElement(%v) = %v want %v", ll, got, want)
	}
}
