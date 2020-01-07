package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Chaincode struct {
}

func (c *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Init Success!"))
}

func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	cByte, err := json.Marshal(time.Now())
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("key", cByte)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("Invoke Success!"))
}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		fmt.Printf("Start Chaincode Failed, error: %s", err)
	}
}
