package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/iamtbay/go-auction/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)



// Login
func (x *Auth) Login(userInfo *models.LoginModel) (*models.UserInfoModel, error) {
	collection := client.Database(dbName).Collection("users")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()
	//db
	var dbInfo models.GetUserDBModel
	err := collection.FindOne(ctx, bson.M{"email": userInfo.Email}).Decode(&dbInfo)
	if err != nil {
		fmt.Println("email couldn't find error")
		return nil, errors.New("email or password incorrect")
	}
	//check password true or not
	err = bcrypt.CompareHashAndPassword([]byte(dbInfo.Password), []byte(userInfo.Password))
	if err != nil {
		fmt.Println("compare password error")
		return nil, errors.New("email or password incorrect")
	}

	return &models.UserInfoModel{
		ID:       dbInfo.ID,
		Email:    dbInfo.Email,
		Username: dbInfo.Username,
	}, nil

}

// Register
func (x *Auth) Register(userInfo *models.RegisterModel) error {
	collection := client.Database(dbName).Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()
	//check email is exist
	err := checkEmail(true, userInfo.Email, primitive.NilObjectID)
	if err != nil {
		return err
	}
	//check username is exist
	err = checkUsername(true, userInfo.Username, primitive.NilObjectID)
	if err != nil {
		return err
	}
	//hash password
	hashedPassword, err := hashPassword(userInfo.Password)
	if err != nil {
		return err
	}
	userInfo.Password = hashedPassword

	//save to db
	_, err = collection.InsertOne(ctx, userInfo)
	if err != nil {
		return err
	}
	//
	return nil
}

// Update
func (x *Auth) Update(userID primitive.ObjectID, userInfo *models.RegisterModel) error {
	collection := client.Database(dbName).Collection("users")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()
	//db ops

	err := checkID(userID)
	if err != nil {
		return err
	}
	err = checkEmail(false, userInfo.Email, userID)
	if err != nil {
		return err
	}
	err = checkUsername(false, userInfo.Username, userID)
	if err != nil {
		return err
	}
	//hash pass
	hashedPassword, err := hashPassword(userInfo.Password)
	if err != nil {
		return err
	}
	userInfo.Password = hashedPassword

	//update
	update := bson.M{
		"$set": bson.M{
			"email":    userInfo.Email,
			"username": userInfo.Username,
			"password": userInfo.Password,
		},
	}
	updateRes, err := collection.UpdateOne(ctx, bson.M{"_id": userID}, update)
	if err != nil {
		return err
	}
	fmt.Println(updateRes)

	return nil
}

// HELPERS
func checkID(val primitive.ObjectID) error {
	collection := client.Database(dbName).Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()
	count, err := collection.CountDocuments(ctx, bson.M{"_id": val})
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("ups, user couldn't find")
	}
	return nil
}

type UsStr struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
}

// SOME BUG CAN BE HERE CHECK AFTER
// check email
func checkEmail(isRegister bool, email string, userID primitive.ObjectID) error {
	collection := client.Database(dbName).Collection("users")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()

	if isRegister {
		count, err := collection.CountDocuments(ctx, bson.M{"email": email})
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("email already in use please try another email")
		}
	} else {
		var userIDDB UsStr
		err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&userIDDB)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil
			}
			return err
		}
		if userID != userIDDB.ID {
			return errors.New("Email already in use")
		}
		return nil
	}
	return nil
}

// check username
func checkUsername(isRegister bool, username string, userID primitive.ObjectID) error {
	collection := client.Database(dbName).Collection("users")
	//ctx
	ctx, cancel := context.WithTimeout(context.TODO(), dbTimeout)
	defer cancel()

	if isRegister {
		count, err := collection.CountDocuments(ctx, bson.M{"username": username})
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("username already in use please try another username")
		}
	} else {
		var userIDDB UsStr
		err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&userIDDB)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil
			}
			return err
		}
		if userIDDB.ID != userID {
			return errors.New("username already in use")
		}
		return nil

	}
	return nil
}

// hash pass
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
