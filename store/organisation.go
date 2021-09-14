package store

import (
	"context"
	"fmt"
	"log"

	"github.com/practice-application/app-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Organisation

func (s *Store) AddOrg(o model.Org) {
	insertResult, err := s.Org.InsertOne(context.Background(), o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) GetOrg(id string) (model.Org, error) {
	var o model.Org
	err := s.Org.FindOne(
		context.Background(),
		bson.M{"id": id},
	).Decode(&o)
	if err != nil {
		return model.Org{}, err
	}

	return o, nil
}
func (s *Store) GetOrganisations(on, ot, searchText string, limit, skip *int64) ([]model.Org, error) {

	filter := bson.M{}

	if on != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"organisationName": on},
		}}
	}

	if ot != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"orgType": ot},
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
		Sort:  bson.M{"organisationName": -1},
	}

	mctx := context.Background()
	cursor, err := s.Org.Find(mctx, filter, &opt)
	if err != nil {
		return []model.Org{}, nil
	}

	// unpack results
	var org []model.Org
	if err := cursor.All(mctx, &org); err != nil {
		return []model.Org{}, nil
	}

	return org, nil
}

func (s *Store) UpdateOrg(id string, o model.Org) {
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
