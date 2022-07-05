package notes

import (
	"fmt"
	"github.com/orionlab42/parmtracker/mysql"
	"time"
)

type Item struct {
	ItemId         int       `json:"item_id"`
	NoteId         int       `json:"note_id"`
	ItemText       string    `json:"item_text"`
	ItemIsComplete bool      `json:"item_is_complete"`
	ItemDate       string    `json:"item_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Items []Item

// Load item
func (i *Item) Load(id int) error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from note_items where item_id = ?`)
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
		e := rows.Scan(&i.ItemId, &i.NoteId, &i.ItemText, &i.ItemIsComplete, &i.ItemDate, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading id %v: %s", id, e.Error())
			return e
		}
		i.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		i.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return nil
}

// Insert a new item
func (i *Item) Insert() error {
	if i.CreatedAt.IsZero() {
		i.CreatedAt = time.Now().UTC()
	}
	if i.UpdatedAt.IsZero() {
		i.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`insert note_items set item_id=?, note_id=?, item_text=?, item_is_complete=?, item_date=?, created_at=?, updated_at=?`)
	defer stmt.Close()

	res, e := stmt.Exec(i.ItemId, i.NoteId, i.ItemText, i.ItemIsComplete, i.ItemDate, i.CreatedAt, i.UpdatedAt)
	if e != nil {
		fmt.Printf("Error when inserting new item: %s", e.Error())
		return e
	}
	id, _ := res.LastInsertId()
	i.ItemId = int(id)
	return nil
}

func (i *Item) Save() error {
	if i.UpdatedAt.IsZero() {
		i.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`update note_items set note_id=?, item_text=?, item_is_complete=?, item_date=?, created_at=?, updated_at=? where item_id=?`)
	defer stmt.Close()

	_, e := stmt.Exec(i.NoteId, i.ItemText, i.ItemIsComplete, i.ItemDate, i.CreatedAt, i.UpdatedAt, i.ItemId)
	if e != nil {
		fmt.Printf("Error when saving item: %s", e.Error())
		return e
	}
	return nil
}

func (i *Item) Delete() error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`delete from note_items where item_id=?`)
	defer stmt.Close()
	_, e := stmt.Exec(i.ItemId)
	if e != nil {
		fmt.Printf("Error when deleting item: %s", e.Error())
		return e
	}
	return e
}

func GetItems() Items {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from note_items order by item_id asc;`)
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		fmt.Printf("Error when preparing stmt in getting all items: %s", e.Error())
		return Items{}
	}
	defer rows.Close()
	items := Items{}
	for rows.Next() {
		item := Item{}
		var createdAt string
		var updatedAt string
		e := rows.Scan(&item.ItemId, &item.NoteId, &item.ItemText, &item.ItemIsComplete, &item.ItemDate, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading items: %s", e.Error())
			return Items{}
		}
		item.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		item.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		items = append(items, item)
	}
	return items
}

func GetItemsByNoteId(noteId int) Items {
	allItems := GetItems()
	var items Items
	for _, item := range allItems {
		if item.NoteId == noteId {
			items = append(items, item)
		}
	}
	return items
}
