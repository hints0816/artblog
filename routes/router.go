package routes

import (
	"hello/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.GET("/:id", api.GetUserInfo)

	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run()
	// router.Run(":3000") for a hard coded port

}
