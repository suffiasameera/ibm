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
																				Payment Services
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/

func CreatePaymentTable(stub shim.ChaincodeStubInterface) ([]byte, error) {

	fmt.Println("Creating Payment Table ...")

	// Create Payment table
	err := stub.CreateTable("PAYMENT", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "payment_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "payment_record", Type: shim.ColumnDefinition_STRING, Key: true},
	})

	if err != nil {
		return nil, fmt.Errorf("Failed creating Payment table.")
	}

	fmt.Println("Payment table initialization done successfully... !!! ")

	return nil, nil
}

/*
	Create Payment record
*/
func CreatePayment(stub shim.ChaincodeStubInterface, paymentJSON string) ([]byte, error) {

	fmt.Println("Creating Payment record ...")

	if paymentJSON == "" {
		return nil, fmt.Errorf("Payment record not sent")
	}

	paymentRecord := new(data.PaymentDO)
	paymentByteArray := []byte(paymentJSON)
	unmarshalErr := json.Unmarshal(paymentByteArray,paymentRecord)

	if (unmarshalErr != nil) {
		return nil, fmt.Errorf("Error in unmarshalling JSON string to Incident record.")
	}

	success, err := stub.InsertRow("PAYMENT", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: paymentRecord.PaymentID}},
			&shim.Column{Value: &shim.Column_String_{String_: paymentJSON}},
		},
	})

	if ((err != nil) || (!success)) {
		return nil, fmt.Errorf("Error in creating Payment record.")
	}

	// if (!(success && (err == nil))) {
	// 	return nil, fmt.Errorf("Error in creating Payment record.")
	// }

	fmt.Println("Payment record created. Payment Id : [%s]", string(paymentRecord.PaymentID))

	return nil, err
}


func UpdatePayment(stub shim.ChaincodeStubInterface, paymentJSON string) ([]byte, error) {
	fmt.Println("Updating Payment record ...")

	if paymentJSON == "" {
		return nil, fmt.Errorf("Payment record not sent")
	}

	paymentRecord := new(data.PaymentDO)
	paymentByteArray := []byte(paymentJSON)
	unmarshalErr := json.Unmarshal(paymentByteArray,paymentRecord)

	if (unmarshalErr != nil) {
		return nil, fmt.Errorf("Error in unmarshalling JSON string to Incident record.")
	}

	success, err := stub.ReplaceRow("PAYMENT", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: paymentRecord.PaymentID}},
			&shim.Column{Value: &shim.Column_String_{String_: paymentJSON}},
		},
	})

	if ((err != nil) || (!success)) {
		return nil, fmt.Errorf("Error in updating Payment record.")
	}

	// if (!(success && (err == nil))) {
	// 	return nil, fmt.Errorf("Error in updating Payment record.")
	// }

	fmt.Println("Payment record updated. Payment Id : [%s]", string(paymentRecord.PaymentID))

	return nil, err
}

/*
 Retrieve Payment record
*/
func RetrievePayment(stub shim.ChaincodeStubInterface, paymentId string) (string, error) {

	fmt.Println("Retrieveing Payment record. Payment Id : [%s]", string(paymentId))

	var columns []shim.Column
	paymentIdColumn := shim.Column{Value: &shim.Column_String_{String_: paymentId}}
	columns = append(columns, paymentIdColumn)
	row, err := stub.GetRow("PAYMENT", columns)

	if err != nil {
		fmt.Printf("Error retriving Payment record [%s]: [%s]", paymentId, err)
		fmt.Println()
		return "", fmt.Errorf("Error retriving Payment record [%s]: [%s]", paymentId, err)
	}

	fmt.Printf("Row - [%s]", row)
	fmt.Println()

	var jsonRespBuffer bytes.Buffer
	jsonRespBuffer.WriteString(row.Columns[1].GetString_())

	return jsonRespBuffer.String(), nil
}
