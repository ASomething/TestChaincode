package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type BadChaincode struct {
}

func (t *BadChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("success"))
}

func (t *BadChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	tByte, err := json.Marshal(time.Now())
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("key", tByte)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("success"))
}

func main() {
	if err := shim.Start(new(BadChaincode)); err != nil {
		fmt.Printf("Error starting BadChaincode chaincode: %s", err)
	}
}
