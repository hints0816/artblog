package main

import (
	"hello/model"

	"hello/routes"
)

func main() {
	model.InitDb()
	model.InitMinio()
	routes.InitRouter()
}
