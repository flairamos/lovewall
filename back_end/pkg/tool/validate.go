package tool

import "regexp"

func ValidateEmail(email string) bool {
	regex := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	return regex.MatchString(email)
}
