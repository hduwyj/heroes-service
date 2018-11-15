package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserConroller(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
