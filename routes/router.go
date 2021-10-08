package routes

import (
	"hello/api"
	"hello/middleware"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("index", "web/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode("debug")

	r := gin.Default()
	r.HTMLRender = createMyRender()
	r.Use(middleware.Cors())

	r.Static("/static", "./web/dist/static")
	r.StaticFile("/favicon.ico", "/web/dist/favicon.ico")

	router := r.Group("api")
	{
		// 登录控制模块
		router.POST("loginfront", api.LoginFront)
	}

	r.Run()
}
