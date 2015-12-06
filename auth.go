package main

import (
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"time"
)

// Authenticate is a wrapper function for ensuring the request comes with a
// valid session key
func authenticate(pass http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("SessionKey")
		if !validateAuthToken(token) {
			login(w, r)
		} else {
			pass(w, r)
		}
	}
}

func validateAuthToken(token string) bool {
	return token == "qwertzui"
}

// ValidateLogin checks whether the login request should be accepted or denied
func validateLogin(username, password string) bool {
	return username == "me" && password == "secret"
}

func logmein(username, password string) (sessionID string) {
	sessionID = newToken()
	return
}

func newToken() (token string) {
	hasher := sha1.New()
	raw := []byte(time.Now().String())
	hasher.Write(raw)
	token = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return
}
