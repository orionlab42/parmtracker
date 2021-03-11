package categories

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCategorySaveAndFetch(t *testing.T) {
	cat := Category{
		CategoryName: "groceries",
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
	e := cat.Insert()
	assert.Nil(t, e)
	categories := GetCategories()
	assert.NotEqual(t, len(categories), 0)
	e = cat.Delete()
	assert.Nil(t, e)
}

func TestCategorySave(t *testing.T) {
	cat := Category{
		CategoryName: "groceries",
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
	e := cat.Insert()
	assert.Nil(t, e)
	cat.CategoryName = "gift"
	cat.Save()
	assert.Equal(t, cat.CategoryName, "gift")
	e = cat.Delete()
	assert.Nil(t, e)
}

func TestCategoriesSaveSeedData(t *testing.T) {
	table := []Category{
		{CategoryName: "groceries"},
		{CategoryName: "restaurant"},
		{CategoryName: "gift"},
		{CategoryName: "health"}}

	for _, row := range table {
		row.Insert()
	}
	entries := GetCategories()
	assert.NotEqual(t, len(entries), 0)
}
