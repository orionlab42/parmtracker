package notes

import (
	"fmt"
	"github.com/orionlab42/parmtracker/mysql"
	"time"
)

type NoteUser struct {
	NoteUserId int       `json:"note_user_id"`
	NoteId     int       `json:"note_id"`
	UserId     int       `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type NoteUsers []NoteUser

// Load note_user
func (n *NoteUser) Load(id int) error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from notes_users where note_user_id = ?`)
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
		e := rows.Scan(&n.NoteUserId, &n.NoteId, &n.UserId, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading id %v: %s", id, e.Error())
			return e
		}
		n.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		n.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return nil
}

// Insert a new note_user
func (n *NoteUser) Insert() error {
	if n.CreatedAt.IsZero() {
		n.CreatedAt = time.Now().UTC()
	}
	if n.UpdatedAt.IsZero() {
		n.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`insert notes_users set note_user_id=?, note_id=?, user_id=?, created_at=?, updated_at=?`)
	defer stmt.Close()

	res, e := stmt.Exec(n.NoteUserId, n.NoteId, n.UserId, n.CreatedAt, n.UpdatedAt)
	if e != nil {
		fmt.Printf("Error when inserting new note_user: %s", e.Error())
		return e
	}
	id, _ := res.LastInsertId()
	n.NoteUserId = int(id)
	return nil
}

func (n *NoteUser) Save() error {
	if n.UpdatedAt.IsZero() {
		n.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`update notes_users set note_id=?, user_id=?, created_at=?, updated_at=? where note_user_id=?`)
	defer stmt.Close()

	_, e := stmt.Exec(n.NoteId, n.UserId, n.CreatedAt, n.UpdatedAt, n.NoteUserId)
	if e != nil {
		fmt.Printf("Error when saving note_user: %s", e.Error())
		return e
	}
	return nil
}

func (n *NoteUser) Delete() error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`delete from notes_users where note_user_id=?`)
	defer stmt.Close()
	_, e := stmt.Exec(n.NoteUserId)
	if e != nil {
		fmt.Printf("Error when deleting note_user: %s", e.Error())
		return e
	}
	return e
}

func GetNotesUsers() NoteUsers {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from notes_users order by note_user_id desc;`)
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		fmt.Printf("Error when preparing stmt in getting all note_users: %s", e.Error())
		return NoteUsers{}
	}
	defer rows.Close()
	noteUsers := NoteUsers{}
	for rows.Next() {
		n := NoteUser{}
		var createdAt string
		var updatedAt string
		e := rows.Scan(&n.NoteUserId, &n.NoteId, &n.UserId, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading note_users: %s", e.Error())
			return NoteUsers{}
		}
		n.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		n.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		noteUsers = append(noteUsers, n)
	}
	return noteUsers
}
