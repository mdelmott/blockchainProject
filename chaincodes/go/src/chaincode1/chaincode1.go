/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package chaincode1

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/mdelmott/blockchainProject/chaincodes/go/src/chaincode1/deploy"
	"github.com/mdelmott/blockchainProject/chaincodes/go/src/chaincode1/invokeElement"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	deploy := chaincode1.Deploy{}
	return deploy.Init(stub, args)
}

/*func (t *SimpleChaincode) createTable(stub *shim.ChaincodeStub) error {
	var columnDefsTable []*shim.ColumnDefinition
	columnOneTableDef := shim.ColumnDefinition{Name: "colOne",
		Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoTableDef := shim.ColumnDefinition{Name: "colTwo",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnThreeTableDef := shim.ColumnDefinition{Name: "colThree",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnFourTableDef := shim.ColumnDefinition{Name: "colFour",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnDefsTable = append(columnDefsTable, &columnOneTableDef)
	columnDefsTable = append(columnDefsTable, &columnTwoTableDef)
	columnDefsTable = append(columnDefsTable, &columnThreeTableDef)
	columnDefsTable = append(columnDefsTable, &columnFourTableDef)
	return stub.CreateTable("table", columnDefsTable)
}*/

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	function = args[0]

	switch function {
	case "add" : return t.add(stub, args)
	case "delete" : return t.delete(stub, args)
	case "transaction" : return t.transaction(stub, args)
	case "fusion" : return t.fusion(stub, args)
	case "insertTable" : return t.insertTable(stub, args)
	default: return nil, errors.New("Incorrect name of function")
	}
}

func (t *SimpleChaincode) insertTable(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	var attr1, attr2, attr3, attr4 string

	attr1 = "aaa"
	attr2 = args[1]
	attr3 = args[2]
	attr4 = args[3]

	var columns []*shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: attr1}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: attr2}}
	col3 := shim.Column{Value: &shim.Column_String_{String_: attr3}}
	col4 := shim.Column{Value: &shim.Column_String_{String_: attr4}}
	columns = append(columns, &col1)
	columns = append(columns, &col2)
	columns = append(columns, &col3)
	columns = append(columns, &col4)

	row := shim.Row{Columns: columns}
	ok, err := stub.InsertRow("table", row)
	if err != nil {
		return nil, fmt.Errorf("insertRowTable operation failed. %s", err)
	}
	if !ok {
		return nil, errors.New("insertRowTable operation failed. Row with given key already exists")
	}

	return nil, nil
}

func (t *SimpleChaincode) add(stub *shim.ChaincodeStub, args []string) ([]byte, error){
	var A string    // Entities
	var Aval int // Asset holdings
	var err error

	A = args[1];
	Aval,err = strconv.Atoi(args[2]);
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *SimpleChaincode) transaction(stub *shim.ChaincodeStub, args []string) ([]byte, error){

	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	A = args[1]
	B = args[2]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Avalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Bvalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[3])
	if err != nil {
		return nil, errors.New("Invalid transaction amount, expecting a integer value")
	}
	Aval = Aval - X
	Bval = Bval + X
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *SimpleChaincode) fusion(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var A, B string
	var Aval, Bval int

	A = args[1]
	B = args[2]

	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Avalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Bvalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	Aval = Aval + Bval

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.DelState(B)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}


// Deletes an entity from state
func (t *SimpleChaincode) delete(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	A := args[1]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	function = args[0];

	switch function {
	case "queryElement" : return t.queryElement(stub, args)
	case "queryTable" : return t.queryTable(stub, args)
	default: return nil, fmt.Errorf("Incorrect name of function, function received : %s", function)
	}

}

func (t *SimpleChaincode) queryElement(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var A string // Entities
	var err error

	A = args[1]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}

func (t *SimpleChaincode) queryTable(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	var columns []shim.Column

	col1Val := "aaa"
	col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
	columns = append(columns, col1)

	rowChannel, err := stub.GetRows("table", columns)
	var rows []shim.Row
	if err != nil {
		return nil, fmt.Errorf("getTable operation failed. %s", err)
	}

	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	jsonRows, err := json.Marshal(rows)
	if err != nil {
		return nil, fmt.Errorf("getTable operation failed. Error marshaling JSON: %s", err)
	}

	return jsonRows, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
