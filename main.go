package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Hello world!")
	r := mux.NewRouter()

	// for static pages like home
	fs := http.StripPrefix("/home/", http.FileServer(http.Dir("./stat/")))
	//http.Handle("/home/", fs)
	r.PathPrefix("/home/").Handler(fs).Methods("GET")

	// for dynamic routing
	r.HandleFunc("/expenses", handler)
	//http.HandleFunc("/hello", handler)
	http.ListenAndServe(":80", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	fmt.Fprintf(w, "What did you spend?")
}
