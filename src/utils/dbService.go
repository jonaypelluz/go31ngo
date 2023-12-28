package utils

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type DBService interface {
	InsertOne(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, collection string, documents []interface{}) (*mongo.InsertManyResult, error)
	DeleteOne(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)
	UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	FindBy(ctx context.Context, collection string, filter interface{}) (*mongo.SingleResult, error)
}

type MongoService struct {
	client       *mongo.Client
	databaseName string
}

func TheMongoService(client *mongo.Client) *MongoService {
	db := os.Getenv("MONGO_DB_NAME")
	return &MongoService{
		client:       client,
		databaseName: db,
	}
}

func (m *MongoService) InsertOne(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	return m.client.Database(m.databaseName).Collection(collection).InsertOne(ctx, document)
}

func (m *MongoService) InsertMany(ctx context.Context, collection string, documents []interface{}) (*mongo.InsertManyResult, error) {
	return m.client.Database(m.databaseName).Collection(collection).InsertMany(ctx, documents)
}

func (m *MongoService) DeleteOne(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	return m.client.Database(m.databaseName).Collection(collection).DeleteOne(ctx, filter)
}

func (m *MongoService) DeleteMany(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	return m.client.Database(m.databaseName).Collection(collection).DeleteMany(ctx, filter)
}

func (m *MongoService) UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return m.client.Database(m.databaseName).Collection(collection).UpdateOne(ctx, filter, update)
}

func (m *MongoService) FindBy(ctx context.Context, collection string, filter interface{}) (*mongo.SingleResult, error) {
	return m.client.Database(m.databaseName).Collection(collection).FindOne(ctx, filter), nil
}
