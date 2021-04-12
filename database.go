package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
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
	ListCollectionNames(
		ctx context.Context,
		filter interface{},
		opts ...*options.ListCollectionsOptions,
	) ([]string, error)
	ListCollections(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) (Cursor, error)
	ListCollectionSpecifications(
		ctx context.Context,
		filter interface{},
		opts ...*options.ListCollectionsOptions,
	) ([]*mongo.CollectionSpecification, error)
	Name() string
	ReadConcern() *readconcern.ReadConcern
	ReadPreference() *readpref.ReadPref
	RunCommand(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) SingleResult
	RunCommandCursor(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) (Cursor, error)
	Watch(
		ctx context.Context,
		pipeline interface{},
		opts ...*options.ChangeStreamOptions,
	) (ChangeStream, error)
	WriteConcern() *writeconcern.WriteConcern

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

func (d *database) ListCollectionNames(
	ctx context.Context,
	filter interface{},
	opts ...*options.ListCollectionsOptions,
) ([]string, error) {
	return d.db.ListCollectionNames(ctx, filter, opts...)
}

func (d *database) ListCollections(
	ctx context.Context,
	filter interface{},
	opts ...*options.ListCollectionsOptions,
) (Cursor, error) {
	cr, err := d.db.ListCollections(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

func (d *database) ListCollectionSpecifications(
	ctx context.Context,
	filter interface{},
	opts ...*options.ListCollectionsOptions,
) ([]*mongo.CollectionSpecification, error) {
	return d.db.ListCollectionSpecifications(ctx, filter, opts...)
}

func (d *database) Name() string {
	return d.db.Name()
}

func (d *database) ReadConcern() *readconcern.ReadConcern {
	return d.db.ReadConcern()
}

func (d *database) ReadPreference() *readpref.ReadPref {
	return d.db.ReadPreference()
}

func (d *database) RunCommand(
	ctx context.Context,
	runCommand interface{},
	opts ...*options.RunCmdOptions,
) SingleResult {
	return wrapSingleResult(d.db.RunCommand(ctx, runCommand, opts...))
}

func (d *database) RunCommandCursor(
	ctx context.Context,
	runCommand interface{},
	opts ...*options.RunCmdOptions,
) (Cursor, error) {
	cr, err := d.db.RunCommandCursor(ctx, runCommand, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

func (d *database) Watch(
	ctx context.Context,
	pipeline interface{},
	opts ...*options.ChangeStreamOptions,
) (ChangeStream, error) {
	cs, err := d.db.Watch(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return wrapChangeStream(cs), nil
}

func (d *database) WriteConcern() *writeconcern.WriteConcern {
	return d.db.WriteConcern()
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
