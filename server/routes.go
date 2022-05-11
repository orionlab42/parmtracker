package server

import (
	"github.com/orionlab42/parmtracker/server/api"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Auth        bool
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	//Route{"Index", http.MethodGet, "/", api.Index},
	Route{"Expenses", http.MethodGet, "/api/expenses", true, api.Expenses},
	Route{"ChartsExpensesByDate", http.MethodGet, "/api/charts-expenses-by-date", true, api.ChartsExpensesByDate},
	Route{"ChartsExpensesByCategory", http.MethodGet, "/api/charts-expenses-by-category/{filter}", true, api.ChartsExpensesByCategory},
	Route{"ChartsExpensesByWeek", http.MethodGet, "/api/charts-expenses-by-week/{filter}", true, api.ChartsExpensesByWeek},
	Route{"ChartsExpensesByMonth", http.MethodGet, "/api/charts-expenses-by-month/{filter}", true, api.ChartsExpensesByMonth},
	Route{"ChartsPieExpensesByCategory", http.MethodGet, "/api/charts-pie-expenses-by-category/{filter}", true, api.ChartsPieExpensesByCategory},
	Route{"EntryNew", http.MethodPost, "/api/expenses", true, api.EntryNew},
	Route{"EntryGet", http.MethodGet, "/api/expenses/{id}", true, api.EntryGet},
	Route{"EntryUpdate", http.MethodPut, "/api/expenses/{id}", true, api.EntryUpdate},
	Route{"EntryDelete", http.MethodDelete, "/api/expenses/{id}", true, api.EntryDelete},
	Route{"Categories", http.MethodGet, "/api/categories/{id}", true, api.Categories},
	Route{"CategoryNew", http.MethodPost, "/api/categories", true, api.CategoryNew},
	Route{"CategoryDelete", http.MethodDelete, "/api/categories/{id}", true, api.CategoryDelete},

	Route{"UserRegister", http.MethodPost, "/api/register", false, api.UserRegister},
	Route{"UserLogin", http.MethodPost, "/api/login", false, api.UserLogin},
	Route{"User", http.MethodGet, "/api/user", false, api.User},
	Route{"Logout", http.MethodPost, "/api/logout", false, api.Logout},

	Route{"Users", http.MethodGet, "/api/all-users", true, api.Users},
}
