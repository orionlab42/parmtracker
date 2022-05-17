package categories

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCategorySaveAndFetch(t *testing.T) {
	cat := Category{
		CategoryName:  "groceries",
		CategoryColor: "#dfc6c6",
		CategoryIcon:  "mdi-home-plus-outline",
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
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
		CategoryName:  "groceries",
		CategoryColor: "#dfc6c6",
		CategoryIcon:  "mdi-home-plus-outline",
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
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
		{CategoryName: "groceries", CategoryColor: "#dfc6c6", CategoryIcon: "mdi-food-apple"},
		{CategoryName: "restaurant", CategoryColor: "#a4a4a4", CategoryIcon: "mdi-food-fork-drink"},
		{CategoryName: "gift", CategoryColor: "#f08080", CategoryIcon: "mdi-gift"},
		{CategoryName: "housing", CategoryColor: "#badcea", CategoryIcon: "mdi-home"},
		{CategoryName: "transportation", CategoryColor: "#fff68f", CategoryIcon: "mdi-tram"},
		{CategoryName: "utilities", CategoryColor: "#ffc000", CategoryIcon: "mdi-hand-water"},
		{CategoryName: "insurance", CategoryColor: "#817171", CategoryIcon: "mdi-shield-home"},
		{CategoryName: "saving", CategoryColor: "#b0a368", CategoryIcon: "mdi-bank"},
		{CategoryName: "services", CategoryColor: "#e6e6fa", CategoryIcon: "mdi-home-plus-outline"},
		{CategoryName: "healthcare", CategoryColor: "#c39797", CategoryIcon: "mdi-medical-bag"}}
	for _, row := range table {
		row.Insert()
	}
	entries := GetCategories()
	assert.NotEqual(t, len(entries), 0)
}

func TestGetFilledCategoriesIds(t *testing.T) {
	categories := GetFilledCategoriesIds()
	for i, row := range categories {
		fmt.Printf("%+v\t", i)
		fmt.Printf("%+v\n", row)
	}
}

func TestGetFilledCategories(t *testing.T) {
	categories := GetFilledCategories()
	for _, row := range categories {
		fmt.Printf("%+v\n", row)
	}
}
