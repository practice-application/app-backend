package store

import (
	"context"
	"fmt"
	"github.com/practice-application/app-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (s *Store) AddProduct(prd model.Product) {
	insertResult, err := s.prodColl.InsertOne(context.Background(), prd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) GetProduct(id string) (model.Product, error) {
	var prd model.Product

	err := s.prodColl.FindOne(
		context.Background(),
		bson.M{"id": id},
	).Decode(&prd)
	if err != nil {
		return model.Product{}, err
	}

	return prd, nil
}

func (s *Store) GetProducts(nm, ctg, searchText string, limit, skip *int64) (model.ProductPage, error) {

	filter := bson.M{}

	if nm != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"name": nm},
		}}
	}

	if ctg != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"category": ctg},
		}}
	}

	if searchText != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"$text": bson.M{"$search": searchText}},
		}}
	}

	opt := options.FindOptions{
		Skip:  skip,
		Limit: limit,
		Sort:  bson.M{"name": -1},
	}

	mctx := context.Background()
	cursor, err := s.prodColl.Find(mctx, filter, &opt)
	if err != nil {
		return model.ProductPage{}, err
	}

	// unpack results
	var pg model.ProductPage
	if err := cursor.All(mctx, &pg.Data); err != nil {
		return model.ProductPage{}, err
	}
	if pg.Matches, err = s.prodColl.CountDocuments(mctx, filter); err != nil {
		return model.ProductPage{}, err
	}
	return pg, nil
}

func (s *Store) UpdateProduct(id string, prd model.Product) {
	insertResult, err := s.prodColl.ReplaceOne(context.Background(), bson.M{"id": id}, prd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) DeleteProduct(id string) error {
	removeResult, err := s.prodColl.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nRemoved a Single Document: %v\n", removeResult)
	return nil
}
