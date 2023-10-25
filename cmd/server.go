package main

import (
	"context"
	"fmt"
	"os"

	"github.com/iamtbay/go-auction/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		fmt.Println("db start error!")
		panic(err)
	}
	fmt.Println("Db Started")
	database.InitClient(client)
	fmt.Println("Client ready to connection")
}
