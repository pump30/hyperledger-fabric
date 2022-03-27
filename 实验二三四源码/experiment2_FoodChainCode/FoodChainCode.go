package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type FoodChainCode struct {
}

// FoodInfo food数据结构体
type FoodInfo struct {
	FoodID      string    `json:"FoodID"`      //食品ID
	FoodProInfo ProInfo   `json:"FoodProInfo"` //生产信息
	FoodIngInfo []IngInfo `json:"FoodIngInfo"` //配料信息
	FoodLogInfo LogInfo   `json:"FoodLogInfo"` //物流信息
}
type FoodAllInfo struct {
	FoodID      string    `json:"FoodID"`
	FoodProInfo ProInfo   `json:"FoodProInfo"`
	FoodIngInfo []IngInfo `json:"FoodIngInfo"`
	FoodLogInfo []LogInfo `json:"FoodLogInfo"`
}

// ProInfo 生产信息
type ProInfo struct {
	FoodName     string `json:"FoodName"`     //食品名称
	FoodSpec     string `json:"FoodSpec"`     //食品规格
	FoodMFGDate  string `json:"FoodMFGDate"`  //食品出产日期
	FoodEXPDate  string `json:"FoodEXPDate"`  //食品保质期
	FoodLOT      string `json:"FoodLOT"`      //食品批次号
	FoodQSID     string `json:"FoodQSID"`     //食品生产许可证编号
	FoodMFRSName string `json:"FoodMFRSName"` //食品生产商名称
	FoodProPrice string `json:"FoodProPrice"` //食品生产价格
	FoodProPlace string `json:"FoodProPlace"` //食品生产所在地
}
type IngInfo struct {
	IngID   string `json:"IngID"`   //配料ID
	IngName string `json:"IngName"` //配料名称
}
type LogInfo struct {
	LogDepartureTm string `json:"LogDepartureTm"` //出发时间
	LogArrivalTm   string `json:"LogArrivalTm"`   //到达时间
	LogMission     string `json:"LogMission"`     //处理业务（储存or运输）
	LogDeparturePl string `json:"LogDeparturePl"` //出发地
	LogDest        string `json:"LogDest"`        //目的地
	LogToSeller    string `json:"LogToSeller"`    //销售商
	LogStorageTm   string `json:"LogStorageTm"`   //存储时间
	LogMOT         string `json:"LogMOT"`         //运送方式
	LogCopName     string `json:"LogCopName"`     //物流公司名称
	LogCost        string `json:"LogCost"`        //费用
}

func (a *FoodChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}
func (a *FoodChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	switch fn {
	case "addProInfo":
		return addProInfo(stub, args)
	case "addIngInfo":
		return addIngInfo(stub, args)
	case "getFoodInfo":
		return getFoodInfo(stub, args)
	case "addLogInfo":
		return addLogInfo(stub, args)
	case "getProInfo":
		return getProInfo(stub, args)
	case "getLogInfo":
		return getLogInfo(stub, args)
	case "getIngInfo":
		return getIngInfo(stub, args)
	case "getLogInfo_1":
		return getLogInfo_1(stub, args)

	}
	return shim.Error(fmt.Sprintf("unsupported function: %s", fn))
}

//新增生产函数
func addProInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var FoodInfos FoodInfo
	//检查参数个数
	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments.")
	}
	//参数解析
	FoodInfos.FoodID = args[0]
	if FoodInfos.FoodID == "" {
		return shim.Error("FoodID can not be empty")
	}
	FoodInfos.FoodProInfo.FoodName = args[1]
	FoodInfos.FoodProInfo.FoodSpec = args[2]
	FoodInfos.FoodProInfo.FoodMFGDate = args[3]
	FoodInfos.FoodProInfo.FoodEXPDate = args[4]
	FoodInfos.FoodProInfo.FoodLOT = args[5]
	FoodInfos.FoodProInfo.FoodQSID = args[6]
	FoodInfos.FoodProInfo.FoodMFRSName = args[7]
	FoodInfos.FoodProInfo.FoodProPrice = args[8]
	FoodInfos.FoodProInfo.FoodProPlace = args[9]
	//对结构体进行序列化
	ProInfosJSONasBytes, err := json.Marshal(FoodInfos)
	if err != nil {
		return shim.Error(err.Error())
	}
	//保存状态
	err = stub.PutState(FoodInfos.FoodID, ProInfosJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//新增配料信息
func addIngInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var FoodInfos FoodInfo
	var IngInfoitem IngInfo
	//判断参数合法性
	if (len(args)-1)%2 != 0 || len(args) == 1 {
		return shim.Error("Incorrect number of arguments.")
	}
	//参数解析
	FoodID := args[0]
	for i := 1; i < len(args); {
		IngInfoitem.IngID = args[i]
		IngInfoitem.IngName = args[i+1]
		FoodInfos.FoodIngInfo = append(FoodInfos.FoodIngInfo, IngInfoitem)
		i = i + 2
	}
	FoodInfos.FoodID = FoodID
	//结构体
	IngInfoJsonAsBytes, err := json.Marshal(FoodInfos)
	if err != nil {
		return shim.Error(err.Error())
	}
	//状态保存
	err = stub.PutState(FoodInfos.FoodID, IngInfoJsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//获取食品全部信息
func getFoodInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//检查参数个数
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}
	//参数解析
	FoodID := args[0]
	//在世界状态中根据FoodID查询记录
	resultIterator, err := stub.GetHistoryForKey(FoodID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultIterator.Close()

	//迭代获取该FoodID对应的全部信息
	var foodAllinfo FoodAllInfo
	for resultIterator.HasNext() {
		var FoodInfos FoodInfo
		response, err := resultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value, &FoodInfos)
		if FoodInfos.FoodProInfo.FoodName != "" {
			foodAllinfo.FoodProInfo = FoodInfos.FoodProInfo
		} else if FoodInfos.FoodIngInfo != nil {
			foodAllinfo.FoodIngInfo = FoodInfos.FoodIngInfo
		} else if FoodInfos.FoodLogInfo.LogMission != "" {
			foodAllinfo.FoodLogInfo = append(foodAllinfo.FoodLogInfo, FoodInfos.FoodLogInfo)
		}
	}

	//对结果进行序列化
	jsonAsBytes, err := json.Marshal(foodAllinfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	//运行成功并返回
	return shim.Success(jsonAsBytes)
}

//新增物流信息
func addLogInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var FoodInfos FoodInfo
	//判断参数合法性并解析参数
	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments.")
	}
	FoodInfos.FoodID = args[0]
	if FoodInfos.FoodID == "" {
		return shim.Error("FoodID can not be empty")
	}
	FoodInfos.FoodLogInfo.LogDepartureTm = args[1]
	FoodInfos.FoodLogInfo.LogArrivalTm = args[2]
	FoodInfos.FoodLogInfo.LogMission = args[3]
	FoodInfos.FoodLogInfo.LogDeparturePl = args[4]
	FoodInfos.FoodLogInfo.LogDest = args[5]
	FoodInfos.FoodLogInfo.LogToSeller = args[6]
	FoodInfos.FoodLogInfo.LogStorageTm = args[7]
	FoodInfos.FoodLogInfo.LogMOT = args[8]
	FoodInfos.FoodLogInfo.LogCopName = args[9]
	FoodInfos.FoodLogInfo.LogCost = args[10]

	//序列化
	LogInfosJSONasBytes, err := json.Marshal(FoodInfos)
	if err != nil {
		return shim.Error(err.Error())
	}
	//保存状态
	err = stub.PutState(FoodInfos.FoodID, LogInfosJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//获取食品基本生产信息
func getProInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//判断参数合法性并解析参数
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}
	FoodID := args[0]
	//根据FoodID查询信息
	resultIterator, err := stub.GetHistoryForKey(FoodID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultIterator.Close()

	//迭代赋值
	var foodProInfo ProInfo
	for resultIterator.HasNext() {
		var FoodInfos FoodInfo
		response, err := resultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		//反序列化
		json.Unmarshal(response.Value, &FoodInfos)
		if FoodInfos.FoodProInfo.FoodName != "" {
			foodProInfo = FoodInfos.FoodProInfo
			continue
		}
	}
	//对结果进行序列化
	jsonAsBytes, err := json.Marshal(foodProInfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	//返回
	return shim.Success(jsonAsBytes)
}

//获取配料信息
func getIngInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}
	FoodID := args[0]
	resultIterator, err := stub.GetHistoryForKey(FoodID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultIterator.Close()

	var foodIngInfo []IngInfo
	for resultIterator.HasNext() {
		var FoodInfos FoodInfo
		response, err := resultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		//反序列化
		json.Unmarshal(response.Value, &FoodInfos)
		if FoodInfos.FoodIngInfo != nil {
			foodIngInfo = FoodInfos.FoodIngInfo
			continue
		}
	}
	jsonAsBytes, err := json.Marshal(foodIngInfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	//返回
	return shim.Success(jsonAsBytes)
}

//获取全部物流信息
func getLogInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var LogInfos []LogInfo
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}
	FoodID := args[0]
	resultIterator, err := stub.GetHistoryForKey(FoodID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultIterator.Close()

	for resultIterator.HasNext() {
		var FoodInfos FoodInfo
		response, err := resultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		//反序列化
		json.Unmarshal(response.Value, &FoodInfos)
		if FoodInfos.FoodLogInfo.LogMission != "" {
			LogInfos = append(LogInfos, FoodInfos.FoodLogInfo)
		}
	}
	jsonAsBytes, err := json.Marshal(LogInfos)
	if err != nil {
		return shim.Error(err.Error())
	}
	//返回
	return shim.Success(jsonAsBytes)
}

//获取最新物流信息
func getLogInfo_1(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var LogInfo LogInfo
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}
	FoodID := args[0]
	resultIterator, err := stub.GetHistoryForKey(FoodID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultIterator.Close()

	for resultIterator.HasNext() {
		var FoodInfos FoodInfo
		response, err := resultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		//反序列化
		json.Unmarshal(response.Value, &FoodInfos)
		if FoodInfos.FoodLogInfo.LogMission != "" {
			LogInfo = FoodInfos.FoodLogInfo
			continue
		}
	}
	jsonAsBytes, err := json.Marshal(LogInfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	//返回
	return shim.Success(jsonAsBytes)
}
func main() {
	if err := shim.Start(new(FoodChainCode)); err != nil {
		fmt.Printf("Error starting Food ChainCode:%s", err)
	}
}
