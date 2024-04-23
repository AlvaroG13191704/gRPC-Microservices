package models

type Appointment struct {
	ID          string `json:"id" bson:"_id"`
	PatientName string `json:"patientName" bson:"patientName"`
	Description string `json:"description" bson:"description"`
	DoctorID    string `json:"doctorId" bson:"doctorId"`
}
