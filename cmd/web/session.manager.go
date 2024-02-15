package main

import (
	"net/http"
	"time"
)

var SESSION map[string]int

func isConnected(r *http.Request) (int) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return -1
	}
	if cookie.Value == ""{
		return -1
	}
	userId, exist := SESSION[cookie.Value]
	if !exist {
		return -1
	}
	// verify if the user didn't logged in 
	// on another browser
	for key, sess := range SESSION{
		if sess == userId && key != cookie.Value{
			delete(SESSION,key)
		}
	}
	// verify if the cookie didn't expire
	if cookie.Expires.Before(time.Now()){
		delete(SESSION,cookie.Value)
		return -1
	}
	return userId
}
