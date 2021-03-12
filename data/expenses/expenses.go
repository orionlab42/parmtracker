package expenses

import (
	"fmt"
	"github.com/annakallo/parmtracker/mysql"
	"time"
)

type ExpenseEntry struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Category  int       `json:"category"`
	Shop      string    `json:"shop"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Expenses []ExpenseEntry

// Load trade order
func (entry *ExpenseEntry) Load(id int) error {
	db := mysql.GetInstance()
	stmt, _ := db.Prepare(`select * from expenses where id = ?`)
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
		e := rows.Scan(&entry.Id, &entry.Title, &entry.Amount, &entry.Category, &entry.Shop, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading id %v: %s", id, e.Error())
			return e
		}
		entry.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		entry.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return nil
}

// Insert a new trade
func (entry *ExpenseEntry) Insert() error {
	if entry.CreatedAt.IsZero() {
		entry.CreatedAt = time.Now().UTC()
	}
	if entry.UpdatedAt.IsZero() {
		entry.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance()
	stmt, _ := db.Prepare(`insert expenses set id=?, title=?, amount=?, category=?, shop=?, created_at=?, updated_at=?`)
	defer stmt.Close()

	res, e := stmt.Exec(entry.Id, entry.Title, entry.Amount, entry.Category, entry.Shop, entry.CreatedAt, entry.UpdatedAt)
	if e != nil {
		fmt.Printf("Error when inserting expense entry: %s", e.Error())
		return e
	}
	id, _ := res.LastInsertId()
	entry.Id = int(id)
	return nil
}

func (entry *ExpenseEntry) Save() error {
	if entry.UpdatedAt.IsZero() {
		entry.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance()
	stmt, _ := db.Prepare(`update expenses set title=?, amount=?, category=?, shop=?, created_at=?, updated_at=? where id=?`)
	defer stmt.Close()

	_, e := stmt.Exec(entry.Title, entry.Amount, entry.Category, entry.Shop, entry.CreatedAt, entry.UpdatedAt, entry.Id)
	if e != nil {
		fmt.Printf("Error when saving expense entry: %s", e.Error())
		return e
	}
	return nil
}

func (entry *ExpenseEntry) Delete() error {
	db := mysql.GetInstance()
	stmt, _ := db.Prepare(`delete from expenses where id=?`)
	defer stmt.Close()
	_, e := stmt.Exec(entry.Id)
	if e != nil {
		fmt.Printf("Error when deleting expense entry: %s", e.Error())
		return e
	}
	return e
}

func GetExpenseEntries() Expenses {
	db := mysql.GetInstance()
	stmt, _ := db.Prepare(`select * from expenses`)
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		fmt.Printf("Error when preparing stmt in getting all expenses: %s", e.Error())
		return Expenses{}
	}
	defer rows.Close()
	expenses := Expenses{}
	for rows.Next() {
		entry := ExpenseEntry{}
		var createdAt string
		var updatedAt string
		e := rows.Scan(&entry.Id, &entry.Title, &entry.Amount, &entry.Category, &entry.Shop, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading entries: %s", e.Error())
			return Expenses{}
		}
		entry.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		entry.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		expenses = append(expenses, entry)
	}
	return expenses
}

func GetExpenseEntry(entryId int) Expenses {
	db := mysql.GetInstance()
	stmt, _ := db.Prepare(`select * from expenses where id = ?`)
	defer stmt.Close()
	rows, e := stmt.Query(entryId)
	if e != nil {
		fmt.Printf("Error when preparing stmt in getting entry with id %d: %s", entryId, e.Error())
		return Expenses{}
	}
	defer rows.Close()
	expenses := Expenses{}
	for rows.Next() {
		entry := ExpenseEntry{}
		var createdAt string
		var updatedAt string
		e := rows.Scan(&entry.Id, &entry.Title, &entry.Amount, &entry.Category, &entry.Shop, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading entry with id %d: %s", entryId, e.Error())
			return Expenses{}
		}
		entry.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		entry.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		expenses = append(expenses, entry)
	}
	return expenses
}
