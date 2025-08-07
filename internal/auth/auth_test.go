package auth

import "testing"

// Tests to make sure HashPassword doesn't return an error
func TestHashPassword(t *testing.T) {
	samplePassword := "password"
	hashedPassword, err := HashPassword(samplePassword)

	if err != nil {
		t.Errorf("Converting 'password' with HashPassword resulted in: %v", err)
		return
	}

	if hashedPassword == "" {
		t.Errorf("Converting 'password' with HashPassword resulted in empty string")
		return
	}
}

// Tests that a simple hashed password will match with itself
func TestSimpleHashedPasswordMatch(t *testing.T) {
	samplePassword := "password"
	hashedPassword, err := HashPassword(samplePassword)

	if err != nil {
		t.Fatal()
	}

	err = CheckPasswordHash(samplePassword, hashedPassword)

	if err != nil {
		t.Errorf("'password' didn't match itself after hashing")
	}
}

// Tests that a complex hashed password will match with itself
func TestComplexHashedPasswordMatch(t *testing.T) {
	samplePassword := "xVc8903h3m!8F33&6%&*$D/!"
	hashedPassword, err := HashPassword(samplePassword)

	if err != nil {
		t.Fatal()
	}

	err = CheckPasswordHash(samplePassword, hashedPassword)

	if err != nil {
		t.Errorf("'password' didn't match itself after hashing")
	}
}
