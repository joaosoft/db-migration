package services

import (
	"fmt"

	manager "github.com/joaosoft/manager"
	"github.com/labstack/gommon/log"
)

// AppConfig ...
type AppConfig struct {
	DbMigration WatcherConfig `json:"dbmigration"`
}

// WatcherConfig ...
type WatcherConfig struct {
	Host string           `json:"host"`
	Path string           `json:"path"`
	Db   manager.DBConfig `json:"db"`
	Log  struct {
		Level string `json:"level"`
	} `json:"log"`
}

// NewConfig ...
func NewConfig(host string, db manager.DBConfig) *WatcherConfig {
	appConfig := &AppConfig{}
	if _, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig); err != nil {
		log.Error(err.Error())
	}

	appConfig.DbMigration.Host = host

	return &appConfig.DbMigration
}
