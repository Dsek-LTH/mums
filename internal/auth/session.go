package auth

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/config"
	"github.com/Dsek-LTH/mums/pkg/token"
)

type session struct {
	userAccountID int64
	expiresAt     time.Time
}

func newSession(userAccountID int64) *session {
	s := &session{userAccountID: userAccountID}
	s.touch()
	return s
}

func (s *session) isExpired() bool {
	return s.expiresAt.Before(time.Now())
}

func (s *session) touch() {
	s.expiresAt = time.Now().Add(config.SessionExpirationTime)
}

type SessionStore struct {
	sync.RWMutex
	sessions map[string]*session // sessionID -> userAccountID
}

func (ss *SessionStore) CleanupExpiredSessions() {
	ticker := time.NewTicker(config.SessionCleanupInterval)
	defer ticker.Stop()

	for range ticker.C {
		ss.Lock()
		for sid, s := range ss.sessions {
			if s.isExpired() {
				delete(ss.sessions, sid)
			}
		}
	}
}

func NewSessionStore() *SessionStore {
	ss := &SessionStore{sessions: make(map[string]*session)}

	go ss.CleanupExpiredSessions()

	return ss
}

func (ss *SessionStore) Create(userAccountID int64) string {
	// token.GenerateSecure is guaranteed not to return an error
	sid := token.MustGenerateSecure(config.SessionIDLength)
	s := newSession(userAccountID)

	ss.Lock()
	ss.sessions[sid] = s
	ss.Unlock()

	return sid
}

func (ss *SessionStore) Get(sid string) (*session, bool) {
	ss.RLock()
	s, ok := ss.sessions[sid]
	ss.RUnlock()

	return s, ok
}

func (ss *SessionStore) Delete(sid string) {
	ss.Lock()
	delete(ss.sessions, sid)
	ss.Unlock()
}

func SessionMiddleware(ss *SessionStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			setNotLoggedIn := func() error {
			    c.Set(config.CTXKeyIsLoggedIn, false)
				return next(c)
			}

			cookie, err := c.Cookie(config.SessionCookieName)
			if err != nil {
				return setNotLoggedIn()
			}
			sid := cookie.Value

			s, ok := ss.Get(sid)
			if !ok {
				return setNotLoggedIn()
			}

			if s.isExpired() {
				ss.Delete(sid)
				return setNotLoggedIn()
			}
			s.touch()

			c.Set(config.CTXKeySessionID, sid)
			c.Set(config.CTXKeyIsLoggedIn, true)
			c.Set(config.CTXKeyUserAccountID, s.userAccountID)

			return next(c)
		}
	}
}

func GetSessionID(c echo.Context) string {
	sid, ok := c.Get(config.CTXKeySessionID).(string)
	if !ok {
		panic("config.CTXKeyIsLoggedIn is not set in context, was SessionMiddleware not applied?")
	}

	return sid
}

func GetIsLoggedIn(c echo.Context) bool {
	isLoggedIn, ok := c.Get(config.CTXKeyIsLoggedIn).(bool)
	if !ok {
		panic("config.CTXKeyIsLoggedIn is not set in context, was SessionMiddleware not applied?")
	}

	return isLoggedIn
}

func GetUserAccountID(c echo.Context) int64 {
	userAccountID, ok := c.Get(config.CTXKeyUserAccountID).(int64)
	if !ok {
		panic("config.CTXKeyUserAccountID is not set in context, was SessionMiddleware not applied?")
	}

	return userAccountID
}

func RequireSession() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if GetIsLoggedIn(c) {
				return next(c)
			}

			return c.Redirect(http.StatusSeeOther, "/login")
		}
	}
}
