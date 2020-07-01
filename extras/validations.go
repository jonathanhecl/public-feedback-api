package extras

import "regexp"

func ValidateEmail(email string) bool {

	return regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email)

}
