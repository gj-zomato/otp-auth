package validation

import "strings"

// ValidOTP checks if the provided OTP is a valid 6-digit numeric string.
func ValidOTP(otp string) bool {
	if len(otp) != 6 {
		return false
	}
	for _, char := range otp {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

// CheckIdentifier determines if the identifier is an email or phone number.
func DetermineIdentifier(identifier string) string {
	if strings.Contains(identifier, "@") {
		return "email" // Email identifier
	}
	for _, char := range identifier {
		if char < '0' || char > '9' {
			return "unknown" // Not a valid phone number
		}
	}
	return "phone" // Phone number identifier
}

// ValidEmail checks if the provided email is in a valid format.
func ValidEmail(email string) bool {
	if len(email) < 3 || !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
		return false
	}
	domainParts := strings.Split(parts[1], ".")
	if len(domainParts) < 2 || domainParts[0] == "" || domainParts[len(domainParts)-1] == "" {
		return false
	}
	return true
}

// ValidPhoneNumber checks if the provided phone number is a valid numeric string.
func ValidPhoneNumber(phone string) bool {
	if len(phone) < 10 || len(phone) > 15 {
		return false
	}
	for _, char := range phone {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
