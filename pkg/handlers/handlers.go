package handlers

import (
	"net/http"

	"github.com/AnggunPermata/Simple-Web-with-Golang/pkg/render"
)

//This is a homepage
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home-page.html")
}

//This is an about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about-page.html")
}
