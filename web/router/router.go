package router

import (
	"github.com/chainHero/heroes-service/web/controllers"
	"github.com/chainHero/heroes-service/web/middleWare"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *controllers.Application) {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/assets", "web/assets")

	router.GET("/", app.HomeController())
	router.GET("/user", app.UserController())

	router.GET("/admin", app.AdminControllerGet())
	router.POST("/admin", app.AdminControllerPost())
	router.GET("/admin/test", app.AdminPushController())
	r1 := router.Group("/user")

	{
		r1.POST("/register", app.RegisterController())
		r1.POST("/login", app.LoginController())
	}

	r2 := router.Group("/vote")
	r2.Use(middleWare.JWT())
	{
		r2.POST("/", app.VoteHomeHandle())
		r2.POST("/prepare", app.PrepareController())
		r2.GET("/", app.VoteHomeHandle())
		r2.GET("/key", app.KeyController())
		r2.GET("/genKey", app.GenKeyFileController())

		r2.POST("/v", app.VoteV())
		//r2.POST("/insertKey", app.GenKeyFileController())
	}
	router.Run(":8080")

}
