package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authorization := headers.Get(AuthHeader)

	if !strings.HasPrefix(authorization, ApiKeyPrefix) {
		return "", errors.New(InvalidAuthHeaderErrMsg)
	}

	apiKey := strings.TrimPrefix(authorization, ApiKeyPrefix)

	if apiKey == "" {
		return "", errors.New(NoAuthHeaderErrMsg)
	}

	return apiKey, nil
}
