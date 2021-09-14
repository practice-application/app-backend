package store

import (
	"context"
	"fmt"
	"log"

	"github.com/practice-application/app-backend/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Store) AddProduct(prd model.Product) {
	insertResult, err := s.Product.InsertOne(context.Background(), prd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) GetProduct(id string) (model.Product, error) {
	var prd model.Product
	err := s.Product.FindOne(
		context.Background(),
		bson.M{"id": id},
	).Decode(&prd)
	if err != nil {
		return model.Product{}, err
	}

	return prd, nil
}

func (s *Store) UpdateProduct(id string, prd model.Product) {
	insertResult, err := s.Product.ReplaceOne(context.Background(), bson.M{"id": id}, prd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) DeleteProduct(id string) error {
	removeResult, err := s.Product.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nRemoved a Single Document: %v\n", removeResult)
	return nil
}
