package settings

import (
	"github.com/annakallo/parmtracker/log"
	"github.com/annakallo/parmtracker/mysql"
)

const (
	LogPrefix   = "Version settings"
	PackageName = "settings"
)

func UpdateSettingsTable() string {
	version := GetCurrentVersion(PackageName)
	version = updateV1M0(version)
	return version
}

func updateV1M0(version string) string {
	db := mysql.GetInstance().GetConn()
	if version == "" {
		query := `create table if not exists settings (
					k varchar(255)  default '' not null,
					v varchar(255)  default '' not null,
					created_at datetime default current_timestamp() not null,
					updated_at datetime default current_timestamp() not null,
					primary key (k)
					);`
		_, e := db.Exec(query)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble at creating settings table: ", e)
			return version
		}
		log.GetInstance().Infof(LogPrefix, "Table settings created.")

		version = "v1.0-0"
		UpdateVersion(PackageName, version)
	}

	return version
}
