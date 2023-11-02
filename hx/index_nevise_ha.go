package hx

import (
	"html/template"
	"math/rand"
	"net/http"

	"github.com/torbatti/nevise/models"
	"github.com/torbatti/nevise/routes"
)

func IndexNeviseHa(w http.ResponseWriter, r *http.Request) {
	rand := rand.Intn(12)

	var nevise models.Nevise
	var user models.User
	App.Db.Where("id = ?", rand+1).Find(&nevise)
	App.Db.Where("id = ?", nevise.UserRefer).Find(&user)

	// log.Println(rand + 1)
	// log.Println(nevise.Text)
	// log.Println(nevise.Title)

	funcs := template.FuncMap{}
	layout_name := "nevise"
	data := struct {
		UserName string
		Title    string
		Text     string
	}{
		UserName: user.UserName,
		Title:    nevise.Title,
		Text:     nevise.Text,
	}
	tmpl, err := template.New("").Funcs(funcs).ParseFiles("views/components/pages/nevise/base.html")
	routes.Check("Template Parsing Error: %v", err)

	err = tmpl.ExecuteTemplate(w, layout_name, data)
	routes.Check("Template Execution Error: %v", err)
}
