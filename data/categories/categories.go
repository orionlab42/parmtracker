package categories

import (
	"fmt"
	"github.com/orionlab42/parmtracker/data/expenses"
	"github.com/orionlab42/parmtracker/mysql"
	"time"
)

type Category struct {
	Id            int       `json:"id"`
	CategoryName  string    `json:"category_name"`
	CategoryColor string    `json:"category_color"`
	CategoryIcon  string    `json:"category_icon"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Categories []Category

// Load loads a category based on the id from the categories table.
func (cat *Category) Load(id int) error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from categories where id = ?`)
	defer stmt.Close()
	rows, e := stmt.Query(id)
	if e != nil {
		fmt.Printf("Error when preparing stmt id %d: %s", id, e.Error())
		return e
	}
	defer rows.Close()
	if rows.Next() {
		var createdAt string
		var updatedAt string
		e := rows.Scan(&cat.Id, &cat.CategoryName, &cat.CategoryColor, &cat.CategoryIcon, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading id %v: %s", id, e.Error())
			return e
		}
		cat.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		cat.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return nil
}

// Insert inserts a new category to the categories table.
func (cat *Category) Insert() error {
	if cat.CreatedAt.IsZero() {
		cat.CreatedAt = time.Now().UTC()
	}
	if cat.UpdatedAt.IsZero() {
		cat.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`insert categories set id=?, category_name=?, category_color=?, category_icon=?, created_at=?, updated_at=?`)
	defer stmt.Close()

	res, e := stmt.Exec(cat.Id, cat.CategoryName, &cat.CategoryColor, &cat.CategoryIcon, cat.CreatedAt, cat.UpdatedAt)
	if e != nil {
		fmt.Printf("Error when inserting category: %s", e.Error())
		return e
	}
	id, _ := res.LastInsertId()
	cat.Id = int(id)
	return nil
}

// Save saves a change to a category already inserted to the categories table.
func (cat *Category) Save() error {
	if cat.UpdatedAt.IsZero() {
		cat.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`update categories set category_name=?, category_color=?, category_icon=?, created_at=?, updated_at=? where id=?`)
	defer stmt.Close()

	_, e := stmt.Exec(cat.CategoryName, &cat.CategoryColor, &cat.CategoryIcon, cat.CreatedAt, cat.UpdatedAt, cat.Id)
	if e != nil {
		fmt.Printf("Error when saving category: %s", e.Error())
		return e
	}
	return nil
}

// Delete deletes a category from the categories table.
func (cat *Category) Delete() error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`delete from categories where id=?`)
	defer stmt.Close()
	_, e := stmt.Exec(cat.Id)
	if e != nil {
		fmt.Printf("Error when deleting category: %s", e.Error())
		return e
	}
	return e
}

// GetCategories returns a slice of struct Category with all existing categories in the categories table.
func GetCategories() Categories {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from categories order by category_name ASC`)
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		fmt.Printf("Error when preparing stmt in getting all expenses: %s", e.Error())
		return Categories{}
	}
	defer rows.Close()
	categories := Categories{}
	for rows.Next() {
		cat := Category{}
		var createdAt string
		var updatedAt string
		e := rows.Scan(&cat.Id, &cat.CategoryName, &cat.CategoryColor, &cat.CategoryIcon, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading categories: %s", e.Error())
			return Categories{}
		}
		cat.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		cat.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		categories = append(categories, cat)
	}
	return categories
}

// GetFilledCategoriesIds creates a map with only the categories which already have entries saved. The key will be the id
// of the category and the value is the count of how many entries are of this category.
func GetFilledCategoriesIds() map[int]int {
	expensesAll := expenses.GetExpenseEntries()
	ids := make(map[int]int)
	for _, val := range expensesAll {
		ids[val.Category]++
	}
	return ids
}

// GetFilledCategories returns a slice of struct Category with only the categories which already have entries saved.
func GetFilledCategories() Categories {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from categories order by category_name ASC`)
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		fmt.Printf("Error when preparing stmt in getting all expenses: %s", e.Error())
		return Categories{}
	}
	defer rows.Close()
	categories := Categories{}
	ids := GetFilledCategoriesIds()
	for rows.Next() {
		cat := Category{}
		var createdAt string
		var updatedAt string
		e := rows.Scan(&cat.Id, &cat.CategoryName, &cat.CategoryColor, &cat.CategoryIcon, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading categories: %s", e.Error())
			return Categories{}
		}
		cat.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		cat.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		for id, _ := range ids {
			if cat.Id == id {
				categories = append(categories, cat)
			}
		}
	}
	return categories
}
