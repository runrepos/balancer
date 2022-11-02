package redirect_url

import (
	"fmt"
	"log"
	"superapp/internal/utils/tools"
	"math/rand"
	"strings"
	"errors"
	"time"
)

var CDN_HOST="http://localhost:3200"
var SERVER_ID="s1"
var REDIRECT_PERCENT = .29  // %
var HISTORY_BALANCE_PERCENT = .2 // функция балансировки истории, при изменении %
var BALANCE_AUTOMATICALLY_FORGOTTEN = .7 // %
var BALANCE_AUTOMATICALLY_FORGETTEN_TIME = 5 // seconds

var total_requests=[]float64{0,0}

func WorkerBlanaceForgotten(){
	for {
		time.Sleep(time.Duration(BALANCE_AUTOMATICALLY_FORGETTEN_TIME) * time.Second) // pause
		CorrectBalance() ; log.Println("Correct balance", 
									 BALANCE_AUTOMATICALLY_FORGOTTEN, 
									 total_requests[0]+total_requests[1] ) // сокращение истории
	}
}

func CorrectBalance(){ //HISTORY_BALANCE_PERCENT float64){ // понижение накомленного количества
	total_requests[1] *= BALANCE_AUTOMATICALLY_FORGOTTEN 
	total_requests[0] *= BALANCE_AUTOMATICALLY_FORGOTTEN
}

// Функция баланси
func SetSettings(cdn_host string, 
				 percent float64, 
				 history_balance float64){
	REDIRECT_PERCENT = percent*1.05
	if percent< 0.2 { REDIRECT_PERCENT = percent*1.2}
	CDN_HOST = cdn_host
	HISTORY_BALANCE_PERCENT = history_balance // update
	total_requests[1] *= HISTORY_BALANCE_PERCENT ; total_requests[0]*= HISTORY_BALANCE_PERCENT //.2 // fix % пропорции
}

// счетчик %
func MaxRequstsPercent() bool {
	total_requests[0] += 1
	rand.Seed(42)
	var flag = false
	err := tools.Try_catch(func(){
		r := float64(rand.Intn(100000)/100000)
		if r < REDIRECT_PERCENT && total_requests[1]/total_requests[0] < REDIRECT_PERCENT {
			total_requests[1]+=1
			flag = true
        }
	}) ; if err != nil { log.Println("MaxRequstsPercent", err) } ; return flag }

// парсинг url
func ParseRequestUrl(url *string) (string, string) {
	if strings.Contains(*url, "/video/") {
		splited_url := strings.Split(*url, "/video/")
		server_id := strings.Split(strings.Split(splited_url[0],"://")[1],".")[0]
		if len(splited_url[1]) > 1 { return server_id, splited_url[1] }
	}
	return "","" }

// формат полученного угол
func FormatVideoRedirectUrl(redirect_url string) (string, error) {
    server_id, right_part := ParseRequestUrl(&redirect_url)
    if len(right_part)<3 { return "", errors.New("low len right part") } // Not found
 	res := fmt.Sprintf("%s/%s/video/%s", CDN_HOST, server_id, right_part) ; return res, nil } // ok

// редирект
func RootRedirect(url string) (string, error) {
	if len(url)<3 { return url, nil } // not found
	if MaxRequstsPercent() {
		url, err := FormatVideoRedirectUrl(url) ; 
		if err != nil { return url, err } ; return url, nil
	} ; return url, nil
}

// api редирект
// func RootRedirectUrl(c *fiber.Ctx) error {
// 	url := c.Query("video")
// 	log.Println(url)
// 	url, err := RootRedirect(url) ; if err != nil { return c.Status(404).SendString(err.Error()) } 
// 	log.Println(url)
// 	c.Redirect(url, 301) ; return nil
// 	//return c.Status(fiber.StatusOK).JSON( &url )
// }
