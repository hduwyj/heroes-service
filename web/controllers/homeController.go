package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeConroller(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}
