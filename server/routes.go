package server

import (
	"github.com/annakallo/parmtracker/server/api"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", http.MethodGet, "/", api.Index},
	Route{"Expenses", http.MethodGet, "/api/expenses", api.Expenses},
	Route{"EntryNew", http.MethodPost, "/api/expenses", api.EntryNew},
	Route{"Categories", http.MethodGet, "/api/categories", api.Categories},
	//Route{"EntryNew", "POST", "/api/expenses/new", api.EntryNew},
	//Route{"EntryShow", "GET", "/expenses/{id}", api.EntryShow},

}
