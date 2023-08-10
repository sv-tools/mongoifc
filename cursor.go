package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
)

// Cursor is an interface for `mongo.Cursor` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Cursor
type Cursor interface {
	Current() bson.Raw
	All(ctx context.Context, results interface{}) error
	Close(ctx context.Context) error
	Decode(val interface{}) error
	Err() error
	ID() int64
	Next(ctx context.Context) bool
	RemainingBatchLength() int
	SetBatchSize(batchSize int32)
	TryNext(ctx context.Context) bool
}

type cursor struct {
	cr *mongo.Cursor
}

func (c *cursor) Current() bson.Raw {
	return c.cr.Current
}

func (c *cursor) All(ctx context.Context, results interface{}) error {
	return c.cr.All(ctx, results)
}

func (c *cursor) Close(ctx context.Context) error {
	return c.cr.Close(ctx)
}

func (c *cursor) Decode(val interface{}) error {
	return c.cr.Decode(val)
}

func (c *cursor) Err() error {
	return c.cr.Err()
}

func (c *cursor) ID() int64 {
	return c.cr.ID()
}

func (c *cursor) Next(ctx context.Context) bool {
	return c.cr.Next(ctx)
}

func (c *cursor) RemainingBatchLength() int {
	return c.cr.RemainingBatchLength()
}

func (c *cursor) TryNext(ctx context.Context) bool {
	return c.cr.TryNext(ctx)
}

func (c *cursor) SetBatchSize(batchSize int32) {
	c.cr.SetBatchSize(batchSize)
}

func wrapCursor(cr *mongo.Cursor) Cursor {
	return &cursor{cr: cr}
}

// NewCursorFromDocuments is a wrapper for NewCursorFromDocuments function of the mongodb to return Cursor
// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#NewCursorFromDocuments
func NewCursorFromDocuments(documents []interface{}, err error, registry *bsoncodec.Registry) (Cursor, error) {
	cr, err := mongo.NewCursorFromDocuments(documents, err, registry)
	if err != nil {
		return nil, err
	}
	return wrapCursor(cr), nil
}
