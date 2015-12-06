package main

import (
	"fmt"
	"net/http"
)

const (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "calendar-app"
)

func main() {
	// register handlers
	http.HandleFunc("/", info(authenticate(root)))
	http.HandleFunc("/login", login)
	http.HandleFunc("/show", info(authenticate(show)))
	http.HandleFunc("/create", info(authenticate(create)))

	// start serving
	http.ListenAndServe(":8080", nil)
}

// Info shows a basic information of where the request comes from..
func info(pass http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[INFO] Requst incoming: remote=%s, url=%s, method=%s\n",
			r.RemoteAddr, r.URL.String(), r.Method)
		pass(w, r)
	}
}
