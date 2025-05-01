package simple_test

import (
	"context"
	"fmt"
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

var expectedUsers = []simple.User{
	{Name: "foo", Active: true, IsAdmin: true},
	{Name: "bar", Active: true, IsAdmin: true},
}

func TestUsersWorkflow(t *testing.T) {
	t.Parallel()

	t.Run("mockery", func(t *testing.T) {
		t.Parallel()

		cur := &mockeryMocks.Cursor{}
		defer cur.AssertExpectations(t)
		cur.On("All", t.Context(), mock.Anything).Run(func(args mock.Arguments) {
			users := args[1].(*[]simple.User)
			*users = append(*users, expectedUsers...)
		}).Return(nil)

		col := &mockeryMocks.Collection{}
		defer col.AssertExpectations(t)
		col.On("InsertMany", t.Context(), mock.Anything).Return(
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
		col.On("Find", t.Context(), mock.AnythingOfType("User")).Return(cur, nil)
		col.On("DeleteMany", t.Context(), mock.AnythingOfType("primitive.M")).Return(
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
		cur.EXPECT().All(t.Context(), gomock.Any()).Do(func(ctx context.Context, arg interface{}) {
			users := arg.(*[]simple.User)
			*users = append(*users, expectedUsers...)
		}).Return(nil)

		col := gomockMocks.NewMockCollection(ctrl)
		col.EXPECT().InsertMany(t.Context(), gomock.Any()).Return(
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
		col.EXPECT().Find(t.Context(), gomock.Any()).Return(cur, nil)
		col.EXPECT().DeleteMany(t.Context(), gomock.Any()).Return(
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

		opt := options.Client().ApplyURI(MongoUri)
		cl, err := mongoifc.Connect(t.Context(), opt)
		require.NoError(t, err)
		require.NotNil(t, cl)
		t.Cleanup(func() {
			require.NoError(t, cl.Disconnect(context.Background())) //nolint:usetesting
		})

		db := cl.Database(fmt.Sprintf("simple_%d", time.Now().Unix()))
		usersWorkflow(t, db)
	})
}

func usersWorkflow(tb testing.TB, db mongoifc.Database) {
	tb.Helper()

	ids, err := simple.Create(tb.Context(), db,
		simple.User{Name: "blocked admin", Active: false, IsAdmin: true},
		simple.User{Name: "active non-admin", Active: true, IsAdmin: false},
		expectedUsers[0],
		expectedUsers[1],
	)
	require.NoError(tb, err)
	require.Len(tb, ids, 4)

	users, err := simple.GetAdmins(tb.Context(), db)
	require.NoError(tb, err)
	for i, u := range users {
		u.ID = ""
		users[i] = u
	}
	require.Equal(tb, expectedUsers, users)

	require.NoError(tb, simple.Delete(tb.Context(), db, ids...))
}
