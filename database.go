package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Database is an interface for `mongo.Database` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#Database
type Database interface {
	Aggregate(
		ctx context.Context,
		pipeline any,
		opts ...options.Lister[options.AggregateOptions],
	) (Cursor, error)
	Client() Client
	Collection(name string, opts ...options.Lister[options.CollectionOptions]) Collection
	CreateCollection(
		ctx context.Context,
		name string,
		opts ...options.Lister[options.CreateCollectionOptions],
	) error
	CreateView(
		ctx context.Context,
		viewName, viewOn string,
		pipeline any,
		opts ...options.Lister[options.CreateViewOptions],
	) error
	Drop(ctx context.Context) error
	GridFSBucket(opts ...options.Lister[options.BucketOptions]) GridFSBucket
	ListCollectionNames(
		ctx context.Context,
		filter any,
		opts ...options.Lister[options.ListCollectionsOptions],
	) ([]string, error)
	ListCollections(
		ctx context.Context,
		filter any,
		opts ...options.Lister[options.ListCollectionsOptions],
	) (Cursor, error)
	ListCollectionSpecifications(
		ctx context.Context,
		filter any,
		opts ...options.Lister[options.ListCollectionsOptions],
	) ([]mongo.CollectionSpecification, error)
	Name() string
	RunCommand(
		ctx context.Context,
		runCommand any,
		opts ...options.Lister[options.RunCmdOptions],
	) SingleResult
	RunCommandCursor(
		ctx context.Context,
		runCommand any,
		opts ...options.Lister[options.RunCmdOptions],
	) (Cursor, error)
	Watch(
		ctx context.Context,
		pipeline any,
		opts ...options.Lister[options.ChangeStreamOptions],
	) (ChangeStream, error)
}

type database struct {
	db *mongo.Database
	cl *client
}

// Aggregate is a wrapper for `mongo.Database.Aggregate` method
func (d *database) Aggregate(
	ctx context.Context,
	pipeline any,
	opts ...options.Lister[options.AggregateOptions],
) (Cursor, error) {
	cr, err := d.db.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

// Client is a wrapper for `mongo.Database.Client` method
func (d *database) Client() Client {
	return d.cl
}

// Collection is a wrapper for `mongo.Database.Collection` method
func (d *database) Collection(
	name string,
	opts ...options.Lister[options.CollectionOptions],
) Collection {
	return wrapCollection(d.db.Collection(name, opts...), d)
}

// CreateCollection is a wrapper for `mongo.Database.CreateCollection` method
func (d *database) CreateCollection(
	ctx context.Context,
	name string,
	opts ...options.Lister[options.CreateCollectionOptions],
) error {
	return d.db.CreateCollection(ctx, name, opts...)
}

// CreateView is a wrapper for `mongo.Database.CreateView` method
func (d *database) CreateView(
	ctx context.Context,
	viewName, viewOn string,
	pipeline any,
	opts ...options.Lister[options.CreateViewOptions],
) error {
	return d.db.CreateView(ctx, viewName, viewOn, pipeline, opts...)
}

// Drop is a wrapper for `mongo.Database.Drop` method
func (d *database) Drop(ctx context.Context) error {
	return d.db.Drop(ctx)
}

// GridFSBucket is a wrapper for `mongo.Database.GridFSBucket` method
func (d *database) GridFSBucket(opts ...options.Lister[options.BucketOptions]) GridFSBucket {
	return wrapGridFSBucket(d.db.GridFSBucket(opts...))
}

// ListCollectionNames is a wrapper for `mongo.Database.ListCollectionNames` method
func (d *database) ListCollectionNames(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.ListCollectionsOptions],
) ([]string, error) {
	return d.db.ListCollectionNames(ctx, filter, opts...)
}

// ListCollections is a wrapper for `mongo.Database.ListCollections` method
func (d *database) ListCollections(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.ListCollectionsOptions],
) (Cursor, error) {
	cr, err := d.db.ListCollections(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

// ListCollectionSpecifications is a wrapper for `mongo.Database.ListCollectionSpecifications` method
func (d *database) ListCollectionSpecifications(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.ListCollectionsOptions],
) ([]mongo.CollectionSpecification, error) {
	return d.db.ListCollectionSpecifications(ctx, filter, opts...)
}

// Name is a wrapper for `mongo.Database.Name` method
func (d *database) Name() string {
	return d.db.Name()
}

// RunCommand is a wrapper for `mongo.Database.RunCommand` method
func (d *database) RunCommand(
	ctx context.Context,
	runCommand any,
	opts ...options.Lister[options.RunCmdOptions],
) SingleResult {
	return wrapSingleResult(d.db.RunCommand(ctx, runCommand, opts...))
}

// RunCommandCursor is a wrapper for `mongo.Database.RunCommandCursor` method
func (d *database) RunCommandCursor(
	ctx context.Context,
	runCommand any,
	opts ...options.Lister[options.RunCmdOptions],
) (Cursor, error) {
	cr, err := d.db.RunCommandCursor(ctx, runCommand, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

// Watch is a wrapper for `mongo.Database.Watch` method
func (d *database) Watch(
	ctx context.Context,
	pipeline any,
	opts ...options.Lister[options.ChangeStreamOptions],
) (ChangeStream, error) {
	cs, err := d.db.Watch(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return wrapChangeStream(cs), nil
}

func wrapDatabase(db *mongo.Database, cl *client) Database {
	return &database{
		db: db,
		cl: cl,
	}
}

// WrapDatabase returns an instance of Database interface for given mongo.Database object
func WrapDatabase(db *mongo.Database) Database {
	return wrapDatabase(db, WrapClient(db.Client()).(*client))
}

// UnWrapDatabase returns original mongo.Database
func UnWrapDatabase(db Database) *mongo.Database {
	return db.(*database).db
}
