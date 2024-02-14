package main

import (
	"errors"
	"net/http"
	"path/filepath"
	"text/template"

	"forum.01/internal/models"
)

type TemplateData struct {
	PostInfo       *models.PostInfo
	CommentsInfo   []*models.CommentInfo
	Categores      []*models.Category
	PostsInfo      []*models.PostInfo
	BadRequestForm bool
}

func cachingTemplate() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		tf, err := template.ParseFiles(
			"./ui/html/base.tmpl",
			"./ui/html/portions/header.tmpl",
			page,
		)
		if err != nil {
			return nil, err
		}
		name := filepath.Base(page)
		cache[name] = tf
	}

	return cache, nil
}

func (app *application) render(w http.ResponseWriter, page string, data *TemplateData) {
	extention := ".tmpl"
	tmpl, exist := app.cacheTemplate[page+extention]
	if !exist {
		err := errors.New("This template not found " + page)
		app.serverError(w, err)
		return
	}
	w.WriteHeader(200)
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
