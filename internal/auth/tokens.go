package auth

import (
	"crypto/rand"
	"encoding/hex"
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
		return uuid.Nil, errors.New(InterpretClaimsErrMsg)
	}

	ID, err := uuid.Parse(tokenClaims.Subject)

	if err != nil {
		return uuid.Nil, err
	}

	return ID, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	token := headers.Get(AuthHeader)

	if !strings.HasPrefix(token, BearerPrefix) {
		return "", errors.New(InvalidAuthHeaderErrMsg)
	}

	cleanedToken := strings.TrimPrefix(token, BearerPrefix)

	if cleanedToken == "" {
		return "", errors.New(NoAuthHeaderErrMsg)
	}

	return cleanedToken, nil
}

func MakeRefreshToken() (string, error) {
	randomString := make([]byte, 32)
	rand.Read(randomString)
	hexString := hex.EncodeToString(randomString)
	return hexString, nil
}
