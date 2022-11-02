package config

import (
		"os"
		"time"
		"strconv"
		"fmt"
)

type Cfg struct {
	Local *time.Location
	MYSQL_USER string
	MYSQL_HOST string
	MYSQL_PORT string
	MYSQL_DATABASE string
	MYSQL_PASSWORD string
	MYSQL_CONNECT_STRING string

	HISTORY_BALANCE_PERCENT float64
	BALANCE_AUTOMATICALLY_FORGOTTEN float64
	BALANCE_AUTOMATICALLY_FORGETTEN_TIME float64
}
// export $(grep -v '^#' .env | xargs) ; printenv

var Config Cfg
func InitConfig(){
	Config = Cfg{
		//Local: Local, //dt.Local,
		MYSQL_HOST: os.Getenv("MYSQL_HOST"),
		MYSQL_PORT: os.Getenv("MYSQL_PORT"),
		MYSQL_DATABASE: os.Getenv("MYSQL_DATABASE"),
		MYSQL_USER: "root",
		MYSQL_PASSWORD: os.Getenv("MYSQL_ROOT_PASSWORD"),
	}
	Config.MYSQL_CONNECT_STRING = fmt.Sprintf( 
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.MYSQL_USER, 
		Config.MYSQL_PASSWORD, 
		Config.MYSQL_HOST, 
		Config.MYSQL_PORT,
		Config.MYSQL_DATABASE,
	) 
	n1, err := strconv.ParseFloat(os.Getenv("HISTORY_BALANCE_PERCENT_GLOBAL"), 32)
	Config.HISTORY_BALANCE_PERCENT = n1 ; if err!= nil { fmt.Println(n1, err ,"HISTORY_BALANCE_PERCENT_GLOBAL") }
	n1, err = strconv.ParseFloat(os.Getenv("BALANCE_AUTOMATICALLY_FORGOTTEN"), 32)
	Config.BALANCE_AUTOMATICALLY_FORGOTTEN = n1 ; if err!= nil { fmt.Println(n1, err ,"BALANCE_AUTOMATICALLY_FORGOTTEN") }
	n1, err = strconv.ParseFloat(os.Getenv("BALANCE_AUTOMATICALLY_FORGETTEN_TIME_INTERVAL"), 32)
	Config.BALANCE_AUTOMATICALLY_FORGETTEN_TIME = n1 ; if err!= nil { fmt.Println(n1, err , "BALANCE_AUTOMATICALLY_FORGETTEN_TIME_INTERVAL") }
	fmt.Println(Config.MYSQL_CONNECT_STRING )
}
