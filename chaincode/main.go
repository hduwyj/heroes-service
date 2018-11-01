package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Voter struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	IdCard  string `json:"idCard"`
	IsVoted bool   `json:"isVoted"` //投票人是否已经投票,每人一票
}

type Candidate struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	IdCard    string `json:"idCard"`
	Content   string `json:"content"`
	VoteCount int    `json:"voteCount"` //候选人得票数
}

//数据从数据库导出？？？？？？？？？？
var voters = []*Voter{
	{1, "zhangsan", "123456", false},
	{2, "lisi", "123456", false},
	{3, "wangwu", "123456", false},
}

var candidates = []*Candidate{
	{1, "奥巴马", "男", "123456789", "请投奥巴马一票", 0},
	{2, "特朗普", "男", "123456789", "请投特朗普一票", 0},
	{3, "希拉里", "女", "123456789", "请投希拉里一票", 0},
}

type SmartContract struct {
}

const (
	VOTER      = "voter"
	CANDIDATES = "candidate"
)

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	voterAsBytes, err := json.Marshal(voters)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(VOTER, voterAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	candidatesAsBytes, err := json.Marshal(candidates)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(CANDIDATES, candidatesAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("Init Success"))
}
func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	if funcName == "queryAllCandidates" {
		return s.queryAllCandidates(stub, args)
	} else if funcName == "vote" {
		return s.vote(stub, args)
	}
	return shim.Error("Invoke Failed")
}

func (s *SmartContract) queryAllCandidates(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	candidatesAsBytes, err := stub.GetState(CANDIDATES)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(candidatesAsBytes)
}
func (s *SmartContract) vote(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	voteName := args[0]
	if len(args) != 1 {
		return shim.Error("vote:Incorrect number of arguments!")
	}
	candidatesAsBytes, err := stub.GetState(CANDIDATES)
	if err != nil {
		return shim.Error(err.Error())
	}
	cs := make([]*Candidate, 5)
	json.Unmarshal(candidatesAsBytes, &cs)
	for _, c := range cs {
		if c.Name == voteName {
			c.VoteCount++
		}
	}
	candidatesAsBytes, err = json.Marshal(cs)
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(CANDIDATES, candidatesAsBytes)
	fmt.Println(string(candidatesAsBytes))
	stub.SetEvent("eventInvoke", nil)
	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
