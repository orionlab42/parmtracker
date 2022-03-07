package expenses

import (
	"fmt"
	"github.com/annakallo/parmtracker/mysql"
	"time"
)

type ExpenseEntry struct {
	Id        int       `json:"id"`
	Name      string    `json:"entry_name"`
	Amount    float64   `json:"amount"`
	Category  int       `json:"category"`
	UserId    int       `json:"user_id"`
	Shared    bool      `json:"shared"`
	Date      time.Time `json:"entry_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Expenses []ExpenseEntry

// Load expense entry
func (entry *ExpenseEntry) Load(id int) error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from expenses where id = ?`)
	defer stmt.Close()
	rows, e := stmt.Query(id)
	if e != nil {
		fmt.Printf("Error when preparing stmt id %d: %s", id, e.Error())
		return e
	}
	defer rows.Close()
	if rows.Next() {
		var date string
		var createdAt string
		var updatedAt string
		e := rows.Scan(&entry.Id, &entry.Name, &entry.Amount, &entry.Category, &entry.UserId, &entry.Shared, &date, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading id %v: %s", id, e.Error())
			return e
		}
		entry.Date, _ = time.Parse(mysql.MysqlDateFormat, date)
		entry.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		entry.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return nil
}

// Insert a new entry
func (entry *ExpenseEntry) Insert() error {
	if entry.CreatedAt.IsZero() {
		entry.CreatedAt = time.Now().UTC()
	}
	if entry.UpdatedAt.IsZero() {
		entry.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`insert expenses set id=?, entry_name=?, amount=?, category=?, user_id=?, shared=?, entry_date=?, created_at=?, updated_at=?`)
	defer stmt.Close()

	res, e := stmt.Exec(entry.Id, entry.Name, entry.Amount, entry.Category, entry.UserId, entry.Shared, entry.Date, entry.CreatedAt, entry.UpdatedAt)
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
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`update expenses set entry_name=?, amount=?, category=?, user_id=?, shared=?, entry_date=?, created_at=?, updated_at=? where id=?`)
	defer stmt.Close()

	_, e := stmt.Exec(entry.Name, entry.Amount, entry.Category, entry.UserId, entry.Shared, entry.Date, entry.CreatedAt, entry.UpdatedAt, entry.Id)
	if e != nil {
		fmt.Printf("Error when saving expense entry: %s", e.Error())
		return e
	}
	return nil
}

func (entry *ExpenseEntry) Delete() error {
	db := mysql.GetInstance().GetConn()
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
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from expenses order by entry_date asc;`)
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
		var date string
		var createdAt string
		var updatedAt string
		e := rows.Scan(&entry.Id, &entry.Name, &entry.Amount, &entry.Category, &entry.UserId, &entry.Shared, &date, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading entries: %s", e.Error())
			return Expenses{}
		}
		entry.Date, _ = time.Parse(mysql.MysqlDateFormat, date)
		entry.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		entry.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		// correcting the entries that have an incorrect date, by adding the date when it was updated the entry instead
		// the date was generated automatically
		layout := "2006-01-02T15:04:05.000Z"
		strStart := "1000-01-01T08:00:00.371Z"
		firstDate, _ := time.Parse(layout, strStart)
		if entry.Date.Before(firstDate) {
			entry.Date = entry.UpdatedAt
			entry.Save()
		}
		expenses = append(expenses, entry)
	}
	return expenses
}

func GetExpenseEntry(entryId int) Expenses {
	db := mysql.GetInstance().GetConn()
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
		var date string
		var createdAt string
		var updatedAt string
		e := rows.Scan(&entry.Id, &entry.Name, &entry.Amount, &entry.Category, &entry.UserId, &entry.Shared, &date, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading entry with id %d: %s", entryId, e.Error())
			return Expenses{}
		}
		entry.Date, _ = time.Parse(mysql.MysqlDateFormat, date)
		entry.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		entry.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		expenses = append(expenses, entry)
	}
	return expenses
}

func GetExpenseEntriesMergedByDate() Expenses {
	expenses := GetExpenseEntries()
	var expensesNew Expenses
	for _, val := range expenses {
		isSaved := false
		for i, _ := range expensesNew {
			if val.Date == expensesNew[i].Date {
				expensesNew[i].Name = "Total expenses of the day: " + fmt.Sprint(expensesNew[i].Date)
				expensesNew[i].Amount = expensesNew[i].Amount + val.Amount
				isSaved = true
				break
			}
		}
		if isSaved == false {
			expensesNew = append(expensesNew, val)
		}
	}
	return expensesNew
}

func GetExpenseEntriesMergedByCategory() Expenses {
	expenses := GetExpenseEntries()
	var expensesNew Expenses
	for _, val := range expenses {
		isSaved := false
		for i, _ := range expensesNew {
			if val.Category == expensesNew[i].Category {
				expensesNew[i].Amount = expensesNew[i].Amount + val.Amount
				isSaved = true
				break
			}
		}
		if isSaved == false {
			val.Name = "Total expenses of " + fmt.Sprint(val.Category)
			expensesNew = append(expensesNew, val)
		}
	}
	return expensesNew
}

func GetExpenseEntriesMergedByWeek() Expenses {
	expenses := GetExpenseEntries()
	var expensesNew Expenses
	for _, val := range expenses {
		isSaved := false
		year1, week1 := val.Date.ISOWeek()
		for i, _ := range expensesNew {
			year2, week2 := expensesNew[i].Date.ISOWeek()
			if year1 == year2 && week1 == week2 {
				expensesNew[i].Amount = expensesNew[i].Amount + val.Amount
				isSaved = true
				break
			}
		}
		if isSaved == false {
			val.Name = "Week nr." + fmt.Sprint(week1) + "/" + fmt.Sprint(year1)
			expensesNew = append(expensesNew, val)
		}
	}
	return expensesNew
}

func GetExpenseEntriesMergedByMonth() Expenses {
	expenses := GetExpenseEntries()
	var expensesNew Expenses
	for _, val := range expenses {
		isSaved := false
		for i, _ := range expensesNew {
			if val.Date.Month() == expensesNew[i].Date.Month() && val.Date.Year() == expensesNew[i].Date.Year() {
				expensesNew[i].Amount = expensesNew[i].Amount + val.Amount
				isSaved = true
				break
			}
		}
		if isSaved == false {
			val.Name = "Total expenses of " + fmt.Sprint(val.Date.Month()) + " " + fmt.Sprint(val.Date.Year())
			expensesNew = append(expensesNew, val)
		}
	}
	return expensesNew
}

func GetExpenseEntriesPieByMonth() Expenses {
	expensesByCategory := GetExpenseEntriesMergedByCategory()
	var totalExpenses float64
	for _, val := range expensesByCategory {
		totalExpenses += val.Amount
	}
	for i, _ := range expensesByCategory {
		expensesByCategory[i].Amount = expensesByCategory[i].Amount * 100 / totalExpenses
	}
	return expensesByCategory
}
