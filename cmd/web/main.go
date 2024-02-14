package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"forum.01/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	connDB        *models.ConnDB
	cacheTemplate map[string]*template.Template
}

func init() {
	initError := log.New(os.Stderr, "ERROR INIT\t", log.Lmicroseconds)
	dbFiles, err := filepath.Glob("./db/*.db")
	if err != nil {
		log.Fatalln(initError)
		return
	}
	if len(dbFiles) == 1 {
		os.RemoveAll(dbFiles[0])
	}
}

func main() {
	PORT := flag.String("addr", ":4000", "enter port")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Lmicroseconds)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Lmicroseconds|log.Lshortfile)

	db, err := openDB("./db/DB_forum.db")
	if err != nil {
		errorLog.Fatalln("Open DB error\t", err.Error())
		return
	}

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		connDB:        &models.ConnDB{DB: db},
	}

	err = app.createTable()
	if err != nil {
		errorLog.Fatal("Sql migration error --->", err.Error())
		return
	}
	infoLog.Println("<<Successfully>> Creation Table")
	
	err = app.setData()
	if err != nil {
		errorLog.Fatal("Sql set data error --->", err.Error())
		return
	}
	infoLog.Println("<<Successfully>> Set Data")
	
	cache, err := cachingTemplate()
	if err != nil {
		errorLog.Fatalln("caching Template error\t", err.Error())
	}
	app.cacheTemplate = cache

	srv := &http.Server{
		Addr:     *PORT,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}
	
	infoLog.Printf("server on http://localhost%s", *PORT)
	err = srv.ListenAndServe()
	errorLog.Fatalln(err.Error())
}
