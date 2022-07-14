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
	if version == "v1.0-0" {
		query := `create table if not exists note_items (
					item_id int(11) unsigned not null auto_increment,
					note_id int(11) not null,
					item_text varchar(255) not null,
					item_is_complete boolean  not null default false,
					item_date datetime  not null,
					created_at datetime not null default now(),
					updated_at datetime not null default now(),
					PRIMARY KEY (item_id)
					);`
		_, e := db.Exec(query)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble at creating note_items table: ", e)
			return version
		}
		log.GetInstance().Infof(LogPrefix, "Table note_items created.")
		version = "v1.0-1"
		settings.UpdateVersion(PackageName, version)
	}

	if version == "v1.0-1" {
		query := `create table if not exists notes_users (
					note_user_id int(11) unsigned not null auto_increment,
					note_id int(11) not null,
					user_id int(11) not null,
					created_at datetime not null default now(),
					updated_at datetime not null default now(),
					PRIMARY KEY (note_user_id)
					);`
		_, e := db.Exec(query)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble at creating notes_users table: ", e)
			return version
		}
		log.GetInstance().Infof(LogPrefix, "Table notes_users created.")
		version = "v1.0-2"
		settings.UpdateVersion(PackageName, version)
	}
	return version
}
