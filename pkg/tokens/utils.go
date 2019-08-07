package tokens

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"

	//"github.com/dgrijalva/jwt-go/request"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

var Logger *log.Logger

var (
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
)

func Error(args ...interface{}) {
	Logger.SetPrefix("ERROR ")
	Logger.Println(args...)
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("./web/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(writer, "layout", data)
	if err != nil {
		Error(err)
		return
	}
}

// TODO: implement context
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := c.Value
	claims := &Claims{}

	//r.Header.Set("Authorization", "Bearer " + tokenString)

	/*token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
	func(token *jwt.Token) (interface{}, error) {
		return VerifyKey, nil
	})
	*/

	//validate token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return VerifyKey, nil
	})

	if err == nil {
		if token.Valid {
			context.Set(r, "email", claims.Email)
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w, r, "/login", http.StatusFound)
			fmt.Fprint(w, "Token is not valid. Need to refresh.")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		http.Redirect(w, r, "/login", http.StatusFound)
		fmt.Fprint(w, "Unauthorised access to this resource")
	}

}
