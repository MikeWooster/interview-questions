package main

// GetDuplicateChars returns the characters that are
// repeated in the given string, in the order they
// first appear in the string.
func GetDuplicateChars(chars string) []string {
	var repeated []string

	charCount := make(map[string]int)

	for _, c := range chars {
		s := string(c)

		// Ignore whitespace
		if s == " " {
			continue
		}

		// the map will default to 0, so can safely increment on
		// the map.
		charCount[s]++

	}

	// As a map is created without any guarantee of order,
	// to get the chars in an the order they first appear,
	// you re-iterate over the original string to get the order,
	// removing from the count map to avoid duplicating in the output.
	for _, c := range chars {
		s := string(c)
		if charCount[s] > 1 {
			repeated = append(repeated, s)
		}
		charCount[s] = 0
	}

	return repeated

}
