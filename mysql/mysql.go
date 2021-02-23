package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt time.Time
}

func OpenConnection() {
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/parmtracker")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//query := `
	//CREATE TABLE users (
	//    id INT AUTO_INCREMENT,
	//    username TEXT NOT NULL,
	//    password TEXT NOT NULL,
	//    created_at DATETIME,
	//    PRIMARY KEY (id)
	//);`
	//_, e := db.Exec(query)
	//if e != nil {
	//	panic(e)
	//}

	username := "johndoe1"
	password := "secret1"
	createdAt := time.Now()

	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	fmt.Println(result)
	if err != nil {
		panic(err)
	}
	userID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(userID)

	// Query the database
	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		var createdAtString string
		err := rows.Scan(&u.Id, &u.Username, &u.Password, &createdAtString)
		if err != nil {
			panic(err)
		}
		u.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAtString)
		users = append(users, u)
	}

	for _, user := range users {
		fmt.Println(user.Id, user.Username, user.CreatedAt)
	}
}
