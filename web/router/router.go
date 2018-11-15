package router

import (
	"github.com/chainHero/heroes-service/web/controllers"
	"github.com/chainHero/heroes-service/web/middleWare"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/assets", "web/assets")

	router.GET("/", controllers.HomeConroller)
	router.GET("/user", controllers.UserConroller)

	r1 := router.Group("/user")

	{
		r1.POST("/register", controllers.RegisterController)
		r1.POST("/login", controllers.LoginController)
	}

	r2 := router.Group("/vote")
	r2.Use(middleWare.JWT())
	{
		r2.POST("/", controllers.VoteHomeHandle)
		r2.GET("/", controllers.VoteHomeHandle)
		r2.POST("/v", controllers.VoteV)
		r2.POST("/insertKey", controllers.InsertKeyController)
	}
	router.Run(":8080")

}

//6048726614496508713878300278815397381906177242929895369650895101732324435675
//6829401635473773037320850146839216518456048656358077053368809942868569262482
//111117949253145538666741223747812423626384789221178933318006058448485265292409
