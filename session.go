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
	SnapshotTime() bson.Timestamp

	// Functions to modify mutable session properties.

	AdvanceClusterTime(d bson.Raw) error
	AdvanceOperationTime(ts *bson.Timestamp) error
}

type session struct {
	ss *mongo.Session
	cl *client
}

// StartTransaction is a wrapper for `mongo.Session.StartTransaction` method
func (s *session) StartTransaction(opts ...options.Lister[options.TransactionOptions]) error {
	return s.ss.StartTransaction(opts...)
}

// AbortTransaction is a wrapper for `mongo.Session.AbortTransaction` method
func (s *session) AbortTransaction(ctx context.Context) error {
	return s.ss.AbortTransaction(ctx)
}

// CommitTransaction is a wrapper for `mongo.Session.CommitTransaction` method
func (s *session) CommitTransaction(ctx context.Context) error {
	return s.ss.CommitTransaction(ctx)
}

// WithTransaction is a wrapper for `mongo.Session.WithTransaction` method
func (s *session) WithTransaction(
	ctx context.Context,
	fn func(ctx context.Context) (any, error),
	opts ...options.Lister[options.TransactionOptions],
) (any, error) {
	return s.ss.WithTransaction(ctx, fn, opts...)
}

// EndSession is a wrapper for `mongo.Session.EndSession` method
func (s *session) EndSession(ctx context.Context) {
	s.ss.EndSession(ctx)
}

// ClusterTime is a wrapper for `mongo.Session.ClusterTime` method
func (s *session) ClusterTime() bson.Raw {
	return s.ss.ClusterTime()
}

// OperationTime is a wrapper for `mongo.Session.OperationTime` method
func (s *session) OperationTime() *bson.Timestamp {
	return s.ss.OperationTime()
}

// Client is a wrapper for `mongo.Session.Client` method
func (s *session) Client() Client {
	return s.cl
}

// ID is a wrapper for `mongo.Session.ID` method
func (s *session) ID() bson.Raw {
	return s.ss.ID()
}

// SnapshotTime is a wrapper for `mongo.Session.SnapshotTime` method
func (s *session) SnapshotTime() bson.Timestamp {
	return s.ss.SnapshotTime()
}

// AdvanceClusterTime is a wrapper for `mongo.Session.AdvanceClusterTime` method
func (s *session) AdvanceClusterTime(d bson.Raw) error {
	return s.ss.AdvanceClusterTime(d)
}

// AdvanceOperationTime is a wrapper for `mongo.Session.AdvanceOperationTime` method
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

// SessionFromContext is a wrapper for `mongo.SessionFromContext` function to return the object as `Session` interface
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#SessionFromContext
func SessionFromContext(ctx context.Context) Session {
	ss := mongo.SessionFromContext(ctx)
	if ss == nil {
		return nil
	}
	return WrapSession(ss)
}
