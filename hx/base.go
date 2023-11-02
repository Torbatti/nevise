package hx

import (
	"html/template"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/torbatti/nevise/core"
	"github.com/torbatti/nevise/routes"
)

var App *core.App
var JWT_SECRET string
var AUTH_SECRET string
var TokenAuth *jwtauth.JWTAuth

func Empty(w http.ResponseWriter, r *http.Request) {
	funcs := template.FuncMap{}
	layout_name := "empty"
	data := struct{}{}
	tmpl, err := template.New("").Funcs(funcs).ParseFiles("views/components/pages/base/empty.html")
	routes.Check("Template Parsing Error: %v", err)

	err = tmpl.ExecuteTemplate(w, layout_name, data)
	routes.Check("Template Execution Error: %v", err)
}
