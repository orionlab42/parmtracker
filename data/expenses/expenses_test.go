package expenses

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExpenseInsertAndFetch(t *testing.T) {
	entry := ExpenseEntry{
		Title:     "Dinner A",
		Amount:    25,
		Category:  1,
		Shop:      "Lidl",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	e := entry.Insert()
	assert.Nil(t, e)
	entries := GetExpenseEntries()
	assert.NotEqual(t, len(entries), 0)
	e = entry.Delete()
	assert.Nil(t, e)
}

func TestExpenseSave(t *testing.T) {
	entry := ExpenseEntry{
		Title:     "Dinner A",
		Amount:    25,
		Category:  1,
		Shop:      "Lidl",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	e := entry.Insert()
	assert.Nil(t, e)
	entry.Shop = "Mercadona"
	entry.Save()
	assert.Equal(t, entry.Shop, "Mercadona")
	e = entry.Delete()
	assert.Nil(t, e)
}

func TestExpenseSaveSeedData(t *testing.T) {
	table := []ExpenseEntry{
		{Title: "Weekly big food", Amount: 11, Shop: "Mercadona"},
		{Title: "Weekly big food", Amount: 111, Category: 1, Shop: "Lidl"},
		{Title: "Sushi Wednesday", Amount: 1, Category: 2, Shop: "Koyo"},
		{Title: "Lunch", Amount: 11, Category: 2, Shop: "KFC"},
		{Title: "B-day gift Roser", Amount: 111, Category: 3, Shop: "Yves Rocher"},
		{Title: "Weekly big food", Amount: 111, Category: 1, Shop: "Corte Ingles"},
		{Title: "Weekly big food", Amount: 1.1, Category: 1, Shop: "Lidl"},
		{Title: "Burger Wednesday", Amount: 11, Category: 2, Shop: "Goiko"},
		{Title: "Lunch", Amount: 1, Category: 2, Shop: "Burger King"},
		{Title: "B-day gift Gabri", Amount: 1111, Category: 3, Shop: "Artisanal beers"}}

	for _, row := range table {
		row.Insert()
	}
	entries := GetExpenseEntries()
	assert.Equal(t, len(entries), 10)
}
