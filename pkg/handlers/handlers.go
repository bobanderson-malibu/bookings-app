package handlers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.renderTemplate(w, "about.page.html")
}
