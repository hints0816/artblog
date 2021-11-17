package routes

import (
	"hello/api"
	"hello/middleware"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("front", "./blog/dist/spa/index.html")
	return p
}

func InitRouter() {
	gin.SetMode("debug")

	r := gin.Default()
	r.HTMLRender = createMyRender()
	r.Use(middleware.Cors())

	r.Static("/css", "./blog/dist/spa/css")
	r.Static("/fonts", "./blog/dist/spa/fonts")
	r.Static("/icons", "./blog/dist/spa/icons")
	r.Static("/js", "./blog/dist/spa/js")
	r.StaticFile("/favicon.ico", "./blog/dist/spa/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	auth := r.Group("api")
	auth.Use(middleware.JwtToken())
	{
		auth.GET("admin/users", api.GetUsers)
	}

	router := r.Group("api")
	{
		// 登录控制模块
		router.POST("login", api.Login)
		router.POST("loginfront", api.LoginFront)
		router.POST("article/add", api.AddArticle)

		router.GET("article/info/:id", api.GetArtInfo)
	}

	r.Run(":3000")
}
