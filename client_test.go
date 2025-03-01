package mongoifc_test

import (
	"context"
	"errors"
	"fmt"
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

	require.NotEmpty(tb, MongoUri)
	opt := options.Client().ApplyURI(MongoUri)

	cl, err := mongoifc.Connect(tb.Context(), opt)
	require.NoError(tb, err)
	tb.Cleanup(func() {
		require.NoError(tb, cl.Disconnect(context.WithoutCancel(tb.Context())))
	})

	err = cl.Ping(tb.Context(), readpref.Primary())
	require.NoError(tb, err)

	return cl
}

func TestNewClient(t *testing.T) {
	t.Parallel()

	opt := options.Client().ApplyURI(MongoUri)
	cl, err := mongoifc.NewClient(opt)
	require.NoError(t, err)
	require.NotNil(t, cl)

	err = cl.Connect(t.Context())
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, cl.Disconnect(context.WithoutCancel(t.Context())))
	})

	opt2 := options.Client().ApplyURI("fake")
	cl2, err := mongoifc.NewClient(opt2)
	require.Error(t, err)
	require.Nil(t, cl2)
}

func TestConnect(t *testing.T) {
	t.Parallel()

	opt := options.Client().ApplyURI(MongoUri)
	cl, err := mongoifc.Connect(t.Context(), opt)
	require.NoError(t, err)
	require.NotNil(t, cl)

	t.Cleanup(func() {
		require.NoError(t, cl.Disconnect(context.WithoutCancel(t.Context())))
	})

	err = cl.Ping(t.Context(), readpref.Primary())
	require.NoError(t, err)

	opt2 := options.Client().ApplyURI("fake")
	cl2, err := mongoifc.Connect(t.Context(), opt2)
	require.Error(t, err)
	require.Nil(t, cl2)
}

func TestWithSession(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	sess, err := cl.StartSession()
	require.NoError(t, err)
	t.Cleanup(func() {
		sess.EndSession(context.WithoutCancel(t.Context()))
	})

	err = mongoifc.WithSession(t.Context(), sess, func(sc mongoifc.SessionContext) error {
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
	names, err := cl.ListDatabaseNames(t.Context(), bson.M{})
	require.NoError(t, err)
	t.Logf("database names: %v", names)
	require.NotEmpty(t, names)
	require.Contains(t, names, "admin")
}

func TestClient_ListDatabases(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	dbs, err := cl.ListDatabases(t.Context(), bson.M{})
	require.NoError(t, err)
	require.NotZero(t, dbs.TotalSize)
	require.NotEmpty(t, dbs.Databases)
}

func TestClient_NumberSessionsInProgress(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	sess, err := cl.StartSession()
	require.NoError(t, err)
	t.Cleanup(func() {
		sess.EndSession(context.WithoutCancel(t.Context()))
	})

	require.NotZero(t, cl.NumberSessionsInProgress())
}

func TestClient_UseSession(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	err := cl.UseSession(t.Context(), func(sc mongoifc.SessionContext) error {
		require.NotNil(t, sc.ID())
		return nil
	})
	require.NoError(t, err)

	err = cl.UseSession(t.Context(), func(sc mongoifc.SessionContext) error {
		return testErr
	})
	require.ErrorIs(t, err, testErr)
}

func TestClient_UseSessionWithOptions(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	err := cl.UseSessionWithOptions(
		t.Context(),
		options.Session(),
		func(sc mongoifc.SessionContext) error {
			require.NotNil(t, sc.ID())
			return nil
		},
	)
	require.NoError(t, err)

	err = cl.UseSessionWithOptions(
		t.Context(),
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
		sess.EndSession(context.WithoutCancel(t.Context()))
	})
}

func TestWrapClient_UnWrapClient(t *testing.T) {
	t.Parallel()

	cl, err := mongo.Connect(t.Context())
	require.NoError(t, err)
	require.NotNil(t, cl)

	wcl := mongoifc.WrapClient(cl)
	require.Equal(t, cl, mongoifc.UnWrapClient(wcl))
}

func TestClient_Watch(t *testing.T) {
	t.Parallel()

	cl := connect(t)
	cur, err := cl.Watch(t.Context(), mongo.Pipeline{})
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, cur.Close(context.WithoutCancel(t.Context())))
	})
	require.NotZero(t, cur.ID())
}
