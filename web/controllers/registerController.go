package controllers

import (
	"github.com/chainHero/heroes-service/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterController(c *gin.Context) {
	var msg string
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")

	if username == "" || password == "" {
		msg = "注册失败，账户名密码不能为空"
		c.HTML(http.StatusOK, "login.html", gin.H{
			"msg": msg,
		})
	} else {
		err := models.InsertVoter(username, password)
		if err != nil {
			log.Printf("%v", err)
		}
		msg = "注册成功，请登录"
		c.HTML(http.StatusOK, "login.html", gin.H{
			"msg": msg,
		})
	}

}
