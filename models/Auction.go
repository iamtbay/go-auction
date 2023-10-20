package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewAuction struct {
	ProductID  primitive.ObjectID `json:"product_id" bson:"product_id"`
	SellerID   primitive.ObjectID `json:"seller_id" bson:"seller_id"`
	StartPrice float64            `json:"start_price" bson:"start_price"`
	Duration   time.Time          `json:"duration" bson:"duration"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
}

type GetAuction struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	ProductID  primitive.ObjectID `json:"product_id" bson:"product_id"`
	SellerID   primitive.ObjectID `json:"seller_id" bson:"seller_id"`
	Duration   time.Time          `json:"duration" bson:"duration"`
	StartPrice float64            `json:"start_price" bson:"start_price"`
}
