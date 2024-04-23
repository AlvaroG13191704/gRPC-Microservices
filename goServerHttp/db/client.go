package db

import (
	"context"
	"fmt"
	"log"
	"sync"

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
