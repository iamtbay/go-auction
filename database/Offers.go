package database

import (
	"context"
	"errors"
	"math"

	"github.com/iamtbay/go-auction/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NEW OFFER
func (x *Offers) NewOffer(offer *models.OfferInfo) error {
	collection := client.Database(dbName).Collection("offers")
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()

	//find max
	opts := options.FindOne().SetSort(bson.M{"offer": -1})
	var maxBid struct {
		Offer float64 `json:"offer" bson:"offer"`
	}

	err := collection.FindOne(ctx, bson.M{"auction_id": offer.AuctionID}, opts).Decode(&maxBid)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}
	if maxBid.Offer > 0 && math.Abs(maxBid.Offer-offer.Offer) < 1 {
		return errors.New("please raise your bid at least more $1 than current bid")
	}
	//add
	_, err = collection.InsertOne(ctx, offer)
	if err != nil {
		return err
	}

	return nil
}
