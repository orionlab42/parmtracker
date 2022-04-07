package users

import (
	"fmt"
	"github.com/orionlab42/parmtracker/mysql"
	"time"
)

type User struct {
	UserId    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	Password  []byte    `json:"-"`
	Email     string    `json:"email"`
	UserColor string    `json:"user_color"`
	DarkMode  bool      `json:"dark_mode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

// Load user
func (u *User) Load(id int) error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from users where user_id = ?`)
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
		e := rows.Scan(&u.UserId, &u.UserName, &u.Password, &u.Email, &u.UserColor, &u.DarkMode, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading id %v: %s", id, e.Error())
			return e
		}
		u.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		u.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return nil
}

// LoadByName user
func (u *User) LoadByName(name string) error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from users where user_name = ?`)
	defer stmt.Close()
	rows, e := stmt.Query(name)
	if e != nil {
		fmt.Printf("Error when preparing stmt id %v: %s", name, e.Error())
		return e
	}
	defer rows.Close()
	if rows.Next() {
		var createdAt string
		var updatedAt string
		e := rows.Scan(&u.UserId, &u.UserName, &u.Password, &u.Email, &u.UserColor, &u.DarkMode, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading name %v: %s", name, e.Error())
			return e
		}
		u.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		u.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return nil
}

// LoadByEmail user
func (u *User) LoadByEmail(email string) error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from users where email = ?`)
	defer stmt.Close()
	rows, e := stmt.Query(email)
	if e != nil {
		fmt.Printf("Error when preparing stmt id %v: %s", email, e.Error())
		return e
	}
	defer rows.Close()
	if rows.Next() {
		var createdAt string
		var updatedAt string
		e := rows.Scan(&u.UserId, &u.UserName, &u.Password, &u.Email, &u.UserColor, &u.DarkMode, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading email %v: %s", email, e.Error())
			return e
		}
		u.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		u.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return nil
}

// Insert a new user
func (u *User) Insert() error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now().UTC()
	}
	if u.UpdatedAt.IsZero() {
		u.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`insert users set user_id=?, user_name=?, password=?, email=?, user_color=?, dark_mode=?, created_at=?, updated_at=?`)
	defer stmt.Close()

	res, e := stmt.Exec(u.UserId, u.UserName, u.Password, u.Email, u.UserColor, u.DarkMode, u.CreatedAt, u.UpdatedAt)
	if e != nil {
		fmt.Printf("Error when inserting new user: %s", e.Error())
		return e
	}
	id, _ := res.LastInsertId()
	u.UserId = int(id)
	return nil
}

func (u *User) Save() error {
	if u.UpdatedAt.IsZero() {
		u.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`update users set user_name=?, password=?, email=?, user_color=?, dark_mode=?, created_at=?, updated_at=? where user_id=?`)
	defer stmt.Close()

	_, e := stmt.Exec(u.UserName, u.Password, u.Email, u.UserColor, u.DarkMode, u.CreatedAt, u.UpdatedAt, u.UserId)
	if e != nil {
		fmt.Printf("Error when saving user: %s", e.Error())
		return e
	}
	return nil
}

func (u *User) Delete() error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`delete from users where user_id=?`)
	defer stmt.Close()
	_, e := stmt.Exec(u.UserId)
	if e != nil {
		fmt.Printf("Error when deleting user: %s", e.Error())
		return e
	}
	return e
}

func GetUsers() Users {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from users order by user_id asc;`)
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		fmt.Printf("Error when preparing stmt in getting all users: %s", e.Error())
		return Users{}
	}
	defer rows.Close()
	users := Users{}
	for rows.Next() {
		u := User{}
		var createdAt string
		var updatedAt string
		e := rows.Scan(&u.UserId, &u.UserName, &u.Password, &u.Email, &u.UserColor, &u.DarkMode, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading users: %s", e.Error())
			return Users{}
		}
		u.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		u.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		users = append(users, u)
	}
	return users
}
