package mongoifc_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/sv-tools/mongoifc/v2"
)

func TestSession_WithTransaction(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	sess, err := cl.StartSession()
	require.NoError(t, err)
	t.Cleanup(func() {
		sess.EndSession(context.WithoutCancel(t.Context()))
	})
	name := fmt.Sprintf("test_%d", time.Now().Unix())
	res, err := sess.WithTransaction(t.Context(), func(ctx context.Context) (any, error) {
		return cl.
			Database(name).
			Collection("test-with").
			InsertOne(ctx, bson.M{"foo": "bar"})
	})
	require.NoError(t, err)
	require.NotNil(t, res.(*mongo.InsertOneResult).InsertedID)

	n, err := cl.
		Database(name).
		Collection("test-with").
		CountDocuments(t.Context(), bson.M{"foo": "bar"})
	require.NoError(t, err)
	require.Equal(t, int64(1), n)
}

func TestSession_StartAndAbortTransaction(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	name := fmt.Sprintf("test_%d", time.Now().Unix())
	err := cl.UseSession(t.Context(), func(ctx context.Context) error {
		sc := mongo.SessionFromContext(ctx)
		err := sc.StartTransaction()
		require.NoError(t, err)

		res, err := cl.
			Database(name).
			Collection("test-start-abort").
			InsertOne(ctx, bson.M{"foo": "bar"})
		require.NoError(t, err)
		require.NotNil(t, res.InsertedID)

		return sc.AbortTransaction(ctx)
	})
	require.NoError(t, err)

	n, err := cl.
		Database(name).
		Collection("test-start-abort").
		CountDocuments(t.Context(), bson.M{"foo": "bar"})
	require.NoError(t, err)
	require.Zero(t, n)
}

func TestWrapSession_UnWrapSession(t *testing.T) {
	t.Parallel()
	cl := connect(t)
	mcl := mongoifc.UnWrapClient(cl)
	orig, err := mcl.StartSession()
	require.NoError(t, err)
	wrapped := mongoifc.WrapSession(orig)
	require.Equal(t, orig, mongoifc.UnWrapSession(wrapped))
}
