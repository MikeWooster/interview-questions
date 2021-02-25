package main

import "testing"

func TestRemoveDuplicateNumbers(t *testing.T) {
	cases := []struct {
		numbers []int
		want    []int
	}{
		{numbers: []int{1, 2, 3, 3, 4, 5, 2, 6, 7, 1}, want: []int{1, 2, 3, 4, 5, 6, 7}},
		{numbers: []int{1, 2, 3, 4, 2, 5}, want: []int{1, 2, 3, 4, 5}},
		{numbers: []int{10, 1001, 100, 1000, 0}, want: []int{10, 1001, 100, 1000, 0}},
	}

	for _, c := range cases {
		got := RemoveDuplicateNumbers(c.numbers)

		if cmpIntSlice(got, c.want) == false {
			t.Errorf("RemoveDuplicateNumbers(%v) = %v, want %v", c.numbers, got, c.want)
		}
	}
}
