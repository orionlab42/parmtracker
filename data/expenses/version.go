package expenses

import (
	"github.com/annakallo/parmtracker/log"
	"github.com/annakallo/parmtracker/mysql"
)

const (
	LogPrefix = "Version expenses"
)

func UpdateExpensesTable() {

	query := `create table if not exists expenses (
	id int(11) unsigned not null auto_increment,
	entry_name varchar(255)  not null,
	amount float not null,
	category int(11),
	entry_date datetime,
	created_at datetime not null default now(),
	updated_at datetime not null default now(),
	PRIMARY KEY (id)
	);`

	db := mysql.GetInstance().GetConn()
	_, e := db.Exec(query)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "Trouble at creating expenses table: ", e)
	}
	log.GetInstance().Infof(LogPrefix, "Table expenses created.")
}
