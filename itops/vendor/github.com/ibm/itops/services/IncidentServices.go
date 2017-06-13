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

var mapIncident map[string]data.IncidentDO

/*
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
																				Incident Services
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/

/*func CreateIncidentTable(stub shim.ChaincodeStubInterface) (bool, error) {

	fmt.Println("Creating Incident Table ...")

	// Create Incident table
	err := stub.CreateTable("INCIDENT", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "incident_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "incident_record", Type: shim.ColumnDefinition_STRING, Key: false}})

	if err != nil {
		return false, fmt.Errorf("Failed creating Incident table.")
	}

	fmt.Println("Incident table initialization done successfully... !!! ")

	return true, nil
}*/

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
	fmt.Println("Incident record is:  ", incidentJSON)

	/*err1 := stub.PutState(incidentRecord.IncidentID, incidentRecordBytes)
    	if err1 != nil {
       	   fmt.Println("Could not save changes", err1)
		return false, fmt.Errorf("Error in storing.")
    	}*/


	success, err := stub.InsertRow("INCIDENT", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.IncidentID}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.IncidentTitle}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.IncidentType}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.Severity}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.Status}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.RefIncidentID}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.OriginalIncidentID}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.ParticipantIDFrom}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.ParticipantIDTo}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.ContactEmail}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.CreatedDate}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.ExpectedCloseDate}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.ActualCloseDate}}}})

	if (err != nil) {
		return false, fmt.Errorf("Error in creating Incident record.")
	}

	if (!success) {
		fmt.Printf("Error in creating Incident record. Row with given key already exists! Updating...")
		success, err := UpdateIncident(stub, incidentRecord)
		if ((!success) || (err != nil)) {
	 		return false, fmt.Errorf("Error in updating Incident record.")
		}

	}

	// if (!(success && (err == nil))) {
	// 	return nil, fmt.Errorf("Error in creating Incident record.")
	// }

	fmt.Println("Incident record created/updated. Incident Id : [%s]", string(incidentRecord.IncidentID))
	fmt.Println("Adding to map!")
	mapIncident[incidentRecord.IncidentID] = incidentRecord
	mapIncident[incidentRecord.IncidentTitle] = incidentRecord
	mapIncident[incidentRecord.Severity] = incidentRecord
	mapIncident[incidentRecord.Status] = incidentRecord
	mapIncident[incidentRecord.ContactEmail] = incidentRecord
	fmt.Println("Added to map!")

	fmt.Println("Querying it!")
	n := mapIncident[incidentRecord.Status]
	fmt.Println("Value of status is: ", n)
	
	return true, nil
}


func UpdateIncident(stub shim.ChaincodeStubInterface, incidentRecord data.IncidentDO) (bool, error) {
	fmt.Println("Updating Incident record ...")

	incidentRecordBytes, marshalErr := json.Marshal(incidentRecord)

	if (marshalErr != nil) {
		return false, fmt.Errorf("Error in marshalling Incident record.")
	}

	incidentJSON := string(incidentRecordBytes)
	fmt.Println("Incident record is:  ", incidentJSON)

	success, err := stub.ReplaceRow("INCIDENT", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.IncidentID}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.IncidentTitle}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.IncidentType}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.Severity}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.Status}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.RefIncidentID}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.OriginalIncidentID}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.ParticipantIDFrom}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.ParticipantIDTo}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.ContactEmail}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.CreatedDate}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.ExpectedCloseDate}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.ActualCloseDate}}}})


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

	fmt.Println("Retrieving Incident record. Incident Id : [%s]", string(incidentId))

	var columns []shim.Column
	incidentIdColumn := shim.Column{Value: &shim.Column_String_{String_: incidentId}}
	columns = append(columns, incidentIdColumn)
	row, err := stub.GetRow("INCIDENT", columns)

	if err != nil {
		fmt.Printf("Error retrieving Incident record [%s]: [%s]", string(incidentId), err)
		fmt.Println()
		return "", fmt.Errorf("Error retrieving Incident record [%s]: [%s]", string(incidentId), err)
	}

	fmt.Printf("Row - [%s]", row)
	fmt.Println()

	fmt.Println("Printing the map!")
	for k, v := range mapIncident {
		fmt.Println("Key-value pairs: ", k, " : ", v)
	}

	var jsonRespBuffer bytes.Buffer
	jsonRespBuffer.WriteString(row.Columns[1].GetString_())

	return jsonRespBuffer.String(), nil

	/*bytes, err := stub.GetState(incidentId)
	if err != nil {
		fmt.Printf("Could not fetch record with id " + incidentId + " from ledger", err)
		return "", fmt.Errorf("Error in fetching : ", err)
	}
	return string(bytes), nil*/

}



func CreateIncidentTable(stub shim.ChaincodeStubInterface) ([]byte, error) {

	fmt.Println("Creating Incident Table ...")

	/*if len(args) != 0 {
		return nil, fmt.Errorf("Incorrect number of arguments. Expecting 0")
	}*/

	// Create Incident table
	err := stub.CreateTable("INCIDENT", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "incident_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "incident_title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "incident_type", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "severity", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ref_incident_id", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "original_incident_id", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "participant_id_from", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "participant_id_to", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "contact_email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "created_date", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "expected_close_date", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "actual_close_date", Type: shim.ColumnDefinition_STRING, Key: false},
	})

	if err != nil {
		return nil, fmt.Errorf("Failed creating Incident table.")
	}

	fmt.Println("Incident table initialization done successfully... !!! ")

	return nil, nil
}
