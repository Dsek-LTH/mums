package config

import "time"

const (
	SessionIDLength        = 128
	SessionExpirationTime  = 7 * 24 * time.Hour
	SessionCookieName      = "session_id"
	SessionCleanupInterval = 10 * time.Minute
)
