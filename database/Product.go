package database

import (
	"context"
	"errors"

	"github.com/iamtbay/go-auction/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NEW PRODUCT
func (x *Product) New(productInfo *models.NewProductInfo) error {
	collection := client.Database(dbName).Collection("products")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()

	//check slug is exist or not
	collection.FindOne()

	//save to db

	//
	_, err := collection.InsertOne(ctx, productInfo)
	if err != nil {
		return err
	}

	return nil
}
//fix here
func checkSlug(isNew bool, slug string) error {
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
		var User struct {
			ID primitive.ObjectID `json:"_id" bson:"_id"`
		}

	}

}
