package mgoutil

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

// NewMgoDocDB returns an instance of MgoDocDB
func NewMgoDocDB(ctx context.Context, cli *mongo.Client, db *mongo.Database) (*MgoDocDB, error) {
	mdb := &MgoDocDB{
		cli:       cli,
		db:        db,
		insertOpt: options.InsertMany().SetOrdered(false),
		aggOpt:    options.Aggregate(),
	}

	mdb.cols.m = make(map[string]*mongo.Collection)
	return mdb, nil
}

// MgoDocDB is an implementation of the common.DocumentDB, based on mongo
type MgoDocDB struct {
	cli *mongo.Client
	db  *mongo.Database

	insertOpt *options.InsertManyOptions
	aggOpt    *options.AggregateOptions

	cols struct {
		sync.RWMutex
		m map[string]*mongo.Collection
	}
}

// Connect attempts to establish a client with the given dsn
func Connect(ctx context.Context, dsn string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	//client, err := mongo.NewClient(options.Client().ApplyURI(dsn).SetAppName("bell"))
	if err != nil {
		return nil, fmt.Errorf("new client: %w", err)
	}

	return client, nil
}
