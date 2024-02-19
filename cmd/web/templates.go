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
	UserInfo       *models.User
	CommentsInfo   []*models.CommentInfo
	Categores      []*models.Category
	PostsInfo      []*models.PostInfo
	Disconnected   bool
	BadRequestForm bool
}

func cachingTemplate() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	var tf *template.Template
	for _, page := range pages {
		name := filepath.Base(page)
		if name != "login.tmpl" && name != "register.tmpl" && name != "error.tmpl" {
			tf, err = template.ParseFiles(
				"./ui/html/base.tmpl",
				"./ui/html/portions/header.tmpl",
				page,
			)
		} else {
			tf, err = template.ParseFiles(
				"./ui/html/baseLogRegis.tmpl",
				"./ui/html/portions/header.tmpl",
				page,
			)
		}
		if err != nil {
			return nil, err
		}
		cache[name] = tf
	}

	return cache, nil
}

func (app *application) render(w http.ResponseWriter, r *http.Request, layout, page string, data *TemplateData) {
	var err error
	extention := ".tmpl"
	tmpl, exist := app.cacheTemplate[page+extention]
	if !exist {
		err := errors.New("This template not found " + page)
		app.serverError(w, err)
		return
	}
	w.WriteHeader(200)

	//Comme que tout le temps on a le meme user
	if layout == "base" {
		userId, err := app.validSession(r)
		if err == nil {
			userInfo, err := app.connDB.GetUser(userId)
			if err != nil {
				app.serverError(w, err)
				return
			}
			userInfo.LikeCounter, err = app.connDB.GetLikeNumberByUser(userId)
			if err != nil {
				app.serverError(w, err)
				return
			}
			userInfo.CommentCounter, err = app.connDB.GetCommentNumberByUser(userId)
			if err != nil {
				app.serverError(w, err)
				return
			}

			data.UserInfo = userInfo
		}
	}
	err = tmpl.ExecuteTemplate(w, layout, data)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
