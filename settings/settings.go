package settings

import (
	"github.com/annakallo/parmtracker/log"
	"github.com/annakallo/parmtracker/mysql"
	"time"
)

const (
	TableName = "settings"
)

// Get the value of key in table
func GetCurrentVersion(key string) string {
	db := mysql.GetInstance().GetConn()
	stmt, e := db.Prepare("select v from " + TableName + " where k = ?")
	if e != nil {
		return ""
	}
	defer stmt.Close()

	rows, e := stmt.Query(key)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "when get %s: %s", key, e.Error())
		return ""
	}
	defer rows.Close()

	var value string
	if rows.Next() {
		_ = rows.Scan(&value)
	}

	return value
}

// Set value of key in table
func UpdateVersion(key string, value string) {
	db := mysql.GetInstance().GetConn()
	timestamp := time.Now().UTC()

	stmt, _ := db.Prepare("insert into " + TableName + " (k, v, updated_at) values(?, ?, ?) on duplicate key update v = ?, updated_at = ?")
	defer stmt.Close()

	q, e := stmt.Query(key, value, timestamp, value, timestamp)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "when set %s value %s: %s", key, value, e.Error())
		return
	}
	defer q.Close()
}

// Delete key in table
func Delete(key string) {
	db := mysql.GetInstance().GetConn()

	stmt, _ := db.Prepare("Delete from " + TableName + " where k = ?")
	defer stmt.Close()
	q, e := stmt.Query(key)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "when delete %s: %s", key, e.Error())
		return
	}
	defer q.Close()
}
