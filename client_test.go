package mongoifc_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/sv-tools/mongoifc"
)

var testErr = errors.New("test")

func connect(tb testing.TB) mongoifc.Client {
	tb.Helper()

	uri := os.Getenv("MONGO_URI")
	require.NotEmpty(tb, uri)
	opt := options.Client().ApplyURI(uri)

	cl, err := mongoifc.Connect(context.Background(), opt)
	require.NoError(tb, err)
	tb.Cleanup(func() {
		require.NoError(tb, cl.Disconnect(context.Background()))
	})

	err = cl.Ping(context.Background(), readpref.Primary())
	require.NoError(tb, err)

	return cl
}

func TestNewClient(t *testing.T) {
	t.Parallel()

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

	opt2 := options.Client().ApplyURI("fake")
	cl2, err := mongoifc.NewClient(opt2)
	require.Error(t, err)
	require.Nil(t, cl2)
}

func TestConnect(t *testing.T) {
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

	err = cl.Ping(context.Background(), readpref.Primary())
	require.NoError(t, err)

	opt2 := options.Client().ApplyURI("fake")
	cl2, err := mongoifc.Connect(context.Background(), opt2)
	require.Error(t, err)
	require.Nil(t, cl2)
}

func TestWithSession(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	sess, err := cl.StartSession()
	require.NoError(t, err)
	t.Cleanup(func() {
		sess.EndSession(context.Background())
	})

	err = mongoifc.WithSession(context.Background(), sess, func(sc mongoifc.SessionContext) error {
		require.NotNil(t, sc.ID())
		return nil
	})
	require.NoError(t, err)
}

func TestClient_Database(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	name := fmt.Sprintf("test_%d", time.Now().Unix())
	db := cl.Database(name)
	require.NotNil(t, db)
	require.Equal(t, name, db.Name())
}

func TestClient_ListDatabaseNames(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	names, err := cl.ListDatabaseNames(context.Background(), bson.M{})
	require.NoError(t, err)
	t.Logf("database names: %v", names)
	require.NotZero(t, len(names))
	require.Contains(t, names, "admin")
}

func TestClient_ListDatabases(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	dbs, err := cl.ListDatabases(context.Background(), bson.M{})
	require.NoError(t, err)
	require.NotZero(t, dbs.TotalSize)
	require.NotZero(t, len(dbs.Databases))
}

func TestClient_NumberSessionsInProgress(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	sess, err := cl.StartSession()
	require.NoError(t, err)
	t.Cleanup(func() {
		sess.EndSession(context.Background())
	})

	require.NotZero(t, cl.NumberSessionsInProgress())
}

func TestClient_UseSession(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	err := cl.UseSession(context.Background(), func(sc mongoifc.SessionContext) error {
		require.NotNil(t, sc.ID())
		return nil
	})
	require.NoError(t, err)

	err = cl.UseSession(context.Background(), func(sc mongoifc.SessionContext) error {
		return testErr
	})
	require.ErrorIs(t, err, testErr)
}

func TestClient_UseSessionWithOptions(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	err := cl.UseSessionWithOptions(
		context.Background(),
		options.Session(),
		func(sc mongoifc.SessionContext) error {
			require.NotNil(t, sc.ID())
			return nil
		},
	)
	require.NoError(t, err)

	err = cl.UseSessionWithOptions(
		context.Background(),
		options.Session(),
		func(sc mongoifc.SessionContext) error {
			return testErr
		},
	)
	require.ErrorIs(t, err, testErr)
}

func TestClient_StartSession(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	sess, err := cl.StartSession()
	require.NoError(t, err)
	t.Cleanup(func() {
		sess.EndSession(context.Background())
	})

	cl2, err := mongoifc.NewClient()
	require.NoError(t, err)
	require.NotNil(t, cl2)

	sess2, err := cl2.StartSession()
	require.Error(t, err)
	require.Nil(t, sess2)
}

func TestWrapClient(t *testing.T) {
	t.Parallel()

	cl, err := mongo.NewClient()
	require.NoError(t, err)
	require.NotNil(t, cl)

	wcl := mongoifc.WrapClient(cl)
	require.Equal(t, cl, mongoifc.UnWrapClient(wcl))
}

func TestClient_Watch(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	cur, err := cl.Watch(context.Background(), mongo.Pipeline{})
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, cur.Close(context.Background()))
	})
	require.NotZero(t, cur.ID())
}
