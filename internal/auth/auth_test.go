package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

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

// Test the creation of a JWT with valid input
func TestMakeJWTValid(t *testing.T) {
	userID := uuid.New()
	secretString := "secret"
	expiresIn := time.Minute
	tokenString, err := MakeJWT(userID, secretString, expiresIn)

	if err != nil {
		t.Errorf("MakeJWT with 'secret' and 1 minute expiry failed to create: %v", err)
	}

	validatedUUID, err := ValidateJWT(tokenString, secretString)

	if err != nil {
		t.Errorf("MakeJWT with 'secret' and 1 minute expiry failed to validate: %v", err)
	}

	if validatedUUID != userID {
		t.Errorf("MakeJWT with 'secret' and 1 minute expiry did not validate uuid. Expected %s, got %s", userID, validatedUUID)
	}
}

// Test a valid Bearer token
func TestGetBearerTokenValid(t *testing.T) {
	headers := http.Header{}
	tokenString := "asldkfjaslkdfrjal"
	headers.Set("Authorization", "Bearer "+tokenString)
	token, err := GetBearerToken(headers)

	if err != nil {
		t.Errorf("Header 'Authorization: Bearer asldkfjaslkdfrjal' failed to validate: %v", err)
	}

	if token != tokenString {
		t.Errorf("Token didn't parse correctly. Expected: %v, Got: ", err)
	}
}

// Test an invalid Bearer token
func TestGetBearerTokenInvalid(t *testing.T) {
	headers := http.Header{}
	tokenString := "asldkfjaslkdfrjal"
	headers.Set("Authorization", "BBearerr "+tokenString)
	token, err := GetBearerToken(headers)

	if err == nil {
		t.Errorf("Header 'Authorization: BBearerr asldkfjaslkdfrjal' failed to generate error. Returned: %s", token)
	}
}

// Test an empty Bearer token
func TestGetBearerTokenEmpty(t *testing.T) {
	headers := http.Header{}
	tokenString := ""
	headers.Set("Authorization", "Bearer "+tokenString)
	token, err := GetBearerToken(headers)

	if err == nil {
		t.Errorf("Header 'Authorization: Bearer ' failed to generate error. Returned: %s", token)
	}
}
