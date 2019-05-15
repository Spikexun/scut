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
/*
 *每一个chaincode需要实现Chaincode接口，其方法是用于响应接收到的transaction。
 *当chaincode接收到instantiate或者upgrade transaction时Init方法被调用了，以便
 *chaincode能够执行任何必要的初始化，包括application state的初始化。
 *当chaincode接收到invoke transaction时调用invoke方法，用于处理transaction
 *proposal。
*/
package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"fmt"//格式化IO包，实现了输入，输出函数
	"strconv"//包含一系列函数和方法包括ParseBool将字符串转变为布尔值，FormatBool将布尔值转换为字符串“true"或"false"等类型转换函数

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {//定义名为SimpleChaincode的空结构体
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	var AC, CA,EA,NA,SA,SHA,SIA,WA,XA string    // Entities(实体)
	var ACval, CAval,EAval,NAval,SAval,SHAval,SIAval,WAval,XAval int // Asset holdings(资产持有量)
	var err error

	//if len(args) != 18 {
	//	return shim.Error("Incorrect number of arguments. Expecting 4")
	//}

	// Initialize the chaincode
	AC = args[0]
	ACval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	CA = args[2]
	CAval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	EA = args[4]
	EAval, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	NA = args[6]
	NAval, err = strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	SA = args[8]
	SAval, err = strconv.Atoi(args[9])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	SHA = args[10]
	SHAval, err = strconv.Atoi(args[11])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	SIA = args[12]
	SIAval, err = strconv.Atoi(args[13])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	WA = args[14]
	WAval, err = strconv.Atoi(args[15])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	XA = args[16]
	XAval, err = strconv.Atoi(args[17])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}


	fmt.Printf("ACval = %d, CAval = %d\n", ACval, CAval)

	// Write the state to the ledger
	err = stub.PutState(AC, []byte(strconv.Itoa(ACval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(CA, []byte(strconv.Itoa(CAval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(EA, []byte(strconv.Itoa(EAval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(NA, []byte(strconv.Itoa(NAval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(SA, []byte(strconv.Itoa(SAval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(SHA, []byte(strconv.Itoa(SHAval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(SIA, []byte(strconv.Itoa(SIAval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(WA, []byte(strconv.Itoa(WAval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(XA, []byte(strconv.Itoa(XAval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
/*实现Invoke函数，负责处理传递的交易信息
 *并且在Inovke函数中，还有invoke,delete,query三个函数，用于细化交易的不同情况分支
 *这三个函数会在下面同样被实现
*/
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X, Y int          // Transaction value
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	Y, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}

	Aval = Aval - X
	Bval = Bval + Y
	fmt.Printf("%s = %d, %s = %d\n", A ,Aval,B, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}


        jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
        }

       func main() {      //主函数在实例化时启动容器中的链码
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
