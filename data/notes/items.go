package notes

import (
	"fmt"
	"github.com/orionlab42/parmtracker/mysql"
	"sort"
	"time"
)

type Item struct {
	ItemId         int       `json:"item_id"`
	NoteId         int       `json:"note_id"`
	ItemText       string    `json:"item_text"`
	ItemIsComplete bool      `json:"item_is_complete"`
	ItemDate       time.Time `json:"item_date"`
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
		var itemDate string
		var createdAt string
		var updatedAt string
		e := rows.Scan(&i.ItemId, &i.NoteId, &i.ItemText, &i.ItemIsComplete, &itemDate, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading id %v: %s", id, e.Error())
			return e
		}
		i.ItemDate, _ = time.Parse(mysql.MysqlDateFormat, itemDate)
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
		var itemDate string
		var createdAt string
		var updatedAt string
		e := rows.Scan(&item.ItemId, &item.NoteId, &item.ItemText, &item.ItemIsComplete, &itemDate, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading items: %s", e.Error())
			return Items{}
		}
		item.ItemDate, _ = time.Parse(mysql.MysqlDateFormat, itemDate)
		item.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		item.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		items = append(items, item)
	}
	return items
}

func (p Items) Len() int { return len(p) }

func (p Items) Less(i, j int) bool { return p[i].ItemDate.Before(p[j].ItemDate) }

func (p Items) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func GetItemsByNoteId(noteId int) Items {
	allItems := GetItems()
	var items Items
	for _, item := range allItems {
		if item.NoteId == noteId {
			items = append(items, item)
		}
	}
	if len(items) > 0 {
		layout := "2006-01-02T15:04:05.000Z"
		strStart := "0001-01-01 00:00:00 +0000 UTC 2022-07-08 10:32:51 +0000 UTC"
		t, _ := time.Parse(layout, strStart)
		if items[0].ItemDate != t {
			sort.Sort(items)
		}
	}
	return items
}

func CreateItemsByNoteId(noteId int, startDate string, endDate string) Items {
	items := GetItemsByNoteId(noteId)
	if startDate == "null" || endDate == "null" {
		return items
	}

	start, err := time.Parse("2006-01-02T15:04:05Z07:00", startDate)
	if err != nil {
		panic(err)
	}

	end, err := time.Parse("2006-01-02T15:04:05Z07:00", endDate)
	if err != nil {
		panic(err)
	}

	if len(items) == 0 {
		for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
			i := Item{
				NoteId:         noteId,
				ItemText:       "",
				ItemIsComplete: false,
				ItemDate:       d,
				CreatedAt:      time.Now().UTC(),
				UpdatedAt:      time.Now().UTC(),
			}
			e := i.Insert()
			if e != nil {
				fmt.Printf("Error when adding agenda item: %s", e.Error())
			}
		}
	} else {
		for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
			notSaved := true
			for _, item := range items {
				if item.ItemDate.Before(start) || item.ItemDate.After(end) {
					var i Item
					e := i.Load(item.ItemId)
					if e != nil {
						fmt.Printf("Error when loading agenda item: %s", e.Error())
					}
					e = i.Delete()
					if e != nil {
						fmt.Printf("Error when deleting agenda item: %s", e.Error())
					}
				}
				if item.ItemDate.Equal(d) {
					notSaved = false
				}
			}
			if notSaved {
				i := Item{
					NoteId:         noteId,
					ItemText:       "",
					ItemIsComplete: false,
					ItemDate:       d,
					CreatedAt:      time.Now().UTC(),
					UpdatedAt:      time.Now().UTC(),
				}
				e := i.Insert()
				if e != nil {
					fmt.Printf("Error when adding agenda item: %s", e.Error())
				}
			}
		}
	}
	items = GetItemsByNoteId(noteId)
	return items
}

func DeleteByNoteId(noteId int) {
	allItems := GetItemsByNoteId(noteId)
	for _, item := range allItems {
		var i Item
		e := i.Load(item.ItemId)
		if e != nil {
			fmt.Printf("Error when deleting items: %s", e.Error())
			return
		}
		e = i.Delete()
		if e != nil {
			fmt.Printf("Error when deleting items: %s", e.Error())
			return
		}
	}
}
