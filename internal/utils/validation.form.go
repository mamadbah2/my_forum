package utils

import (
	"regexp"
	"strings"
)

func UsernameValidation(username string) bool {
	if len(username) == 0 || len(username) > 15 {
		return false
	}
	for _, us := range username {
		if us < 'A' || us > 'Z' && us < 'a' || us > 'z' {
			return false
		}
	}
	return true
}

func EmailValidation(email string) bool {
	rx := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return rx.MatchString(email)
}

func PasswordValidation(password string) bool {
	if strings.TrimSpace(password) == "" {
		return false
	}
	return true
}