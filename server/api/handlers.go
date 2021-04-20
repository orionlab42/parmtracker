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
	"strconv"
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

// CategoryNew is a handler for: /api/categories
func CategoryNew(w http.ResponseWriter, r *http.Request) {
	var cat categories.Category
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	fmt.Printf("%+v\n", string(body))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &cat); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Printf("%+v\n", cat)
	cat.Insert()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(cat); err != nil {
		panic(err)
	}
}

// CategoryDelete is a handler for: /api/categories/{id}
func CategoryDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["id"]
	var cat categories.Category
	id, err := strconv.Atoi(categoryId)
	if err == nil {
		fmt.Println(id)
	}
	//fmt.Println("Id: ", id)
	//fmt.Println("Loading the entry with the Id: ",entry.Load(id))

	if err := cat.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The category with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			panic(err)
		}
	}

	cat.Delete()
	categories := categories.GetCategories()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		panic(err)
	}
}

// EntryNew is a handler for: /api/expenses
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(entry); err != nil {
		panic(err)
	}
}

// EntryShow is a handler for: /expenses/{id}
func EntryShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryId := vars["id"]
	var entry expenses.ExpenseEntry
	id, err := strconv.Atoi(entryId)
	if err == nil {
		fmt.Println(id)
	}

	if err := entry.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The entry with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(entry); err != nil {
		panic(err)
	}
}

// EntryUpdate is a handler for: /api/expenses
func EntryUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryId := vars["id"]
	var entry expenses.ExpenseEntry
	id, err := strconv.Atoi(entryId)
	if err == nil {
		fmt.Println(id)
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	fmt.Printf("%+v\n", string(body))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := entry.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The entry with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			panic(err)
		}
	}

	if err := json.Unmarshal(body, &entry); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Printf("%+v\n", entry)
	entry.Save()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(entry); err != nil {
		panic(err)
	}
}

// EntryDelete is a handler for: /api/expenses/{id}
func EntryDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryId := vars["id"]
	var entry expenses.ExpenseEntry
	id, err := strconv.Atoi(entryId)
	if err == nil {
		fmt.Println(id)
	}
	//fmt.Println("Id: ", id)
	//fmt.Println("Loading the entry with the Id: ",entry.Load(id))

	if err := entry.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The entry with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			panic(err)
		}
	}

	entry.Delete()
	expenses := expenses.GetExpenseEntries()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(expenses); err != nil {
		panic(err)
	}
}
