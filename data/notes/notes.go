package notes

import (
	"fmt"
	"github.com/orionlab42/parmtracker/mysql"
	"time"
)

type Note struct {
	NoteId    int       `json:"note_id"`
	OwnerId   int       `json:"owner_id"`
	NoteType  int       `json:"note_type"`
	NoteTitle string    `json:"note_title"`
	NoteText  string    `json:"note_text"`
	NoteEmpty bool      `json:"note_empty"`
	NoteItems []Item    `json:"note_items"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Notes []Note

// Load note
func (n *Note) Load(id int) error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from notes where note_id = ?`)
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
		e := rows.Scan(&n.NoteId, &n.OwnerId, &n.NoteType, &n.NoteTitle, &n.NoteText, &n.NoteEmpty, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading id %v: %s", id, e.Error())
			return e
		}
		n.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		n.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		n.NoteItems = GetItemsByNoteId(n.NoteId)
	}
	return nil
}

// Insert a new note
func (n *Note) Insert() error {
	if n.CreatedAt.IsZero() {
		n.CreatedAt = time.Now().UTC()
	}
	if n.UpdatedAt.IsZero() {
		n.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`insert notes set note_id=?, owner_id=?, note_type=?, note_title=?, note_text=?, note_empty=?, created_at=?, updated_at=?`)
	defer stmt.Close()

	res, e := stmt.Exec(n.NoteId, n.OwnerId, n.NoteType, n.NoteTitle, n.NoteText, n.NoteEmpty, n.CreatedAt, n.UpdatedAt)
	if e != nil {
		fmt.Printf("Error when inserting new note: %s", e.Error())
		return e
	}
	id, _ := res.LastInsertId()
	n.NoteId = int(id)
	return nil
}

func (n *Note) Save() error {
	if n.UpdatedAt.IsZero() {
		n.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`update notes set owner_id=?, note_type=?, note_title=?, note_text=?, note_empty=?, created_at=?, updated_at=? where note_id=?`)
	defer stmt.Close()

	_, e := stmt.Exec(n.OwnerId, n.NoteType, n.NoteTitle, n.NoteText, n.NoteEmpty, n.CreatedAt, n.UpdatedAt, n.NoteId)
	if e != nil {
		fmt.Printf("Error when saving note: %s", e.Error())
		return e
	}
	return nil
}

func (n *Note) Delete() error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`delete from notes where note_id=?`)
	defer stmt.Close()
	_, e := stmt.Exec(n.NoteId)
	if e != nil {
		fmt.Printf("Error when deleting note: %s", e.Error())
		return e
	}
	return e
}

func GetNotes() Notes {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from notes order by note_id desc;`)
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		fmt.Printf("Error when preparing stmt in getting all notes: %s", e.Error())
		return Notes{}
	}
	defer rows.Close()
	notes := Notes{}
	for rows.Next() {
		note := Note{}
		var createdAt string
		var updatedAt string
		e := rows.Scan(&note.NoteId, &note.OwnerId, &note.NoteType, &note.NoteTitle, &note.NoteText, &note.NoteEmpty, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading notes: %s", e.Error())
			return Notes{}
		}
		note.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		note.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		note.NoteItems = GetItemsByNoteId(note.NoteId)
		notes = append(notes, note)
	}
	return notes
}

func GetNotesByIds(noteUsers NoteUsers) Notes {
	var notes Notes
	allNotes := GetNotes()
	for _, noteUser := range noteUsers {
		for _, note := range allNotes {
			if noteUser.NoteId == note.NoteId {
				notes = append(notes, note)
			}
		}
	}
	return notes
}
