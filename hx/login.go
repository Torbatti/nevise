package hx

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/torbatti/nevise/models"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Get Username and Email Errors
	email_exist_err := App.Db.Where("email = ?", r.FormValue("email")).First(&user).Error
	// Validate Username and Email Errors
	if email_exist_err != gorm.ErrRecordNotFound {
		email := r.FormValue("email")
		password := r.FormValue("password")

		hash := sha256.New()
		hash.Write([]byte(email + AUTH_SECRET + password))
		hashString := fmt.Sprintf("%x", hash.Sum(nil))

		if user.Hash == hashString {
			log.Println("User Successfully Logged In")
			_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{"user_id": user.ID})

			cookie := http.Cookie{
				Name:     "jwt",
				Value:    tokenString,
				Path:     "/",
				MaxAge:   3600,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			}
			http.SetCookie(w, &cookie)

			platformParam := chi.URLParam(r, "platform")
			if platformParam == "m" {
				w.Header().Add("HX-REDIRECT", "/m-index")
			} else {
				w.Header().Add("HX-REDIRECT", "/")
			}
		}

	} else {
		if email_exist_err == gorm.ErrRecordNotFound {
			log.Println("Email is not a valid value !")
		} else {
			log.Println("SOMETHING WENT WRONG and IDK WHATITIS !")
		}
	}

}
