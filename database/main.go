package database

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client *mongo.Client
)
var dbTimeout = time.Second * 10
var dbName = "go-auction"

func InitClient(cli *mongo.Client) {
	client = cli
}
