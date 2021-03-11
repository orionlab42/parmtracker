package data

import (
	"fmt"
	"github.com/annakallo/parmtracker/data/categories"
	"github.com/annakallo/parmtracker/data/expenses"
	"time"
)

var currentId int

var AllExpenses expenses.Expenses
var AllCategories categories.Categories

func convertTimeType(timeString string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	t, _ := time.Parse(layout, timeString)
	return t
}

// Give us some seed data
func init() {
	RepoCreateEntry(expenses.ExpenseEntry{Title: "Weekly big food", Amount: 11, Category: 1, Shop: "Mercadona", CreatedAt: convertTimeType("2021-03-01T19:04:28.809Z")})
	RepoCreateEntry(expenses.ExpenseEntry{Title: "Weekly big food", Amount: 111, Category: 1, Shop: "Lidl", CreatedAt: convertTimeType("2021-01-12T19:04:28.809Z")})
	RepoCreateEntry(expenses.ExpenseEntry{Title: "Sushi Wednesday", Amount: 1, Category: 2, Shop: "Koyo", CreatedAt: convertTimeType("2021-02-25T19:04:28.809Z")})
	RepoCreateEntry(expenses.ExpenseEntry{Title: "Lunch", Amount: 11, Category: 2, Shop: "KFC", CreatedAt: convertTimeType("2021-02-17T19:04:28.809Z")})
	RepoCreateEntry(expenses.ExpenseEntry{Title: "B-day gift Roser", Amount: 111, Category: 3, Shop: "Yves Rocher", CreatedAt: convertTimeType("2021-02-16T19:04:28.809Z")})
	RepoCreateEntry(expenses.ExpenseEntry{Title: "Weekly big food", Amount: 111, Category: 1, Shop: "Corte Ingles", CreatedAt: convertTimeType("2021-02-03T19:04:28.809Z")})
	RepoCreateEntry(expenses.ExpenseEntry{Title: "Weekly big food", Amount: 1.1, Category: 1, Shop: "Lidl", CreatedAt: convertTimeType("2021-02-12T19:04:28.809Z")})
	RepoCreateEntry(expenses.ExpenseEntry{Title: "Burger Wednesday", Amount: 11, Category: 2, Shop: "Goiko", CreatedAt: convertTimeType("2021-03-01T19:04:28.809Z")})
	RepoCreateEntry(expenses.ExpenseEntry{Title: "Lunch", Amount: 1, Category: 2, Shop: "Burger King", CreatedAt: convertTimeType("2021-01-24T19:04:28.809Z")})
	RepoCreateEntry(expenses.ExpenseEntry{Title: "B-day gift Gabri", Amount: 1111, Category: 3, Shop: "Artisanal beers", CreatedAt: convertTimeType("2021-02-25T19:04:28.809Z")})

	RepoCreateCategory(categories.Category{CategoryName: "groceries"})
	RepoCreateCategory(categories.Category{CategoryName: "restaurant"})
	RepoCreateCategory(categories.Category{CategoryName: "gift"})
	RepoCreateCategory(categories.Category{CategoryName: "health"})
}

func RepoFindEntry(id int) expenses.ExpenseEntry {
	for _, entry := range AllExpenses {
		if entry.Id == id {
			return entry
		}
	}
	// return empty To do if not found
	return expenses.ExpenseEntry{}
}

func RepoCreateEntry(entry expenses.ExpenseEntry) expenses.ExpenseEntry {
	currentId += 1
	entry.Id = currentId
	AllExpenses = append(AllExpenses, entry)
	return entry
}

func RepoDestroyEntry(id int) error {
	for i, t := range AllExpenses {
		if t.Id == id {
			AllExpenses = append(AllExpenses[:i], AllExpenses[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not find Todo with id of %d to delete", id)
}

func RepoCreateCategory(category categories.Category) categories.Category {
	currentId += 1
	category.Id = currentId
	AllCategories = append(AllCategories, category)
	return category
}
