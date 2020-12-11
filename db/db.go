package db

import (
	"sync"
	"log"
	"context"
	"fmt"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

var instance *mongo.Database
var m sync.Mutex

func GetDB() *mongo.Database {
	if instance == nil {
        m.Lock()
        defer m.Unlock()
        if instance == nil {
			uri := "mongodb://192.168.219.201:27017"
			clientOptions := options.Client().ApplyURI(uri)
		
			// Connect to MongoDB
			client, err := mongo.Connect(context.TODO(), clientOptions)
		
			if err != nil {
				log.Fatal(err)
			}
			
			// defer func() {
			// 	if err = client.Disconnect(context.TODO()); err != nil {
			// 		panic(err)
			// 	}
			// }()
		
			// Check the connection
			err = client.Ping(context.TODO(), nil)
		
			if err != nil {
				log.Fatal(err)
			}

			db := client.Database("avi")
			fmt.Println("Connected to MongoDB!")
            instance = db
        }
    }
    return instance
}