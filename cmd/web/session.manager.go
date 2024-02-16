package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) validSession(r *http.Request) (int, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, err
	}
	if cookie.Value == "" {
		return 0, fmt.Errorf("validSession() : cookie don't contain value")
	}
	userId, exist := app.Session[cookie.Value]
	if !exist {
		return 0, fmt.Errorf("validSession() :cookie value not found in session map")
	}

	// verify if the user didn't logged in
	// on another browser
	for key, sess := range app.Session {
		if sess == userId && key != cookie.Value {
			delete(app.Session, key)
		}
	}
	// verify if the cookie didn't expire
	if cookie.Expires.After(time.Now()) {
		delete(app.Session, cookie.Value)
		return 0, fmt.Errorf("cookie expire")
	}
	return userId, nil
}
