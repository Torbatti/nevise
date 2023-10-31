package routes

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var funcs = template.FuncMap{}

	// Data
	// data := struct {
	// 	Title string
	// }{
	// 	Title: "My Page",
	// }
	data := struct {
		Title       string
		KeyWords    string
		Description string
	}{
		Title:       "",
		KeyWords:    "",
		Description: "",
	}

	TemplateMaker(w, &funcs, "views/pages/index.html", data)
}
