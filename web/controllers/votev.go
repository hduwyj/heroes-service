package controllers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/chainHero/heroes-service/allType"
	"github.com/gin-gonic/gin"
	"github.com/hduwyj/urs"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"regexp"
	"strings"
)

func (app *Application) VoteV() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, _ := c.GetPostForm("name")

		//name := "wangyujiang"
		fmt.Println("------------", name)

		privateKey, publicKeyRing := getPublicKeyRingAndPrivateKey(app, c)
		rs, err := urs.Sign(rand.Reader, privateKey, publicKeyRing, []byte(name))
		if err != nil {
			log.Printf("%s", err)
		}

		//TODO 智能合约验证签名的有效性
		//验证签名的有效性
		b := urs.Verify(publicKeyRing, []byte(name), rs)

		if b {
			s := rs.ToBase58()
			isVoted := app.Fabric.IsVoted(s)
			if isVoted {
				c.JSON(http.StatusBadRequest, "已经投票")
			} else {
				_, err = app.Fabric.Vote([]byte(name))

				if err != nil {
					log.Printf("%v", err)
				}
				//将所有候选人的签名 hx 全部提交到区块链上
				err := app.putAllHx(s, name, privateKey, publicKeyRing)
				if err != nil {
					log.Printf("%v", err)
				}
				c.Redirect(http.StatusPermanentRedirect, "/vote")
			}

		} else {
			c.JSON(http.StatusBadRequest, "投票失败")
		}

	}
}

//将所有的候选人签名，并将签名信息的X上传至区块链，防止用户多次投票
func (app *Application) putAllHx(ringSign, name string, pk *ecdsa.PrivateKey, pkr *urs.PublicKeyRing) error {
	candidateAsBytes, err := app.Fabric.QueryAllCandidate()
	if err != nil {
		return err
	}
	builder := strings.Builder{}
	builder.WriteString(ringSign + "_" + name)
	var candidates []allType.Candidate
	json.Unmarshal(candidateAsBytes, &candidates)
	for _, c := range candidates {
		rs, _ := urs.Sign(rand.Reader, pk, pkr, []byte(c.Name))
		builder.WriteString("_" + rs.X.String())
	}
	_, err = app.Fabric.PutRingSign(builder.String())
	if err != nil {
		return err
	}
	return nil

}

//获取密钥信息和公钥环
func getPublicKeyRingAndPrivateKey(app *Application, c *gin.Context) (*ecdsa.PrivateKey, *urs.PublicKeyRing) {
	file, _, err := c.Request.FormFile("vote")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败"})
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
	}

	privateKeyD, publicKeyX, publicKeyY, err := getPrivateKeyString(content)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	D := new(big.Int)
	D.SetString(privateKeyD, 10)
	X := new(big.Int)
	X.SetString(publicKeyX, 10)
	Y := new(big.Int)
	Y.SetString(publicKeyY, 10)
	fmt.Println(D, X, Y)
	priv := &ecdsa.PrivateKey{ecdsa.PublicKey{elliptic.P256(), X, Y}, D}
	publicKeyRing := urs.NewPublicKeyRing(100)
	publicKeyRing.Add(priv.PublicKey)
	//从区块链上读取公钥信息
	bytes, err := app.Fabric.QueryAllPK()
	if err != nil {
		log.Println(err)
	}
	var pk []ecdsa.PublicKey
	err = json.Unmarshal(bytes, &pk)
	if err != nil {
		log.Println(err)
	}
	//因为elliptic.P256不能被json解析，故手动加上
	for i := 0; i < len(pk); i++ {
		pk[i].Curve = elliptic.P256()
	}
	for _, k := range pk {
		publicKeyRing.Add(k)
	}
	return priv, publicKeyRing
}

func getPrivateKeyString(content []byte) (d, x, y string, err error) {
	c := regexp.MustCompile("[^:][0-9]*[,$]")
	allString := c.FindAllString(string(content), -1)
	err = nil
	if len(allString) != 3 {
		err = fmt.Errorf("%s", "密钥格式错误")
	}
	d = strings.TrimRight(allString[0], ",")
	x = strings.TrimRight(allString[1], ",")
	y = strings.TrimRight(allString[2], ",")
	return
}
