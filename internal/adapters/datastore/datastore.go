package datastore

import (
	"context"
)

// Datastore defines the interface that abstracts data storage/access implementations.
type Datastore interface {
	WithContext(ctx context.Context) Datastore
	Create(value interface{}) Datastore
	Where(query interface{}, args ...interface{}) Datastore
	First(value interface{}) Datastore

	Error() error
}
