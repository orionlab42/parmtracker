package config

import (
	"encoding/json"
	"github.com/annakallo/parmtracker/log"
	"os"
	"sync"
)

const (
	LogPrefix  = "Config"
	File       = "config.json"
	FileCustom = "config.custom.json"
)

type Config struct {
	log         *log.Logger
	LogLevel    int
	LogFile     string
	MysqlIP     string
	MysqlPort   string
	MysqlUser   string
	MysqlPass   string
	MysqlDB     string
	WebPort     uint
	WebPrefix   string
	WebUsername string
	WebPassword string
	Client      string
	Static      string
	Template    string
}

var instance *Config
var once sync.Once

// Get config and subscribe them with flags
func NewConfig() *Config {
	// Get defaults
	config := &Config{}
	config.log = log.GetInstance()

	// Get configuration from configuration file
	file, err := os.Open(FileCustom)
	if err != nil {
		file, err = os.Open(File)
		if err != nil {
			config.log.Error(LogPrefix, "Configuration file not found")
			os.Exit(1)
		} else {
			config.log.Debug(LogPrefix, "Using default configuration file")
		}
	} else {
		config.log.Debug(LogPrefix, "Using custom configuration file")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		config.log.Error(LogPrefix, "Syntax error in configuration file")
		os.Exit(1)
	}
	config.log.Debug(LogPrefix, "Configuration loaded")
	return config
}

// Transforming Config into a Singleton
func GetInstance() *Config {
	once.Do(func() {
		instance = NewConfig()
	})
	return instance
}
