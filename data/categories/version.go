package categories

import (
	"github.com/orionlab42/parmtracker/log"
	"github.com/orionlab42/parmtracker/mysql"
	"github.com/orionlab42/parmtracker/settings"
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
					category_color varchar(255),
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
		defaultCategories := []Category{
			{CategoryName: "groceries", CategoryColor: "#dfc6c6", CategoryIcon: "mdi-food-apple"},
			{CategoryName: "restaurant", CategoryColor: "#a4a4a4", CategoryIcon: "mdi-food-fork-drink"},
			{CategoryName: "gift", CategoryColor: "#f08080", CategoryIcon: "mdi-gift"},
			{CategoryName: "housing", CategoryColor: "#badcea", CategoryIcon: "mdi-home"},
			{CategoryName: "transportation", CategoryColor: "#fff68f", CategoryIcon: "mdi-tram"},
			{CategoryName: "utilities", CategoryColor: "#ffc000", CategoryIcon: "mdi-hand-water"},
			{CategoryName: "insurance", CategoryColor: "#817171", CategoryIcon: "mdi-shield-home"},
			{CategoryName: "saving", CategoryColor: "#b0a368", CategoryIcon: "mdi-bank"},
			{CategoryName: "services", CategoryColor: "#e6e6fa", CategoryIcon: "mdi-home-plus-outline"},
			{CategoryName: "healthcare", CategoryColor: "#c39797", CategoryIcon: "mdi-medical-bag"}}
		for _, cat := range defaultCategories {
			stmt, _ := db.Prepare(`insert categories set category_name=?, category_color=?, category_icon=?`)
			_, e := stmt.Exec(cat.CategoryName, cat.CategoryColor, cat.CategoryIcon)
			if e != nil {
				log.GetInstance().Errorf(LogPrefix, "Trouble when inserting default category in categories table: ", e.Error())
				return version
			}
			stmt.Close()
		}

		log.GetInstance().Infof(LogPrefix, "Table categories updated.")
		version = "v1.0-1"
		settings.UpdateVersion(PackageName, version)
	}

	if version == "v1.0-1" {
		cat := Category{}
		e := cat.Load(11)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble when adding category icons in categories table: ", e.Error())
			return version
		}
		cat.CategoryIcon = "mdi-airballoon"
		cat.Save()
		e = cat.Load(12)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble when adding category icons in categories table: ", e.Error())
			return version
		}
		cat.CategoryIcon = "mdi-tshirt-crew"
		cat.Save()

		log.GetInstance().Infof(LogPrefix, "Table categories updated.")
		version = "v1.0-2"
		settings.UpdateVersion(PackageName, version)
	}
	return version
}
