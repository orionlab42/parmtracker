package categories

import (
	"fmt"
	"github.com/annakallo/parmtracker/mysql"
	"time"
)

type Category struct {
	Id           int       `json:"id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Categories []Category

// Load trade order
func (cat *Category) Load(id int) error {
	db := mysql.GetInstance()
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
		e := rows.Scan(&cat.Id, &cat.CategoryName, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading id %v: %s", id, e.Error())
			return e
		}
		cat.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		cat.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return nil
}

// Insert a new trade
func (cat *Category) Insert() error {
	if cat.CreatedAt.IsZero() {
		cat.CreatedAt = time.Now().UTC()
	}
	if cat.UpdatedAt.IsZero() {
		cat.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance()
	stmt, _ := db.Prepare(`insert categories set id=?, category_name=?, created_at=?, updated_at=?`)
	defer stmt.Close()

	res, e := stmt.Exec(cat.Id, cat.CategoryName, cat.CreatedAt, cat.UpdatedAt)
	if e != nil {
		fmt.Printf("Error when inserting category: %s", e.Error())
		return e
	}
	id, _ := res.LastInsertId()
	cat.Id = int(id)
	return nil
}

func (cat *Category) Save() error {
	if cat.UpdatedAt.IsZero() {
		cat.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance()
	stmt, _ := db.Prepare(`update categories set category_name=?, created_at=?, updated_at=? where id=?`)
	defer stmt.Close()

	_, e := stmt.Exec(cat.CategoryName, cat.CreatedAt, cat.UpdatedAt, cat.Id)
	if e != nil {
		fmt.Printf("Error when saving category: %s", e.Error())
		return e
	}
	return nil
}

func (cat *Category) Delete() error {
	db := mysql.GetInstance()
	stmt, _ := db.Prepare(`delete from categories where id=?`)
	defer stmt.Close()
	_, e := stmt.Exec(cat.Id)
	if e != nil {
		fmt.Printf("Error when deleting category: %s", e.Error())
		return e
	}
	return e
}

func GetCategories() Categories {
	db := mysql.GetInstance()
	stmt, _ := db.Prepare(`select * from categories order by id DESC`)
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		fmt.Printf("Error when preparing stmt in getting all expenses: %s", e.Error())
		return Categories{}
	}
	defer rows.Close()
	categories := Categories{}
	if rows.Next() {
		cat := Category{}
		var createdAt string
		var updatedAt string
		e := rows.Scan(&cat.Id, &cat.CategoryName, &createdAt, &updatedAt)
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
