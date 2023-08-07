package simple_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/mock/gomock"

	"github.com/sv-tools/mongoifc"
	"github.com/sv-tools/mongoifc/examples/simple"
	gomockMocks "github.com/sv-tools/mongoifc/mocks/gomock"
	mockeryMocks "github.com/sv-tools/mongoifc/mocks/mockery"
)

var (
	expectedUsers = []simple.User{
		{Name: "foo", Active: true, IsAdmin: true},
		{Name: "bar", Active: true, IsAdmin: true},
	}
)

func TestUsersWorkflow(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("mockery", func(t *testing.T) {
		t.Parallel()

		cur := &mockeryMocks.Cursor{}
		defer cur.AssertExpectations(t)
		cur.On("All", ctx, mock.Anything).Run(func(args mock.Arguments) {
			users := args[1].(*[]simple.User)
			*users = append(*users, expectedUsers...)
		}).Return(nil)

		col := &mockeryMocks.Collection{}
		defer col.AssertExpectations(t)
		col.On("InsertMany", ctx, mock.Anything).Return(
			&mongo.InsertManyResult{
				InsertedIDs: []interface{}{
					primitive.NewObjectID(),
					primitive.NewObjectID(),
					primitive.NewObjectID(),
					primitive.NewObjectID(),
				},
			},
			nil,
		)
		col.On("Find", ctx, mock.AnythingOfType("User")).Return(cur, nil)
		col.On("DeleteMany", ctx, mock.AnythingOfType("primitive.M")).Return(
			&mongo.DeleteResult{
				DeletedCount: 4,
			},
			nil,
		)

		db := &mockeryMocks.Database{}
		defer db.AssertExpectations(t)
		db.On("Collection", simple.UsersCollection).Return(col)

		usersWorkflow(t, db)
	})

	t.Run("gomock", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cur := gomockMocks.NewMockCursor(ctrl)
		cur.EXPECT().All(ctx, gomock.Any()).Do(func(ctx context.Context, arg interface{}) {
			users := arg.(*[]simple.User)
			*users = append(*users, expectedUsers...)
		}).Return(nil)

		col := gomockMocks.NewMockCollection(ctrl)
		col.EXPECT().InsertMany(ctx, gomock.Any()).Return(
			&mongo.InsertManyResult{
				InsertedIDs: []interface{}{
					primitive.NewObjectID(),
					primitive.NewObjectID(),
					primitive.NewObjectID(),
					primitive.NewObjectID(),
				},
			},
			nil,
		)
		col.EXPECT().Find(ctx, gomock.Any()).Return(cur, nil)
		col.EXPECT().DeleteMany(ctx, gomock.Any()).Return(
			&mongo.DeleteResult{
				DeletedCount: 4,
			},
			nil,
		)

		db := gomockMocks.NewMockDatabase(ctrl)
		db.EXPECT().Collection(simple.UsersCollection).Return(col).AnyTimes()

		usersWorkflow(t, db)
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
		usersWorkflow(t, db)
	})
}

func usersWorkflow(t testing.TB, db mongoifc.Database) {
	ctx := context.Background()
	ids, err := simple.Create(ctx, db,
		simple.User{Name: "blocked admin", Active: false, IsAdmin: true},
		simple.User{Name: "active non-admin", Active: true, IsAdmin: false},
		expectedUsers[0],
		expectedUsers[1],
	)
	require.NoError(t, err)
	require.Len(t, ids, 4)

	users, err := simple.GetAdmins(ctx, db)
	require.NoError(t, err)
	for i, u := range users {
		u.ID = ""
		users[i] = u
	}
	require.Equal(t, expectedUsers, users)

	require.NoError(t, simple.Delete(ctx, db, ids...))
}
