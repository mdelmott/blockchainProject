package chaincode1

import (
	"errors"
	"strconv"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type InvokeElement struct {
}

func (ie *InvokeElement) add(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return nil, nil
}

func main() {
	//new(InvokeElement)
}
