package controllers

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/chainHero/heroes-service/models"
	"github.com/chainHero/heroes-service/web/controllers/util"
	"github.com/gin-gonic/gin"
	"github.com/hduwyj/urs"
	"log"
	"net/http"
)

var candidates = make([]*util.Candidate, 5)

//func VoteHomeHandle(c *gin.Context) {
//
//	allCandidateAsBytes, _ := util.App.Fabric.QueryAllCandidate()
//	json.Unmarshal(allCandidateAsBytes, &candidates)
//
//	op := c.Query("op")
//	if op == "vote" {
//		name := c.Query("name")
//		for _, candidate := range candidates {
//			if candidate.Name == name {
//				transactionID, err := util.App.Fabric.Vote([]string{name})
//				if err != nil {
//					fmt.Println(err)
//				}
//				c.HTML(http.StatusOK, "voteSuccess.html", gin.H{
//					"transactionID": transactionID,
//				})
//			}
//		}
//	} else {
//		c.HTML(http.StatusOK, "voteHome.html", gin.H{
//			"candidates": candidates,
//		})
//	}
//
//}
func VoteHomeHandle(c *gin.Context) {
	candidates := models.Query()
	op := c.Query("op")

	switch op {
	case "vote":
		name := c.Query("name")
		//for _,candidate:=range candidates{
		//	if candidate.Name==name{
		//		candidate.VoteCount++
		//	}
		//}
		priv := c.PostForm("privateKey")
		fmt.Println("privateKey:", priv)
		for i := 0; i < len(candidates); i++ {
			if candidates[i].Name == name {
				candidates[i].VoteCount++
			}
		}
		c.HTML(http.StatusOK, "voteHome.html", gin.H{
			"candidates": candidates,
			"priv":       priv,
		})
	case "genKey":
		priv, err := urs.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			log.Printf("%v", err)
		}
		c.HTML(http.StatusOK, "voteHome.html", gin.H{
			"candidates":  candidates,
			"priv":        priv.D,
			"publicKey_x": priv.X,
			"publicKey_y": priv.Y,
		})

	default:
		c.HTML(http.StatusOK, "voteHome.html", gin.H{
			"candidates": candidates,
		})
	}

}
