package database

import (
	"context"
	"errors"
	"time"

	"github.com/iamtbay/go-auction/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (x *Auction) New(currentUserID primitive.ObjectID, auctionInfo *models.NewAuction) error {
	collection := client.Database(dbName).Collection("auctions")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()
	//
	//check product exist or not
	var productInfo *models.GetProductAndOwner
	err := client.Database(dbName).Collection("products").FindOne(ctx, bson.M{"_id": auctionInfo.ProductID}).Decode(&productInfo)
	if err != nil {
		return errors.New("product couldn't find")
	}
	//check product belong current user or not
	if productInfo.SellerID != currentUserID {
		return errors.New("product isn't belong to you")
	}
	//
	auctionInfo.CreatedAt = time.Now()
	auctionInfo.Duration = time.Now().Add(48 * time.Hour)
	_, err = collection.InsertOne(ctx, auctionInfo)
	if err != nil {
		return err
	}
	return nil
}

func (x *Auction) Get(productID primitive.ObjectID) (*models.GetAuction, error) {
	//collection
	collection := client.Database(dbName).Collection("auctions")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()
	//
	var auctionInfo *models.GetAuction
	err := collection.FindOne(ctx, bson.M{"_id": productID}).Decode(&auctionInfo)
	if err != nil {
		return nil, err
	}
	return auctionInfo, nil
}
