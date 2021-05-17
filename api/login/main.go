package login

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
)

type Dictionary map[string]interface{}

var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

// Verifies whether a request has a correct token or not.
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	jsonData := simplejson.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")

		// Validates if token header is set
		if reqToken != "" {

			// Extract token
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]

			// Try to parse the token with HMAC enc algorithm and with signing key
			token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("[-] There was an error trying to parse the token with HMAC algorithm.")
				}

				return mySigningKey, nil
			})

			if err != nil {
				w.WriteHeader(403)

				// Set the JSON Body values
				w.Header().Set("Content-Type", "application/json")
				jsonData.Set("message", "error")
				jsonData.Set("description", "invalid signature")
			}

			// If token is valid, render the homepage
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(401)

			// Set the JSON Body values
			w.Header().Set("Content-Type", "application/json")
			jsonData.Set("message", "error")
			jsonData.Set("description", "not authorized")
		}

		payload, err := jsonData.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}

		w.Write(payload)
	})
}

func GenerateJWT(email string, password string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went wrong %s", err.Error())
		return "", err
	}

	return tokenString, err
}

// Extract claims from a JWT
func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("[-] Invalid JWT Token.")
		return nil, false
	}
}
