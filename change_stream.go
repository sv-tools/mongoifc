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

func (c *changeStream) Current() bson.Raw {
	return c.cs.Current
}

func (c *changeStream) Close(ctx context.Context) error {
	return c.cs.Close(ctx)
}

func (c *changeStream) Decode(val any) error {
	return c.cs.Decode(val)
}

func (c *changeStream) Err() error {
	return c.cs.Err()
}

func (c *changeStream) ID() int64 {
	return c.cs.ID()
}

func (c *changeStream) Next(ctx context.Context) bool {
	return c.cs.Next(ctx)
}

func (c *changeStream) RemainingBatchLength() int { return c.cs.RemainingBatchLength() }

func (c *changeStream) ResumeToken() bson.Raw {
	return c.cs.ResumeToken()
}

func (c *changeStream) SetBatchSize(size int32) {
	c.cs.SetBatchSize(size)
}

func (c *changeStream) TryNext(ctx context.Context) bool {
	return c.cs.TryNext(ctx)
}

func wrapChangeStream(cs *mongo.ChangeStream) ChangeStream {
	return &changeStream{cs: cs}
}
