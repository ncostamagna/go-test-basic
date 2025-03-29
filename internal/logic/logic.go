package logic

import "regexp"

func Sum(n ...int) int {
	var total int
	for _, v := range n {
		total += v
	}
	return total
}

func CheckEmail(email string) bool {
	// Basic pattern for validating email addresses
	// This regex checks for:
	// - At least one character before the @ symbol
	// - At least one character for domain name
	// - At least one dot in the domain part
	// - At least two characters after the last dot (TLD)
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	// Check if the email matches the pattern
	return re.MatchString(email)
}
