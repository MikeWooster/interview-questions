package main

// GetDuplicateNumbers aims to find all duplicated numbers in a slice of
// numbers. I have assumed the order is not important.
// Solution O(n)
// This makes use of Go's named return values.
func GetDuplicateNumbers(numbers []int) (dupes []int) {

	// keep track of numbers in the array
	count := make(map[int]struct{})

	for _, n := range numbers {
		if _, ok := count[n]; ok {
			// Number is in map - it is repeated.
			dupes = append(dupes, n)
		} else {
			// First occurrence of the number, store in set.
			count[n] = struct{}{}
		}
	}

	return dupes
}

// GetDuplicateNumbersInOrder aims to find all duplicated numbers in a slice of
// numbers, in the order they originally appeared in the array
// Solution O(n)
// This makes use of Go's named return values.
func GetDuplicateNumbersInOrder(numbers []int) (dupes []int) {

	// keep track of numbers in the array
	count := make(map[int]struct{})
	// Another "Set" to store the duplicates.
	dupeMap := make(map[int]struct{})

	for _, n := range numbers {
		if _, ok := count[n]; ok {
			// Number is in map - it is repeated.
			dupeMap[n] = struct{}{}
		} else {
			// First occurrence of the number, store in set.
			count[n] = struct{}{}
		}
	}

	// Go through the numbers one more time to get the dupes in order
	for _, n := range numbers {
		if _, ok := dupeMap[n]; ok {
			// Remove from the map so we don't get dupes in the actual
			// return slice.
			delete(dupeMap, n)
			dupes = append(dupes, n)
		}
	}

	return dupes
}
