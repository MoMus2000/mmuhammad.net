package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func createJWT(w http.ResponseWriter, Email string) error {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "smsAccessToken",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/",
	})

	return nil
}

func refreshJWT(w http.ResponseWriter, r *http.Request) error {
	// TODO :
	return nil
}

func SignoutJWT(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "smsAccessToken",
		MaxAge: -1}
	http.SetCookie(w, &c)
}

func ValidateJWT(r *http.Request) bool {
	token, err := r.Cookie("smsAccessToken")
	if err != nil {
		return false
	}

	tokenString := token.Value

	claims := &Claims{}

	result, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !result.Valid {
		return false
	}

	return true
}

func GetEmailFromJwt(r *http.Request) *Claims {
	token, err := r.Cookie("smsAccessToken")
	if err != nil {
		return nil
	}

	tokenString := token.Value

	claims := &Claims{}

	_, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil
	}

	return claims
}
