package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//Setting port
const port = ":8080"

//This is a homepage
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home-page.html")
}

//This is an about page
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about-page.html")
}

//This is renderTemplate's function
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template")
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(port, nil)
}
