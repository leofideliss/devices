package repository

import (
	"context"
	"log"
    "os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}
    
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("DB_HOST")))
	if err != nil {
		log.Fatal(err)
	}
    
	DB = client.Database(os.Getenv("DATABASE"))
}

type MongoRepository[T any] struct{
    collection *mongo.Collection
}

func NewMongoRepository[T any](db *mongo.Database , collectionName string) *MongoRepository[T]{
    return &MongoRepository[T]{collection:db.Collection(collectionName)}
}

func (r *MongoRepository[T]) FindById( ctx context.Context , id string) (*T , error){
    var entity T
    err := r.collection.FindOne(ctx,bson.M{"_id":id}).Decode(&entity)
    return &entity , err
}

func (r *MongoRepository[T]) Insert(ctx context.Context , entity *T) (*mongo.InsertOneResult , error){
    result , err := r.collection.InsertOne(ctx,entity)
    return result,err
}
