package main

import "net/http"

func main() {
	// register handlers
	http.HandleFunc("/", root)
	http.HandleFunc("/login", login)
	http.HandleFunc("/show", show)
	http.HandleFunc("/create", create)

	// start serving
	http.ListenAndServe("localhost:80", nil)
}

func root(w http.ResponseWriter, r *http.Request) {

}

func login(w http.ResponseWriter, r *http.Request) {

}

func show(w http.ResponseWriter, r *http.Request) {

}

func create(w http.ResponseWriter, r *http.Request) {

}
