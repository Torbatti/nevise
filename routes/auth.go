package routes

import (
	"html/template"
	"net/http"
)

func MSignup(w http.ResponseWriter, r *http.Request) {
	var funcs = template.FuncMap{}

	// Making Template
	layout_name := "base"
	data := struct {
		Title       string
		KeyWords    string
		Description string
	}{
		Title:       "ثبت نام",
		KeyWords:    "",
		Description: "",
	}
	tmpl, err := template.New("").Funcs(funcs).ParseFiles(
		"views/pages/auth/m-signup.html",
		"views/layouts/base.html")
	Check("Template Parsing Error: %v", err)

	err = tmpl.ExecuteTemplate(w, layout_name, data)
	Check("Template Execution Error: %v", err)
}

func Signup(w http.ResponseWriter, r *http.Request) {

}

func MLogin(w http.ResponseWriter, r *http.Request) {

}
func Login(w http.ResponseWriter, r *http.Request) {

}
