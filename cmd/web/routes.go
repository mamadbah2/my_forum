package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./ui/static"))
	router.Handle("/static/", http.StripPrefix("/static", fs))

	// Here all of routes
	router.HandleFunc("/", app.home)
	router.HandleFunc("/create", app.create)
	router.HandleFunc("/comment", app.comment)

	return app.panicRecover(app.logRequest(router))
}
