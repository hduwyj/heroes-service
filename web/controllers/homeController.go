package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *Application) HomeController() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	}
}
