package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewRepository is factory
func NewRepository() *Repository {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	return &Repository{
		client: client,
	}
}

// Repository is struct
type Repository struct {
	client *mongo.Client
}

// Put is to create item
func (repository *Repository) Put(id string, data interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := repository.client.Database("testing").Collection("numbers")
	collection.InsertOne(ctx, bson.M{"id": id, "data": data})

	fmt.Println("Putting data using mongdb database")
	fmt.Printf("Id: %s , data %v \n", id, data)
}

// Find is to find item
func (repository *Repository) Find(id string) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := repository.client.Database("testing").Collection("numbers")
	data := collection.FindOne(ctx, bson.M{"id": id})

	fmt.Println("Finding data using mongdb database")
	fmt.Printf("Id: %s , data %v \n", id, data)

	return data
}
