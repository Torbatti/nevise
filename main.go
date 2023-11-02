package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"

	"github.com/torbatti/nevise/apis"
	"github.com/torbatti/nevise/core"
	"github.com/torbatti/nevise/hx"
	"github.com/torbatti/nevise/middlewares"
	"github.com/torbatti/nevise/models"
	"github.com/torbatti/nevise/routes"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var app *core.App

var PortString string
var AUTH_SECRET string
var JWT_SECRET string

var tokenAuth *jwtauth.JWTAuth

func init() {
	// Loading .env Variables
	godotenv.Load(".env") //os.Setenv(port, "8000")

	// Getting .env Variables
	PortString = os.Getenv("PORT")
	AUTH_SECRET = os.Getenv("AUTH_SECRET")
	JWT_SECRET = os.Getenv("JWT_SECRET")

	middlewares.JWT_SECRET = JWT_SECRET

	// TOKEN AUTH
	tokenAuth = jwtauth.New("HS256", []byte(JWT_SECRET), nil)

	middlewares.TokenAuth = tokenAuth
	hx.TokenAuth = tokenAuth
}

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
	db.AutoMigrate(&models.User{}, &models.Nazar{}, &models.Nevise{})

	// Connecting
	app.Router = chi.NewRouter()
	app.Db = db

	return app
}

func main() {
	app := makeApp()
	// middlewares.Register()

	// Root
	root := chi.NewRouter()

	// Public Routes
	root.Group(func(r chi.Router) {
		middlewares.Cors(app.Router)

		r.Get("/m-signup", routes.MSignup)
		r.Get("/m-login", routes.MLogin)
		r.Post("/hx/signup/{platform}", hx.Signup)
		r.Post("/hx/login/{platform}", hx.Login)

		// Public
		public := http.FileServer(http.Dir("./public"))
		r.Mount("/", public)
	})

	// Private Routes
	root.Group(func(r chi.Router) {
		middlewares.Cors(app.Router)

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		//Views
		r.Get("/", routes.Index)
		r.Get("/m-index", routes.MIndex)
		r.Post("/hx/nevise/{IsNewOrSave}", hx.Nevise)
		r.Post("/hx/spawn_nevise", hx.SpawnNevise)
		r.Post("/hx/empty", hx.Empty)
		r.Post("/hx/index_nevise_ha", hx.IndexNeviseHa)

	})
	app.Router.Mount("/", root)

	// Initial server
	server := &http.Server{
		Handler: app.Router,
		Addr:    ":" + PortString,
	}

	// Listen And Serve
	log.Println("Listening On " + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
