package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/chainHero/heroes-service/allType"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (app *Application) VoteHomeHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		//username, _ := c.Cookie("user")

		//candidates := models.Query()
		candidatesAsBytes, err := app.Fabric.QueryAllCandidate()
		if err != nil {
			log.Printf("%v", err)
		}
		var candidates []allType.Candidate
		err = json.Unmarshal(candidatesAsBytes, &candidates)
		if err != nil {
			log.Printf("%v", err)
		}
		for _, c := range candidates {
			fmt.Println(c.VoteCount)
		}
		//op := c.Query("op")
		//
		//switch op {
		//case "vote":
		//	name := c.Query("name")
		//	bytes, err := app.Fabric.Vote([]byte(name))
		//	if err!=nil{
		//		fmt.Println(err)
		//	}
		//	json.Unmarshal(bytes,&candidates)
		//	c.HTML(http.StatusOK, "voteHome.html", gin.H{
		//		"candidates": candidates,
		//	})
		//case "genKey":
		//	isGenerateKey := models.IsGenerateKey(username)
		//	if !isGenerateKey{
		//		priv, err := urs.GenerateKey(elliptic.P256(), rand.Reader)
		//		if err != nil {
		//			log.Printf("%v", err)
		//		}
		//		err = models.UpdateGenerateKey(username)
		//		if err!=nil{
		//			fmt.Println(err)
		//		}
		//		c.HTML(http.StatusOK, "voteHome.html", gin.H{
		//			"candidates":  candidates,
		//			"priv":        priv.D,
		//			"publicKey_x": priv.X,
		//			"publicKey_y": priv.Y,
		//		})
		//	}else {
		//		c.JSON(http.StatusBadRequest,"你已经生成过密钥")
		//	}
		//
		//
		//default:
		//	c.HTML(http.StatusOK, "voteHome.html", gin.H{
		//		"candidates": candidates,
		//	})
		//}

		c.HTML(http.StatusOK, "voteHome.html", gin.H{
			"candidates": candidates,
		})

	}

}
