package mysql

import (
	"database/sql"
	"fmt"
	"github.com/annakallo/parmtracker/config"
	"github.com/annakallo/parmtracker/log"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

const (
	LogPrefix       = "Mysql"
	MysqlDateFormat = "2006-01-02 15:04:05"
)

type Mysql struct {
	Driver string
	Ip     string
	Port   string
	User   string
	Pass   string
	Db     string
	Log    *log.Logger
}

var instance *Mysql
var once sync.Once
var conn *sql.DB

func NewServer() *Mysql {
	conf := config.GetInstance()
	mysql := Mysql{
		Driver: "mysql",
		Ip:     conf.MysqlIP,
		Port:   conf.MysqlPort,
		User:   conf.MysqlUser,
		Pass:   conf.MysqlPass,
		Db:     conf.MysqlDB,
		Log:    log.GetInstance(),
	}
	return &mysql
}

// Transforming Mysql into a Singleton
func GetInstance() *Mysql {
	once.Do(func() {
		instance = NewServer()
	})
	return instance
}

// Get connection
func (m *Mysql) GetConn() *sql.DB {
	if conn == nil {
		var e error
		conn, e = sql.Open(m.Driver, m.User+":"+m.Pass+"@tcp("+m.Ip+":"+m.Port+")/"+m.Db)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "could not connect: ", e.Error())
		}
	}
	return conn
}

// Gets connection
func OpenConnection() {
	mysql := GetInstance()
	c := mysql.GetConn()
	err := c.Ping()
	if err != nil {
		log.GetInstance().Errorf(LogPrefix, "could not open connection: ", err)
	}
}

// Close the connection because why not to implement it
func (m *Mysql) CloseConn() {
	_ = conn.Close()
	conn = nil
}

// Check errors and panic
func (m *Mysql) PanicIfError(prefix string, err error) {
	if err != nil {
		log.GetInstance().Errorf(prefix, "panic! error appeared: %s", err.Error())
		panic(fmt.Sprintf("panic! error appeared: %s", err.Error()))
	}
}
