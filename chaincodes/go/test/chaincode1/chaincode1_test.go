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

package chaincode1_test

import (
	"../../src/chaincode1"
	"testing"
	//"github.com/hyperledger/fabric/core/chaincode/shim"
	//"github.com/stretchr/testify/mock"
	//"github.com/stretchr/testify/assert"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


func TestInit(t *testing.T) {
	sc := shim.Start(new(chaincode1.SimpleChaincode))
	t.Log(sc)

}

/*func testCreateTable(t *testing.T) {

}

func testInvoke(t *testing.T) {

}

func testInsertTable(t *testing.T) {

}

func testAdd(t *testing.T) {

}

func testTransaction(t *testing.T) {

}

func testFusion(t *testing.T) {

}

func testDelete(t *testing.T) {

}

func testQuery(t *testing.T) {

}

func testQueryElement(t *testing.T) {

}

func testQueryTable(t *testing.T) {

}*/

/*func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}*/
