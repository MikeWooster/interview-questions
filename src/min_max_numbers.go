package main

type MinMaxResult struct {
	Min, Max int
}

// GetMinMaxNumber find the min and max number in a slice of numbers.
// This involves a single iteration of the number slice, resulting in
// an O(n) solution.
func GetMinMaxNumber(numbers []int) MinMaxResult {
	// First we create the initial result, setting both min
	// and max to the first element.
	result := MinMaxResult{Min: numbers[0], Max: numbers[0]}

	// Go through the numbers once to set the min/max
	for _, n := range numbers {
		if n < result.Min {
			result.Min = n
		}
		if n > result.Max {
			result.Max = n
		}
	}
	return result
}
