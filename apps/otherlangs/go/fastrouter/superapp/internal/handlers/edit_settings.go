package handlers

import (
	//"github.com/gofiber/fiber/v2"
	"net/http"
	"github.com/razonyang/fastrouter"
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

func Init_SettingsRoutes(app *fastrouter.Router){ //app string) { //go service.SomeWorker(); 
	PreloadSettings_And_Update() // при старте ⭐
	app.Get("/api/get_settings", http.HandlerFunc(
		GetSettings,
	)) 
	app.Post("/api/edit_settings", http.HandlerFunc(
		EditSettings,
	)) 
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("query users"))
	// })

    // app.Get("/api/get_settings", GetSettings) 
	// app.Post("/api/edit_settings", EditSettings) 
} // some events

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


func GetSettings(w http.ResponseWriter, r *http.Request) { //c *fiber.Ctx) error {
	settings := Storage.Mysql.SelectAll("settings")
	repack_settings := Repack(settings)

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(repack_settings)
	//return c.Status(fiber.StatusOK).JSON( &repack_settings)
}

func EditSettings(w http.ResponseWriter, r *http.Request){ //c *fiber.Ctx) error {
	new_setting := models.SettingsUnpack{}
	// if err := c.BodyParser(&new_setting); err != nil {
	// 	//return c.Status(503).SendString(err.Error()) 
	// 	}    
    err := json.NewDecoder(r.Body).Decode(&new_setting)
    if err != nil { 
		http.Error(w, err.Error(), http.StatusBadRequest) ;   return
    }
	toSave := RepackBackOne(new_setting)
	updatedRows, insertedId := Storage.Mysql.UpdateOrCreate("settings", &toSave)
	if insertedId < 1 && updatedRows<1 { 
		//eturn c.Status(500).SendString("Fail")  
		http.Error(w, "Fail", http.StatusBadRequest) ; return
	} else {
		PreloadSettings_And_Update() // При обновлении загружать настройки ⭐
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if insertedId > 0 { 

		json.NewEncoder(w).Encode(map[string]int{"insertedId": insertedId, }) ; return
		//http.Error(w, "Ok Inserted", http.StatusBadRequest) ;
		//return c.Status(201).JSON( ) 
	} 
	json.NewEncoder(w).Encode(map[string]int64{"updatedCount": updatedRows, }) ; return
	//return c.Status(201).JSON( map[string]int64{"updatedCount": updatedRows, })
}
