package expenses

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExpenseInsertAndFetch(t *testing.T) {
	entry := ExpenseEntry{
		Name:      "Dinner A",
		Amount:    25,
		Category:  1,
		Date:      time.Now().UTC(),
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
		Name:      "Dinner A",
		Amount:    25,
		Category:  1,
		Date:      time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	e := entry.Insert()
	assert.Nil(t, e)
	entry.Name = "Mercadona"
	entry.Save()
	assert.Equal(t, entry.Name, "Mercadona")
	e = entry.Delete()
	assert.Nil(t, e)
}

func TestExpenseSaveSeedData(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	strStart := "2020-12-12T08:00:00.371Z"
	t1, _ := time.Parse(layout, strStart)
	table := []ExpenseEntry{
		{Name: "Weekly big food", Amount: 11, Date: t1},
		{Name: "Weekly big food", Amount: 111, Category: 1, Date: time.Now().UTC()},
		{Name: "Sushi Wednesday", Amount: 1, Category: 2, Date: time.Now().UTC()},
		{Name: "Lunch", Amount: 11, Category: 2, Date: time.Now().UTC()},
		{Name: "B-day gift Roser", Amount: 111, Category: 3, Date: time.Now().UTC()},
		{Name: "Weekly big food", Amount: 111, Category: 1, Date: time.Now().UTC()},
		{Name: "Weekly big food", Amount: 1.1, Category: 1, Date: time.Now().UTC()},
		{Name: "Burger Wednesday", Amount: 11, Category: 2, Date: time.Now().UTC()},
		{Name: "Lunch", Amount: 1, Category: 2, Date: time.Now().UTC()},
		{Name: "B-day gift Gabri", Amount: 1111, Category: 3, Date: time.Now().UTC()}}

	for _, row := range table {
		row.Insert()
	}
	entries := GetExpenseEntries()
	assert.Equal(t, len(entries), 10)
}
