package core

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type App struct {
	Router *chi.Mux
	Db     *gorm.DB
}
