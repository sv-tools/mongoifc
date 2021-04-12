package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IndexView is an interface for `mongo.IndexView` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#IndexView
type IndexView interface {
	CreateMany(ctx context.Context, models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error)
	CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error)
	DropAll(ctx context.Context, opts ...*options.DropIndexesOptions) (bson.Raw, error)
	DropOne(ctx context.Context, name string, opts ...*options.DropIndexesOptions) (bson.Raw, error)
	List(ctx context.Context, opts ...*options.ListIndexesOptions) (Cursor, error)
	ListSpecifications(ctx context.Context, opts ...*options.ListIndexesOptions) ([]*mongo.IndexSpecification, error)

	WrappedIndexView() *mongo.IndexView
}

type indexView struct {
	iv *mongo.IndexView
}

func (i *indexView) CreateMany(
	ctx context.Context,
	models []mongo.IndexModel,
	opts ...*options.CreateIndexesOptions,
) ([]string, error) {
	return i.iv.CreateMany(ctx, models, opts...)
}

func (i *indexView) CreateOne(
	ctx context.Context,
	model mongo.IndexModel,
	opts ...*options.CreateIndexesOptions,
) (string, error) {
	return i.iv.CreateOne(ctx, model, opts...)
}

func (i *indexView) DropAll(ctx context.Context, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	return i.iv.DropAll(ctx, opts...)
}

func (i *indexView) DropOne(ctx context.Context, name string, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	return i.iv.DropOne(ctx, name, opts...)
}

func (i *indexView) List(ctx context.Context, opts ...*options.ListIndexesOptions) (Cursor, error) {
	cr, err := i.iv.List(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

func (i *indexView) ListSpecifications(
	ctx context.Context,
	opts ...*options.ListIndexesOptions,
) ([]*mongo.IndexSpecification, error) {
	return i.iv.ListSpecifications(ctx, opts...)
}

func (i *indexView) WrappedIndexView() *mongo.IndexView {
	return i.iv
}

func wrapIndexView(iv *mongo.IndexView) IndexView {
	return &indexView{iv: iv}
}
