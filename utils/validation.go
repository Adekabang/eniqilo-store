package utils

import (
	"regexp"
)

func ValidatePhoneNumber(phoneNumber string) bool {
	if phoneNumber == "" || len(phoneNumber) < 10 || len(phoneNumber) > 16 || phoneNumber[0] != '+' {
		return false
	}
	pattern := `^\+[1-9]\d{1,14}$`
	matched, err := regexp.MatchString(pattern, phoneNumber)
	if err != nil || !matched {
		return false
	}

	return true
}

// ValidateEmail checks if the email is in the correct format
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword checks if the password meets length requirements
func ValidatePassword(password string) bool {
	return len(password) >= 5 && len(password) <= 15
}

// ValidateName checks if the name meets length requirements
func ValidateName(name string) bool {
	return len(name) >= 5 && len(name) <= 50
}
