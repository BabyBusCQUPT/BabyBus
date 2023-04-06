package dao

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)

func InitMysql() {
	dns := "root:pwd@/dbname?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	DB = db
}

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
	_, err := RDB.Ping().Result()
	if err != nil {
		log.Fatal(err)
		return
	}
}
