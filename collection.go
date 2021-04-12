package mongoifc

import "go.mongodb.org/mongo-driver/mongo"

// Collection is an interface for `mongo.Collection` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection
type Collection interface {
	WrappedCollection() *mongo.Collection
}

type collection struct {
	co *mongo.Collection
	db *database
}

func (c *collection) WrappedCollection() *mongo.Collection {
	return c.co
}

func wrapCollection(co *mongo.Collection, db *database) Collection {
	return &collection{co: co, db: db}
}
