package controllers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/chainHero/heroes-service/models"
	"github.com/gin-gonic/gin"
	"github.com/hduwyj/urs"
	"log"
	"math/big"
	"net/http"
)

func VoteV(c *gin.Context) {
	name, _ := c.GetPostForm("name")
	privateKey, publicKeyRing := getPublicKeyRingAndPrivateKey(c)

	voteCount := 0
	cs := models.Query()
	for _, v := range cs {
		if v.Name == name {
			voteCount = v.VoteCount
			break
		}
	}
	rs, err := urs.BlindSign(rand.Reader, privateKey, publicKeyRing, []byte(name))
	if err != nil {
		log.Printf("%s", err)
	}
	b := urs.BlindVerify(publicKeyRing, []byte(name), rs)
	fmt.Println(b)
	if b == true {
		err := models.UpdateCandidate(voteCount, name)
		if err != nil {
			log.Printf("%v", err)
		}
		c.Redirect(http.StatusPermanentRedirect, "/vote")
	} else {
		c.JSON(http.StatusBadRequest, "投票失败")
	}

}

func getPublicKeyRingAndPrivateKey(c *gin.Context) (*ecdsa.PrivateKey, *urs.PublicKeyRing) {
	privateKeyD, _ := c.GetPostForm("privateKey")
	publicKeyX, _ := c.GetPostForm("publicKeyX")
	publicKeyY, _ := c.GetPostForm("publicKeyY")
	D := new(big.Int)
	D.SetString(privateKeyD, 10)
	X := new(big.Int)
	X.SetString(publicKeyX, 10)
	Y := new(big.Int)
	Y.SetString(publicKeyY, 10)
	priv := &ecdsa.PrivateKey{ecdsa.PublicKey{elliptic.P256(), X, Y}, D}
	publicKeyRing := urs.NewPublicKeyRing(100)
	pk, _ := models.GetAllPk()
	for _, k := range pk {
		publicKeyRing.Add(k)
	}
	return priv, publicKeyRing
}
