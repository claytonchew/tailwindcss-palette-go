package version

// Version information
var (
	Version    = "0.2.0"
	CommitHash = "unknown"
	BuildDate  = "unknown"
)

func Info() string {
	return Version + " (" + CommitHash + ") built on " + BuildDate
}
