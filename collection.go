package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Collection is an interface for `mongo.Collection` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#Collection
type Collection interface {
	Aggregate(ctx context.Context, pipeline any, opts ...options.Lister[options.AggregateOptions]) (Cursor, error)
	BulkWrite(
		ctx context.Context,
		models []mongo.WriteModel,
		opts ...options.Lister[options.BulkWriteOptions],
	) (*mongo.BulkWriteResult, error)
	Clone(opts ...options.Lister[options.CollectionOptions]) Collection
	CountDocuments(
		ctx context.Context,
		filter any,
		opts ...options.Lister[options.CountOptions],
	) (int64, error)
	Database() Database
	DeleteMany(
		ctx context.Context,
		filter any,
		opts ...options.Lister[options.DeleteManyOptions],
	) (*mongo.DeleteResult, error)
	DeleteOne(
		ctx context.Context,
		filter any,
		opts ...options.Lister[options.DeleteOneOptions],
	) (*mongo.DeleteResult, error)
	Distinct(
		ctx context.Context,
		fieldName string,
		filter any,
		opts ...options.Lister[options.DistinctOptions],
	) DistinctResult
	Drop(ctx context.Context) error
	EstimatedDocumentCount(
		ctx context.Context,
		opts ...options.Lister[options.EstimatedDocumentCountOptions],
	) (int64, error)
	Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (Cursor, error)
	FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) SingleResult
	FindOneAndDelete(
		ctx context.Context,
		filter any,
		opts ...options.Lister[options.FindOneAndDeleteOptions],
	) SingleResult
	FindOneAndReplace(
		ctx context.Context,
		filter any,
		replacement any,
		opts ...options.Lister[options.FindOneAndReplaceOptions],
	) SingleResult
	FindOneAndUpdate(
		ctx context.Context,
		filter any,
		update any,
		opts ...options.Lister[options.FindOneAndUpdateOptions],
	) SingleResult
	Indexes() IndexView
	InsertMany(
		ctx context.Context,
		documents []any,
		opts ...options.Lister[options.InsertManyOptions],
	) (*mongo.InsertManyResult, error)
	InsertOne(
		ctx context.Context,
		document any,
		opts ...options.Lister[options.InsertOneOptions],
	) (*mongo.InsertOneResult, error)
	Name() string
	ReplaceOne(
		ctx context.Context,
		filter any,
		replacement any,
		opts ...options.Lister[options.ReplaceOptions],
	) (*mongo.UpdateResult, error)
	SearchIndexes() SearchIndexView
	UpdateByID(
		ctx context.Context,
		id any,
		update any,
		opts ...options.Lister[options.UpdateOneOptions],
	) (*mongo.UpdateResult, error)
	UpdateMany(
		ctx context.Context,
		filter any,
		update any,
		opts ...options.Lister[options.UpdateManyOptions],
	) (*mongo.UpdateResult, error)
	UpdateOne(
		ctx context.Context,
		filter any,
		update any,
		opts ...options.Lister[options.UpdateOneOptions],
	) (*mongo.UpdateResult, error)
	Watch(
		ctx context.Context,
		pipeline any,
		opts ...options.Lister[options.ChangeStreamOptions],
	) (ChangeStream, error)
}

type collection struct {
	co *mongo.Collection
	db *database
}

// Aggregate is a wrapper for `mongo.Collection.Aggregate` method
func (c *collection) Aggregate(
	ctx context.Context,
	pipeline any,
	opts ...options.Lister[options.AggregateOptions],
) (Cursor, error) {
	cr, err := c.co.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

// BulkWrite is a wrapper for `mongo.Collection.BulkWrite` method
func (c *collection) BulkWrite(
	ctx context.Context,
	models []mongo.WriteModel,
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	return c.co.BulkWrite(ctx, models, opts...)
}

// Clone is a wrapper for `mongo.Collection.Clone` method
func (c *collection) Clone(opts ...options.Lister[options.CollectionOptions]) Collection {
	return wrapCollection(c.co.Clone(opts...), c.db)
}

// CountDocuments is a wrapper for `mongo.Collection.CountDocuments` method
func (c *collection) CountDocuments(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.CountOptions],
) (int64, error) {
	return c.co.CountDocuments(ctx, filter, opts...)
}

// Database is a wrapper for `mongo.Collection.Database` method
func (c *collection) Database() Database {
	return c.db
}

// DeleteMany is a wrapper for `mongo.Collection.DeleteMany` method
func (c *collection) DeleteMany(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.DeleteManyOptions],
) (*mongo.DeleteResult, error) {
	return c.co.DeleteMany(ctx, filter, opts...)
}

// DeleteOne is a wrapper for `mongo.Collection.DeleteOne` method
func (c *collection) DeleteOne(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.DeleteOneOptions],
) (*mongo.DeleteResult, error) {
	return c.co.DeleteOne(ctx, filter, opts...)
}

// Distinct is a wrapper for `mongo.Collection.Distinct` method
func (c *collection) Distinct(
	ctx context.Context,
	fieldName string,
	filter any,
	opts ...options.Lister[options.DistinctOptions],
) DistinctResult {
	return wrapDistinctResult(c.co.Distinct(ctx, fieldName, filter, opts...))
}

// Drop is a wrapper for `mongo.Collection.Drop` method
func (c *collection) Drop(ctx context.Context) error {
	return c.co.Drop(ctx)
}

// EstimatedDocumentCount is a wrapper for `mongo.Collection.EstimatedDocumentCount` method
func (c *collection) EstimatedDocumentCount(
	ctx context.Context,
	opts ...options.Lister[options.EstimatedDocumentCountOptions],
) (int64, error) {
	return c.co.EstimatedDocumentCount(ctx, opts...)
}

// Find is a wrapper for `mongo.Collection.Find` method
func (c *collection) Find(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.FindOptions],
) (Cursor, error) {
	cr, err := c.co.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

// FindOne is a wrapper for `mongo.Collection.FindOne` method
func (c *collection) FindOne(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.FindOneOptions],
) SingleResult {
	return wrapSingleResult(c.co.FindOne(ctx, filter, opts...))
}

// FindOneAndDelete is a wrapper for `mongo.Collection.FindOneAndDelete` method
func (c *collection) FindOneAndDelete(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.FindOneAndDeleteOptions],
) SingleResult {
	return wrapSingleResult(c.co.FindOneAndDelete(ctx, filter, opts...))
}

// FindOneAndReplace is a wrapper for `mongo.Collection.FindOneAndReplace` method
func (c *collection) FindOneAndReplace(
	ctx context.Context,
	filter any,
	replacement any,
	opts ...options.Lister[options.FindOneAndReplaceOptions],
) SingleResult {
	return wrapSingleResult(c.co.FindOneAndReplace(ctx, filter, replacement, opts...))
}

// FindOneAndUpdate is a wrapper for `mongo.Collection.FindOneAndUpdate` method
func (c *collection) FindOneAndUpdate(
	ctx context.Context,
	filter any,
	update any,
	opts ...options.Lister[options.FindOneAndUpdateOptions],
) SingleResult {
	return wrapSingleResult(c.co.FindOneAndUpdate(ctx, filter, update, opts...))
}

// Indexes is a wrapper for `mongo.Collection.Indexes` method
func (c *collection) Indexes() IndexView {
	iv := c.co.Indexes()
	return wrapIndexView(&iv)
}

// InsertMany is a wrapper for `mongo.Collection.InsertMany` method
func (c *collection) InsertMany(
	ctx context.Context,
	documents []any,
	opts ...options.Lister[options.InsertManyOptions],
) (*mongo.InsertManyResult, error) {
	return c.co.InsertMany(ctx, documents, opts...)
}

// InsertOne is a wrapper for `mongo.Collection.InsertOne` method
func (c *collection) InsertOne(
	ctx context.Context,
	document any,
	opts ...options.Lister[options.InsertOneOptions],
) (*mongo.InsertOneResult, error) {
	return c.co.InsertOne(ctx, document, opts...)
}

// Name is a wrapper for `mongo.Collection.Name` method
func (c *collection) Name() string {
	return c.co.Name()
}

// ReplaceOne is a wrapper for `mongo.Collection.ReplaceOne` method
func (c *collection) ReplaceOne(
	ctx context.Context,
	filter any,
	replacement any,
	opts ...options.Lister[options.ReplaceOptions],
) (*mongo.UpdateResult, error) {
	return c.co.ReplaceOne(ctx, filter, replacement, opts...)
}

// SearchIndexes is a wrapper for `mongo.Collection.SearchIndexes` method
func (c *collection) SearchIndexes() SearchIndexView {
	siv := c.co.SearchIndexes()
	return wrapSearchIndexView(&siv)
}

// UpdateByID is a wrapper for `mongo.Collection.UpdateByID` method
func (c *collection) UpdateByID(
	ctx context.Context,
	id any,
	update any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.co.UpdateByID(ctx, id, update, opts...)
}

// UpdateMany is a wrapper for `mongo.Collection.UpdateMany` method
func (c *collection) UpdateMany(
	ctx context.Context,
	filter any,
	update any,
	opts ...options.Lister[options.UpdateManyOptions],
) (*mongo.UpdateResult, error) {
	return c.co.UpdateMany(ctx, filter, update, opts...)
}

// UpdateOne is a wrapper for `mongo.Collection.UpdateOne` method
func (c *collection) UpdateOne(
	ctx context.Context,
	filter any,
	update any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.co.UpdateOne(ctx, filter, update, opts...)
}

// Watch is a wrapper for `mongo.Collection.Watch` method
func (c *collection) Watch(
	ctx context.Context,
	pipeline any,
	opts ...options.Lister[options.ChangeStreamOptions],
) (ChangeStream, error) {
	cs, err := c.co.Watch(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return wrapChangeStream(cs), nil
}

func wrapCollection(co *mongo.Collection, db *database) Collection {
	return &collection{co: co, db: db}
}

// WrapCollection returns an instance of Collection interface for given mongo.Collection object
func WrapCollection(co *mongo.Collection) Collection {
	return wrapCollection(co, WrapDatabase(co.Database()).(*database))
}

// UnWrapCollection returns original mongo.Collection
func UnWrapCollection(co Collection) *mongo.Collection {
	return co.(*collection).co
}
