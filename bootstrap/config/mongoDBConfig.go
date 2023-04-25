package config

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	ctx context.Context
}

func NewMongoDBConfig() *MongoDBConfig {
	return &MongoDBConfig{
		ctx: context.TODO(),
	}
}

func (mongoDBConfig MongoDBConfig) NewMongoDBClient() (*mongo.Client) {
	clientOptions := options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27117")
	client, err := mongo.Connect(mongoDBConfig.ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(mongoDBConfig.ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (mongoDBConfig MongoDBConfig) Context() (context.Context) {
	return mongoDBConfig.ctx
}