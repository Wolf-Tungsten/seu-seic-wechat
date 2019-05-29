package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"wechat-bind/secret"
)

var (
	Client *mongo.Client
)

func Connect() {

	client, err := mongo.NewClient(options.Client().ApplyURI(secret.MongodbUrl))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	Client = client

}
