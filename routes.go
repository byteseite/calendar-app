package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"/var/www/byteseite/calendar-app/login.html",
))

// Root checks whether the visitor is already signed in or not and redirects
// either to a basic home page or to the login prompt
func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "You are at the root directory.")
}

// Login shows a login prompt and validates post requests
func login(w http.ResponseWriter, r *http.Request) {
	cxt := loginContext{}

	if r.PostFormValue("username") != "" {
		user := r.PostFormValue("username")
		pass := r.PostFormValue("password")
		fmt.Printf("[INFO] Login request: username=%s, password=%s\n", user, pass)
		if validateLogin(user, pass) {
			logmein(user, pass)
			show(w, r)
		} else {
			fmt.Printf("[INFO] login rejected: username=%s, password=%s\n",
				user, pass)
		}
	}

	fmt.Println("[INFO] Rendering template: login.html")
	templates.ExecuteTemplate(w, "login.html", &cxt)
}

// Show renders a calendar either in month, week or day view
func show(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Soon, here you will find your calendar")
}

// Create validates a new event and possibly saves it to the calendar
func create(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Soon, here you will be creating new appointments and alikes.")
}

// LoginContext is givven to the template login.html
type loginContext struct {
	Error string
}
