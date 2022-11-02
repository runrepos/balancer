package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

var Funcs map[string]interface{}

func InitServer(settings map[string]interface{}, funcs map[string]interface{}){ 
	PreSettings(settings) ;	Funcs = funcs // исполняемые функции
	app := fiber.New() // инициализация
	
	// инициализация роутов
	Init_RootRedirectUrl(app) // handlers
	Init_SettingsRoutes(app)

	log.Fatal(app.Listen(":3000"))  // запуск 
}