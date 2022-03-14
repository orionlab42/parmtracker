package api

import (
	"encoding/json"
	"fmt"
	"github.com/annakallo/parmtracker/data/categories"
	"github.com/annakallo/parmtracker/data/expenses"
	"github.com/annakallo/parmtracker/data/users"
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
	exp := expenses.GetExpenseEntries()
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ChartsExpensesByDate is a handler for: /api/charts-expenses-by-date
func ChartsExpensesByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := expenses.GetExpenseEntriesMergedByDate()
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ChartsExpensesByCategory is a handler for: /api/charts-expenses-by-category
func ChartsExpensesByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := expenses.GetExpenseEntriesMergedByCategory()
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ChartsExpensesByWeek is a handler for: /api/charts-expenses-by-week
func ChartsExpensesByWeek(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := expenses.GetExpenseEntriesMergedByWeek()
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ChartsExpensesByMonth is a handler for: /api/charts-expenses-by-month
func ChartsExpensesByMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := expenses.GetExpenseEntriesMergedByMonth()
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ChartsPieExpensesByCategory is a handler for: /api/charts-pie-expenses-by-category
func ChartsPieExpensesByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := expenses.GetExpenseEntriesPieByMonth()
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// Categories is a handler for: /api/categories
func Categories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	cat := categories.GetCategories()
	if err := json.NewEncoder(w).Encode(cat); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// CategoryNew is a handler for: /api/categories
func CategoryNew(w http.ResponseWriter, r *http.Request) {
	var cat categories.Category
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := json.Unmarshal(body, &cat); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	err = cat.Insert()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// CategoryDelete is a handler for: /api/categories/{id}
func CategoryDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["id"]
	var cat categories.Category
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := cat.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The category with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	err = cat.Delete()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// EntryNew is a handler for: /api/expenses
func EntryNew(w http.ResponseWriter, r *http.Request) {
	var entry expenses.ExpenseEntry
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := json.Unmarshal(body, &entry); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	err = entry.Insert()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// EntryGet is a handler for: /expenses/{id}
func EntryGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryId := vars["id"]
	var entry expenses.ExpenseEntry
	id, err := strconv.Atoi(entryId)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := entry.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The entry with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(entry); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// EntryUpdate is a handler for: /api/expenses/{id}
func EntryUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryId := vars["id"]
	var entry expenses.ExpenseEntry
	id, err := strconv.Atoi(entryId)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := entry.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The entry with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	if err := json.Unmarshal(body, &entry); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	err = entry.Save()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// EntryDelete is a handler for: /api/expenses/{id}
func EntryDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryId := vars["id"]
	var entry expenses.ExpenseEntry
	id, err := strconv.Atoi(entryId)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := entry.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The entry with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	err = entry.Delete()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Users is a handler for: /api/user
func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := users.GetUsers()
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// UserNew is a handler for: /api/user
func UserNew(w http.ResponseWriter, r *http.Request) {
	var user users.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	err = user.Insert()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	jasonWebToken := "blablabla2"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(jasonWebToken); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// UserGet is a handler for: /api/user/{id}
func UserGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	var user users.User
	id, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := user.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The user with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// UserAuth is a handler for: /api/auth
func UserAuth(w http.ResponseWriter, r *http.Request) {
	var user users.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	var u users.User
	_ = u.LoadByName(user.UserName)
	if u.UserName == "" {
		_ = u.LoadByEmail(user.Email)
	}
	jasonWebToken := "blablabla"
	if user.Password == u.Password {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jasonWebToken); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
}

// UserUpdate is a handler for: /api/user/{id}
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	var user users.User
	id, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := user.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The user with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	err = user.Save()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// UserDelete is a handler for: /api/user/{id}
func UserDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	var user users.User
	id, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := user.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The user with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	err = user.Delete()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
