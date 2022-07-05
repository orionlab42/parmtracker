package notes

import (
	"github.com/orionlab42/parmtracker/log"
	"github.com/orionlab42/parmtracker/mysql"
	"github.com/orionlab42/parmtracker/settings"
)

const (
	LogPrefix   = "Version notes"
	PackageName = "notes"
)

func UpdateNotesTable() string {
	version := settings.GetCurrentVersion(PackageName)
	version = updateV1M0(version)
	return version
}

func updateV1M0(version string) string {
	db := mysql.GetInstance().GetConn()

	if version == "" {
		query := `create table if not exists notes (
					note_id int(11) unsigned not null auto_increment,
					user_id int(11) not null,
					note_type int(11) unsigned not null,
					note_title varchar(255)  not null,
					note_text varchar(255)  not null,
					note_empty boolean  not null default true,
					created_at datetime not null default now(),
					updated_at datetime not null default now(),
					PRIMARY KEY (note_id)
					);`
		_, e := db.Exec(query)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble at creating notes table: ", e)
			return version
		}
		log.GetInstance().Infof(LogPrefix, "Table notes created.")

		version = "v1.0-0"
		settings.UpdateVersion(PackageName, version)
	}
	return version
}
