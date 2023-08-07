package simple_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/mock/gomock"

	"github.com/sv-tools/mongoifc"
	"github.com/sv-tools/mongoifc/examples/simple"
	gomockMocks "github.com/sv-tools/mongoifc/mocks/gomock"
	mockeryMocks "github.com/sv-tools/mongoifc/mocks/mockery"
)

func TestCollectionsWorkflow(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("mockery", func(t *testing.T) {
		t.Parallel()

		col := &mockeryMocks.Collection{}
		defer col.AssertExpectations(t)
		col.On("Drop", ctx).Return(nil)

		db := &mockeryMocks.Database{}
		defer db.AssertExpectations(t)
		db.On("Collection", mock.Anything).Return(col)
		db.On("CreateCollection", ctx, mock.Anything).Return(nil)
		db.On("ListCollectionNames", ctx, mock.AnythingOfType("primitive.M")).Return([]string{"fake"}, nil).Once()
		db.On("ListCollectionNames", ctx, mock.AnythingOfType("primitive.M")).Return([]string{}, nil).Twice()

		collectionsWorkflow(t, db)
	})

	t.Run("gomock", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		col := gomockMocks.NewMockCollection(ctrl)
		col.EXPECT().Drop(ctx).Return(nil)

		db := gomockMocks.NewMockDatabase(ctrl)
		db.EXPECT().Collection(gomock.Any()).Return(col)
		db.EXPECT().CreateCollection(ctx, gomock.Any()).Return(nil)
		db.EXPECT().ListCollectionNames(ctx, gomock.Any()).Return([]string{"fake"}, nil)
		db.EXPECT().ListCollectionNames(ctx, gomock.Any()).Return([]string{}, nil).Times(2)

		collectionsWorkflow(t, db)
	})

	t.Run("docker", func(t *testing.T) {
		t.Parallel()

		uri := os.Getenv("MONGO_URI")
		require.NotEmpty(t, uri)

		opt := options.Client().ApplyURI(uri)
		cl, err := mongoifc.NewClient(opt)
		require.NoError(t, err)
		require.NotNil(t, cl)

		err = cl.Connect(ctx)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, cl.Disconnect(ctx))
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

func collectionsWorkflow(t testing.TB, db mongoifc.Database) {
	ctx := context.Background()
	name := randSeq(42)

	require.NoError(t, simple.CreateCollection(ctx, db, name))

	res, err := simple.CollectionExists(ctx, db, name)
	require.NoError(t, err)
	require.True(t, res)

	res, err = simple.CollectionExists(ctx, db, name+"42")
	require.NoError(t, err)
	require.False(t, res)

	require.NoError(t, simple.DropCollection(ctx, db, name))

	res, err = simple.CollectionExists(ctx, db, name)
	require.NoError(t, err)
	require.False(t, res)
}
