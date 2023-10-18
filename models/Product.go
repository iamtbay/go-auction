package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewProductInfo struct {
	Name         string             `json:"name"`
	Brand        string             `json:"brand"`
	Category     string             `json:"category"`
	Info         string             `json:"info"`
	Slug         string             `json:"slug"`
	Photos       []string           `json:"photos"`
	SellerID     primitive.ObjectID `json:"seller_id"`
	IsOpenToSell bool               `json:"is_open_to_sell"`
	CreatedAt    time.Time          `json:"created_at"`
}

type GetProductInfo struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	NewProductInfo
}
