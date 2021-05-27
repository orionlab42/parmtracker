package main

import (
	"encoding/json"
	"github.com/annakallo/parmtracker/config"
	"github.com/annakallo/parmtracker/data/categories"
	"github.com/annakallo/parmtracker/data/expenses"
	"github.com/annakallo/parmtracker/log"
	"github.com/annakallo/parmtracker/server"
	"github.com/annakallo/parmtracker/settings"
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
