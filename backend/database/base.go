package database

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func connect() *gorm.DB {
	conn, err := gorm.Open("mysql", "root@tcp(db:3306)/default?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		return connect()
	}
	return conn
}

func init() {
	db = connect()
	db.Debug().AutoMigrate(&User{}, &Equipment{})
}

func GetDB() *gorm.DB {
	return db
}
