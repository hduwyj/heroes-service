package controllers

import (
	"github.com/chainHero/heroes-service/web/middleWare"
	"github.com/gin-gonic/gin"
	"jwtAuth/model"
	"log"
	"net/http"
)

func LoginController(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	log.Printf("username:%s", username)
	password, _ := c.GetPostForm("password")
	log.Printf("password:%s", password)

	//验证
	if model.IsValidateVoter(username, password) {
		token, _ := middleWare.GenerateToken(username, password)
		c.SetCookie("token", token, 1000000, "/", "", false, false)
		c.Redirect(http.StatusTemporaryRedirect, "/vote")
	} else {
		msg := "用户名或者密码错误，请重新登录"
		c.HTML(http.StatusOK, "login.html", gin.H{
			"msg": msg,
		})
	}

}
