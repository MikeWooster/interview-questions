package main

// RemoveDuplicateNumbers takes an array and removes any duplicates
// if they are present. This removes any subsequent dupes.
func RemoveDuplicateNumbers(numbers []int) []int {

	// Create a map to store numbers as they appear in array.
	count := make(map[int]struct{})

	var out []int

	for _, n := range numbers {
		if _, ok := count[n]; ok {
			// Number is a duplicate - ignore it.
			continue
		}
		out = append(out, n)
		count[n] = struct{}{}
	}

	return out
}
