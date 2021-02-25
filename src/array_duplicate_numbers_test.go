package main

import "testing"

func TestGetDuplicateNumbers(t *testing.T) {
	cases := []struct {
		numbers []int
		want    []int
	}{
		{numbers: []int{1, 2, 3, 3, 4, 5, 2, 6, 7, 1}, want: []int{3, 2, 1}},
		{numbers: []int{1, 2, 3, 4, 2, 5}, want: []int{2}},
		{numbers: []int{10, 1001, 100, 1000, 0}, want: []int{}},
	}

	for _, c := range cases {
		got := GetDuplicateNumbers(c.numbers)

		if cmpIntSlice(got, c.want) == false {
			t.Errorf("GetDuplicateNumbers(%v) = %v, want %v", c.numbers, got, c.want)
		}
	}
}

func TestGetDuplicateNumbersInOrder(t *testing.T) {
	cases := []struct {
		numbers []int
		want    []int
	}{
		{numbers: []int{1, 2, 3, 3, 4, 5, 2, 6, 7, 1}, want: []int{1, 2, 3}},
		{numbers: []int{1, 2, 3, 4, 2, 5}, want: []int{2}},
		{numbers: []int{10, 1001, 100, 1000, 0}, want: []int{}},
	}

	for _, c := range cases {
		got := GetDuplicateNumbersInOrder(c.numbers)

		if cmpIntSlice(got, c.want) == false {
			t.Errorf("GetDuplicateNumbersInOrder(%v) = %v, want %v", c.numbers, got, c.want)
		}
	}
}
