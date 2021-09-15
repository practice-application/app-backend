package store

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	persColl *mongo.Collection
	orgColl  *mongo.Collection
	prodColl *mongo.Collection
}

func Connect() *Store {
	clientOptions := options.Client().ApplyURI("mongodb+srv://databaseUser:b6kz42hs@data.zugzp.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("app")

	return &Store{
		persColl: db.Collection("person"),
		orgColl:  db.Collection("organisation"),
		prodColl: db.Collection("product"),
	}
}
