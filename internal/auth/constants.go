package auth

const (
	AuthHeader = "Authorization"

	BearerPrefix = "Bearer "
	ApiKeyPrefix = "ApiKey "

	// Error Messages
	InterpretClaimsErrMsg   = "Unable to interpret claims"
	InvalidAuthHeaderErrMsg = "Invalid Authorization header format"
	NoAuthHeaderErrMsg      = "No authorization header found"
)
