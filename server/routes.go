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
	Route{"Index", "GET", "/", api.Index},
	Route{"Expenses", "GET", "/api/expenses", api.Expenses},
	Route{"Categories", "GET", "/api/categories", api.Categories},
	Route{"EntryShow", "GET", "/expenses/{id}", api.EntryShow},
	Route{"EntryNew", "POST", "/expenses/new", api.EntryNew},
}
