package controllers

import (
	"github.com/chainHero/heroes-service/models"
	"github.com/chainHero/heroes-service/web/middleWare"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (app *Application) UserController() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func (app *Application) RegisterController() gin.HandlerFunc {
	return func(c *gin.Context) {
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

}

func (app *Application) LoginController() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _ := c.GetPostForm("username")
		log.Printf("username:%s", username)
		password, _ := c.GetPostForm("password")
		log.Printf("password:%s", password)

		//验证
		if models.IsValidateVoter(username, password) {
			token, _ := middleWare.GenerateToken(username, password)
			c.SetCookie("token", token, 1000000, "/", "", false, false)
			c.SetCookie("user", username, 1000000, "/", "", false, false)
			c.Redirect(http.StatusTemporaryRedirect, "/vote/prepare")
		} else {
			msg := "用户名或者密码错误，请重新登录"
			c.HTML(http.StatusOK, "login.html", gin.H{
				"msg": msg,
			})
		}

	}
}

func (app *Application) PrepareController() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "prepare.html", nil)

	}
}
