package constants

import "time"

const (
	DefaultHTTPWriteTimeout = 2 * time.Second
	DefaultHTTPReadTimeout  = 2 * time.Second

	DefaultSessionExpiration = 24 * time.Hour
	DefaultTTLUserBalance    = 30 * time.Minute
)
