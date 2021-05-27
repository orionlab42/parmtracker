package categories

import (
	"github.com/annakallo/parmtracker/log"
	"github.com/annakallo/parmtracker/mysql"
	"github.com/annakallo/parmtracker/settings"
)

const (
	LogPrefix   = "Version categories"
	PackageName = "categories"
)

func UpdateCategoriesTable() string {
	version := settings.GetCurrentVersion(PackageName)
	version = updateV1M0(version)
	return version
}

func updateV1M0(version string) string {
	db := mysql.GetInstance().GetConn()

	if version == "" {
		query := `create table if not exists categories (
					id int(11) unsigned not null auto_increment,
					category_name varchar(255)  not null,
					category_icon varchar(255),
					created_at datetime not null default now(),
					updated_at datetime not null default now(),
					PRIMARY KEY (id)
					);`
		_, e := db.Exec(query)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble at creating categories table: ", e)
			return version
		}
		log.GetInstance().Infof(LogPrefix, "Table categories created.")

		version = "v1.0-0"
		settings.UpdateVersion(PackageName, version)
	}

	if version == "v1.0-0" {
		query := `alter table categories add category_color varchar(255) after category_name;`
		_, e := db.Exec(query)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble at updating categories table: ", e)
			return version
		}
		log.GetInstance().Infof(LogPrefix, "Table categories updated.")
		version = "v1.0-1"
		settings.UpdateVersion(PackageName, version)
	}

	return version
}
