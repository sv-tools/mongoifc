package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is an interface for `mongo.Collection` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection
type Collection interface {
	Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (Cursor, error)
	BulkWrite(
		ctx context.Context,
		models []mongo.WriteModel,
		opts ...*options.BulkWriteOptions,
	) (*mongo.BulkWriteResult, error)
	Clone(opts ...*options.CollectionOptions) (Collection, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	Database() Database
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Distinct(
		ctx context.Context,
		fieldName string,
		filter interface{},
		opts ...*options.DistinctOptions,
	) ([]interface{}, error)
	Drop(ctx context.Context) error
	EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) SingleResult
	FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) SingleResult
	FindOneAndReplace(
		ctx context.Context,
		filter interface{},
		replacement interface{},
		opts ...*options.FindOneAndReplaceOptions,
	) SingleResult
	FindOneAndUpdate(
		ctx context.Context,
		filter interface{},
		update interface{},
		opts ...*options.FindOneAndUpdateOptions,
	) SingleResult
	Indexes() IndexView
	InsertMany(
		ctx context.Context,
		documents []interface{},
		opts ...*options.InsertManyOptions,
	) (*mongo.InsertManyResult, error)
	InsertOne(
		ctx context.Context,
		document interface{},
		opts ...*options.InsertOneOptions,
	) (*mongo.InsertOneResult, error)
	Name() string
	ReplaceOne(
		ctx context.Context,
		filter interface{},
		replacement interface{},
		opts ...*options.ReplaceOptions,
	) (*mongo.UpdateResult, error)
	UpdateByID(
		ctx context.Context,
		id interface{},
		update interface{},
		opts ...*options.UpdateOptions,
	) (*mongo.UpdateResult, error)
	UpdateMany(
		ctx context.Context,
		filter interface{},
		update interface{},
		opts ...*options.UpdateOptions,
	) (*mongo.UpdateResult, error)
	UpdateOne(
		ctx context.Context,
		filter interface{},
		update interface{},
		opts ...*options.UpdateOptions,
	) (*mongo.UpdateResult, error)
	Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (ChangeStream, error)

	WrappedCollection() *mongo.Collection
}

type collection struct {
	co *mongo.Collection
	db *database
}

func (c *collection) Aggregate(
	ctx context.Context,
	pipeline interface{},
	opts ...*options.AggregateOptions,
) (Cursor, error) {
	cr, err := c.co.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

func (c *collection) BulkWrite(
	ctx context.Context,
	models []mongo.WriteModel,
	opts ...*options.BulkWriteOptions,
) (*mongo.BulkWriteResult, error) {
	return c.co.BulkWrite(ctx, models, opts...)
}

func (c *collection) Clone(opts ...*options.CollectionOptions) (Collection, error) {
	co, err := c.co.Clone(opts...)
	if err != nil {
		return nil, err
	}

	return wrapCollection(co, c.db), nil
}

func (c *collection) CountDocuments(
	ctx context.Context,
	filter interface{},
	opts ...*options.CountOptions,
) (int64, error) {
	return c.co.CountDocuments(ctx, filter, opts...)
}

func (c *collection) Database() Database {
	return c.db
}

func (c *collection) DeleteMany(
	ctx context.Context,
	filter interface{},
	opts ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	return c.co.DeleteMany(ctx, filter, opts...)
}

func (c *collection) DeleteOne(
	ctx context.Context,
	filter interface{},
	opts ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	return c.co.DeleteOne(ctx, filter, opts...)
}

func (c *collection) Distinct(
	ctx context.Context,
	fieldName string,
	filter interface{},
	opts ...*options.DistinctOptions,
) ([]interface{}, error) {
	return c.co.Distinct(ctx, fieldName, filter, opts...)
}

func (c *collection) Drop(ctx context.Context) error {
	return c.co.Drop(ctx)
}

func (c *collection) EstimatedDocumentCount(
	ctx context.Context,
	opts ...*options.EstimatedDocumentCountOptions,
) (int64, error) {
	return c.co.EstimatedDocumentCount(ctx, opts...)
}

func (c *collection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	cr, err := c.co.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

func (c *collection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) SingleResult {
	return wrapSingleResult(c.co.FindOne(ctx, filter, opts...))
}

func (c *collection) FindOneAndDelete(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOneAndDeleteOptions,
) SingleResult {
	return wrapSingleResult(c.co.FindOneAndDelete(ctx, filter, opts...))
}

func (c *collection) FindOneAndReplace(
	ctx context.Context,
	filter interface{},
	replacement interface{},
	opts ...*options.FindOneAndReplaceOptions,
) SingleResult {
	return wrapSingleResult(c.co.FindOneAndReplace(ctx, filter, replacement, opts...))
}

func (c *collection) FindOneAndUpdate(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.FindOneAndUpdateOptions,
) SingleResult {
	return wrapSingleResult(c.co.FindOneAndUpdate(ctx, filter, update, opts...))
}

func (c *collection) Indexes() IndexView {
	iv := c.co.Indexes()
	return wrapIndexView(&iv)
}

func (c *collection) InsertMany(
	ctx context.Context,
	documents []interface{},
	opts ...*options.InsertManyOptions,
) (*mongo.InsertManyResult, error) {
	return c.co.InsertMany(ctx, documents, opts...)
}

func (c *collection) InsertOne(
	ctx context.Context,
	document interface{},
	opts ...*options.InsertOneOptions,
) (*mongo.InsertOneResult, error) {
	return c.co.InsertOne(ctx, document, opts...)
}

func (c *collection) Name() string {
	return c.co.Name()
}

func (c *collection) ReplaceOne(
	ctx context.Context,
	filter interface{},
	replacement interface{},
	opts ...*options.ReplaceOptions,
) (*mongo.UpdateResult, error) {
	return c.co.ReplaceOne(ctx, filter, replacement, opts...)
}

func (c *collection) UpdateByID(
	ctx context.Context,
	id interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return c.co.UpdateByID(ctx, id, update, opts...)
}

func (c *collection) UpdateMany(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return c.co.UpdateMany(ctx, filter, update, opts...)
}

func (c *collection) UpdateOne(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return c.co.UpdateOne(ctx, filter, update, opts...)
}

func (c *collection) Watch(
	ctx context.Context,
	pipeline interface{},
	opts ...*options.ChangeStreamOptions,
) (ChangeStream, error) {
	cs, err := c.co.Watch(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return wrapChangeStream(cs), nil
}

func (c *collection) WrappedCollection() *mongo.Collection {
	return c.co
}

func wrapCollection(co *mongo.Collection, db *database) Collection {
	return &collection{co: co, db: db}
}
