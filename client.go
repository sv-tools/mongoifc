package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// Client is an interface for `mongo.Client` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#Client
type Client interface {
	Database(name string, opts ...options.Lister[options.DatabaseOptions]) Database
	Disconnect(ctx context.Context) error
	ListDatabaseNames(
		ctx context.Context,
		filter any,
		opts ...options.Lister[options.ListDatabasesOptions],
	) ([]string, error)
	ListDatabases(
		ctx context.Context,
		filter any,
		opts ...options.Lister[options.ListDatabasesOptions],
	) (mongo.ListDatabasesResult, error)
	NumberSessionsInProgress() int
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	StartSession(opts ...options.Lister[options.SessionOptions]) (Session, error)
	UseSession(
		ctx context.Context,
		fn func(ctx context.Context) error,
	) error
	UseSessionWithOptions(
		ctx context.Context,
		opts *options.SessionOptionsBuilder,
		fn func(ctx context.Context) error,
	) error
	Watch(
		ctx context.Context,
		pipeline any,
		opts ...options.Lister[options.ChangeStreamOptions],
	) (ChangeStream, error)
}

type client struct {
	cl *mongo.Client
}

func (c *client) Database(name string, opts ...options.Lister[options.DatabaseOptions]) Database {
	return wrapDatabase(c.cl.Database(name, opts...), c)
}

func (c *client) Disconnect(ctx context.Context) error {
	return c.cl.Disconnect(ctx)
}

func (c *client) ListDatabaseNames(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.ListDatabasesOptions],
) ([]string, error) {
	return c.cl.ListDatabaseNames(ctx, filter, opts...)
}

func (c *client) ListDatabases(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.ListDatabasesOptions],
) (mongo.ListDatabasesResult, error) {
	return c.cl.ListDatabases(ctx, filter, opts...)
}

func (c *client) NumberSessionsInProgress() int {
	return c.cl.NumberSessionsInProgress()
}

func (c *client) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return c.cl.Ping(ctx, rp)
}

func (c *client) StartSession(opts ...options.Lister[options.SessionOptions]) (Session, error) {
	ss, err := c.cl.StartSession(opts...)
	if err != nil {
		return nil, err
	}

	return wrapSession(ss, c), nil
}

func (c *client) UseSession(ctx context.Context, fn func(ctx context.Context) error) error {
	return c.cl.UseSession(ctx, fn)
}

func (c *client) UseSessionWithOptions(
	ctx context.Context,
	opts *options.SessionOptionsBuilder,
	fn func(ctx context.Context) error,
) error {
	return c.cl.UseSessionWithOptions(ctx, opts, fn)
}

func (c *client) Watch(
	ctx context.Context,
	pipeline any,
	opts ...options.Lister[options.ChangeStreamOptions],
) (ChangeStream, error) {
	cs, err := c.cl.Watch(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return wrapChangeStream(cs), nil
}

// WrapClient returns an instance of Client interface for given mongo.Client object
func WrapClient(cl *mongo.Client) Client {
	return &client{cl: cl}
}

// UnWrapClient returns original mongo.Client
func UnWrapClient(cl Client) *mongo.Client {
	return cl.(*client).cl
}

// Connect is a wrapper for `mongo.Connect` function to return the object as `Client` interface
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#Connect
func Connect(opts ...*options.ClientOptions) (Client, error) {
	cl, err := mongo.Connect(opts...)
	if err != nil {
		return nil, err
	}

	return WrapClient(cl), nil
}

// WithSession is a wrapper for `mongo.WithSession` function to call then `mongo.WithSession` function
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#WithSession
func WithSession(ctx context.Context, sess Session, fn func(ctx context.Context) error) error {
	return mongo.WithSession(ctx, sess.(*session).ss, fn)
}
