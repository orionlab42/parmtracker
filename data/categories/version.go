package categories

import (
	"fmt"
	"github.com/annakallo/parmtracker/mysql"
)

func UpdateCategoriesTable() {

	query := `create table if not exists categories (
	id int(11) unsigned not null auto_increment,
	category_name varchar(255)  not null,
	created_at datetime not null default now(),
	updated_at datetime not null default now(),
	PRIMARY KEY (id)
	);`

	db := mysql.GetInstance()
	_, e := db.Exec(query)
	if e != nil {
		panic(e)
	}
	fmt.Println("Table categories created.")
}
