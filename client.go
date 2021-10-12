package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client is an interface for `mongo.Client` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Client
type Client interface {
	Connect(ctx context.Context) error
	Database(name string, opts ...*options.DatabaseOptions) Database
	Disconnect(ctx context.Context) error
	ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error)
	ListDatabases(
		ctx context.Context,
		filter interface{},
		opts ...*options.ListDatabasesOptions,
	) (mongo.ListDatabasesResult, error)
	NumberSessionsInProgress() int
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	StartSession(opts ...*options.SessionOptions) (Session, error)
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error
	Watch(
		ctx context.Context,
		pipeline interface{},
		opts ...*options.ChangeStreamOptions,
	) (ChangeStream, error)
	WrappedClient() *mongo.Client
}

type client struct {
	cl *mongo.Client
}

func (c *client) Connect(ctx context.Context) error {
	return c.cl.Connect(ctx)
}

func (c *client) Database(name string, opts ...*options.DatabaseOptions) Database {
	return wrapDatabase(c.cl.Database(name, opts...), c)
}

func (c *client) Disconnect(ctx context.Context) error {
	return c.cl.Disconnect(ctx)
}

func (c *client) ListDatabaseNames(
	ctx context.Context,
	filter interface{},
	opts ...*options.ListDatabasesOptions,
) ([]string, error) {
	return c.cl.ListDatabaseNames(ctx, filter, opts...)
}

func (c *client) ListDatabases(
	ctx context.Context,
	filter interface{},
	opts ...*options.ListDatabasesOptions,
) (mongo.ListDatabasesResult, error) {
	return c.cl.ListDatabases(ctx, filter, opts...)
}

func (c *client) NumberSessionsInProgress() int {
	return c.cl.NumberSessionsInProgress()
}

func (c *client) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return c.cl.Ping(ctx, rp)
}

func (c *client) StartSession(opts ...*options.SessionOptions) (Session, error) {
	ss, err := c.cl.StartSession(opts...)
	if err != nil {
		return nil, err
	}

	return wrapSession(ss, c), nil
}

func (c *client) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return c.cl.UseSession(ctx, fn)
}

func (c *client) UseSessionWithOptions(
	ctx context.Context,
	opts *options.SessionOptions,
	fn func(mongo.SessionContext) error,
) error {
	return c.cl.UseSessionWithOptions(ctx, opts, fn)
}

func (c *client) Watch(
	ctx context.Context,
	pipeline interface{},
	opts ...*options.ChangeStreamOptions,
) (ChangeStream, error) {
	cs, err := c.cl.Watch(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return wrapChangeStream(cs), nil
}

func (c *client) WrappedClient() *mongo.Client {
	return c.cl
}

// WrapClient an original object of `mongo.Client` type and returns the object as `Client` interface
func WrapClient(cl *mongo.Client) Client {
	return &client{cl: cl}
}

// Connect is a wrapper for `mongo.Connect` function to return the object as `Client` interface
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Connect
func Connect(ctx context.Context, opts ...*options.ClientOptions) (Client, error) {
	cl, err := mongo.Connect(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return WrapClient(cl), nil
}

// NewClient is a wrapper for `mongo.NewClient` function to return the object as `Client` interface
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#NewClient
func NewClient(opts ...*options.ClientOptions) (Client, error) {
	cl, err := mongo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return WrapClient(cl), nil
}

// WithSession is a wrapper for `mongo.WithSession` function to call then `mongo.WithSession` function
// *WARNING*: There is no simple way to wrap a SessionContext, so the original client and session will be used
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#WithSession
func WithSession(ctx context.Context, sess Session, fn func(mongo.SessionContext) error) error {
	return mongo.WithSession(ctx, sess.WrappedSession(), fn)
}
