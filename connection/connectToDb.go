package connection

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

const databaseurl = "mongodb://localhost:27017/drywave/?timeoutMS=5000"

func ConnectToDb() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(databaseurl))

	if err != nil {
		return err
	}

	database = client.Database("drywave")
	return nil

}

func GetDBCollection(collection string) *mongo.Collection {
	return database.Collection(collection)
}
