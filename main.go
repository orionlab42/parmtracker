package main

import (
	"fmt"
	"github.com/annakallo/parmtracker/mysql"
	"github.com/annakallo/parmtracker/server"
	"net/http"
)

func main() {
	mysql.OpenConnection()

	fmt.Println("Hello world!")
	//r := mux.NewRouter()
	//
	//// for static pages like home
	//fs := http.StripPrefix("/home", http.FileServer(http.Dir("./stat/")))
	////http.Handle("/home/", fs)
	//r.PathPrefix("/home").Handler(fs).Methods("GET")
	//
	//// for dynamic routing
	//r.HandleFunc("/expenses", handler)
	//r.HandleFunc("/api/expenses", handlerApi)

	// other tutorial
	r := server.NewRouter()
	//r.HandleFunc("/todos", TodoIndex)
	//r.HandleFunc("/todos/{todoId1}", TodoShow)
	//http.HandleFunc("/hello", handler)
	http.ListenAndServe(":80", r)
	//log.Fatal(http.ListenAndServe(":12345", r))
}

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
//	fmt.Fprintf(w, "What did you spend?")
//}
//
//func handlerApi(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
//	fmt.Fprintf(w, "What did you spend?")
//}
