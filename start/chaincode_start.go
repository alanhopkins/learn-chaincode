/*
Copyright IBM Corp 2016 All Rights Reserved.

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

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) ([]byte, error) {

    _, args := stub.GetFunctionAndParameters()

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting 1")
    }

    err := stub.PutState("hello_world", []byte(args[0]))
    if err != nil {
        return nil, err
    }

    return nil, nil

}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) ([]byte, error) {

    function, args := stub.GetFunctionAndParameters()

    fmt.Println("invoke is running " + function)
    // Handle different functions
    if function == "init" {
        return t.Init(stub)
    } else if function == "write" {
        return t.write(stub, args)
    }
    fmt.Println("invoke did not find func: " + function)

    return nil, errors.New("Received unknown function invocation: " + function)

}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface) ([]byte, error) {
    fmt.Println("query entry")

    function, args := stub.GetFunctionAndParameters()
    fmt.Println("query is running " + function)
    fmt.Println("num args = %d", len(args))

    // Handle different functions
    if function == "read" {                            //read a variable
        return t.read(stub)
    }
    fmt.Println("query did not find func: " + function)

    return nil, errors.New("Received unknown function query: " + function)

}

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface) ([]byte, error) {
    var key, jsonResp string
    var err error
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("read func = : " + function)

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    key = args[0]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }

    return valAsbytes, nil
}

func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, value string
    var err error
    fmt.Println("running write()")

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }

    key = args[0]                            //rename for fun
    value = args[1]
    err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
        return nil, err
    }
    return nil, nil
}

