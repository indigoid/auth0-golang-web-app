package user

import (
	"github.com/indigoid/auth0-golang-web-app/app"
	templates "github.com/indigoid/auth0-golang-web-app/routes"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.RenderTemplate(w, "user", session.Values["profile"])
}
