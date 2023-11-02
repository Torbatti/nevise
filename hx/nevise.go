package hx

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/torbatti/nevise/models"
	"github.com/torbatti/nevise/routes"
)

func Nevise(w http.ResponseWriter, r *http.Request) {
	// cookie, err := r.Cookie("jwt")
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(cookie.Value)

	_, claims, _ := jwtauth.FromContext(r.Context())
	// log.Println(claims["user_id"])

	var nevise models.Nevise
	user_id := claims["user_id"]
	nevise_title := r.FormValue("nevise_title")
	nevise_text := r.FormValue("nevise_text")

	saveParam := chi.URLParam(r, "IsNewOrSave")

	data := struct {
		HaveMsg bool
		Msg     string
	}{}

	if len(nevise_title) < 10 || len(nevise_text) < 10 {
		data.Msg = "short"
		data.HaveMsg = true
	} else if len(nevise_title) > 10 || len(nevise_text) > 10 {
		if saveParam == "new" {
			nevise = models.Nevise{
				Title:     nevise_title,
				Text:      nevise_text,
				IsSave:    false,
				UserRefer: uint(user_id.(float64)),
			}
			data.Msg = "new"
			data.HaveMsg = true
		} else if saveParam == "save" {
			nevise = models.Nevise{
				Title:     nevise_title,
				Text:      nevise_text,
				IsSave:    true,
				UserRefer: uint(user_id.(float64)),
			}
			data.Msg = "save"
			data.HaveMsg = true
		}

		App.Db.Create(&nevise)
	}

	// Making Template
	funcs := template.FuncMap{"contains": strings.Contains}
	layout_name := "new"

	tmpl, err := template.New("").Funcs(funcs).ParseFiles(
		// Components
		"views/components/pages/new/m-new.html",
	)
	routes.Check("Template Parsing Error: %v", err)

	err = tmpl.ExecuteTemplate(w, layout_name, data)
	routes.Check("Template Execution Error: %v", err)
}

func SpawnNevise(w http.ResponseWriter, r *http.Request) {

	// Making Template
	funcs := template.FuncMap{"contains": strings.Contains}
	layout_name := "new"
	data := struct {
		HaveMsg bool
		Msg     string
	}{
		HaveMsg: false,
		Msg:     "",
	}
	tmpl, err := template.New("").Funcs(funcs).ParseFiles(
		// Components
		"views/components/pages/new/m-new.html",
	)
	routes.Check("Template Parsing Error: %v", err)

	err = tmpl.ExecuteTemplate(w, layout_name, data)
	routes.Check("Template Execution Error: %v", err)
}
