package main

import (
	//"log"
	"superapp/internal"
	"superapp/internal/handlers"
	//msql "superapp/internal/storage/mysql"
	//msql "superapp/internal/storage/mysql_gorm"
	//"superapp/config"
	//"superapp/internal/models"
)

func main() {
	app_settings := internal.InitConfig() //Config := app_settings["config"].(*config.Cfg)
	handlers.InitServer(app_settings, map[string]interface{}{})
}
