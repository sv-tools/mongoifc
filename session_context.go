package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// SessionContext is an interface emulates `mongo.SessionContext`
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#SessionContext
type SessionContext interface {
	context.Context
	Session
}

type sessionContext struct {
	context.Context
	Session
}

func wrapSessionContext(sc mongo.SessionContext, cl *client) SessionContext {
	return &sessionContext{
		Context: sc,
		Session: wrapSession(sc, cl),
	}
}

func wrapFn1(
	fn func(sc SessionContext) error,
	cl *client,
) func(sc mongo.SessionContext) error {
	return func(sc mongo.SessionContext) error {
		return fn(wrapSessionContext(sc, cl))
	}
}

func wrapFn2(
	fn func(sc SessionContext) (interface{}, error),
	cl *client,
) func(sc mongo.SessionContext) (interface{}, error) {
	return func(sc mongo.SessionContext) (interface{}, error) {
		return fn(wrapSessionContext(sc, cl))
	}
}

// NewSessionContext is wrapper for `mongo.NewSessionContext`
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#NewSessionContext
func NewSessionContext(ctx context.Context, sess Session) SessionContext {
	ms := mongo.NewSessionContext(ctx, UnWrapSession(sess))
	return wrapSessionContext(ms, sess.Client().(*client))
}

// SessionFromContext for `mongo.SessionFromContext`
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#SessionFromContext
func SessionFromContext(ctx context.Context) Session {
	ms := mongo.SessionFromContext(ctx)
	return WrapSession(ms)
}
