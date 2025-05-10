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

func (d *database) Client() Client {
	return d.cl
}

func (d *database) Collection(
	name string,
	opts ...options.Lister[options.CollectionOptions],
) Collection {
	return wrapCollection(d.db.Collection(name, opts...), d)
}

func (d *database) CreateCollection(
	ctx context.Context,
	name string,
	opts ...options.Lister[options.CreateCollectionOptions],
) error {
	return d.db.CreateCollection(ctx, name, opts...)
}

func (d *database) CreateView(
	ctx context.Context,
	viewName, viewOn string,
	pipeline any,
	opts ...options.Lister[options.CreateViewOptions],
) error {
	return d.db.CreateView(ctx, viewName, viewOn, pipeline, opts...)
}

func (d *database) Drop(ctx context.Context) error {
	return d.db.Drop(ctx)
}
func (d *database) GridFSBucket(opts ...options.Lister[options.BucketOptions]) GridFSBucket {
	return wrapGridFSBucket(d.db.GridFSBucket(opts...))
}

func (d *database) ListCollectionNames(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.ListCollectionsOptions],
) ([]string, error) {
	return d.db.ListCollectionNames(ctx, filter, opts...)
}

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

func (d *database) ListCollectionSpecifications(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.ListCollectionsOptions],
) ([]mongo.CollectionSpecification, error) {
	return d.db.ListCollectionSpecifications(ctx, filter, opts...)
}

func (d *database) Name() string {
	return d.db.Name()
}

func (d *database) RunCommand(
	ctx context.Context,
	runCommand any,
	opts ...options.Lister[options.RunCmdOptions],
) SingleResult {
	return wrapSingleResult(d.db.RunCommand(ctx, runCommand, opts...))
}

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
