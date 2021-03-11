package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	MysqlDateFormat = "2006-01-02 15:04:05"
)

var conn *sql.DB

// Get connection
func GetInstance() *sql.DB {
	if conn == nil {
		var e error
		conn, e = sql.Open("mysql", "root:admin@tcp(localhost:3306)/parmtracker")
		if e != nil {
			fmt.Println(e) // needs a logger
			panic(e)

		}
	}
	return conn
}

func OpenConnection() {
	c := GetInstance()
	err := c.Ping()
	if err != nil {
		panic(err)
	}

}
