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

//var mapIncident = map[string]data.IncidentDO{}


/*
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
																				Incident Services
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/

func CreateIncidentTable(stub shim.ChaincodeStubInterface) (bool, error) {

	fmt.Println("Creating Incident Table ...")

	// Create Incident table
	err := stub.CreateTable("INCIDENT", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "incident_key_column", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "incident_record", Type: shim.ColumnDefinition_STRING, Key: false}})

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
	fmt.Println("Incident record is:  ", incidentJSON)

	success1, err1 := stub.InsertRow("INCIDENT", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.IncidentID}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
		},
	})
	
	if (err1 != nil) {
		return false, fmt.Errorf("Error in creating Incident record.")
	}

	if (!success1) {
		fmt.Printf("Error in creating Incident record. Row with given key already exists!")
		fmt.Printf("Retrieving the existing record.")
		incidentRecordJSONOld, errR := RetrieveIncident(stub, incidentRecord.IncidentID)

		if (errR != nil)  {
			return false, fmt.Errorf("[ITOpsChaincode]: Error in retrieving Incident record.")
		}
		
		//unmarshal incidentRecordJSONOld into struct and update the rows
		var incidentRecordOld data.IncidentDO
		unmarshalErr := json.Unmarshal([]byte(incidentRecordJSONOld), &incidentRecordOld)

		if (unmarshalErr != nil) {
			return false, fmt.Errorf("[ITOpsChaincode]: CreateIncident - Error in unmarshalling JSON string to Incident record.")
		}
		
		//update the rows with the same incident record individually
		
		success11, err11 := UpdateIncident(stub, incidentRecordOld, incidentRecord, "IncidentID")
		if ((!success11) || (err11 != nil)) {
	 		return false, fmt.Errorf("Error in updating Incident record.")
		}
		
		
		success22, err22 := UpdateIncident(stub, incidentRecordOld, incidentRecord, "IncidentTitle")
		if ((!success22) || (err22 != nil)) {
	 		return false, fmt.Errorf("Error in updating Incident record.")
		}
		
		success33, err33 := UpdateIncident(stub, incidentRecordOld, incidentRecord, "Severity")
		if ((!success33) || (err33 != nil)) {
	 		return false, fmt.Errorf("Error in updating Incident record.")
		}
		
		success44, err44 := UpdateIncident(stub, incidentRecordOld, incidentRecord, "Status")
		if ((!success44) || (err44 != nil)) {
	 		return false, fmt.Errorf("Error in updating Incident record.")
		}
		
		success55, err55 := UpdateIncident(stub, incidentRecordOld, incidentRecord, "ContactEmail")
		if ((!success55) || (err55 != nil)) {
	 		return false, fmt.Errorf("Error in updating Incident record.")
		}
		

	}	
	
	success2, err2 := stub.InsertRow("INCIDENT", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.IncidentTitle}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
		},
	})
	
	if ((err2 != nil) || (!success2)) {
		return false, fmt.Errorf("Error in creating Incident record.")
	}
	
	success3, err3 := stub.InsertRow("INCIDENT", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.Severity}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
		},
	})	
	
	if ((err3 != nil) || (!success3)) {
		return false, fmt.Errorf("Error in creating Incident record.")
	}
	
	success4, err4 := stub.InsertRow("INCIDENT", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.Status}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
		},
	})	

	if ((err4 != nil) || (!success4)) {
		return false, fmt.Errorf("Error in creating Incident record.")
	}
	
	success5, err5 := stub.InsertRow("INCIDENT", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: incidentRecord.ContactEmail}},
			&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
		},
	})
	
	if ((err5 != nil) || (!success5)) {
		return false, fmt.Errorf("Error in creating Incident record.")
	}
	

	fmt.Println("Incident record created/updated. Incident Id : [%s]", string(incidentRecord.IncidentID))
	
	/*var customEvent = "{eventType: 'Creation', description:" + incidentRecord.IncidentID + "' Successfully created'}"
	errE := stub.SetEvent("evtSender", []byte(customEvent))
    	if errE != nil {
		return false, fmt.Errorf("Error in event 'create'.")
    	}*/

    	fmt.Println("Successfully saved changes")

	return true, nil
}

func UpdateIncident(stub shim.ChaincodeStubInterface, incidentRecordOld data.IncidentDO, incidentRecord data.IncidentDO, option string) (bool, error) {
	fmt.Println("Updating incident record...")

	incidentRecordBytes, marshalErr := json.Marshal(incidentRecord)

	if (marshalErr != nil) {
		return false, fmt.Errorf("Error in marshalling Incident record.")
	}

	incidentJSON := string(incidentRecordBytes)
		
	switch option {
		
	case "IncidentID":
		
		success, err := stub.ReplaceRow("INCIDENT", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: incidentRecordOld.IncidentID}},
				&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
			},
		})

		if ((err != nil) || (!success)) {
			return false, fmt.Errorf("Error in updating Incident record.")
		}			
		
	case "IncidentTitle":
		
		success, err := stub.ReplaceRow("INCIDENT", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: incidentRecordOld.IncidentTitle}},
				&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
			},
		})

		if ((err != nil) || (!success)) {
			return false, fmt.Errorf("Error in updating Incident record.")
		}	
		
	case "Severity":
		
		success, err := stub.ReplaceRow("INCIDENT", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: incidentRecordOld.Severity}},
				&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
			},
		})

		if ((err != nil) || (!success)) {
			return false, fmt.Errorf("Error in updating Incident record.")
		}
		
	case "Status":
		
		success, err := stub.ReplaceRow("INCIDENT", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: incidentRecordOld.Status}},
				&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
			},
		})

		if ((err != nil) || (!success)) {
			return false, fmt.Errorf("Error in updating Incident record.")
		}
		
	case "ContactEmail":
		
		success, err := stub.ReplaceRow("INCIDENT", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: incidentRecordOld.ContactEmail}},
				&shim.Column{Value: &shim.Column_String_{String_: incidentJSON}},
			},
		})

		if ((err != nil) || (!success)) {
			return false, fmt.Errorf("Error in updating Incident record.")
		}
		
	default:		
		
		fmt.Println("Invalid option")
		return false, fmt.Errorf("Error in updating Incident record.")		
	
	}

	fmt.Println("Incident record updated. Incident Id : [%s]", string(incidentRecord.IncidentID))
	
	/*var customEvent = "{eventType: 'Update', description:" + incidentRecord.IncidentID + "' Successfully updated status'}"
	errE := stub.SetEvent("evtSender", []byte(customEvent))
	if errE != nil {
		return false, fmt.Errorf("Error in event 'update'.")
	}*/
	
	fmt.Println("Successfully updated changes")

	return true, nil	
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
	
	var jsonRespBuffer bytes.Buffer
	jsonRespBuffer.WriteString(row.Columns[1].GetString_())

	
	
	col2 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[1].GetString_()}}
    	columns = append(columns, col2)	
	
	rowChannel, err := stub.GetRows("INCIDENT", columns)
  	if err != nil {
    		return "", fmt.Errorf("getRows operation failed. ")
  	}
	
	var rows []shim.Row
  	for {
    		select {
    			case row, ok := <-rowChannel:
      				if !ok {
        				rowChannel = nil
      				} else {
					fmt.Printf("Row - [%s]", row)
					fmt.Printf("")
        				rows = append(rows, row)
      				}
    		}
    		if rowChannel == nil {
     			break
    		}
  	}
	
	jsonRows, err := json.Marshal(rows)
  	if err != nil {
    		return "", fmt.Errorf("getRows operation failed. Error marshaling JSON:")
  	}
	
	return string(jsonRows), nil
	
	/*
	row, err := stub.GetRow("INCIDENT", columns)

	if err != nil {
		fmt.Printf("Error retrieving Incident record [%s]: [%s]", string(incidentId), err)
		fmt.Println()
		return "", fmt.Errorf("Error retrieving Incident record [%s]: [%s]", string(incidentId), err)
	}

	fmt.Printf("Row - [%s]", row)
	fmt.Println()
	
	var jsonRespBuffer bytes.Buffer
	jsonRespBuffer.WriteString(row.Columns[1].GetString_())

	return jsonRespBuffer.String(), nil*/

}

/*
Creating the Incident table
*/

/*func CreateIncidentTable(stub shim.ChaincodeStubInterface) ([]byte, error) {

	fmt.Println("Creating Incident Table ...")

	//if len(args) != 0 {
	//	return nil, fmt.Errorf("Incorrect number of arguments. Expecting 0")
	//}

	
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
}*/
