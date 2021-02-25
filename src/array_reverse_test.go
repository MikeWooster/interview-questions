package main

import "testing"

func TestReverseArray(t *testing.T) {
	cases := []struct {
		numbers []int
		want    []int
	}{
		{numbers: []int{1, 2, 3, 3, 4, 5, 2, 6, 7, 1}, want: []int{1, 7, 6, 2, 5, 4, 3, 3, 2, 1}},
		{numbers: []int{1, 2, 3, 4, 2, 5}, want: []int{5, 2, 4, 3, 2, 1}},
		{numbers: []int{10, 1001, 100, 1000, 0}, want: []int{0, 1000, 100, 1001, 10}},
	}

	for _, c := range cases {
		got := ReverseArray(c.numbers)

		if cmpIntSlice(got, c.want) == false {
			t.Errorf("ReverseArray(%v) = %v, want %v", c.numbers, got, c.want)
		}
	}
}
