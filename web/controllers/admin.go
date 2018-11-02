package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminControllerGet(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", nil)
}

func AdminControllerPost(c *gin.Context) {
	fmt.Println(c.PostForm("id"))
	c.HTML(http.StatusOK, "admin.html", nil)
}
