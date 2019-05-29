package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	UserId  string             `bson:"userId"`
	Token   string             `bson:"token"`
	OpenId  string             `bson:"openId"`
	Name    string             `bson:"name"`
	Cardnum string             `bson:"cardnum"`
}
