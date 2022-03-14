package main

import (
	"encoding/json"
	"github.com/orionlab42/parmtracker/config"
	"github.com/orionlab42/parmtracker/data/categories"
	"github.com/orionlab42/parmtracker/data/expenses"
	"github.com/orionlab42/parmtracker/data/users"
	"github.com/orionlab42/parmtracker/log"
	"github.com/orionlab42/parmtracker/server"
	"github.com/orionlab42/parmtracker/settings"
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
}

func main() {
	initializeConfigAndLogger()
	log.GetInstance().Error(LogPrefix, "Started application service")

	UpdateTablesVersion()

	r := server.NewRouter()
	e := http.ListenAndServe(":12345", r)
	// @TODO needs to be tested this error handling
	responseJson, _ := json.Marshal(e)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "Error in listen and serve %s", string(responseJson))
	}
}
