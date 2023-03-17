package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func InitMysql() {
	dns := "root:pwd@/dbname?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
