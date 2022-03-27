package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type FinanceChainCode struct {
}

// User 用户
type User struct {
	Name       string   `json:"name"`
	Uid        string   `json:"uid"`
	CompactIDs []string `json:"compactIDs"`
}

// Compact 合同
type Compact struct {
	Timestamp        int64  `json:"timestamp"`
	Uid              string `json:"uid"`
	LoanAmount       string `json:"loanAmount"`
	ApplyDate        string `json:"applyDate"`
	CompactStartDate string `json:"compactStartDate"`
	CompactEndDate   string `json:"compactEndDate"`
	ID               string `json:"id"`
}

func (t *FinanceChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	if len(args) != 0 {
		return shim.Error("Parameter error while Init")
	}
	return shim.Success(nil)
}
func (t *FinanceChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	functionName, args := stub.GetFunctionAndParameters()
	switch functionName {
	case "userRegister":
		return userRegister(stub, args)
	case "loan":
		return loan(stub, args)
	case "queryCompact":
		return queryCompact(stub, args)
	case "queryUser":
		return queryUser(stub, args)
	default:
		return shim.Error("Invalid Smart Contract function name.")
	}
}

//用户注册
func userRegister(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//检查参数个数
	if len(args) != 2 {
		return shim.Error("Not enough args")
	}
	//检查参数正确性
	name := args[0]
	id := args[1]
	if name == "" || id == "" {
		return shim.Error("Invalid args")
	}
	//检查数据是否存在
	if userBytes, err := stub.GetState(id); err != nil || len(userBytes) != 0 {
		return shim.Error("User already exists")
	}
	//写入状态
	var user = User{Name: name, Uid: id}
	//序列化对象
	userBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error(fmt.Sprintf("marshal user error %s", err))
	}
	err = stub.PutState(id, userBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("put user error %s", err))
	}
	return shim.Success(nil)
}

//查询用户
func queryUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	userID := args[0]
	if userID == "" {
		return shim.Error("Invalid args")
	}
	userBytes, err := stub.GetState(userID)
	if err != nil || len(userBytes) == 0 {
		return shim.Error("user not found")
	}
	return shim.Success(userBytes)
}

//记录贷款数据
func loan(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//检查参数个数
	/**
	**
	**
	 */

	//检查参数值
	/**
	**
	**
	 */

	//检查该ID的贷款记录是否存在，贷款用户ID是否存在
	/**
	**
	**
	 */

	//保存合同信息
	/**
	**
	**
	 */

	//更新用户数据
	owner := new(User)
	if err := json.Unmarshal(ownerBytes, owner); err != nil {
		return shim.Error(fmt.Sprintf("unmarshal user error %s", err))
	}
	owner.CompactIDs = append(owner.CompactIDs, compact.ID)
	ownerBytes, err = json.Marshal(owner)
	if err != nil {
		return shim.Error(fmt.Sprintf("marshal user error %s", err))
	}

	//保存用户数据
	/**
	**
	**
	 */
	return shim.Success([]byte("记录贷款数据成功"))
}

//查询电子合同
func queryCompact(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//检查参数个数
	/**
	**
	**
	 */

	//检查参数值
	/**
	**
	**
	 */

	//读取合同信息（如果合同不存在，返回对应提示）
	/**
	**
	**
	 */
	return shim.Success(compactBytes)
}

func main() {
	if err := shim.Start(new(FinanceChainCode)); err != nil {
		fmt.Printf("Error creating new Smart Contart:%s", err)
	}
}
