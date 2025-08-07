package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) error {
	passBytes := []byte(password)
	hashBytes := []byte(hash)
	err := bcrypt.CompareHashAndPassword(hashBytes, passBytes)

	return err
}
