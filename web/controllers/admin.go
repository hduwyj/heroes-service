package controllers

import (
	"github.com/chainHero/heroes-service/allType"
	"github.com/chainHero/heroes-service/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (app *Application) AdminControllerGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		op := c.Query("op")
		if op == "query" {
			cs := models.Query()
			keys, _ := models.GetAllPk()
			c.HTML(http.StatusOK, "cadidates.html", gin.H{
				"title":      "美国总统选举",
				"candidates": cs,
				"keys":       keys,
			})
			return
		} else if op == "del" {
			name := c.Query("name")
			err := models.DeleteCandidate(name)
			if err != nil {
				log.Fatal(err)
			}
			cs := models.Query()
			c.HTML(http.StatusOK, "cadidates.html", gin.H{
				"candidates": cs,
			})
			return
		}
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"title": "投票",
		})
	}
}

func (app *Application) AdminControllerPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.PostForm("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.HTML(http.StatusNotFound, "404.html", nil)
			return
		}
		name := c.PostForm("name")
		gender := c.PostForm("gender")
		content := c.PostForm("content")
		idCard := c.PostForm("idCard")
		candidate := allType.Candidate{idInt, name, gender, idCard, content, 0}

		//插入到数据库
		err = models.InsertCandidate(candidate)
		if err != nil {

			log.Fatal(err)
		}
		c.HTML(http.StatusOK, "admin.html", nil)
	}

}
