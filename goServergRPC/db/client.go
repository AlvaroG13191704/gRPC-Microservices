package db

import (
	confproto "conferencia/goClientgRPC/appointmentProto"
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var lock = &sync.Mutex{}

var mongoClient *mongo.Client

func Connect() *mongo.Client {

	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
		panic(err)
	}

	return client
}

func GetMongoInstance() *mongo.Client {

	if mongoClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if mongoClient == nil {
			fmt.Println("Creating single mongo instance now.")
			mongoClient = Connect()
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return mongoClient
}

type AppointmentCollection struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func (pc *AppointmentCollection) Connect() {
	pc.Client = GetMongoInstance()
	pc.Coll = pc.Client.Database("clinic").Collection("appointments")
}

// Function to get the data from the database
func (ap *AppointmentCollection) GetAppointments(doctorID string) ([]*confproto.Appointment, error) {
	ap.Connect()

	// var appointments []models.Appointment
	var appointments []*confproto.Appointment

	filter := bson.D{{Key: "doctor_id", Value: doctorID}}

	cursor, err := ap.Coll.Find(context.Background(), filter)
	if err != nil {
		log.Panic(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var appointment *confproto.Appointment
		cursor.Decode(&appointment)
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}
