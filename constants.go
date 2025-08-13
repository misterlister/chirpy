package main

const (
	// Port
	Port = "8080"

	// Path strings
	AppPrefix         = "/app"
	ApiPrefix         = "/api"
	AdminPrefix       = "/admin"
	HealthPath        = "/healthz"
	MetricPath        = "/metrics"
	ResetPath         = "/reset"
	ValidateChirpPath = "/validate_chirp"
	ChirpsPath        = "/chirps"
	UsersPath         = "/users"
	LoginPath         = "/login"
	RefreshPath       = "/refresh"
	RevokePath        = "/revoke"
	PolkaPath         = "/polka"
	WebhooksPath      = "/webhooks"

	// Request strings
	GetReq    = "GET "
	PutReq    = "PUT "
	DeleteReq = "DELETE "
	PostReq   = "POST "

	// Header strings
	ContentType = "Content-Type"
	TextHeader  = "text/plain; charset=utf-8"
	HtmlHeader  = "text/html"
	JsonHeader  = "application/json"

	// Defined numbers
	MaxChirpLength = 140

	// Path variables

	ChirpID = "chirpID"

	// Error message strings
	UnknownErrMsg           = "Something went wrong"
	TooLongErrMsg           = "Chirp is too long"
	MissingParamErrMsg      = "User did not provide required content"
	DatabaseErrMsg          = "Error connecting to the database"
	DatabaseInitErrMsg      = "Error initializing database connection"
	GetChirpsErrMsg         = "Error retrieving Chirps from database"
	NoChirpFoundErrMsg      = "No matching Chirp found"
	UUIDErrMsg              = "Invalid UUID"
	UnauthorizedErrMsg      = "Unauthorized action attempted"
	PasswordFailErrMsg      = "Email and Password do not match"
	TokenFailErrMsg         = "Error generating login token"
	DeleteAuthErrMsg        = "Cannot delete another user's Chirp"
	UnrecognizedEventErrMsg = "Unrecognized event"
	CensorSymbol            = "****"

	// Success message strings
	ResetMsg = "All users deleted. Hits reset to 0"

	// Webhook events
	UpgradeUser = "user.upgraded"
)

// Words that must be censored
var BadWords = []string{
	"kerfuffle",
	"sharbert",
	"fornax",
}
