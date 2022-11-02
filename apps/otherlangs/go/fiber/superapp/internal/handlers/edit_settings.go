package handlers

import (
	"github.com/gofiber/fiber/v2"
	"superapp/internal/models"
	"encoding/json"
	"log"
)

// обзновление настроек
func PreloadSettings_And_Update(){
	settings := Storage.Mysql.SelectAll("settings")
	repack_settings := Repack(settings)

	for _, el := range repack_settings {
		if el.Tag == "VIDEO:CDN" {
			SetSettings(el.Value[0].Host, 
				el.Value[0].Percent,
				Config.HISTORY_BALANCE_PERCENT) // пока из настроек .env
			break
		}
	}

}

func Init_SettingsRoutes(app *fiber.App) { //go service.SomeWorker(); 
	PreloadSettings_And_Update() // при старте ⭐
	
    app.Get("/api/get_settings", GetSettings) 
	app.Post("/api/edit_settings", EditSettings) } // some events

func Repack(input []models.Settings) []models.SettingsUnpack {
	new_settings := []models.SettingsUnpack{}

	for _,el := range(input) {
		valueUnpacked := []models.SettingsValueUnpacked{}  
		err := json.Unmarshal([]byte(el.Value), &valueUnpacked) ; 
		if err != nil { log.Println(err) } //else{
			new_settings=append(new_settings, models.SettingsUnpack{
				Tag: el.Tag,
				Value: valueUnpacked,
			})
		//}
	}
	return new_settings
}
func RepackBack(input []models.SettingsUnpack) []models.Settings {
	new_settings := []models.Settings{}

	for _,el := range(input) {
		valuePacked, err := json.Marshal(el.Value) ; 
		if err != nil { log.Println(err) } //else{
			new_settings=append(new_settings, models.Settings{
				Tag: el.Tag,
				Value: string(valuePacked),
			})
		//}
	}
	return new_settings
}
func RepackBackOne(input models.SettingsUnpack) models.Settings {
	return RepackBack([]models.SettingsUnpack{input})[0] }
func RepackOne(input models.Settings) models.SettingsUnpack {
	return Repack([]models.Settings{input})[0] }


func GetSettings(c *fiber.Ctx) error {
	settings := Storage.Mysql.SelectAll("settings")
	repack_settings := Repack(settings)
	return c.Status(fiber.StatusOK).JSON( &repack_settings)
}

func EditSettings(c *fiber.Ctx) error {
	new_setting := models.SettingsUnpack{}
	if err := c.BodyParser(&new_setting); err != nil {
		return c.Status(503).SendString(err.Error()) }    
	toSave := RepackBackOne(new_setting)
	updatedRows, insertedId := Storage.Mysql.UpdateOrCreate("settings", &toSave)
	if insertedId < 1 && updatedRows<1 { return c.Status(500).SendString("Fail")  
	} else {
		PreloadSettings_And_Update() // При обновлении загружать настройки ⭐
	}
	if insertedId > 0 { return c.Status(201).JSON( map[string]int{"insertedId": insertedId, }) } 
	return c.Status(201).JSON( map[string]int64{"updatedCount": updatedRows, })
}
