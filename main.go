package main

import (
	"hello/model"

	"hello/routes"
)

func main() {
	model.InitDb()
	model.InitMinio()
	model.InitQiniu()
	routes.InitRouter()
}
