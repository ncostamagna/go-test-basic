package logic

import "regexp"

var emailRegex *regexp.Regexp
var pattern = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9][a-zA-Z0-9.\-]*\.[a-zA-Z]{2,}$`

func init() {
	var err error
	emailRegex, err = regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
}

func Sum(n ...int) int {
	return internalSum(n...)
}

func internalSum(n ...int) int {
	var total int
	for _, v := range n {
		total += v
	}
	return total
}

func CheckEmail(email string) bool {
	if email == "" {
		return false
	}
	return emailRegex.MatchString(email)
}
