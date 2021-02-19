package main

import "testing"

func TestGetMinMaxNumber(t *testing.T) {
	cases := []struct {
		numbers []int
		want    MinMaxResult
	}{
		{numbers: []int{1, -92, 5, 9, 55, 90, 99, 100, 55}, want: MinMaxResult{Min: -92, Max: 100}},
		{numbers: []int{5, 4, 3, 2, 1}, want: MinMaxResult{Min: 1, Max: 5}},
		{numbers: []int{1000, 1001, 100, 1000, 0}, want: MinMaxResult{Min: 0, Max: 1001}},
	}

	for _, c := range cases {
		got := GetMinMaxNumber(c.numbers)

		if got != c.want {
			t.Errorf("GetDuplicateNumber(%v) = %v, want %v", c.numbers, got, c.want)
		}
	}
}
