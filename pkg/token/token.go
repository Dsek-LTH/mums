package token

import (
	"crypto/rand"
	"encoding/base64"
)

func MustGenerateSecure(size int) string {
	token := make([]byte, size)
	if _, err := rand.Read(token); err != nil {
		panic("failed to generate secure random token: " + err.Error())
	}
	return base64.URLEncoding.EncodeToString(token)
}
