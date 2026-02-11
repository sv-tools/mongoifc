package simple_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/mock/gomock"

	"github.com/sv-tools/mongoifc/v2"
	"github.com/sv-tools/mongoifc/v2/examples/simple"
	gomockMocks "github.com/sv-tools/mongoifc/v2/mocks/gomock"
	mockeryMocks "github.com/sv-tools/mongoifc/v2/mocks/mockery"
)

func TestCollectionsWorkflow(t *testing.T) {
	t.Parallel()

	t.Run("mockery", func(t *testing.T) {
		t.Parallel()

		col := &mockeryMocks.Collection{}
		defer col.AssertExpectations(t)
		col.On("Drop", t.Context(), mock.Anything).Return(nil)

		db := &mockeryMocks.Database{}
		defer db.AssertExpectations(t)
		db.On("Collection", mock.Anything).Return(col)
		db.On("CreateCollection", t.Context(), mock.Anything).Return(nil)
		db.On("ListCollectionNames", t.Context(), mock.AnythingOfType("bson.M")).
			Return([]string{"fake"}, nil).
			Once()
		db.On("ListCollectionNames", t.Context(), mock.AnythingOfType("bson.M")).
			Return([]string{}, nil).
			Twice()

		collectionsWorkflow(t, db)
	})

	t.Run("gomock", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		col := gomockMocks.NewMockCollection(ctrl)
		col.EXPECT().Drop(t.Context(), gomock.Any()).Return(nil)

		db := gomockMocks.NewMockDatabase(ctrl)
		db.EXPECT().Collection(gomock.Any()).Return(col)
		db.EXPECT().CreateCollection(t.Context(), gomock.Any()).Return(nil)
		db.EXPECT().ListCollectionNames(t.Context(), gomock.Any()).Return([]string{"fake"}, nil)
		db.EXPECT().ListCollectionNames(t.Context(), gomock.Any()).Return([]string{}, nil).Times(2)

		collectionsWorkflow(t, db)
	})

	t.Run("docker", func(t *testing.T) {
		t.Parallel()

		opt := options.Client().ApplyURI(MongoUri)
		cl, err := mongoifc.Connect(opt)
		require.NoError(t, err)
		require.NotNil(t, cl)
		t.Cleanup(func() {
			require.NoError(t, cl.Disconnect(context.Background()))
		})

		db := cl.Database(fmt.Sprintf("simple_%d", time.Now().Unix()))
		collectionsWorkflow(t, db)
	})
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func collectionsWorkflow(tb testing.TB, db mongoifc.Database) {
	tb.Helper()

	name := randSeq(42)

	require.NoError(tb, simple.CreateCollection(tb.Context(), db, name))

	res, err := simple.CollectionExists(tb.Context(), db, name)
	require.NoError(tb, err)
	require.True(tb, res)

	res, err = simple.CollectionExists(tb.Context(), db, name+"42")
	require.NoError(tb, err)
	require.False(tb, res)

	require.NoError(tb, simple.DropCollection(tb.Context(), db, name))

	res, err = simple.CollectionExists(tb.Context(), db, name)
	require.NoError(tb, err)
	require.False(tb, res)
}
