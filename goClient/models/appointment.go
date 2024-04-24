package models

type Appointment struct {
	ID          string `json:"id" bson:"_id"`
	PatientName string `json:"patient_name" bson:"patient_name"`
	Description string `json:"description" bson:"description"`
	DoctorID    string `json:"doctor_id" bson:"doctor_id"`
}
