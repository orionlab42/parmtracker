package api

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/orionlab42/parmtracker/config"
	"github.com/orionlab42/parmtracker/data/categories"
	"github.com/orionlab42/parmtracker/data/expenses"
	"github.com/orionlab42/parmtracker/data/filters"
	"github.com/orionlab42/parmtracker/data/notes"
	"github.com/orionlab42/parmtracker/data/users"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Index is a handler for: /
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// IsAuthorized checks with the jwt in the cookies if the user is logged in for sending different data
func IsAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetInstance().JWTSecret), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if token.Valid {
			endpoint(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
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

// ExpensesByDate is a handler for: /api/expenses/by-date/{filter}
func ExpensesByDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter := vars["filter"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := filters.GetExpenseEntriesByDate(filter)
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ChartsExpensesByDate is a handler for: /api/charts-expenses-by-date
func ChartsExpensesByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := filters.GetExpenseEntriesMergedByDate()
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ChartsExpensesByWeek is a handler for: /api/charts-expenses-by-week/{filter}
func ChartsExpensesByWeek(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter := vars["filter"]
	categoryId, err := strconv.Atoi(filter)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if categoryId != 0 {
		var category categories.Category
		if err := category.Load(categoryId); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(404) // not found
			message := "The filter category with the given ID not found."
			if err := json.NewEncoder(w).Encode(message); err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := filters.GetExpenseEntriesMergedByWeek(categoryId)
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ChartsExpensesByMonth is a handler for: /api/charts-expenses-by-month/{filter}
func ChartsExpensesByMonth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter := vars["filter"]
	categoryId, err := strconv.Atoi(filter)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if categoryId != 0 {
		var category categories.Category
		if err := category.Load(categoryId); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(404) // not found
			message := "The filter category with the given ID not found."
			if err := json.NewEncoder(w).Encode(message); err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := filters.GetExpenseEntriesMergedByMonth(categoryId)
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ChartsExpensesByCategory is a handler for: /api/charts-expenses-by-category/{filter}
func ChartsExpensesByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter := vars["filter"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := filters.GetExpenseEntriesMergedByCategory(filter, expenses.Expenses{})
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ChartsExpensesByCategoryAndUser is a handler for: /api/charts-expenses-by-category-and-user/{filter}
func ChartsExpensesByCategoryAndUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter := vars["filter"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := filters.GetExpenseEntriesToSeriesByUser(filter)
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ChartsPieExpensesByCategory is a handler for: /api/charts-pie-expenses-by-category
func ChartsPieExpensesByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter := vars["filter"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	exp := filters.GetExpenseEntriesPieByCategory(filter)
	if err := json.NewEncoder(w).Encode(exp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// Categories is a handler for: /api/categories/{id}
func Categories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var cat categories.Categories
	if id == "all" {
		cat = categories.GetCategories()
	} else {
		cat = categories.GetFilledCategories()
	}
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

// UserRegister is a handler for: /api/user
func UserRegister(w http.ResponseWriter, r *http.Request) {
	var userFake UserFake
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := json.Unmarshal(body, &userFake); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(userFake.Password), 14)
	var user users.User
	user.UserName = userFake.UserName
	user.Password = password
	user.Email = userFake.Email
	user.UserColor = userFake.UserColor
	err = user.Insert()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode("Success"); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

type UserFake struct {
	UserId    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	UserColor string    `json:"user_color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserLogin is a handler for: /api/login
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var userFake UserFake
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := json.Unmarshal(body, &userFake); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}

	u := users.User{}
	_ = u.LoadByName(userFake.UserName)
	if u.UserName == "" {
		_ = u.LoadByEmail(userFake.UserName)
	}
	if u.UserId == 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound) // unprocessable entity
		if err := json.NewEncoder(w).Encode("User not found."); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(userFake.Password)); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Wrong password.")
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(u.UserId),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(), // expires in 1 month
	})
	conf := config.GetInstance()
	token, err := claims.SignedString([]byte(conf.JWTSecret))
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode("Could not login."); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24 * 30), // cookie is valid a month, needs to be checked later
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	// where do you set the cookie, to the writer ir to the reader?
	//r.AddCookie(&cookie)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Success!"); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// User is a handler for: /api/user
func User(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		if err == http.ErrNoCookie {
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetInstance().JWTSecret), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user users.User
	issuer, _ := strconv.Atoi(claims.Issuer)
	_ = user.Load(issuer)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// Logout is a handler for /api/logout
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Success logout!"); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// UpdateUser is a handler for: "/user/update-settings/{id}"
func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

// Notes is a handler for: /api/notes/{id}
func Notes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	noteUsers := notes.GetNotesByUserId(userId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	n := notes.GetNotesByIds(noteUsers)
	if err := json.NewEncoder(w).Encode(n); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

type NoteFrontEnd struct {
	NoteId    string `json:"note_id"`
	NoteType  int    `json:"note_type"`
	NoteTitle string `json:"note_title"`
	NoteText  string `json:"note_text"`
	NoteEmpty bool   `json:"note_empty"`
}

// NoteNew is a handler for: /api/notes
func NoteNew(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	var noteFE NoteFrontEnd
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := json.Unmarshal(body, &noteFE); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	var note notes.Note
	note.NoteType = noteFE.NoteType
	note.NoteTitle = noteFE.NoteTitle
	note.NoteText = noteFE.NoteText
	note.NoteEmpty = noteFE.NoteEmpty
	err = note.Insert()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	var noteUser notes.NoteUser
	noteUser.NoteId = note.NoteId
	noteUser.UserId = userId
	err = noteUser.Insert()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// NoteUpdate is a handler for: /api/notes/{id}
func NoteUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["id"]
	var note notes.Note
	id, err := strconv.Atoi(noteId)
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
	if err := note.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The note with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	if err := json.Unmarshal(body, &note); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	err = note.Save()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// NoteDelete is a handler for: /api/notes/{noteId}/{userId}
func NoteDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idNote := vars["noteId"]
	var note notes.Note
	noteId, err := strconv.Atoi(idNote)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	idUser := vars["userId"]
	userId, err := strconv.Atoi(idUser)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	noteUsers := notes.GetNoteByNoteAndUserId(noteId, userId)
	if len(noteUsers) > 0 {
		noteUser := noteUsers[0]
		err = noteUser.Delete()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	noteUsers = notes.GetNotesByNoteId(noteId)

	if len(noteUsers) == 0 {
		if err := note.Load(noteId); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(404) // not found
			message := "The note with the given ID not found."
			if err := json.NewEncoder(w).Encode(message); err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
		}
		err = note.Delete()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Items is a handler for: /api/notes/items/{id}  -here we are having the note id
func Items(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["id"]
	id, err := strconv.Atoi(noteId)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	i := notes.GetItemsByNoteId(id)
	if err := json.NewEncoder(w).Encode(i); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ItemsAgenda is a handler for: /api/notes/items/parameters:note_id, start_date, end_date
func ItemsAgenda(w http.ResponseWriter, r *http.Request) {
	noteId := r.URL.Query().Get("note_id")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	id, err := strconv.Atoi(noteId)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	i := notes.CreateItemsByNoteId(id, startDate, endDate)
	if err := json.NewEncoder(w).Encode(i); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// ItemsDelete is a handler for: /api/notes/items/{id}  -here we are having the note id
func ItemsDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["id"]
	id, err := strconv.Atoi(noteId)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	notes.DeleteByNoteId(id)
	if err := json.NewEncoder(w).Encode("Deleted all items from note: " + noteId); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

type ItemFrontEnd struct {
	ItemId         string    `json:"item_id"`
	NoteId         int       `json:"note_id"`
	ItemText       string    `json:"item_text"`
	ItemIsComplete bool      `json:"item_is_complete"`
	ItemDate       time.Time `json:"item_date"`
}

// ItemNew is a handler for: /api/notes/items
func ItemNew(w http.ResponseWriter, r *http.Request) {
	var itemFE ItemFrontEnd
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := json.Unmarshal(body, &itemFE); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	var item notes.Item
	item.NoteId = itemFE.NoteId
	item.ItemText = itemFE.ItemText
	item.ItemIsComplete = itemFE.ItemIsComplete
	item.ItemDate = itemFE.ItemDate
	err = item.Insert()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// ItemUpdate is a handler for: /api/notes/items/{id}
func ItemUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["id"]
	var item notes.Item
	id, err := strconv.Atoi(itemId)
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
	if err := item.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The item with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	if err := json.Unmarshal(body, &item); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	err = item.Save()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// ItemDelete is a handler for: /api/notes/items/{id}
func ItemDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["id"]
	var item notes.Item
	id, err := strconv.Atoi(itemId)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err := item.Load(id); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // not found
		message := "The note with the given ID not found."
		if err := json.NewEncoder(w).Encode(message); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	err = item.Delete()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// NoteUsers is a handler for: /api/notes_user/{noteId}
func NoteUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["noteId"]
	noteId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	usersOfNote := notes.GetUsersOfNote(noteId)
	if err := json.NewEncoder(w).Encode(usersOfNote); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

// NoteUserNew is a handler for: /api/notes_user/{noteId}/{userId}
func NoteUserNew(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idNote := vars["noteId"]
	noteId, err := strconv.Atoi(idNote)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	idUser := vars["userId"]
	userId, err := strconv.Atoi(idUser)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	noteUsers := notes.GetNoteByNoteAndUserId(noteId, userId)
	if len(noteUsers) > 0 {
		noteUser := noteUsers[0]
		err = noteUser.Delete()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

	}
	if len(noteUsers) == 0 {
		userNote := notes.NoteUser{
			NoteId: noteId,
			UserId: userId,
		}
		err = userNote.Insert()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
