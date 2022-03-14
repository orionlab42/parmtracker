package users

import (
	"github.com/orionlab42/parmtracker/log"
	"github.com/orionlab42/parmtracker/mysql"
	"github.com/orionlab42/parmtracker/settings"
)

const (
	LogPrefix   = "Version users"
	PackageName = "users"
)

func UpdateUsersTable() string {
	version := settings.GetCurrentVersion(PackageName)
	version = updateV1M0(version)
	return version
}

func updateV1M0(version string) string {
	db := mysql.GetInstance().GetConn()

	if version == "" {
		query := `create table if not exists users (
					user_id int(11) unsigned not null auto_increment,
					user_name varchar(255)  not null,
					password varchar(255)  not null,
					email varchar(255)  not null,
    				user_color varchar(255),
					created_at datetime not null default now(),
					updated_at datetime not null default now(),
					PRIMARY KEY (user_id)
					);`
		_, e := db.Exec(query)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble at creating users table: ", e)
			return version
		}
		log.GetInstance().Infof(LogPrefix, "Table users created.")

		version = "v1.0-0"
		settings.UpdateVersion(PackageName, version)
	}
	return version
}
