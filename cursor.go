package mongoifc

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

// Cursor is an interface for `mongo.Cursor` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#Cursor
type Cursor interface {
	Current() bson.Raw
	All(ctx context.Context, results any) error
	Close(ctx context.Context) error
	Decode(val any) error
	Err() error
	ID() int64
	Next(ctx context.Context) bool
	RemainingBatchLength() int
	SetBatchSize(batchSize int32)
	SetComment(comment any)
	SetMaxAwaitTime(dur time.Duration)
	TryNext(ctx context.Context) bool
}

type cursor struct {
	cr *mongo.Cursor
}

func (c *cursor) Current() bson.Raw {
	return c.cr.Current
}

func (c *cursor) All(ctx context.Context, results any) error {
	return c.cr.All(ctx, results)
}

func (c *cursor) Close(ctx context.Context) error {
	return c.cr.Close(ctx)
}

func (c *cursor) Decode(val any) error {
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

func (c *cursor) SetComment(comment any) {
	c.cr.SetComment(comment)
}

func (c *cursor) SetMaxAwaitTime(dur time.Duration) {
	c.cr.SetMaxAwaitTime(dur)
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
// https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#NewCursorFromDocuments
func NewCursorFromDocuments(documents []any, err error, registry *bson.Registry) (Cursor, error) {
	cr, err := mongo.NewCursorFromDocuments(documents, err, registry)
	if err != nil {
		return nil, err
	}
	return wrapCursor(cr), nil
}
