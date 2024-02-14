package main

import (
	"fmt"
	"net/http"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s - %s - %s", r.RemoteAddr,  r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}
func (app *application) panicRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			panic := recover()
			if panic != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", panic))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

/* func (app *application) choiceHandler(nextA, nextB http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet :
			nextA(w, r)
		case http.MethodPost :
			nextB(w, r)
		default :
			app.serverError(w, fmt.Errorf(http.StatusText(405)))
		}
	})
} */
