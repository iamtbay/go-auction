package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OfferInfo struct {
	AuctionID primitive.ObjectID `json:"auction_id" bson:"auction_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Offer     float64            `json:"offer" bson:"offer"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
