package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

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
