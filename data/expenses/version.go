package expenses

import (
	"fmt"
	"github.com/annakallo/parmtracker/mysql"
)

func UpdateExpensesTable() {

	query := `create table if not exists expenses (
	id int(11) unsigned not null auto_increment,
	title varchar(255)  not null,
	amount float not null,
	category int(11),
	shop varchar(255),
	created_at datetime not null default now(),
	updated_at datetime not null default now(),
	PRIMARY KEY (id)
	);`

	db := mysql.GetInstance()
	_, e := db.Exec(query)
	if e != nil {
		panic(e)
	}
	fmt.Println("Table expenses created.")
}
