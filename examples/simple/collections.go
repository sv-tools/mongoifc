package simple

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sv-tools/mongoifc/v2"
)

func CreateCollection(ctx context.Context, db mongoifc.Database, name string) error {
	return db.CreateCollection(ctx, name)
}

func DropCollection(ctx context.Context, db mongoifc.Database, name string) error {
	return db.Collection(name).Drop(ctx)
}

func CollectionExists(ctx context.Context, db mongoifc.Database, name string) (bool, error) {
	res, err := db.ListCollectionNames(ctx, bson.M{"name": name})
	if err != nil {
		return false, err
	}
	return len(res) == 1, nil
}
