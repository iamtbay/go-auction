package database

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// structs
type Auth struct{}
type Product struct{}

// init funcs
func AuthDBInit() *Auth {
	return &Auth{}
}

func ProductDBInit() *Product {
	return &Product{}
}

var (
	client    *mongo.Client
	dbTimeout = time.Second * 10
	dbName    = "go-auction"
)

func InitClient(cli *mongo.Client) {
	client = cli
}
