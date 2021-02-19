package main

// GetMissingInteger find the missing integer in an array of numbers 1 to 100.
// A naive approach would be to go through the expected numbers 1 to 100 and
// see if it's missing from the array.  That's an O(n^2) solution.  A better
// solution would be to take the sum of the numbers 1 -> 100 and calculate the
// sum of the slice, with the difference being the missing number, resulting
// in a O(n) solution.
func GetMissingInteger(numbers []int) int {
	max := 100

	// Rather than looping through the expected numbers,
	// a simple formula sidesteps this part.
	expectedSum := (max * (max + 1)) / 2

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return expectedSum - sum
}
