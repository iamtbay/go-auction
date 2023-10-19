package database

import (
	"context"
	"errors"

	"github.com/iamtbay/go-auction/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GET PRODUCT

func (x *Product) Get(slug string) (*models.GetProductInfo, error) {
	collection := client.Database(dbName).Collection("products")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()

	//db
	var productInfo *models.GetProductInfo
	err := collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&productInfo)
	if err != nil {
		return nil, err
	}
	return productInfo, nil

}

// NEW PRODUCT
func (x *Product) New(productInfo *models.NewProductInfo) error {
	collection := client.Database(dbName).Collection("products")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()

	//check slug is exist or not
	err := checkSlug(true, productInfo.Slug, primitive.NilObjectID, primitive.NilObjectID)
	if err != nil {
		return err
	}
	//save to db
	_, err = collection.InsertOne(ctx, productInfo)
	if err != nil {
		return err
	}

	return nil
}

// Update Product
func (x *Product) Update(userID primitive.ObjectID, productInfo *models.GetProductInfo) error {
	collection := client.Database(dbName).Collection("products")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()
	//check slug is exist or not
	err := checkSlug(false, productInfo.Slug, userID, productInfo.ID)
	if err != nil {
		return err
	}
	//save to db
	update := bson.M{
		"$set": productInfo,
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": productInfo.ID}, update)
	if err != nil {
		return err
	}

	return nil
}

// Delete Product
func (x *Product) Delete(userID, productID primitive.ObjectID) error {
	collection := client.Database(dbName).Collection("products")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()

	//
	filter := bson.M{
		"_id":       productID,
		"seller_id": userID,
	}
	var productInfo *models.GetProductInfo
	err := collection.FindOne(ctx, filter).Decode(&productInfo)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}

// check slug
func checkSlug(isNew bool, slug string, sellerID, productID primitive.ObjectID) error {
	collection := client.Database(dbName).Collection("products")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()

	//operation
	if isNew {
		count, err := collection.CountDocuments(ctx, bson.M{"slug": slug})
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("slug is in use, try another slug")
		}

	} else {
		var user struct {
			ProductID primitive.ObjectID `json:"_id" bson:"_id"`
			SellerID  primitive.ObjectID `json:"seller_id" bson:"seller_id"`
		}
		//check slug is exist
		err := collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil
			}
			return err
		}
		//if slug exist but belong to another product send err
		if user.ProductID != productID {
			return errors.New("oh, seems like slug in use try another")
		}
		//if anyone try to update product than who created send unauthorized err
		if user.SellerID != sellerID {
			return errors.New("unauthorized operation")
		}
	}
	return nil

}
