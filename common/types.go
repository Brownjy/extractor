package common

import "context"

// DocumentDB is a simple abstraction of a document databse with only insert & delete methods exported
type DocumentDB interface {
	Insert(ctx context.Context, col string, docs []interface{}) (int, error)
	Delete(ctx context.Context, col string, filter interface{}) (int, error)
	Aggregate(ctx context.Context, col string, pipeline interface{}, res interface{}) error
}
