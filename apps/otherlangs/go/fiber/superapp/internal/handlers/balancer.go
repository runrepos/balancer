package handlers

import (
	"github.com/gofiber/fiber/v2"
	"superapp/internal/service/redirect_url"
)

func Init_RootRedirectUrl(app *fiber.App) { //go service.SomeWorker(); 
	
	go redirect_url.WorkerBlanaceForgotten() // start worker thread

    app.Get("/", Go_RootRedirectUrl) } // some events

func Go_RootRedirectUrl(c *fiber.Ctx) error {

	url := c.Query("video") ; // log.Println(url)
	url, err := redirect_url.RootRedirect(url) ; if err != nil { return c.Status(404).SendString(err.Error()) } // log.Println(url)
	c.Redirect(url, 301) ; return nil 	} // перенаправление 
	
	// return c.Status(fiber.StatusOK).JSON( &url )
// interface
func SetSettings(cdn_host string, percent float64, history_balance float64){
	redirect_url.SetSettings(cdn_host, percent, history_balance) }