package mongoifc_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sv-tools/mongoifc"
)

func TestSession_WithTransaction(t *testing.T) {
	t.Parallel()

	uri := os.Getenv("MONGO_URI")
	require.NotEmpty(t, uri)

	opt := options.Client().ApplyURI(uri)
	cl, err := mongoifc.Connect(context.Background(), opt)
	require.NoError(t, err)
	require.NotNil(t, cl)

	t.Cleanup(func() {
		require.NoError(t, cl.Disconnect(context.Background()))
	})

	sess, err := cl.StartSession()
	require.NoError(t, err)
	t.Cleanup(func() {
		sess.EndSession(context.Background())
	})
	res, err := sess.WithTransaction(context.Background(), func(sc mongoifc.SessionContext) (interface{}, error) {
		return cl.
			Database(fmt.Sprintf("test_%d", time.Now().Unix())).
			Collection("test").
			InsertOne(sc, bson.M{"foo": "bar"})
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "Transaction numbers are only allowed on a replica set member or mongos")
	require.Nil(t, res)
}

func TestSession_StartTransaction(t *testing.T) {
	t.Parallel()

	uri := os.Getenv("MONGO_URI")
	require.NotEmpty(t, uri)

	opt := options.Client().ApplyURI(uri)
	cl, err := mongoifc.Connect(context.Background(), opt)
	require.NoError(t, err)
	require.NotNil(t, cl)

	t.Cleanup(func() {
		require.NoError(t, cl.Disconnect(context.Background()))
	})

	err = cl.UseSession(context.Background(), func(sc mongoifc.SessionContext) error {
		err = sc.StartTransaction()
		require.NoError(t, err)

		res, err := cl.
			Database(fmt.Sprintf("test_%d", time.Now().Unix())).
			Collection("test").
			InsertOne(sc, bson.M{"foo": "bar"})
		require.Nil(t, res)
		return err
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "Transaction numbers are only allowed on a replica set member or mongos")
}
