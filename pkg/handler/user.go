package handler

import (
	"PetBook/pkg/models"
	"PetBook/pkg/controllers"
	"PetBook/pkg/tokens"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"regexp"
	"time"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	//"github.com/jmoiron/sqlx"
	//"github.com/subosito/gotenv"
)

type Controller struct {
	UserStore controllers.UserStorer
	PetStore  controllers.PetStorer
}

// func (c Controller) LoginHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		user := &models.User{
// 			Email:     "asdsad@gmail.com",
// 			Login:     "mylogin",
// 			Password:  "123123",
// 			Firstname: "name",
// 			Lastname:  "surname",
// 		}
// 		err := c.UserStore.Login(user)
// 		if err != nil {
// 			log.Println(err)
// 			//json.NewEncoder(w).Encode("There is no such user!")
// 			//w.WriteHeader(http.StatusNotFound)
// 		}
// 		log.Println(user)
// 		//json.NewEncoder(w).Encode(user)
// 	}
// }

const (
	patternEmail    = `^\w+@\w+\.\w+$`
	patternPassword = `^.{6,}$`
	patternAnyChar  = `.*\S.*`
)

func (c *Controller) LoginGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokens.GenerateHTML(w, nil, "login")
	}
}
func (c *Controller) LoginPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")

		user := models.User{
			Email:    email,
			Password: password,
		}

		if err := c.UserStore.Login(&user); err != nil {
			//utils.Error(err)
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		expirationTime := time.Now().Add(30 * time.Minute)

		claims := &tokens.Claims{
			Email: email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

		tokenString, err := token.SignedString(tokens.SignKey)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while signing the token")
			log.Printf("Error signing token: %v\n", err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		_, err = c.UserStore.GetPet(&user)
		if err != nil {
			http.Redirect(w, r, "/petcabinet", http.StatusFound)
			return
		}
		http.Redirect(w, r, "/mypage", http.StatusFound)
	}
}

// TODO: reduce repeating code?
// TODO: do not log duplicate value error
func (c *Controller) RegisterGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokens.GenerateHTML(w, nil, "register")
	}
}
func (c *Controller) RegisterPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		login := r.FormValue("login")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirmPassword")
		//userType := r.FormValue("userType")
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")

		if matched, err := regexp.Match(patternAnyChar, []byte(login)); !matched || err != nil {
			if err != nil {
				tokens.Error(err, "Error occured while trying to match login.\n")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		if matched, err := regexp.Match(patternEmail, []byte(email)); !matched || err != nil {
			if err != nil {
				tokens.Error(err, "Error occured while trying to match email.\n")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		if matched, err := regexp.Match(patternAnyChar, []byte(firstName)); !matched || err != nil {
			if err != nil {
				tokens.Error(err, "Error occured while trying to match first name.\n")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		if matched, err := regexp.Match(patternAnyChar, []byte(lastName)); !matched || err != nil {
			if err != nil {
				tokens.Error(err, "Error occured while trying to match last name.\n")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		if matched, err := regexp.Match(patternPassword, []byte(password)); !matched || err != nil {
			if err != nil {
				tokens.Error(err, "Error occured while trying to match password.\n")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		if password != confirmPassword {
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		user := models.User{
			Email:     email,
			Login:     login,
			Firstname: firstName,
			Lastname:  lastName,
			Password:  password,
		}
		if err := c.UserStore.Register(&user); err != nil {
			tokens.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
