package database

import (
	"context"
	"log"
	"time"

	"github.com/mikkybang/go-cms/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(config.AppConfig.MongoURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	DB = client
}

func GetCollection(collection string) *mongo.Collection {
	return DB.Database(config.AppConfig.DBName).Collection(collection)
}
