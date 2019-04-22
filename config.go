package main

import (
	"github.com/gorilla/sessions"
)

// SessionStore blabla
var SessionsStore = sessions.NewCookieStore([]byte("secret"))

func InitConfig() {
	SessionsStore.Options = &sessions.Options{
		// Domain:   "localhost",
		// Path:     "*",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}
}
