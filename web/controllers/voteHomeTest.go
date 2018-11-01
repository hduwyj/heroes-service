package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/chainHero/heroes-service/web/controllers/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var candidates = make([]*util.Candidate, 5)

func VoteHomeHandle(c *gin.Context) {

	allCandidateAsBytes, _ := util.App.Fabric.QueryAllCandidate()
	json.Unmarshal(allCandidateAsBytes, &candidates)

	op := c.Query("op")
	if op == "vote" {
		name := c.Query("name")
		for _, candidate := range candidates {
			if candidate.Name == name {
				transactionID, err := util.App.Fabric.Vote([]string{name})
				if err != nil {
					fmt.Println(err)
				}
				c.HTML(http.StatusOK, "voteSuccess.html", gin.H{
					"transactionID": transactionID,
				})
			}
		}
	} else {
		c.HTML(http.StatusOK, "voteHome.html", gin.H{
			"candidates": candidates,
		})
	}

}
