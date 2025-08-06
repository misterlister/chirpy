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

	// Error message strings
	UnknownErrMsg      = "Something went wrong"
	TooLongErrMsg      = "Chirp is too long"
	MissingParamErrMsg = "User did not provide required content"
	DatabaseErrMsg     = "Error connecting to the database"
	DatabaseInitErrMsg = "Error initializing database connection"
	GetChirpsErrMsg    = "Error retrieving Chirps from database"
	UnauthorizedErrMsg = "Unauthorized action attempted"
	CensorSymbol       = "****"

	// Success message strings
	ResetMsg = "All users deleted. Hits reset to 0"
)

// Words that must be censored
var BadWords = []string{
	"kerfuffle",
	"sharbert",
	"fornax",
}
