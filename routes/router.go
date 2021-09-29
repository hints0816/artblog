package routes

import (
	"hello/api"
	"hello/middleware"
	"net/http"

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

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Main website",
		})
	})

	router := r.Group("api")
	{
		router.GET("user/:id", api.GetUserInfo)
	}

	r.Run()
}
