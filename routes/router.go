package routes

import (
	"hello/api"
	"hello/middleware"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("front", "blog/dist/spa/index.html")
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

		auth.GET("blog/me/articlelist", api.ListArticleRepository)
		auth.POST("blog/comment", api.AddComment)
		auth.DELETE("blog/article/:id", api.DelArticle)
		auth.POST("blog/avatar/upload", api.UploadAvatarImage)
		auth.POST("blog/emoji", api.EditEmoji)
		auth.GET("blog/articlelist", api.ListArticle)

		auth.PUT("blog/profile", api.UpdateProfile)
	}

	router := r.Group("api")
	{

		router.GET("blog/profile/:id", api.GetProfile)
		router.GET("blog/article/:id", api.GetArtInfo)
		router.POST("login/users", api.Login)
		router.POST("login/sendvalidator", api.SendvalidateCode)
		// router.POST("login/checkvalidator", api.CheckvalidateCode)
		router.POST("blog/upload", api.UploadArtImage)
		router.GET("blog/category", api.ListCategory)
		router.POST("blog/signup", api.SignUpInClaim)
		router.GET("login/checktoken", api.CheckCode)
	}

	r.Run(":3000")
}
