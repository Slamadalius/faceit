package mongoDB

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(url string) (client *mongo.Client, err error) {
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(url)); err != nil {
		return
	}

	err = client.Ping(context.TODO(), nil)
	return
}
