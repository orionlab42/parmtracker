package main

import (
	"github.com/orionlab42/parmtracker/config"
	"github.com/orionlab42/parmtracker/data/categories"
	"github.com/orionlab42/parmtracker/data/expenses"
	"github.com/orionlab42/parmtracker/data/notes"
	"github.com/orionlab42/parmtracker/data/users"
	"github.com/orionlab42/parmtracker/log"
	"github.com/orionlab42/parmtracker/server"
	"github.com/orionlab42/parmtracker/settings"
	"net"
	"net/http"
)

const (
	LogPrefix = "Parmtracker"
)

func initializeConfigAndLogger() {
	conf := config.GetInstance()
	logger := log.GetInstance()
	logger.SetLogFile(conf.LogFile)
	logger.SetLevel(conf.LogLevel)
}

func UpdateTablesVersion() {
	settings.UpdateSettingsTable()
	categories.UpdateCategoriesTable()
	expenses.UpdateExpensesTable()
	users.UpdateUsersTable()
	notes.UpdateNotesTable()
}

func main() {
	initializeConfigAndLogger()
	log.GetInstance().Error(LogPrefix, "Started application service")

	UpdateTablesVersion()

	// old version without certificates
	//r := server.NewRouter()
	//e := http.ListenAndServe(":12345", r)
	//if e != nil {
	//	log.GetInstance().Errorf(LogPrefix, "Error in listen and serve %s", e.Error())
	//}

	// new version with certificates
	r := server.NewRouter()
	listener, e := net.Listen("tcp", ":"+config.GetInstance().WebPort)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "Error when listening to port %d: %s", config.GetInstance().WebPort, e.Error())
	}
	e = http.ServeTLS(listener, r, "fullchain.pem", "privkey.pem")
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "Error while starting server: %s", e.Error())
	}
}
