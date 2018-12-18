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
	"net/http"
)

func (app *Application) UpLoadKeyController() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (app *Application) KeyController() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "keyFile.html", nil)
	}
}

func (app *Application) GenKeyFileController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var privateKey *ecdsa.PrivateKey
		var err error
		username, _ := c.Cookie("user")
		isGenerateKey := models.IsGenerateKey(username)
		if !isGenerateKey {
			privateKey, err = urs.GenerateKey(elliptic.P256(), rand.Reader)
			if err != nil {
				log.Printf("%v", err)
			}
			err = models.UpdateGenerateKey(username)
			if err != nil {
				fmt.Println(err)
			}
			//c.JSON(http.StatusOK,"密钥已生成，请妥善保管")

		} else {
			c.JSON(http.StatusBadRequest, "你已经生成过密钥")
		}

		//file, _ := os.Create("keys")
		//file.WriteString(privateKey.D.String() + "\n" + privateKey.X.String() + "\n" + privateKey.Y.String())
		pKey := "privateKey:" + privateKey.D.String() + "," + "\n"
		pkX := "publicKeyX:" + privateKey.X.String() + "," + "\n"
		pkY := "publicKeyY:" + privateKey.Y.String() + "," + "\n"

		models.InsertPk(privateKey.X.String(), privateKey.Y.String())

		//s:=privateKey.D.String() + "\n" + privateKey.X.String() + "\n" + privateKey.Y.String()
		c.Writer.WriteHeader(http.StatusOK)
		c.Header("Content-Disposition", "attachment; filename=keys.txt")
		c.Header("Content-Type", "application/text/plain")
		c.Header("Accept-Length", fmt.Sprintf("%d", len([]byte(pKey+pkX+pkY))))
		c.Writer.Write([]byte(pKey + pkX + pkY))
		//c.HTML(http.StatusOK,"genKey",nil)
	}
}
