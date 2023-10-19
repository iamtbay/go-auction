package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewProductInfo struct {
	Name         string             `json:"name" bson:"name"`
	Brand        string             `json:"brand" bson:"brand"`
	Category     string             `json:"category" bson:"category"`
	Info         string             `json:"info" bson:"info"`
	Slug         string             `json:"slug" bson:"slug"`
	Photos       []string           `json:"photos" bson:"photos"`
	SellerID     primitive.ObjectID `json:"seller_id" bson:"seller_id"`
	IsOpenToSell bool               `json:"is_open_to_sell" bson:"is_open_to_sell"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
}

type GetProductInfo struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Brand        string             `json:"brand" bson:"brand"`
	Category     string             `json:"category" bson:"category"`
	Info         string             `json:"info" bson:"info"`
	Slug         string             `json:"slug" bson:"slug"`
	Photos       []string           `json:"photos" bson:"photos"`
	SellerID     primitive.ObjectID `json:"seller_id" bson:"seller_id"`
	IsOpenToSell bool               `json:"is_open_to_sell" bson:"is_open_to_sell"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
}

type DeleteProduct struct {
	ID string `json:"_id" bson:"_id"`
}
