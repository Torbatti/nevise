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

func Signup(w http.ResponseWriter, r *http.Request) {

	// log.Println(r.FormValue("email"))
	// log.Println(r.FormValue("password"))
	// log.Println(r.FormValue("user_name"))

	var user models.User

	// Check If password is shorter than 5 characters
	if len(r.FormValue("user_name")) < 5 {
		fmt.Println("Username too short !")
		return
	}

	// Get Username and Email Errors
	username_exist_err := App.Db.Where("user_name = ?", r.FormValue("user_name")).First(&user).Error
	email_exist_err := App.Db.Where("email = ?", r.FormValue("email")).First(&user).Error
	// Validate Username and Email Errors
	if username_exist_err == gorm.ErrRecordNotFound && email_exist_err == gorm.ErrRecordNotFound {
		// if err := App.Db.Where("email = ?", r.FormValue("email")).First(&user).Error; err == nil {
		log.Println("New User Created Successfully !")

		user_name := r.FormValue("user_name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		hash := sha256.New()
		hash.Write([]byte(email + AUTH_SECRET + password))
		hashString := fmt.Sprintf("%x", hash.Sum(nil))
		println(string(hashString))

		user = models.User{
			UserName: user_name,
			Email:    email,
			Hash:     hashString,
		}
		App.Db.Create(&user)

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

	} else {
		if username_exist_err != gorm.ErrRecordNotFound {
			log.Println("Username has already been taken !")
		} else {
			log.Println("Email is not a valid value !")
		}
	}

}
