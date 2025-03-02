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
    if os.Getenv("TEST_ENV") == "true" {
        return
    }
    
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
    return r.collection.InsertOne(ctx,entity)
}

func (r *MongoRepository[T]) Delete(ctx context.Context ,id string) (*mongo.DeleteResult , error) {
    return r.collection.DeleteOne(ctx,bson.M{"_id":id})
}

func (r *MongoRepository[T]) Update(ctx context.Context , entity *T , id string) (*mongo.UpdateResult , error){
    entityBson, err := bson.Marshal(entity)
    if err != nil {
        return nil, err
    }
    var updateData bson.M
    err = bson.Unmarshal(entityBson, &updateData)
    if err != nil {
        return nil, err
    }
    delete(updateData, "_id")
    return r.collection.UpdateOne(ctx,bson.M{"_id":id},bson.M{"$set":updateData })
}

func (r *MongoRepository[T]) FindAll (ctx context.Context , owner string , limit , page int) ([]T,error){
    var results []T
    filter := bson.M{"owner":owner}
    skip := (page -1) * limit

    cursor, err := r.collection.Find(ctx, filter, options.Find().
        SetLimit(int64(limit)).
        SetSkip(int64(skip)).
        SetSort(bson.M{"_id": -1}))
    if err != nil {
        return nil, err
    }
    
    defer cursor.Close(ctx)

    if err = cursor.All(ctx, &results); err != nil {
        return nil, err
    }

    return results, nil
}
