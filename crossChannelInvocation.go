package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Chaincode struct {
}

func (c *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Init Success!"))
}

/*
**假设本通道名字为mychannel，本链码名字为mycc.
**现在通过Invoke()方法调用另一个通道otherchannel上的链码othercc：
 */
func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetArgs()
	response := stub.InvokeChaincode("othercc", args, "otherchannel")
	if response.Status != shim.OK {
		err := fmt.Sprintln("Invoke Chaincode Failed, error: %s", response.Payload)
		return shim.Error(err)
	}
	return shim.Success([]byte("Invoke Chaincode on another Channel Success!"))
}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		fmt.Printf("Start Chaincode Failed, error: %s", err)
	}
}
