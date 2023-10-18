package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterModel struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoModel struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Email    string             `json:"email" bson:"email"`
	Username string             `json:"username" bson:"username"`
}
type GetUserDBModel struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Email    string             `json:"email" bson:"email"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
}
