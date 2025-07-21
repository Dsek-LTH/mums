package config

import "time"

const (
	SessionTokenSize       = 256
	SessionExpirationTime  = 7 * 24 * time.Hour
	SessionCookieName      = "sessionToken"
	SessionCleanupInterval = 10 * time.Minute
)
