package handlers

import (
	"github/h1mmeister/Go/pkg/render"
	"net/http"
)

// Home - home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About - about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
