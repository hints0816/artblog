package main

import (
	"hello/model"
	"hello/routes"
)

func main() {
	model.InitDb()
	model.InitMinio()
	model.InitQiniu()
	model.InitCronTime()
	routes.InitRouter()
}
