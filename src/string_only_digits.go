package main

// StringIsNumber returns true when the string contains only digits.
// O(n)
func StringIsNumber(s string) bool {
	numbers := map[string]struct{}{
		"0": {},
		"1": {},
		"2": {},
		"3": {},
		"4": {},
		"5": {},
		"6": {},
		"7": {},
		"8": {},
		"9": {},
	}

	// store a count of the number of digits
	count := 0
	// Get the total length of the string
	size := 0

	for _, c := range s {
		if _, ok := numbers[string(c)]; ok {
			// String contains a number
			count++
		}
		size++
	}

	// if count == size, the string contains only digits.
	return count == size
}
