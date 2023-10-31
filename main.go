package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/torbatti/nevise/apis"
	"github.com/torbatti/nevise/core"
	"github.com/torbatti/nevise/hx"
	"github.com/torbatti/nevise/middlewares"
	"github.com/torbatti/nevise/routes"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var app *core.App

func makeApp() *core.App {

	app = &core.App{}

	routes.App = app
	hx.App = app
	apis.App = app

	// Database: Opening
	db, err := gorm.Open(sqlite.Open("nevise.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database .\n ", err.Error())
		os.Exit(2)
	} // db.Logger = logger.Default.LogMode(logger.Info)

	// Database: Migrations
	db.AutoMigrate()

	// Connecting
	app.Router = chi.NewRouter()
	app.Db = db

	return app
}
func main() {
	app := makeApp()

	// Setting Up .env
	godotenv.Load(".env") //os.Setenv(port, "8000")
	portString := os.Getenv("PORT")

	// MiddleWares
	middlewares.Cors(app.Router)

	// Root
	root := chi.NewRouter()
	app.Router.Mount("/", root)

	// Routes
	root.Get("/", routes.Index)
	root.Get("/m-signup", routes.MSignup)

	// Public
	public := http.FileServer(http.Dir("./public"))
	root.Mount("/", public)

	// Initial server
	server := &http.Server{
		Handler: app.Router,
		Addr:    ":" + portString,
	}

	// Listen And Serve
	log.Println("Listening On " + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
