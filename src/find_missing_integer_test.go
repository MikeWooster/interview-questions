package main

import "testing"

func TestGetMissingInteger(t *testing.T) {
	cases := []int{1, 5, 9, 55, 90, 99, 100}

	for _, n := range cases {
		numbers := createSliceWithMissingNumber(n)
		got := GetMissingInteger(numbers)
		if got != n {
			t.Errorf("GetFirstNonRepeatedChar(%v) = %v, want %v", numbers, got, n)
		}
	}
}

func createSliceWithMissingNumber(n int) []int {
	var numbers []int
	for i := 1; i <= 100; i++ {
		if i == n {
			continue
		}
		numbers = append(numbers, i)
	}
	return numbers
}
