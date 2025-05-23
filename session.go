package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Session is an interface for `mongo.Session` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#Session
type Session interface {
	// Functions to modify session state.

	StartTransaction(opts ...options.Lister[options.TransactionOptions]) error
	AbortTransaction(ctx context.Context) error
	CommitTransaction(ctx context.Context) error
	WithTransaction(
		ctx context.Context,
		fn func(ctx context.Context) (any, error),
		opts ...options.Lister[options.TransactionOptions],
	) (any, error)
	EndSession(ctx context.Context)

	// Functions to retrieve session properties.

	ClusterTime() bson.Raw
	OperationTime() *bson.Timestamp
	Client() Client
	ID() bson.Raw

	// Functions to modify mutable session properties.

	AdvanceClusterTime(d bson.Raw) error
	AdvanceOperationTime(ts *bson.Timestamp) error
}

type session struct {
	ss *mongo.Session
	cl *client
}

func (s *session) StartTransaction(opts ...options.Lister[options.TransactionOptions]) error {
	return s.ss.StartTransaction(opts...)
}

func (s *session) AbortTransaction(ctx context.Context) error {
	return s.ss.AbortTransaction(ctx)
}

func (s *session) CommitTransaction(ctx context.Context) error {
	return s.ss.CommitTransaction(ctx)
}

func (s *session) WithTransaction(
	ctx context.Context,
	fn func(ctx context.Context) (any, error),
	opts ...options.Lister[options.TransactionOptions],
) (any, error) {
	return s.ss.WithTransaction(ctx, fn, opts...)
}

func (s *session) EndSession(ctx context.Context) {
	s.ss.EndSession(ctx)
}

func (s *session) ClusterTime() bson.Raw {
	return s.ss.ClusterTime()
}

func (s *session) OperationTime() *bson.Timestamp {
	return s.ss.OperationTime()
}

func (s *session) Client() Client {
	return s.cl
}

func (s *session) ID() bson.Raw {
	return s.ss.ID()
}

func (s *session) AdvanceClusterTime(d bson.Raw) error {
	return s.ss.AdvanceClusterTime(d)
}

func (s *session) AdvanceOperationTime(ts *bson.Timestamp) error {
	return s.ss.AdvanceOperationTime(ts)
}

func wrapSession(ss *mongo.Session, cl *client) Session {
	return &session{ss: ss, cl: cl}
}

// WrapSession returns an instance of Session interface for given mongo.Session object
func WrapSession(ss *mongo.Session) Session {
	return wrapSession(ss, WrapClient(ss.Client()).(*client))
}

// UnWrapSession returns original mongo.Session
func UnWrapSession(ss Session) *mongo.Session {
	return ss.(*session).ss
}
