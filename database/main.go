package database

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// structs
type Auth struct{}
type Product struct{}
type Auction struct{}

// init funcs
func AuthDBInit() *Auth       { return &Auth{} }
func ProductDBInit() *Product { return &Product{} }
func AuctionDBInit() *Auction { return &Auction{} }

var (
	client    *mongo.Client
	dbTimeout = time.Second * 10
	dbName    = "go-auction"
)

func InitClient(cli *mongo.Client) {
	client = cli
}
