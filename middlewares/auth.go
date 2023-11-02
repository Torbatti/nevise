package middlewares

import (
	"fmt"

	"github.com/go-chi/jwtauth/v5"
)

var JWT_SECRET string
var TokenAuth *jwtauth.JWTAuth

func Register() {

	_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{"user_id": 123})

	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}
