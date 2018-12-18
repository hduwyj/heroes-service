package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("2", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

//删除了fabric/vendor/golang.org/x/net/trace包
func TestChaincode(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("ex01", scc)
	checkInit(t, stub, nil)
	PutBlindSign := "PutBlindSign"
	checkInvoke(t, stub, [][]byte{[]byte(PutBlindSign), []byte("wangyu"), []byte("wangyu")})
	checkInvoke(t, stub, [][]byte{[]byte(PutBlindSign), []byte("jiang"), []byte("jiang")})
	checkInvoke(t, stub, [][]byte{[]byte(PutBlindSign), []byte("wangyu"), []byte("wangyu")})
	checkInvoke(t, stub, [][]byte{[]byte(PutBlindSign), []byte("wangyu"), []byte("wangyu")})
}
