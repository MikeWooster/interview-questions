package main

// ReverseArray returns a new slice in the reverse order to
// the original.
func ReverseArray(numbers []int) (reversed []int) {

	for i := len(numbers) - 1; i >= 0; i-- {
		reversed = append(reversed, numbers[i])
	}

	return reversed
}
