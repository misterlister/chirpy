package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn).UTC()),
		Subject:   userID.String(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := jwtToken.SignedString([]byte(tokenSecret))

	return tokenString, err
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := jwt.RegisteredClaims{}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	}
	jwtToken, err := jwt.ParseWithClaims(tokenString, &claims, keyFunc)

	if err != nil {
		return uuid.Nil, err
	}

	tokenClaims, ok := jwtToken.Claims.(*jwt.RegisteredClaims)

	if !ok {
		return uuid.Nil, errors.New("Unable to interpret claims")
	}

	ID, err := uuid.Parse(tokenClaims.Subject)

	if err != nil {
		return uuid.Nil, err
	}

	return ID, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	token := headers.Get("Authorization")

	if !strings.HasPrefix(token, "Bearer ") {
		return "", errors.New("Invalid Authorization header format")
	}

	cleanedToken := strings.TrimPrefix(token, "Bearer ")

	if cleanedToken == "" {
		return "", errors.New("No authorization header found")
	}

	return cleanedToken, nil
}
