package auth

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/config"
	"github.com/Dsek-LTH/mums/pkg/token"
)

type Session struct {
    UserAccountID int64
    ExpiresAt time.Time
}

func NewSession(userAccountID int64) *Session {
	return &Session{
		userAccountID,
		time.Now().Add(config.SessionExpirationTime),
	}
}

func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}

type SessionStore struct {
	sync.RWMutex
	sessions map[string]*Session  // sessionID -> userAccountID
}

func NewSessionStore() *SessionStore {
	return &SessionStore{sessions: make(map[string]*Session)}
}

func (s *SessionStore)  Create(userAccountID int64) string {
	sessionID, _ := token.GenerateSecure(config.SessionIDLength)
	session := NewSession(userAccountID)

	s.Lock()
	s.sessions[sessionID] = session
	s.Unlock()
	
	return sessionID
}

func (s *SessionStore) Get(sessionID string) (*Session, bool) {
    s.RLock()
    session, ok := s.sessions[sessionID]
    s.RUnlock()

    return session, ok
}

func (s *SessionStore) Delete(sessionID string) {
    s.Lock()
    delete(s.sessions, sessionID)
    s.Unlock()
}

func loginRedirect(c echo.Context) error {
    return c.Redirect(http.StatusSeeOther, "/login")
}

func SessionMiddleware(sessionStore *SessionStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie(config.SessionCookieName)
			if err != nil {
				return loginRedirect(c)	
			}
			sessionID := cookie.Value

			session, ok := sessionStore.Get(sessionID)
			if !ok  {
				return loginRedirect(c)	
			}

			if session.IsExpired() {
				sessionStore.Delete(sessionID)
				loginRedirect(c)	
			}

			c.Set("userAccountID", session.UserAccountID)

			return next(c)
		}
	}
}

func GetUserAccountID(c echo.Context) (int64, bool) {
	userAccountID, ok := c.Get("userAccountID").(int64)
	return userAccountID, ok
}

