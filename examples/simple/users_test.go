package simple_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sv-tools/mongoifc"
	"github.com/sv-tools/mongoifc/examples/simple"
	gomockMocks "github.com/sv-tools/mongoifc/mocks/gomock"
	mockeryMocks "github.com/sv-tools/mongoifc/mocks/mockery"
)

func TestGetAdmins(t *testing.T) {
	t.Parallel()

	expectedUsers := []*simple.User{
		{Name: "foo", Active: true, IsAdmin: true},
		{Name: "bar", Active: true, IsAdmin: true},
	}
	ctx := context.Background()

	t.Run("mockery", func(t *testing.T) {
		t.Parallel()

		cur := &mockeryMocks.Cursor{}
		cur.On("All", ctx, mock.Anything).Run(func(args mock.Arguments) {
			users := args[1].(*[]*simple.User)
			*users = append(*users, expectedUsers...)
		}).Return(nil)

		col := &mockeryMocks.Collection{}
		col.On("Find", ctx, mock.AnythingOfType("User")).Return(cur, nil)

		db := &mockeryMocks.Database{}
		db.On("Collection", simple.UsersCollection).Return(col)

		users, err := simple.GetAdmins(ctx, db)
		require.NoError(t, err)
		require.Equal(t, expectedUsers, users)
	})

	t.Run("gomock", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cur := gomockMocks.NewMockCursor(ctrl)
		cur.EXPECT().All(ctx, gomock.Any()).Do(func(ctx context.Context, arg interface{}) {
			users := arg.(*[]*simple.User)
			*users = append(*users, expectedUsers...)
		}).Return(nil)

		col := gomockMocks.NewMockCollection(ctrl)
		col.EXPECT().Find(ctx, gomock.Any()).Return(cur, nil)

		db := gomockMocks.NewMockDatabase(ctrl)
		db.EXPECT().Collection(simple.UsersCollection).Return(col)

		users, err := simple.GetAdmins(ctx, db)
		require.NoError(t, err)
		require.Equal(t, expectedUsers, users)
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
		res, err := db.Collection(simple.UsersCollection).InsertMany(ctx, []interface{}{
			&simple.User{Name: "blocked admin", Active: false, IsAdmin: true},
			&simple.User{Name: "active non-admin", Active: true, IsAdmin: false},
			expectedUsers[0],
			expectedUsers[1],
		})
		require.NoError(t, err)
		require.Len(t, res.InsertedIDs, 4)

		users, err := simple.GetAdmins(ctx, db)
		require.NoError(t, err)
		for _, u := range users {
			u.ID = ""
		}
		require.Equal(t, expectedUsers, users)
	})
}
