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
	//Route{"Index", http.MethodGet, "/", api.Index},
	Route{"Expenses", http.MethodGet, "/api/expenses", api.Expenses},
	Route{"EntryNew", http.MethodPost, "/api/expenses", api.EntryNew},
	Route{"EntryDelete", http.MethodGet, "/api/expenses/{id}", api.EntryGet},
	Route{"EntryUpdate", http.MethodPut, "/api/expenses/{id}", api.EntryUpdate},
	Route{"EntryDelete", http.MethodDelete, "/api/expenses/{id}", api.EntryDelete},
	Route{"Categories", http.MethodGet, "/api/categories", api.Categories},
	Route{"CategoryNew", http.MethodPost, "/api/categories", api.CategoryNew},
	Route{"CategoryDelete", http.MethodDelete, "/api/categories/{id}", api.CategoryDelete},
}
