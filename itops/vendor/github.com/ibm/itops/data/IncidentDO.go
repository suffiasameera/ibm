

package data


import (

)

//Data structure for an Incident record

type IncidentDO struct {
	IncidentID 					string `json:"incidentID"`	//Identification of the incident record as referred by the originator
	IncidentTitle 			string `json:"incidentTitle"` //Short description/title
	IncidentType        string `json:"incidentType"` //Type of the Incident
	Severity         		string `json:"severity"` //Sevirty of the Incident
	Status							string `json:"status"` //Status of the incident record
	RefIncidentID 			string `json:"refIncidentID"` //Id of the incident based on which current record is created
	OriginalIncidentID 	string `json:"originalIncidentIDd"` //Id of the original incident record
	ParticipantIDFrom		string `json:"participantIDFrom"` //ID of the participant who originate
	ParticipantIDTo			string `json:"participantIDTo"` //ID of the participant intended to
	ContactEmail 				string `json:"contactEmail"` //Emain ID of the contact
	CreatedDate					string `json:"createdDate"` //Record create date
	ExpectedCloseDate 	string `json:"expectedCloseDate"` //Expected date of closure filled by SLA manager
	ActualCloseDate 		string `json:"actualCloseDate"` //Actual Date of closure
}
