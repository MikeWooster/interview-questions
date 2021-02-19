package main

// ReverseStringWithRecursion using recursion, reverse a string.
func ReverseStringWithRecursion(chars string) string {
	if len(chars) == 0 {
		return chars
	}
	return ReverseStringWithRecursion(chars[1:]) + string(chars[0])
}

// ReverseString performs the same action, but does not use recursion.
func ReverseString(chars string) string {
	reversed := ""
	for i := len(chars) - 1; i >= 0; i-- {
		reversed += string(chars[i])
	}
	return reversed
}
