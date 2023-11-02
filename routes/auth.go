package routes

import (
	"html/template"
	"net/http"
)

func MSignup(w http.ResponseWriter, r *http.Request) {

	// Making Template
	funcs := template.FuncMap{}
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
		// Layouts
		"views/layouts/m-base.html",
		// Pages
		"views/pages/auth/m-signup.html",
		// Components
		"views/components/pages/base/m-top.html",
		"views/components/pages/base/m-bottom.html",
		"views/components/pages/auth/signup.html",
	)
	Check("Template Parsing Error: %v", err)

	err = tmpl.ExecuteTemplate(w, layout_name, data)
	Check("Template Execution Error: %v", err)
}

func Signup(w http.ResponseWriter, r *http.Request) {

}

func MLogin(w http.ResponseWriter, r *http.Request) {
	// Making Template
	funcs := template.FuncMap{}
	layout_name := "base"
	data := struct {
		Title       string
		KeyWords    string
		Description string
	}{
		Title:       "ورود",
		KeyWords:    "",
		Description: "",
	}
	tmpl, err := template.New("").Funcs(funcs).ParseFiles(
		// Layouts
		"views/layouts/m-base.html",
		// Pages
		"views/pages/auth/m-login.html",
		// Components
		"views/components/pages/base/m-top.html",
		"views/components/pages/base/m-bottom.html",
		"views/components/pages/auth/login.html",
	)
	Check("Template Parsing Error: %v", err)

	err = tmpl.ExecuteTemplate(w, layout_name, data)
	Check("Template Execution Error: %v", err)
}
func Login(w http.ResponseWriter, r *http.Request) {

}
