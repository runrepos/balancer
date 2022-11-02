package mysql_gorm

import (
	//"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"superapp/internal/models"
	"log"
  )
  

// https://gorm.io/docs/connecting_to_the_database.html
//var ConnectStr string

type MySq struct { 
	ConStr string
	Gorm *gorm.DB }

func InitSql(ConnectString string) (*MySq, error) { 
	db := MySq{ConStr:ConnectString} 
	err := db.InitCon() ; if err != nil { return &db, err }
	return &db, nil
	} // ConnectStr = ConnectString ; 

func (me *MySq) InitCon() error {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: me.ConStr, // data source name
		// DefaultStringSize: 256, // default size for string fields
		// DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		// DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		// DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		// SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	me.Gorm = db ;	if err != nil { return err } ; return nil }

func InitDriver(sql string) (*MySq, error){ //(*gorm.DB, error) {
	driver, err := InitSql(sql) //Config.MYSQL_CONNECT_STRING)
	if err != nil { return driver, err } ;	return driver, nil ; }

// настройки, можно перефакторить

// interface
func (me *MySq) SelectAll(table string) []models.Settings { //map[string]interface{} {
	settings := []models.Settings{} //[]map[string]interface{}{}
	me.Gorm.Table(table).Find(&settings)
	return settings
}

func (me *MySq) Create(table string, insert *models.Settings) int {
	me.Gorm.Table(table).Create(&insert) ; 
	return insert.Id //log.Println("InsertedId:",new_setting.Id)
}

func (me *MySq) FindSetting(table string, update *models.Settings) (int,error) {
	rows, err := me.Gorm.Table(table).Select("tag = ?", update.Tag).Rows()
	for rows.Next() {
		return 1, err
	}
	return 0, err
}

func (me *MySq) Update(table string, update *models.Settings) int64 {
	find, err := me.FindSetting(table, update)
	if find >= 1 || err != nil{
		result := me.Gorm.Table(table).Where("tag = ?", update.Tag).Update("value", update.Value)
		log.Println(update, table, result, result.RowsAffected)
		if result.Error != nil {log.Println(result.Error) }
		return 1
	}
	return 0
}

func (me *MySq) UpdateOrCreate(table string, update *models.Settings) (int64, int) {
	try := me.Update(table, update)
	if try < 1 { 
		insertedId := me.Create(table, update)
		return try, insertedId
	}
	return try, 0
}

