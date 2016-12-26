/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package services

import (
	"bytes"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/ibm/itops/data"
)


/*
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
																				Incident Services
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/

func CreateIncidentTable(stub shim.ChaincodeStubInterface) (bool, error) {

	fmt.Println("Creating Incident Table ...")

	// Create Incident table
	err := stub.CreateTable("INCIDENT", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "incident_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "incident_record", Type: shim.ColumnDefinition_STRING, Key: true},
	})

	if err != nil {
		return false, fmt.Errorf("Failed creating Incident table.")
	}

	fmt.Println("Incident table initialization done successfully... !!! ")

	return true, nil
}

/*
	Create Incident record
*/
func CreateIncident(stub shim.ChaincodeStubInterface, incidentRecord data.IncidentDO) (bool, error) {

	fmt.Println("Creating Incident record ...")

	incidentRecordBytes, marshalErr := json.Marshal(incidentRecord)

	if (marshalErr != nil) {
		return false, fmt.Errorf("Error in marshalling Incident record.")
	}

	incidentJSON := string(incidentRecordBytes)

	success, err := stub.InsertRow("INCIDENT", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.IncidentID}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
		},
	})

	if ((err != nil) || (!success)) {
		return false, fmt.Errorf("Error in creating Incident record.")
	}

	// if (!(success && (err == nil))) {
	// 	return nil, fmt.Errorf("Error in creating Incident record.")
	// }

	fmt.Println("Incident record created. Incident Id : [%s]", string(incidentRecord.IncidentID))

	return success, nil
}


func UpdateIncident(stub shim.ChaincodeStubInterface, incidentRecord data.IncidentDO) (bool, error) {
	fmt.Println("Updating Incident record ...")

	incidentRecordBytes, marshalErr := json.Marshal(incidentRecord)

	if (marshalErr != nil) {
		return false, fmt.Errorf("Error in marshalling Incident record.")
	}

	incidentJSON := string(incidentRecordBytes)

	success, err := stub.ReplaceRow("INCIDENT", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.IncidentID}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
		},
	})

	if ((err != nil) || (!success)) {
		return false, fmt.Errorf("Error in updating Incident record.")
	}

	// if (!(success && (err == nil))) {
	// 	return nil, fmt.Errorf("Error in updating Incident record.")
	// }

	fmt.Println("Incident record updated. Incident Id : [%s]", string(incidentRecord.IncidentID))

	return success, nil
}



/*
 Retrieve Incident record
*/
func RetrieveIncident(stub shim.ChaincodeStubInterface, incidentId string) (string, error) {

	fmt.Println("Retrieveing Incident record. Incident Id : [%s]", string(incidentId))

	var columns []shim.Column
	incidentIdColumn := shim.Column{Value: &shim.Column_String_{String_: incidentId}}
	columns = append(columns, incidentIdColumn)
	row, err := stub.GetRow("INCIDENT", columns)

	if err != nil {
		fmt.Printf("Error retriving Incident record [%s]: [%s]", string(incidentId), err)
		fmt.Println()
		return "", fmt.Errorf("Error retriving Incident record [%s]: [%s]", string(incidentId), err)
	}

	fmt.Printf("Row - [%s]", row)
	fmt.Println()

	var jsonRespBuffer bytes.Buffer
	jsonRespBuffer.WriteString(row.Columns[1].GetString_())

	return jsonRespBuffer.String(), nil
}


/*

func CreateIncidentTable(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	fmt.Println("Creating Incident Table ...")

	if len(args) != 0 {
		return nil, fmt.Errorf("Incorrect number of arguments. Expecting 0")
	}

	// Create Incident table
	err := stub.CreateTable("INCIDENT", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "incident_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "incident_title", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "incident_type", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "severity", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ref_incident_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "original_incident_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "participant_id_from", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "participant_id_to", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "contact_email", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "created_date", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "expected_close_date", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "actual_close_date", Type: shim.ColumnDefinition_STRING, Key: false}
	})

	if err != nil {
		return nil, fmt.Errorf("Failed creating Incident table.")
	}

	fmt.Println("Incident table initialization done successfully... !!! ")

	return nil, nil
}

*/
