package main

import (
	"BabyBus/api"
	"BabyBus/dao"
)

func main() {
	dao.InitMysql()
	api.Init()
}
