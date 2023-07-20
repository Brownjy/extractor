package mongo

import (
	"context"
	"extractor/conf"
	"extractor/conf/storage"
	"fmt"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DB struct {
	DocDB       *mongoDriver.Database
	Collections map[string]*mongoDriver.Collection
	Client      *mongoDriver.Client
	Config      storage.Options
}

func NewMongoDB(conf *conf.Config) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	client, err := mongoDriver.Connect(ctx, options.Client().ApplyURI(conf.Storage.DSN))
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongodb %s, error: %s", conf.Storage.DSN, err)
	}
	db := client.Database("test")
	return &DB{
		DocDB: db,
		Collections: map[string]*mongoDriver.Collection{
			"messages": db.Collection("messages"),
		},
		Client: client,
		Config: conf.Storage,
	}, nil
}
