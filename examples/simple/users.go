package simple

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sv-tools/mongoifc"
)

const (
	UsersCollection = "users"
)

type User struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Active bool `json:"active,omitempty" bson:"active,omitempty"`
	IsAdmin bool `json:"is_admin,omitempty" bson:"is_admin,omitempty"`
}

func GetAdmins(ctx context.Context, db mongoifc.Database) ([]*User, error) {
	var users []*User
	cur, err := db.Collection(UsersCollection).Find(ctx, User{
		Active: true,
		IsAdmin: true,
	})
	if err != nil {
		return nil, err
	}
	if err := cur.Decode(&users); err != nil {
		return nil, err
	}
	return users, err
}
