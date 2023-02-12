package simple

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sv-tools/mongoifc"
)

const (
	UsersCollection = "users"
)

type User struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	Email   string `json:"email,omitempty" bson:"email,omitempty"`
	Active  bool   `json:"active,omitempty" bson:"active,omitempty"`
	IsAdmin bool   `json:"is_admin,omitempty" bson:"is_admin,omitempty"`
}

func GetAdmins(ctx context.Context, db mongoifc.Database) ([]User, error) {
	var users []User
	cur, err := db.Collection(UsersCollection).Find(ctx, User{
		Active:  true,
		IsAdmin: true,
	})
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, err
}

func Create(ctx context.Context, db mongoifc.Database, users ...User) ([]string, error) {
	documents := make([]interface{}, len(users))
	for i := 0; i < len(users); i++ {
		documents[i] = users[i]
	}
	res, err := db.Collection(UsersCollection).InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}
	ids := make([]string, len(res.InsertedIDs))
	for i := 0; i < len(res.InsertedIDs); i++ {
		ids[i] = res.InsertedIDs[i].(primitive.ObjectID).Hex()
	}
	return ids, nil
}

func Delete(ctx context.Context, db mongoifc.Database, ids ...string) error {
	documents := make([]primitive.ObjectID, len(ids))
	for i := 0; i < len(ids); i++ {
		id, err := primitive.ObjectIDFromHex(ids[i])
		if err != nil {
			return fmt.Errorf("%s: %w", ids[i], err)
		}
		documents[i] = id
	}

	filter := bson.M{"_id": bson.M{"$in": documents}}
	_, err := db.Collection(UsersCollection).DeleteMany(ctx, filter)
	return err
}
