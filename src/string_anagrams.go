package main

// CheckStringAnagrams returns whether two strings are an anagram
// of each-other.  Whitespace is ignored in both cases.
func CheckStringAnagrams(a, b string) bool {
	// First create a counter of the string chars in b
	counter := make(map[string]int)
	for _, c := range b {
		s := string(c)
		// ignore whitespace
		if s == " " {
			continue
		}
		counter[s]++
	}

	// Next go through the string a and reduce the count.
	// Any reduction below 0 means it is not an anagram.
	for _, c := range a {
		s := string(c)
		// ignore whitespace
		if s == " " {
			continue
		}
		if counter[s] == 0 {
			return false
		}
		counter[s]--
	}

	// Reaching this point means its an anagram
	return true
}
