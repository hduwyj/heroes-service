package router

import (
	"fmt"
	"github.com/chainHero/heroes-service/web/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/assets", "web/assets")

	//router.GET("/test",controllers.VoteHome)
	router.GET("/home", controllers.VoteHomeHandle)
	router.GET("/admin", controllers.AdminControllerGet)
	router.POST("/admin", controllers.AdminControllerPost)
	//router.GET("/",controllers.AdminController)
	fmt.Println("127.0.0.1:8080")
	router.Run(":8080")

}
