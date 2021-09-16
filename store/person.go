package store

import (
	"context"
	"fmt"
	"log"

	"github.com/practice-application/app-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Store) AddPerson(p model.Person) {
	insertResult, err := s.persColl.InsertOne(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) GetPerson(id string) (model.Person, error) {

	var p model.Person
	if err := s.persColl.FindOne(
		context.Background(),
		bson.M{"id": id},
	).Decode(&p); err != nil {
		return model.Person{}, err
	}

	return p, nil
}

func (s *Store) GetPeople(fn, ln, searchText string, limit, skip *int64) (model.Page, error) {

	filter := bson.M{}

	if fn != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"firstName": fn},
		}}
	}

	if ln != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"lastName": ln},
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
		Sort:  bson.M{"lastName": -1},
	}

	mctx := context.Background()
	cursor, err := s.persColl.Find(mctx, filter, &opt)
	if err != nil {
		return model.Page{}, err
	}

	// unpack results
	var pg model.Page
	if err := cursor.All(mctx, &pg.Data); err != nil {
		return model.Page{}, err
	}
	if pg.Matches, err = s.persColl.CountDocuments(mctx, filter); err != nil {
		return model.Page{}, err
	}
	return pg, nil
}

func (s *Store) UpdatePerson(id string, p model.Person) {
	insertResult, err := s.persColl.ReplaceOne(context.Background(), bson.M{"id": id}, p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) DeletePerson(id string) error {
	removeResult, err := s.persColl.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {

		return err
	}
	fmt.Printf("\nRemoved a Single Document: %v\n", removeResult)
	return nil
}
