package admin

import (
	"mustafa_m/controllers"
	"mustafa_m/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createJWT(w http.ResponseWriter, admin *models.Admin) error {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: admin.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return nil
}

func refreshJWT(w http.ResponseWriter, r *http.Request) error {
	// TODO :
	return nil
}

func (admin *Admin) SignoutJWT(w http.ResponseWriter, r *http.Request) {
	internalServerError := controllers.InternalServerError()
	if !controllers.ValidateJWT(r) {
		internalServerError.Render(w, nil)
	}
	c := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusFound)
}
