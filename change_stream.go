package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// ChangeStream is an interface for `mongo.ChangeStream` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#ChangeStream
type ChangeStream interface {
	Current() bson.Raw
	Close(ctx context.Context) error
	Decode(val any) error
	Err() error
	ID() int64
	Next(ctx context.Context) bool
	RemainingBatchLength() int
	ResumeToken() bson.Raw
	SetBatchSize(size int32)
	TryNext(ctx context.Context) bool
}

type changeStream struct {
	cs *mongo.ChangeStream
}

// Current is a wrapper for `mongo.ChangeStream.Current` field
func (c *changeStream) Current() bson.Raw {
	return c.cs.Current
}

// Close is a wrapper for `mongo.ChangeStream.Close` method
func (c *changeStream) Close(ctx context.Context) error {
	return c.cs.Close(ctx)
}

// Decode is a wrapper for `mongo.ChangeStream.Decode` method
func (c *changeStream) Decode(val any) error {
	return c.cs.Decode(val)
}

// Err is a wrapper for `mongo.ChangeStream.Err` method
func (c *changeStream) Err() error {
	return c.cs.Err()
}

// ID is a wrapper for `mongo.ChangeStream.ID` method
func (c *changeStream) ID() int64 {
	return c.cs.ID()
}

// Next is a wrapper for `mongo.ChangeStream.Next` method
func (c *changeStream) Next(ctx context.Context) bool {
	return c.cs.Next(ctx)
}

// RemainingBatchLength is a wrapper for `mongo.ChangeStream.RemainingBatchLength` method
func (c *changeStream) RemainingBatchLength() int { return c.cs.RemainingBatchLength() }

// ResumeToken is a wrapper for `mongo.ChangeStream.ResumeToken` method
func (c *changeStream) ResumeToken() bson.Raw {
	return c.cs.ResumeToken()
}

// SetBatchSize is a wrapper for `mongo.ChangeStream.SetBatchSize` method
func (c *changeStream) SetBatchSize(size int32) {
	c.cs.SetBatchSize(size)
}

// TryNext is a wrapper for `mongo.ChangeStream.TryNext` method
func (c *changeStream) TryNext(ctx context.Context) bool {
	return c.cs.TryNext(ctx)
}

func wrapChangeStream(cs *mongo.ChangeStream) ChangeStream {
	return &changeStream{cs: cs}
}
