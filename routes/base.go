package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/torbatti/nevise/core"
)

var App *core.App

type TmplConf struct {
	Wr         http.ResponseWriter
	Layout     string
	Pages      []string
	Components []string

	Data  interface{} // TODO: FIND BETTTER WAY THAN USING INTERFACES
	Funcs *template.FuncMap
}

func Check(detail string, err error) {
	if err != nil {
		log.Println(detail, " ", err)
	}
}

// TODO: ...string and []string problem in parsefiles
func TemplateMaker(w http.ResponseWriter, fmap *template.FuncMap, path string, data interface{}) {
	// var funcs = template.FuncMap{"join": strings.Join}

	tmpl, err := template.New("").Funcs(*fmap).ParseFiles(path, "views/layouts/base.html")
	// tmpl, err := template.New("views/layouts/base.html").ParseFiles(path)
	Check("Template Parsing Error: %v", err)

	err = tmpl.ExecuteTemplate(w, "base", data)
	Check("Template Execution Error: %v", err)
}
