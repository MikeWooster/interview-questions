package main

// GetFirstNonRepeatedChar finds the first non repeated character
// in a string.
func GetFirstNonRepeatedChar(chars string) string {
	// First get a count of all characters in the string.
	counter := make(map[string]int)
	for _, c := range chars {
		s := string(c)

		if s == " " {
			// don't include whitespace in count
			continue
		}

		counter[s]++
	}

	for _, c := range chars {
		if counter[string(c)] == 1 {
			return string(c)
		}
	}

	// No repeated char found, return empty string
	return ""
}
