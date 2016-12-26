

package main


import (
	"errors"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/ibm/itops/data"
	"github.com/ibm/itops/services"
)

type ITOpsChaincode struct {

}

func (self *ITOpsChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Println("[ITOpsChaincode]: Init - Start")

	_,err := services.CreateIncidentTable(stub)

	if (err != nil) {
		return nil, err
	}

	fmt.Println("[ITOpsChaincode]: Init - End")

	return nil, nil

}


func (self *ITOpsChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Println("[ITOpsChaincode]: Invoke - Start")
	fmt.Printf("[ITOpsChaincode]: Invoke - Function to be called - %s",function)
	fmt.Println()

	if (len(args) == 0) {
		return nil, errors.New("[ITOpsChaincode]: Invoke - Function parameters not received")
	}

	if (function == "") {
		return nil, errors.New("[ITOpsChaincode]: Invoke - Function name not specified")
	}

	// Handle different functions
	if function == "addIncident" {
		self.addIncident(stub, args[0])
	} else if function == "updateIncident" {
		self.updateIncident(stub, args[0])
	} else {
		return nil, errors.New("[ITOpsChaincode]: Invoke - Unknown Function invocation")
	}

	fmt.Println("[ITOpsChaincode]: Invoke - End")

	return nil, nil
}

func (self *ITOpsChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	fmt.Println("[ITOpsChaincode]: Query - Start")
	fmt.Printf("[ITOpsChaincode]: Function to be called - %s",function)
	fmt.Println()

	if (len(args) == 0) {
		return nil, errors.New("[ITOpsChaincode]: Query - Function parameters not received")
	}

	if (function == "") {
		return nil, errors.New("[ITOpsChaincode]: Query - Function name not specified")
	}

	// Handle different functions
	if function == "getIncident" {
		self.getIncident(stub, args[0])
	} else {
		return nil, errors.New("[ITOpsChaincode]: Query - Unknown Function invocation")
	}

	fmt.Println("[ITOpsChaincode]: Query - End")

	return nil, nil
}

func main() {
	fmt.Println("[ITOpsChaincode]: main - Start")

	err := shim.Start(new(ITOpsChaincode))
	if err != nil {
		fmt.Printf("[ITOpsChaincode]: main - Error starting ITOps chaincode: %s", err)
		fmt.Println()
	}

	fmt.Println("[ITOpsChaincode]: main - End")
}

func (self *ITOpsChaincode) addIncident(stub shim.ChaincodeStubInterface, incidentJSON string) (bool, error) {

	fmt.Println("[ITOpsChaincode]: addIncident - Start")

	var incidentRecord data.IncidentDO
	unmarshalErr := json.Unmarshal([]byte(string(incidentJSON)), &incidentRecord)

	if (unmarshalErr != nil) {
		return false, fmt.Errorf("[ITOpsChaincode]: addIncident - Error in unmarshalling JSON string to Incident record.")
	}

	success, err := services.CreateIncident(stub, incidentRecord)

	if ((err != nil) || (!success)) {
		return false, fmt.Errorf("[ITOpsChaincode]: addIncident - Error in creating Incident record.")
	}

	fmt.Printf("[ITOpsChaincode]: addIncident - Incident record created. Incident Id : %s", string(incidentRecord.IncidentID))
	fmt.Println()

	fmt.Println("[ITOpsChaincode]: addIncident - End")

	return success, nil
}


func (self *ITOpsChaincode) updateIncident(stub shim.ChaincodeStubInterface, incidentJSON string) (bool, error) {

	fmt.Println("[ITOpsChaincode]: updateIncident - Start")

	var incidentRecord data.IncidentDO
	unmarshalErr := json.Unmarshal([]byte(string(incidentJSON)), &incidentRecord)

	if (unmarshalErr != nil) {
		return false, fmt.Errorf("[ITOpsChaincode]: updateIncident - Error in unmarshalling JSON string to Incident record.")
	}

	success, err := services.UpdateIncident(stub, incidentRecord)

	if ((err != nil) || (!success)) {
		return false, fmt.Errorf("[ITOpsChaincode]: updateIncident - Error in updating Incident record.")
	}

	fmt.Printf("[ITOpsChaincode]: updateIncident - Incident record updated. Incident Id : %s", string(incidentRecord.IncidentID))
	fmt.Println()
	fmt.Println("[ITOpsChaincode]: updateIncident - End")

	return success, nil
}


func (self *ITOpsChaincode) getIncident(stub shim.ChaincodeStubInterface, incidentID string) (string, error) {

	fmt.Println("[ITOpsChaincode]: updateIncident - Start")
	if incidentID == "" {
		return "", errors.New("Incident ID expected")
	}

	incidentRecordJSON, err := services.RetrieveIncident(stub, incidentID)

	if (err != nil)  {
		return "", fmt.Errorf("Error in rtrieving Incident record.")
	}

	fmt.Printf("[ITOpsChaincode]: updateIncident - Incident record retrieved. Incident Id : %s", string(incidentID))
	fmt.Println()
	fmt.Println("[ITOpsChaincode]: updateIncident - Start")
	return incidentRecordJSON, nil
}

/*

//For future use

func (self *ITOpsChaincode) getFunctionDispatcher(grammer *regexp.Regexp, function string) (api.Dispatcher, int, error) {

	spec := re.FindAllStringSubmatch(function, -1)
	if spec == nil {
		return nil, 0, errors.New("Could not parse function name")
	}

	dispatcher, ok := self.dispatchers[spec[0][1]]
	if !ok {
		return nil, 0, errors.New("Interface not found")
	}

	index, err := strconv.Atoi(spec[0][2])
	if err != nil {
		return nil, 0, errors.New("Could not convert function index")
	}

	return dispatcher, index, nil
}

*/
