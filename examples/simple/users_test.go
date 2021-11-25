package simple

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	gomockMocks "github.com/sv-tools/mongoifc/mocks/gomock"
	mockeryMocks "github.com/sv-tools/mongoifc/mocks/mockery"
)

func TestGetAdmins(t *testing.T) {
	expectedUsers := []*User{
		{Name: "foo", Active: true, IsAdmin: true},
		{Name: "bar", Active: true, IsAdmin: true},
	}
	ctx := context.Background()

	t.Run("mockery", func(t *testing.T) {
		cur := &mockeryMocks.Cursor{}
		cur.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			users := args[0].(*[]*User)
			*users = append(*users, expectedUsers...)
		}).Return(nil)

		col := &mockeryMocks.Collection{}
		col.On("Find", ctx, mock.AnythingOfType("User")).Return(cur, nil)

		db := &mockeryMocks.Database{}
		db.On("Collection", UsersCollection).Return(col)

		users, err := GetAdmins(ctx, db)
		require.NoError(t, err)
		require.Equal(t, expectedUsers, users)
	})

	t.Run("gomock", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cur := gomockMocks.NewMockCursor(ctrl)
		cur.EXPECT().Decode(gomock.Any()).Do(func(arg interface{}) {
			users := arg.(*[]*User)
			*users = append(*users, expectedUsers...)
		}).Return(nil)

		col := gomockMocks.NewMockCollection(ctrl)
		col.EXPECT().Find(ctx, gomock.Any()).Return(cur, nil)

		db := gomockMocks.NewMockDatabase(ctrl)
		db.EXPECT().Collection(UsersCollection).Return(col)

		users, err := GetAdmins(ctx, db)
		require.NoError(t, err)
		require.Equal(t, expectedUsers, users)
	})
}
