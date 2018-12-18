package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"log"
)

type SmartContract struct {
}
type Candidate struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	IdCard    string `json:"idCard"`
	Content   string `json:"content"`
	VoteCount int    `json:"voteCount"` //候选人得票数
}

const (
	CANDIDATES = "CANDIDATES"
	PUBLICKEYS = "PUBLICKEYS"
	BLINDSIGN  = "BLINDSIGN"
)

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("init success"))
}

func (s *SmartContract) PutAllCandidates(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	candidatesBytes := []byte(args[0])
	err := stub.PutState(CANDIDATES, candidatesBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.SetEvent(args[1], []byte{})
	return shim.Success([]byte("putAllCandidates success"))
}

func (s *SmartContract) GetAllCandidates(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	candidatesAsBytes, err := stub.GetState(CANDIDATES)
	if err != nil {
		return shim.Error(err.Error())
	}
	log.Printf("%s", candidatesAsBytes)
	return shim.Success(candidatesAsBytes)
}

//投票
//第一个参数为投票信息，第二个参数为eventID
func (s *SmartContract) Vote(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Vote:incorrect number of arguements")
	}
	candidatesAsBytes, err := stub.GetState(CANDIDATES)
	if err != nil {
		return shim.Error(err.Error())
	}
	var candidates []Candidate
	err = json.Unmarshal(candidatesAsBytes, &candidates)
	if err != nil {
		return shim.Error(err.Error())
	}
	for i := 0; i < len(candidates); i++ {
		if candidates[i].Name == args[0] {
			candidates[i].VoteCount++
			break
		}
	}
	bytes, _ := json.Marshal(candidates)

	err = stub.PutState(CANDIDATES, bytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.SetEvent(args[1], []byte{})
	return shim.Success(bytes)

}

func (s *SmartContract) PutAllPK(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	PKBytes := []byte(args[0])
	err := stub.PutState(PUBLICKEYS, PKBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.SetEvent(args[1], []byte{})
	return shim.Success([]byte("PutAllPK success"))
}

func (s *SmartContract) GetAllPK(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	pkAsBytes, err := stub.GetState(PUBLICKEYS)
	if err != nil {
		return shim.Error(err.Error())
	}
	log.Printf("%s", pkAsBytes)
	return shim.Success(pkAsBytes)
}

func (s *SmartContract) PutSign(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	bytes, err := stub.GetState(BLINDSIGN)
	if err != nil {
		return shim.Error(err.Error())
	}
	//args[0]为签名信息 args[1]为投票信息
	rsName := args[0]

	err = stub.PutState(BLINDSIGN, []byte(rsName))
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(string(bytes))
	stub.SetEvent(args[1], []byte{})
	return shim.Success([]byte("putBlindSign success"))
}

func (s *SmartContract) GetAllSign(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	signAsBytes, err := stub.GetState(BLINDSIGN)
	if err != nil {
		return shim.Error(err.Error())
	}
	log.Printf("%s", signAsBytes)
	return shim.Success(signAsBytes)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	if funcName == "GetAllCandidates" {
		return s.GetAllCandidates(stub, args)
	} else if funcName == "PutAllCandidates" {
		return s.PutAllCandidates(stub, args)
	} else if funcName == "Vote" {
		return s.Vote(stub, args)
	} else if funcName == "PutAllPK" {
		return s.PutAllPK(stub, args)
	} else if funcName == "GetAllPK" {
		return s.GetAllPK(stub, args)
	} else if funcName == "PutSign" {
		return s.PutSign(stub, args)
	} else if funcName == "GetAllSign" {
		return s.GetAllSign(stub, args)
	}
	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
