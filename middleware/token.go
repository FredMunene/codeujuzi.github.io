package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// assign a user a token after sign-in
// generate token, store it?
// validate token during sign in or requesting resources
// cookies? how they work??

var secretKey = []byte("secret-key")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": username, // Subject (user identifier)
			"iss": "codeujuzi", // Issuer
			"aud": getRole(username), // Audience (user role)
			"exp": time.Now().Add(time.Hour * 24).Unix(), // EXpiration time
			"iat": time.Now().Unix(), // Issued at
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}


func getRole(username string) string {
	if username == "instructor" {
		return "instructor"
	}

	return "student"
}

func TokenValidation(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
	if tokenString != ""{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("missing token")
		return
	}

	// strip bearer part
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	err := VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("invalid token")
		return
	}

	next(w,r)
	})
}