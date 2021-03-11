package api

import (
	"encoding/json"
	"fmt"
	"github.com/annakallo/parmtracker/data/categories"
	"github.com/annakallo/parmtracker/data/expenses"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

// Index is a handler for: /
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// Expenses is a handler for: /api/expenses
func Expenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	expenses := expenses.GetExpenseEntries()
	if err := json.NewEncoder(w).Encode(expenses); err != nil {
		panic(err)
	}
}

// Categories is a handler for: /api/categories
func Categories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	categories := categories.GetCategories()
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		panic(err)
	}
}

// EntryShow is a handler for: /expenses/{id}
func EntryShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entry := vars["id"]
	_, _ = fmt.Fprintln(w, "The entry what we are searching, has the id:", entry)
}

// EntryNew is a handler for: /api/expenses/new
func EntryNew(w http.ResponseWriter, r *http.Request) {
	var entry expenses.ExpenseEntry
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	fmt.Printf("%+v\n", string(body))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &entry); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Printf("%+v\n", entry)
	entry.Insert()
	expenses := expenses.GetExpenseEntries()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(expenses); err != nil {
		panic(err)
	}
}
