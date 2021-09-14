package store

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Org      *mongo.Collection
	Customer *mongo.Collection
	Product  *mongo.Collection
}

func (s *Store) Connect() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://databaseUser:b6kz42hs@data.zugzp.mongodb.net/")
	//ctx := context.Background()
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	s.Customer = client.Database("customer-details").Collection("data")
	fmt.Print("Connected to Mongo Database!\n")

	s.Org = client.Database("org-details").Collection("data")
	fmt.Print("Connected to Mongo Database!\n")

	s.Product = client.Database("product-details").Collection("data")
	fmt.Print("Connected to Mongo Database!\n")
}
