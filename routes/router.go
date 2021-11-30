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
		auth.GET("blog/getme", api.GetMe)
		auth.POST("blog/article", api.AddArticle)
		auth.GET("blog/comment/list/:id", api.GetCommentListFront)
		auth.POST("blog/comment/digg", api.SetCommentDigg)
		auth.POST("blog/comment/undigg", api.SetCommentUnDigg)

		auth.POST("blog/comment", api.AddComment)
	}

	router := r.Group("api")
	{
		router.GET("blog/articlelist", api.ListArticle)
		router.POST("user/add", api.AddUser)

		// router.GET("blog/comment/:id", api.GetCommentListFront)

		router.GET("blog/profile/:id", api.GetProfile)
		router.GET("blog/article/:id", api.GetArtInfo)
		router.POST("login/users", api.Login)
	}

	r.Run(":3000")
}
