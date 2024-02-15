package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"forum.01/internal/filters"
	"forum.01/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		postsInfo, err := app.connDB.GetAllPostInfo()
		if err != nil {
			app.serverError(w, err)
			return
		}
		categories, err := app.connDB.GetAllCategory()
		if err != nil {
			app.serverError(w, err)
			return
		}
		badRequest := false

		if r.Form.Has("filter") {
			filterCheck := r.Form["filterCheck"]
			if len(filterCheck) > 0 {
				for _, fc := range filterCheck {
					if fc == "Liked-Post" {
						postsInfo = filters.LikedPostFilter(postsInfo)
					}
					if fc == "Created-Post" {
						// Je mets actuel user en attendant de regler les sessions
						postsInfo = filters.CreatedPostFilter(postsInfo, 3)
					}
					if fc != "Created-Post" && fc != "Liked-Post" {
						app.clientError(w, http.StatusBadRequest)
						return
					}
				}

			}
			categoriesCheck := r.Form["filterCategoryCheck"]
			if len(categoriesCheck) > 0 && len(categoriesCheck) <= len(categories) {
				postsInfo = filters.CategoryFilter(postsInfo, categoriesCheck...)
			} else if len(categoriesCheck) > len(categories) {
				app.clientError(w, http.StatusBadRequest)
				return
			}

			if len(postsInfo) == 0 {
				badRequest = true
			}

		}

		data := &TemplateData{Categores: categories, PostsInfo: postsInfo, BadRequestForm: badRequest}

		app.render(w, "base", "home", data)

	case http.MethodPost:
		postId := r.PostForm.Get("postId")
		pId, err := strconv.Atoi(postId)
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}
		if r.PostForm.Has("like") {
			liked := r.PostForm.Get("like")
			l, err := strconv.ParseBool(liked)
			if err != nil {
				app.clientError(w, http.StatusBadRequest)
				return
			}

			_, err = app.connDB.SetLike(3, pId, l)
			if err != nil {
				app.serverError(w, err)
			}
		}
		if r.PostForm.Has("dislike") {
			disliked := r.PostForm.Get("dislike")
			dl, err := strconv.ParseBool(disliked)
			if err != nil {
				app.clientError(w, http.StatusBadRequest)
				return
			}

			_, err = app.connDB.SetDislike(3, pId, dl)
			if err != nil {
				app.serverError(w, err)
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		categories, err := app.connDB.GetAllCategory()
		if err != nil {
			app.serverError(w, err)
			return
		}
		bad := r.URL.Query().Has("bad")
		data := &TemplateData{Categores: categories, BadRequestForm: bad}

		app.render(w, "base", "form", data)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}

		// faire une fonction pour la logique de validation des donnees
		// la fonction retourne une boolean si donne bonne ou pas
		categoryIds := r.Form["categorCheck"]
		content := r.PostForm.Get("content")
		userId, err := strconv.Atoi(r.PostForm.Get("userId"))
		if len(categoryIds) == 0 || content == "" || err != nil {
			// app.clientError(w, http.StatusBadRequest)
			http.Redirect(w, r, "/create?bad", http.StatusSeeOther)
			return
		}
		lastPostId, err := app.connDB.SetPost(content, userId)
		if err != nil {
			app.serverError(w, err)
			return
		}
		for _, categoryId := range categoryIds {
			cId, err := strconv.Atoi(strings.TrimSpace(categoryId))
			if err != nil {
				// app.clientError(w, http.StatusBadRequest)
				http.Redirect(w, r, "/create?bad", http.StatusSeeOther)
				return
			}
			_, err = app.connDB.SetPostCategory(lastPostId, cId)
			if err != nil {
				app.serverError(w, err)
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) comment(w http.ResponseWriter, r *http.Request) {
	// Recuperation de l'id dans l'url
	idPostUrlVal := r.URL.Query()
	if len(idPostUrlVal) != 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	var (
		pId int
		err error
	)
	for key := range idPostUrlVal {
		pId, err = strconv.Atoi(key)
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}
	}

	switch r.Method {
	case http.MethodGet:
		postInfo, err := app.connDB.GetPostInfo(pId)
		if err != nil {
			app.serverError(w, err)
			return
		}

		commentsInfo, err := app.connDB.GetCommentsInfoByPost(pId)
		if err != nil {
			app.serverError(w, err)
			return
		}

		data := &TemplateData{PostInfo: postInfo, CommentsInfo: commentsInfo}
		app.render(w, "base", "comment", data)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
		}

		if r.PostForm.Has("like") {
			liked := r.PostForm.Get("like")
			l, err := strconv.ParseBool(liked)
			if err != nil {
				app.clientError(w, http.StatusBadRequest)
				return
			}

			_, err = app.connDB.SetLike(3, pId, l)
			if err != nil {
				app.serverError(w, err)
			}
		}
		if r.PostForm.Has("dislike") {
			disliked := r.PostForm.Get("dislike")
			dl, err := strconv.ParseBool(disliked)
			if err != nil {
				app.clientError(w, http.StatusBadRequest)
				return
			}

			_, err = app.connDB.SetDislike(3, pId, dl)
			if err != nil {
				app.serverError(w, err)
			}
		}
		if r.PostForm.Has("send-comment") {
			comment := r.PostForm.Get("comment")
			app.infoLog.Println(comment)
			if len(comment) > 0 {
				// je mets le user id 3 en attendant de regler les connexions
				_, err = app.connDB.SetComment(comment, pId, 3)
				if err != nil {
					app.serverError(w, err)
					return
				}
			}
		}
		http.Redirect(w, r, fmt.Sprintf("/comment?%d", pId), http.StatusSeeOther)

	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bad := r.URL.Query().Has("bad")
		data := &TemplateData{BadRequestForm: bad}

		app.render(w, "baseLogRegis", "login", data)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}
		email := r.PostForm.Get("email")
		password := r.PostForm.Get("password")
		if !utils.EmailValidation(email) || !utils.PasswordValidation(password) {
			http.Redirect(w, r, "/login?bad", http.StatusSeeOther)
			return
		}

		user, err := app.connDB.GetUserByMail(email)
		if err != nil {
			http.Redirect(w, r, "/login?bad", http.StatusSeeOther)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			http.Redirect(w, r, "/login?bad", http.StatusSeeOther)
			return
		}
		/*
			Logique de creation de session ici
		*/
		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		app.clientError(w, http.StatusBadRequest)
		return
	}
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bad := r.URL.Query().Has("bad")
		data := &TemplateData{BadRequestForm: bad}
		app.render(w, "baseLogRegis", "register", data)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}

		username := r.PostForm.Get("username")
		email := r.PostForm.Get("email")
		password := r.PostForm.Get("password")
		if !utils.UsernameValidation(username) || !utils.EmailValidation(email) || !utils.PasswordValidation(password) {
			http.Redirect(w, r, "/register?bad", http.StatusSeeOther)
			return
		}

		encryptPass, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			http.Redirect(w, r, "/register?bad", http.StatusSeeOther)
			return
		}
		password = string(encryptPass)
		app.connDB.SetUser(username, email, password)
		/*
			Logique de creation de session ici
		*/
		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		app.clientError(w, http.StatusBadRequest)
		return
	}
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
