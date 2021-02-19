package main

import "errors"

// GetDuplicateNumber aims to find the duplicated number in a slice of numbers.
// This assumes that only one number is duplicated.
// Solution O(n)
func GetDuplicateNumber(numbers []int) (int, error) {

	count := make(map[int]struct{})

	for _, n := range numbers {
		if _, ok := count[n]; ok {
			// Number is in map - it is repeated.
			return n, nil
		} else {
			count[n] = struct{}{}
		}
	}
	// If it reached this point, then the spec given to us did not work.
	// At this point, we should return an error.
	return 0, errors.New("no number was repeated in the slice")
}
