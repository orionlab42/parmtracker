package server

import (
	"github.com/orionlab42/parmtracker/server/api"
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
	Route{"ChartsExpensesByDate", http.MethodGet, "/api/charts-expenses-by-date", api.ChartsExpensesByDate},
	Route{"ChartsExpensesByCategory", http.MethodGet, "/api/charts-expenses-by-category", api.ChartsExpensesByCategory},
	Route{"ChartsExpensesByWeek", http.MethodGet, "/api/charts-expenses-by-week", api.ChartsExpensesByWeek},
	Route{"ChartsExpensesByMonth", http.MethodGet, "/api/charts-expenses-by-month", api.ChartsExpensesByMonth},
	Route{"ChartsPieExpensesByCategory", http.MethodGet, "/api/charts-pie-expenses-by-category", api.ChartsPieExpensesByCategory},
	Route{"EntryNew", http.MethodPost, "/api/expenses", api.EntryNew},
	Route{"EntryGet", http.MethodGet, "/api/expenses/{id}", api.EntryGet},
	Route{"EntryUpdate", http.MethodPut, "/api/expenses/{id}", api.EntryUpdate},
	Route{"EntryDelete", http.MethodDelete, "/api/expenses/{id}", api.EntryDelete},
	Route{"Categories", http.MethodGet, "/api/categories", api.Categories},
	Route{"CategoryNew", http.MethodPost, "/api/categories", api.CategoryNew},
	Route{"CategoryDelete", http.MethodDelete, "/api/categories/{id}", api.CategoryDelete},
	Route{"UserRegister", http.MethodPost, "/api/register", api.UserRegister},
	Route{"UserLogin", http.MethodPost, "/api/login", api.UserLogin},
	Route{"User", http.MethodGet, "/api/user", api.User},
	Route{"Logout", http.MethodPost, "/api/logout", api.Logout},
	Route{"Users", http.MethodGet, "/api/users", api.Users},
	Route{"User", http.MethodGet, "/api/user/{id}", api.UserGet},
	Route{"UserAuth", http.MethodPost, "/api/auth", api.UserAuth},
	Route{"UserUpdate", http.MethodPut, "/api/user/{id}", api.UserUpdate},
	Route{"UserDelete", http.MethodDelete, "/api/user/{id}", api.UserDelete},
}
