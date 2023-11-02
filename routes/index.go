package routes

import (
	"html/template"
	"log"
	"net/http"
)

func MIndex(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Cookie("jwt"))

	// Making Template
	funcs := template.FuncMap{}
	layout_name := "base"
	data := struct {
		Title       string
		KeyWords    string
		Description string
	}{
		Title:       "نویس",
		KeyWords:    "",
		Description: "",
	}
	tmpl, err := template.New("").Funcs(funcs).ParseFiles(
		// Layouts
		"views/layouts/m-base.html",
		// Pages
		"views/pages/m-index.html",
		// Components
		"views/components/pages/base/m-top.html",
		"views/components/pages/base/m-bottom.html",
		"views/components/pages/base/m-new.html",
	)
	Check("Template Parsing Error: %v", err)

	err = tmpl.ExecuteTemplate(w, layout_name, data)
	Check("Template Execution Error: %v", err)
}

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
