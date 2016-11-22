package chaincode1

import (
	"errors"
	//"strconv"
	//"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

type InvokeElement struct {
}

func (ie *InvokeElement) Add(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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


func main() {
	//new(InvokeElement)
}
