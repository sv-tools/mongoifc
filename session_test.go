package mongoifc_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sv-tools/mongoifc"
)

func TestSession_WithTransaction(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	sess, err := cl.StartSession()
	require.NoError(t, err)
	t.Cleanup(func() {
		sess.EndSession(context.Background())
	})
	name := fmt.Sprintf("test_%d", time.Now().Unix())
	res, err := sess.WithTransaction(context.Background(), func(sc mongoifc.SessionContext) (interface{}, error) {
		return cl.
			Database(name).
			Collection("test-with").
			InsertOne(sc, bson.M{"foo": "bar"})
	})
	require.NoError(t, err)
	require.NotNil(t, res.(*mongo.InsertOneResult).InsertedID)

	n, err := cl.
		Database(name).
		Collection("test-with").
		CountDocuments(context.Background(), bson.M{"foo": "bar"})
	require.NoError(t, err)
	require.Equal(t, int64(1), n)
}

func TestSession_StartAndAbortTransaction(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	name := fmt.Sprintf("test_%d", time.Now().Unix())
	err := cl.UseSession(context.Background(), func(sc mongoifc.SessionContext) error {
		err := sc.StartTransaction()
		require.NoError(t, err)

		res, err := cl.
			Database(name).
			Collection("test-start-abort").
			InsertOne(sc, bson.M{"foo": "bar"})
		require.NoError(t, err)
		require.NotNil(t, res.InsertedID)

		return sc.AbortTransaction(sc)
	})
	require.NoError(t, err)

	n, err := cl.
		Database(name).
		Collection("test-start-abort").
		CountDocuments(context.Background(), bson.M{"foo": "bar"})
	require.NoError(t, err)
	require.Zero(t, n)
}
