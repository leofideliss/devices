package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func init(){
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
    
	DB = client.Database("ms_devices")
}

type MongoRepository[T any] struct{
    collection *mongo.Collection
}

func NewMongoRepository[T any](db *mongo.Database , collectionName string) *MongoRepository[T]{
    return &MongoRepository[T]{collection:db.Collection(collectionName)}
}

func (r *MongoRepository[T]) FindById( ctx context.Context , id primitive.ObjectID) (*T , error){
    var entity T
    err := r.collection.FindOne(ctx,bson.M{"_id":id}).Decode(&entity)
    return &entity , err
}
