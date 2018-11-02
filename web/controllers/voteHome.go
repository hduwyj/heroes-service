package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name      string
	Age       int
	Sex       string
	VoteCount int
}

var users = []*User{
	{"wang", 18, "male", 0},
	{"wangyujiang", 28, "female", 0},
	{"wangyujiang", 28, "female", 0},
}

func VoteHome(c *gin.Context) {
	voteName := c.Query("vote")
	for _, user := range users {
		if user.Name == voteName {
			user.VoteCount++
		}
	}
	c.HTML(http.StatusOK, "voteHome.html", gin.H{
		"users": users,
	})
}
