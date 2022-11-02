package internal

import (
	"superapp/config"
	"superapp/internal/storage"
)

type App struct {
	Config *config.Cfg
	Storage *storage.Storage
}

func InitConfig() map[string]interface{} { 
	config.InitConfig() 
	storage:=storage.InitStorage(&config.Config)
	return map[string]interface{}{
		"config": &config.Config,
		"storage": storage,
	}
}