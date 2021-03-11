package main

import (
	"github.com/annakallo/parmtracker/data/categories"
	"github.com/annakallo/parmtracker/data/expenses"
	"github.com/annakallo/parmtracker/mysql"
	"github.com/annakallo/parmtracker/server"
	"log"
	"net/http"
)

func main() {
	mysql.OpenConnection()
	expenses.UpdateExpensesTable()
	categories.UpdateCategoriesTable()
	r := server.NewRouter()
	log.Fatal(http.ListenAndServe(":80", r))
}
