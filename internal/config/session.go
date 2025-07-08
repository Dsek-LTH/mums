package config

import "time"

const (
	SessionIDLength       = 128
	SessionExpirationTime = 24 * time.Hour
	SessionCookieName     = "session_id"
)
