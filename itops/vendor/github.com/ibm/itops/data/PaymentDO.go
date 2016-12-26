

package data


import (

)

type PaymentDO struct {
	PaymentID						string	`json:"paymentID"`	//Identification of service payment record
	ServiceName     		string	`json:"serviceName"`	//Name of the service for which payment record is created
	DebitedFromID	     	string	`json:"debitedFromID"`	//Id of the participant who honours the service payment
	CreditedToID     		string	`json:"creditedToID"`	//Id of the participant who claims the service payment
	ServiceAgreementRef string	`json:"serviceAgreementRef"`	//Reference Id of the service agreement based on which payment claim is made
	PaymentPurpose 			string	`json:"paymentPurpose"`	//Purpose for which payment claim is made
	OrderID							string	`json:"orderID"`	//Id of the order number agains which claim is made
	PaymentDate					string	`json:"paymentDate"`	//Date on which payment claim is settled
	OriginalIncidentID	string	`json:"originalIncidentID"`	//Id of the original incident record for which payment claim is made
}
