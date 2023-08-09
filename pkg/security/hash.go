package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashString string, compareString string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashString), []byte(compareString))
}
