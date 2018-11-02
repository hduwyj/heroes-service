package util

import (
	"fmt"
	"github.com/chainHero/heroes-service/blockchain"
)

type Candidate struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	IdCard    string `json:"idCard"`
	Content   string `json:"content"`
	VoteCount int    `json:"voteCount"` //候选人得票数
}

type Application struct {
	Fabric *blockchain.FabricSetup
}

var App *Application

func (app *Application) GetApp() *Application {
	fmt.Println(app)
	App = app
	return App
}
