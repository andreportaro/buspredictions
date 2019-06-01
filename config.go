package main

import (
	"github.com/gorilla/sessions"
)

// SessionsStore creates a new SessionStore
var SessionsStore = sessions.NewCookieStore([]byte("secret"))

// InitConfig Starts basic session with expiry configuration
func InitConfig() {
	SessionsStore.Options = &sessions.Options{
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}
}
