package blockflow_chaincode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContractBlockflow struct {
}

/* Define Bloflow structure, with 4 properties.
Structure tags are used by encoding/json library
*/
type Workflow struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	Holder    string `json:"holder"`
}
type Programam struct{}
type Controller struct{}
type Port struct{}
type Channel struct{}
type User struct{}
type Association struct{}
type Execution struct{}
type Usage struct{}
type Generation struct{}
type Entity struct{}
type Collection struct{}
type Visualization struct{}
type Data struct{}
type Document struct{}

/* Define Bloflow structure, Relationships
 */
type activity struct{}
type agent struct{}
type entity struct{}
type hadPlan struct{}
type hasDefaultParam struct{}
type hasInput struct{}
type hasOutput struct{}
type hasSubProgram struct{}
type qualifiedAssociation struct{}
type qualifiedGeneration struct{}
type qualifiedUsage struct{}
type used struct{}
type wasAssociatedWith struct{}
type wasGeneratedBy struct{}
type wasInformeBy struct{}
type wasPartOf struct{}

/*
 * The Init method *
 called when the Smart Contract "tuna-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function
 -- see initLedger()
*/
func (s *SmartContractBlockflow) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "tuna-chaincode"
 The app also specifies the specific smart contract function to call with args
*/
func (s *SmartContractBlockflow) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "recordWorkflow" {
		return s.recordWorkflow(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "queryAllWorkflow" {
		return s.queryAllWorkflow(APIstub)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*Vessel
 * The initLedger method *
Will add test data (5 Workflow catches)to our network
*/
func (s *SmartContractBlockflow) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	workflow := []Workflow{
		Workflow{Name: "923F", Timestamp: "1504054225", Holder: "Raiane"},
		Workflow{Name: "85dF", Timestamp: "1504054225", Holder: "Fabio"},
		Workflow{Name: "9897", Timestamp: "1504054225", Holder: "JP"},
		Workflow{Name: "98KL", Timestamp: "1504054225", Holder: "Isadora"},
	}

	i := 0
	for i < len(workflow) {
		fmt.Println("i is ", i)
		tunaAsBytes, _ := json.Marshal(workflow[i])
		APIstub.PutState(strconv.Itoa(i+1), tunaAsBytes)
		fmt.Println("Added", workflow[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordWorkflow method *
Fisherman like Sarah would use to record each of her workflow catches.
This method takes in five arguments (attributes to be saved in the ledger).
*/
func (s *SmartContractBlockflow) recordWorkflow(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	var workflow = Workflow{Name: args[1], Timestamp: args[2], Holder: args[3]}

	workFlowAsBytes, _ := json.Marshal(workflow)
	err := APIstub.PutState(args[0], workFlowAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record workflow catch: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The queryAllWorkflow method *
allows for assessing all the records added to the ledger(all tuna catches)
This method does not take any arguments. Returns JSON string containing results.
*/
func (s *SmartContractBlockflow) queryAllWorkflow(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllWorkflow:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
 * main function *
calls the Start function
The main function starts the chaincode in the container during instantiation.
*/
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContractBlockflow))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
