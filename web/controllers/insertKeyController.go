package controllers

import (
	"github.com/gin-gonic/gin"
	"jwtAuth/model"
	"net/http"
)

func InsertKeyController(c *gin.Context) {

	publicKey_x, _ := c.GetPostForm("publicKeyX")
	publicKey_y, _ := c.GetPostForm("publicKeyY")
	err := model.InsertPk(publicKey_x, publicKey_y)
	if err != nil {
		c.JSON(http.StatusBadRequest, "插入失败")
	}
	c.Redirect(http.StatusPermanentRedirect, "/vote")
	//}
}
