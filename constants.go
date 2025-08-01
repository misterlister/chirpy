package main

const (
	Port              = "8080"
	AppPrefix         = "/app"
	ApiPrefix         = "/api"
	AdminPrefix       = "/admin"
	HealthPath        = "/healthz"
	MetricPath        = "/metrics"
	ResetPath         = "/reset"
	ValidateChirpPath = "/validate_chirp"
	GetReq            = "GET "
	PutReq            = "PUT "
	DeleteReq         = "DELETE "
	PostReq           = "POST "
	TextHeader        = "text/plain; charset=utf-8"
	HtmlHeader        = "text/html"
	JsonHeader        = "application/json"
	ContentType       = "Content-Type"
	MaxChirpLength    = 140
	UnknownErrMsg     = "Something went wrong"
	TooLongErrMsg     = "Chirp is too long"
	CensorSymbol      = "****"
)

var BadWords = []string{
	"kerfuffle",
	"sharbert",
	"fornax",
}
