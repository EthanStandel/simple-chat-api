package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(input string) string {
	output, _ := bcrypt.GenerateFromPassword(
		[]byte(input),
		bcrypt.DefaultCost,
	)
	return string(output)
}

func ValidateAgainstHash(input, passhash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(passhash),
		[]byte(input),
	)

	return err == nil
}

const minLength, maxLength = 15, 50

func ValidateAgainstRequirements(input string) bool {
	if len(input) < minLength {
		return false
	}
	if len(input) > maxLength {
		return false
	}

	return true
}
