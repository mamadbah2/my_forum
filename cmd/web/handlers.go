package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forum.01/internal/filters"
	"forum.01/internal/utils"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w, r)
		return
	}
	// Verification de la session
	var disconnected bool
	actualUser, err := app.validSession(r)
	if err != nil {
		disconnected = true
	}

	err = r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		postsInfo, err := app.connDB.GetAllPostInfo(actualUser)
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
						postsInfo = filters.CreatedPostFilter(postsInfo, actualUser)
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

		data := &TemplateData{Categores: categories, PostsInfo: postsInfo, BadRequestForm: badRequest, Disconnected: disconnected}

		app.render(w, r, "base", "home", data)

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

			_, err = app.connDB.SetLike(actualUser, pId, l)
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

			_, err = app.connDB.SetDislike(actualUser, pId, dl)
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
	actualUser, err := app.validSession(r)
	if err != nil {
		http.Redirect(w, r, "/logout", http.StatusSeeOther)
		return
	}

	// Action selon la methode d'entrée
	switch r.Method {
	case http.MethodGet:
		categories, err := app.connDB.GetAllCategory()
		if err != nil {
			app.serverError(w, err)
			return
		}
		bad := r.URL.Query().Has("bad")
		data := &TemplateData{Categores: categories, BadRequestForm: bad}

		app.render(w, r, "base", "form", data)

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
		escapedContent := html.EscapeString(content)

		if len(categoryIds) == 0 || strings.TrimSpace(escapedContent) == "" {
			http.Redirect(w, r, "/create?bad", http.StatusSeeOther)
			return
		}
		lastPostId, err := app.connDB.SetPost(escapedContent, actualUser)
		if err != nil {
			app.serverError(w, err)
			return
		}

		for _, categoryId := range categoryIds {
			cId, err := strconv.Atoi(strings.TrimSpace(categoryId))
			if err != nil {
				app.clientError(w, http.StatusBadRequest)
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
	// Verification de la session
	actualUser, err := app.validSession(r)
	var disconnected bool
	if err != nil {
		disconnected = true
	}

	// Recuperation de l'id dans l'url
	idPostUrlVal := r.URL.Query()
	if len(idPostUrlVal) != 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	var (
		pId int
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
		postInfo, err := app.connDB.GetPostInfo(pId, actualUser)
		if err != nil {
			app.serverError(w, err)
			return
		}

		commentsInfo, err := app.connDB.GetCommentsInfoByPost(pId)
		if err != nil {
			app.serverError(w, err)
			return
		}

		data := &TemplateData{PostInfo: postInfo, CommentsInfo: commentsInfo, Disconnected: disconnected}
		app.render(w, r, "base", "comment", data)

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

			_, err = app.connDB.SetLike(actualUser, pId, l)
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

			_, err = app.connDB.SetDislike(actualUser, pId, dl)
			if err != nil {
				app.serverError(w, err)
			}
		}
		if r.PostForm.Has("send-comment") {
			comment := r.PostForm.Get("comment")
			escapedComment := html.EscapeString(comment)
			if len(escapedComment) > 0 {
				_, err = app.connDB.SetComment(escapedComment, pId, actualUser)
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

		app.render(w, r, "baseLogRegis", "login", data)

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
		u, err := uuid.NewV4()

		if err != nil {
			app.serverError(w, err)
			return
		}
		cookies := http.Cookie{
			Name:     "session_token",
			Value:    u.String(),
			Secure:   true,
			Expires:  time.Now().Add(60 * time.Minute),
			HttpOnly: true,
		}
		app.Session[u.String()] = user.User_id
		http.SetCookie(w, &cookies)

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
		app.render(w, r, "baseLogRegis", "register", data)

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
		userId, err := app.connDB.SetUser(username, email, password)
		if err != nil {
			if err.Error() == "UNIQUE constraint failed: User.username" {
				http.Redirect(w, r, "/register?bad", http.StatusSeeOther)
				return
			}
			app.serverError(w, err)
			return
		}
		/*
			Logique de creation de session ici
		*/
		u, err := uuid.NewV4()

		if err != nil {
			app.serverError(w, err)
			return
		}
		cookies := http.Cookie{
			Name:     "session_token",
			Value:    u.String(),
			Secure:   true,
			Expires:  time.Now().Add(60 * time.Minute),
			HttpOnly: true,
		}
		app.Session[u.String()] = userId

		http.SetCookie(w, &cookies)
		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		app.clientError(w, http.StatusBadRequest)
		return
	}
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err == nil {
		delete(app.Session, cookie.Value)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
