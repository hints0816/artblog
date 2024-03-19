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
		auth.GET("blog/imgcomment/list/:id", api.GetImgCommentList)
		auth.POST("blog/comment/digg", api.SetCommentDigg)
		auth.POST("blog/comment/undigg", api.SetCommentUnDigg)
		auth.POST("blog/acticle/favour", api.SetArticleFavour)
		auth.POST("blog/acticle/unfavour", api.SetArticleUnFavour)
		auth.GET("blog/article/:id", api.GetArtInfo)

		auth.POST("blog/img/digg", api.SetFavorite)
		auth.POST("blog/img/undigg", api.SetUnFavorite)

		auth.GET("blog/me/articlelist", api.ListArticleRepository)
		auth.POST("blog/postart", api.PostImg)
		auth.DELETE("blog/artimg/:id", api.DelImgInfo)
		auth.POST("blog/comment", api.AddComment)
		auth.POST("blog/imgcomment", api.AddImgComment)

		auth.DELETE("blog/article/:id", api.DelArticle)
		auth.DELETE("blog/comment/:id", api.DelComment)

		auth.POST("blog/avatar/upload", api.UploadAvatarImage)
		auth.POST("blog/emoji", api.EditEmoji)
		auth.GET("blog/articlelist", api.ListArticle)
		auth.GET("blog/artimg/:id", api.GetImgInfo)

		auth.PUT("blog/profile", api.UpdateProfile)
		auth.PUT("blog/category", api.UpdateCategory)
		auth.DELETE("blog/category/:id", api.DeleteCategory)

		auth.GET("blog/setting/comment", api.GetCommentListFrontAdmin)

		auth.GET("blog/logout", api.LoginOut)
	}

	router := r.Group("api")
	{

		router.GET("blog/profile/:id", api.GetProfile)
		router.POST("login/users", api.Login)
		router.POST("login/sendvalidator", api.SendvalidateCode)
		// router.POST("login/checkvalidator", api.CheckvalidateCode)
		router.POST("blog/upload", api.UploadArtImage)
		router.GET("blog/category", api.ListCategory)
		router.POST("blog/signup", api.SignUpInClaim)
		router.GET("login/checktoken", api.CheckCode)
		router.GET("blog/github", api.GithubCallback)
		router.GET("blog/artlist", api.ListImage)
		router.GET("blog/webmaster", api.GetWebmasterInfo)
	}

	r.Run(":3000")
}
