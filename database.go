package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database is an interface for `mongo.Database` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Database
type Database interface {
	Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (Cursor, error)
	Client() Client
	Collection(name string, opts ...*options.CollectionOptions) Collection
	CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error
	CreateView(
		ctx context.Context,
		viewName, viewOn string,
		pipeline interface{},
		opts ...*options.CreateViewOptions,
	) error
	Drop(ctx context.Context) error

	WrappedDatabase() *mongo.Database
}

type database struct {
	db *mongo.Database
	cl *client
}

func (d *database) Aggregate(
	ctx context.Context,
	pipeline interface{},
	opts ...*options.AggregateOptions,
) (Cursor, error) {
	cr, err := d.db.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

func (d *database) Client() Client {
	return d.cl
}

func (d *database) Collection(name string, opts ...*options.CollectionOptions) Collection {
	return wrapCollection(d.db.Collection(name, opts...), d)
}

func (d *database) CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error {
	return d.db.CreateCollection(ctx, name, opts...)
}

func (d *database) CreateView(
	ctx context.Context,
	viewName, viewOn string,
	pipeline interface{},
	opts ...*options.CreateViewOptions,
) error {
	return d.db.CreateView(ctx, viewName, viewOn, pipeline, opts...)
}

func (d *database) Drop(ctx context.Context) error {
	return d.db.Drop(ctx)
}

func (d *database) WrappedDatabase() *mongo.Database {
	return d.db
}

func wrapDatabase(db *mongo.Database, cl *client) Database {
	return &database{
		db: db,
		cl: cl,
	}
}
