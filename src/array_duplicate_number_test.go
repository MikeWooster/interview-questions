package main

import "testing"

func TestGetDuplicateNumber(t *testing.T) {
	cases := []struct {
		numbers []int
		want    int
	}{
		{numbers: []int{1, 5, 9, 55, 90, 99, 100, 55}, want: 55},
		{numbers: []int{1, 2, 3, 4, 2, 5}, want: 2},
		{numbers: []int{1000, 1001, 100, 1000, 0}, want: 1000},
	}

	for _, c := range cases {
		got, err := GetDuplicateNumber(c.numbers)
		if err != nil {
			t.Errorf("GetDuplicateNumber(%v) returned an error %v", c.numbers, err)
		}
		if got != c.want {
			t.Errorf("GetDuplicateNumber(%v) = %v, want %v", c.numbers, got, c.want)
		}
	}
}

func TestGetDuplicateNumberWithError(t *testing.T) {
	cases := []struct {
		numbers []int
	}{
		{numbers: []int{}},
		{numbers: []int{1, 2, 3, 4, 5}},
	}

	for _, c := range cases {
		_, err := GetDuplicateNumber(c.numbers)
		if err == nil {
			t.Errorf("GetDuplicateNumber(%v) DID NOT return an error.", c.numbers)
		}
	}
}
