/* Sales agent and distributor data*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/*
Structure tags are used by encoding/json library
*/
type DistributorDataNew struct {
	TxId              string `json:"txId"`
	ProjectId         string `json:"projectId"`
	Country           string `json:"country"`
	Continent         string `json:"continent"`
	DistributorName   string `json:"distributorName"`
	Role              string `json:"role"`
	ContNo            string `json:"contNo"`
	ContDate          string `json:"contDate"`
	SignedDm          string `json:"signedDm"`
	SignedLF          string `json:"signedLF"`
	WeekStart         string `json:"weekStart"`
	RightsGranted     string `json:"rightsGranted"`
	LicencePeriod     string `json:"licencePeriod"`
	Currency          string `json:"currency"`
	Gross             string `json:"gross"`
	Terms             string `json:"terms"`
	MGPaid            string `json:"mGPaid"`
	MGUnpaid          string `json:"mGUnpaid"`
	WTaxFees          string `json:"wTaxFees"`
	BankCharge        string `json:"bankCharge"`
	WeekendCollection string `json:"weekendCollection"`
	Month             string `json:"month"`
	Year              string `json:"year"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrive the smartcontract function arguments
	function, args := APIstub.GetFunctionAndParameters()

	if function == "newTransaction" {
		return s.newTransaction(APIstub, args)
	} else if function == "showTransaction" {
		return s.showTransaction(APIstub, args)
	} else if function == "updateTransaction" {
		return s.updateTransaction(APIstub, args)
	} else if function == "queryTransactionByProjectId" {
		return s.queryTransactionByProjectId(APIstub, args)
	} else if function == "showAllTransaction" {
		return s.showAllTransaction(APIstub)
	} else if function == "deleteTransaction" {
		return s.deleteTransaction(APIstub, args)
	}
	return shim.Error("Invalid SmartContract Function Name !!!")
}

// Add new account details
func (s *SmartContract) newTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	// var err error
	// Check even any arguments are not empty
	if len(args) != 23 {
		return shim.Error("Incorrect number of arguments. Expecting 23")
	}

	// Create new stakeholder
	var transactionData = DistributorDataNew{TxId: args[0], ProjectId: args[1], Country: args[2],
		Continent: args[3], DistributorName: args[4], Role: args[5], ContNo: args[6],
		ContDate: args[7], SignedDm: args[8], SignedLF: args[9], WeekStart: args[10],
		RightsGranted: args[11], LicencePeriod: args[12], Currency: args[13], Gross: args[14],
		Terms: args[15], MGPaid: args[16], MGUnpaid: args[17], WTaxFees: args[18], BankCharge: args[19], WeekendCollection: args[20], Month: args[21], Year: args[22]}
	// Check if allready that key is present
	transactionDataAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get account details!!" + err.Error())
	} else if transactionDataAsBytes != nil {
		fmt.Println("This key already exists: " + args[0])
		return shim.Error("This key already exists: " + args[0])
	}

	// Store data
	var key = args[0]
	transactionDataStoreAsBytes, _ := json.Marshal(transactionData)
	err1 := APIstub.PutState(key, transactionDataStoreAsBytes)
	if err1 != nil {
		return shim.Error(fmt.Sprintf("Failed to record stakeholder account details !!! %s", args[0]))
	}
	return shim.Success(nil)
}

func (s *SmartContract) showTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	var id, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting id of the transaction")
	}

	id = args[0]
	valAsbytes, err := APIstub.GetState(id) //get the marksheet from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Transaction does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

func (s *SmartContract) showAllTransaction(APIstub shim.ChaincodeStubInterface) sc.Response {
	startKey := "1"
	endKey := "99999"

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
		buffer.WriteString(queryResponse.Key + "\n")
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value) + "\n")
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllAccountData:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) updateTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	transactionAsBytes, _ := APIstub.GetState(args[0])
	transactionData := DistributorDataNew{}
	// Payout{TxId: args[0], ProjectId: args[1], DistributorName: args[2],
	// 	Movie: args[3], TheatreName: args[4], Country: args[5], Role: args[6],
	// 	WeekendNo: args[7], WeekStart: args[8], WeekEnd: args[9], Gross: args[10],
	// 	Take: args[11], WeekendCollection: args[12], ChangeWeekend: args[13], LifetimeGross: args[14],
	// 	GrossPerScreen: args[15], NoOfScreens: args[16], ChangeInRate: args[17], NoOfTheaters: args[18]}

	json.Unmarshal(transactionAsBytes, &transactionData)
	transactionData.ProjectId = args[1]
	transactionData.Country = args[2]
	transactionData.Continent = args[3]
	transactionData.DistributorName = args[4]
	transactionData.Role = args[5]
	transactionData.ContNo = args[6]
	transactionData.ContDate = args[7]
	transactionData.SignedDm = args[8]
	transactionData.SignedLF = args[9]
	transactionData.WeekStart = args[10]
	transactionData.RightsGranted = args[11]
	transactionData.LicencePeriod = args[12]
	transactionData.Currency = args[13]
	transactionData.Gross = args[14]
	transactionData.Terms = args[15]
	transactionData.MGPaid = args[16]
	transactionData.MGUnpaid = args[17]
	transactionData.WTaxFees = args[18]
	transactionData.BankCharge = args[19]
	transactionData.WeekendCollection = args[20]
	transactionData.Month = args[21]
	transactionData.Year = args[22]

	transactionAsBytes, _ = json.Marshal(transactionData)
	APIstub.PutState(args[0], transactionAsBytes)

	return shim.Success(nil)

}

// Query all project camaterm
func (s *SmartContract) queryTransactionByProjectId(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	//projectId := args[1]
	queryString := args[0]
	//queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"camaterm\",\"projectId\":\"%s\"}}", projectId)
	queryResults, err := getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(APIstub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// ===========================================================================================
// constructQueryResponseFromIterator constructs a JSON array containing query results from
// a given result iterator
// ===========================================================================================
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
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

	return &buffer, nil
}

//Delete Transaction data
func (s *SmartContract) deleteTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var jsonResp string

	A := args[0]
	valAsbytes, err := APIstub.GetState(A) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state invoke for " + A + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"transaction id does not exist: " + A + "\"}"
		return shim.Error(jsonResp)
	}
	// // Delete the key from the state in ledger
	err1 := APIstub.DelState(A)
	if err1 != nil {
		return shim.Error("Failed to delete state:" + err1.Error())
	}

	return shim.Success(nil)
}

// Main method
func main() {
	// Create new smartcontract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
