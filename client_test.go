package mongoifc_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/sv-tools/mongoifc"
)

func TestNewClient(t *testing.T) {
	uri := os.Getenv("MONGO_URI")
	require.NotEmpty(t, uri)

	opt := options.Client().ApplyURI(uri)
	cl, err := mongoifc.NewClient(opt)
	require.NoError(t, err)
	require.NotNil(t, cl)

	err = cl.Connect(context.Background())
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, cl.Disconnect(context.Background()))
	})

	err = cl.Ping(context.Background(), readpref.Primary())
	require.NoError(t, err)
}

func TestConnect(t *testing.T) {
	uri := os.Getenv("MONGO_URI")
	require.NotEmpty(t, uri)

	opt := options.Client().ApplyURI(uri)
	cl, err := mongoifc.Connect(context.Background(), opt)
	require.NoError(t, err)
	require.NotNil(t, cl)

	t.Cleanup(func() {
		require.NoError(t, cl.Disconnect(context.Background()))
	})

	err = cl.Ping(context.Background(), readpref.Primary())
	require.NoError(t, err)
}

func TestWithSession(t *testing.T) {
	uri := os.Getenv("MONGO_URI")
	require.NotEmpty(t, uri)

	opt := options.Client().ApplyURI(uri)
	cl, err := mongoifc.Connect(context.Background(), opt)
	require.NoError(t, err)
	require.NotNil(t, cl)

	sess, err := cl.StartSession()
	require.NoError(t, err)

	err = mongoifc.WithSession(context.Background(), sess, func(sessionContext mongo.SessionContext) error {
		require.NotNil(t, sessionContext.ID())
		return nil
	})
	require.NoError(t, err)
}
