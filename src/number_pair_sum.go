package main

// PairSum returns all numbers in an integer array whose sum is
// equal to a given number.
// For this I am assuming that you can only use a number from the array once.
// This involves a single iteration of the array - therefore O(n)
func PairSum(numbers []int, total int) [][2]int {
	var pairs [][2]int

	// keep a total of the occurrence of numbers not yet viable as a pair.
	counter := make(map[int]int)

	for _, n := range numbers {

		remainder := total - n

		if counter[remainder] > 0 {
			// The number required to create a sum is present in the map.
			// We don't need to keep track of this occurrence of the number,
			// and we can decrement the remainder as we won't be using this
			// any more.
			counter[remainder]--
			pairs = append(pairs, [2]int{remainder, n})
		} else {
			// No pair found, add this number to the counter
			counter[n]++
		}
	}

	return pairs
}
