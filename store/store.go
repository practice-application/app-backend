package store

import (
	"context"
	"fmt"
	"log"

	"github.com/el-zacharoo/go-101/data"
	"go.mongodb.org/mongo-driver/bson"
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

// Person

func (s *Store) AddPerson(p data.Person) {
	insertResult, err := s.Customer.InsertOne(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) GetPerson(id string) (data.Person, error) {

	var p data.Person
	err := s.Customer.FindOne(
		context.Background(),
		bson.M{"id": id},
	).Decode(&p)
	if err != nil {
		return data.Person{}, err
	}

	return p, nil
}

func (s *Store) GetPeople(fn, ln, searchText string, limit *int64) ([]data.Person, error) {

	filter := bson.M{}

	if fn != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"firstname": fn},
		}}
	}

	if ln != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"lastname": ln},
		}}
	}

	if searchText != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"$text": bson.M{"$search": searchText}},
		}}
	}

	opt := options.FindOptions{
		// Skip:  offset,
		Limit: limit,
		Sort:  bson.M{"lastname": -1},
	}

	mctx := context.Background()
	cursor, err := s.Customer.Find(mctx, filter, &opt)
	if err != nil {
		return []data.Person{}, nil
	}

	// unpack results
	var ppl []data.Person
	if err := cursor.All(mctx, &ppl); err != nil {
		return []data.Person{}, nil
	}

	return ppl, nil
}

func (s *Store) UpdatePerson(id string, p data.Person) {
	insertResult, err := s.Customer.ReplaceOne(context.Background(), bson.M{"id": id}, p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) DeleteUser(id string) error {
	removeResult, err := s.Customer.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {

		return err
	}
	fmt.Printf("\nRemoved a Single Document: %v\n", removeResult)
	return nil
}

// Organisation

func (s *Store) AddOrg(o data.Org) {
	insertResult, err := s.Org.InsertOne(context.Background(), o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) GetOrg(id string) (data.Org, error) {
	var o data.Org
	err := s.Org.FindOne(
		context.Background(),
		bson.M{"id": id},
	).Decode(&o)
	if err != nil {
		return data.Org{}, err
	}

	return o, nil
}

func (s *Store) UpdateOrg(id string, o data.Org) {
	insertResult, err := s.Org.ReplaceOne(context.Background(), bson.M{"id": id}, o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) DeleteOrg(id string) error {
	removeResult, err := s.Org.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nRemoved a Single Document: %v\n", removeResult)
	return nil
}

// Product

func (s *Store) AddProduct(prd data.Product) {
	insertResult, err := s.Product.InsertOne(context.Background(), prd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) GetProduct(id string) (data.Product, error) {
	var prd data.Product
	err := s.Product.FindOne(
		context.Background(),
		bson.M{"id": id},
	).Decode(&prd)
	if err != nil {
		return data.Product{}, err
	}

	return prd, nil
}

func (s *Store) UpdateProduct(id string, prd data.Product) {
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
