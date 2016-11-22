package chaincode1

import (
	"errors"
	"strconv"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Deploy struct {
}

func (d *Deploy) Init(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var A string    // Entities
	var Aval int // Asset holdings
	var err error

	if len(args)%2 != 0 {
		return nil, errors.New("Incorrect number of arguments")
	}

	for i := 0; i<len(args); i= i+2 {
		A = args[i];
		Aval,err = strconv.Atoi(args[i+1]);
		if err != nil {
			return nil, errors.New("Expecting integer value for asset holding")
		}
		err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
		if err != nil {
			return nil, err
		}

	}

	err = d.createTable(stub)
	if err != nil {
		return nil, fmt.Errorf("Error creating table one during init. %s", err)
	}

	return nil, nil
}

func (d *Deploy) createTable(stub *shim.ChaincodeStub) error {
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
}


func main() {
	//new(Deploy)
}
