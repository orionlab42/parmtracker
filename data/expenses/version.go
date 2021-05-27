package expenses

import (
	"github.com/annakallo/parmtracker/log"
	"github.com/annakallo/parmtracker/mysql"
	"github.com/annakallo/parmtracker/settings"
)

const (
	LogPrefix   = "Version expenses"
	PackageName = "expenses"
)

func UpdateExpensesTable() string {
	version := settings.GetCurrentVersion(PackageName)
	version = updateV1M0(version)
	return version
}

func updateV1M0(version string) string {
	db := mysql.GetInstance().GetConn()

	if version == "" {
		query := `create table if not exists expenses (
					id int(11) unsigned not null auto_increment,
					entry_name varchar(255)  not null,
					amount float not null,
					category int(11),
					created_at datetime not null default now(),
					updated_at datetime not null default now(),
					PRIMARY KEY (id)
					);`
		_, e := db.Exec(query)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble at creating expenses table: ", e)
			return version
		}
		log.GetInstance().Infof(LogPrefix, "Table expenses created.")

		version = "v1.0-0"
		settings.UpdateVersion(PackageName, version)
	}

	if version == "v1.0-0" {
		query := `alter table expenses add entry_date datetime after category;`
		_, e := db.Exec(query)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble at updating expenses table: ", e)
			return version
		}
		log.GetInstance().Infof(LogPrefix, "Table expenses updated.")
		version = "v1.0-1"
		settings.UpdateVersion(PackageName, version)
	}

	return version
}
