package hx

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/torbatti/nevise/core"
)

var App *core.App
var JWT_SECRET string
var AUTH_SECRET string
var TokenAuth *jwtauth.JWTAuth
