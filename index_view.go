package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// IndexView is an interface for `mongo.IndexView` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#IndexView
type IndexView interface {
	CreateMany(
		ctx context.Context,
		models []mongo.IndexModel,
		opts ...options.Lister[options.CreateIndexesOptions],
	) ([]string, error)
	CreateOne(
		ctx context.Context,
		model mongo.IndexModel,
		opts ...options.Lister[options.CreateIndexesOptions],
	) (string, error)
	DropAll(ctx context.Context, opts ...options.Lister[options.DropIndexesOptions]) error
	DropOne(ctx context.Context, name string, opts ...options.Lister[options.DropIndexesOptions]) error
	DropWithKey(ctx context.Context, keySpecDocument any, opts ...options.Lister[options.DropIndexesOptions]) error
	List(ctx context.Context, opts ...options.Lister[options.ListIndexesOptions]) (Cursor, error)
	ListSpecifications(
		ctx context.Context,
		opts ...options.Lister[options.ListIndexesOptions],
	) ([]mongo.IndexSpecification, error)
}

type indexView struct {
	iv *mongo.IndexView
}

// CreateMany is a wrapper for `mongo.IndexView.CreateMany` method
func (i *indexView) CreateMany(
	ctx context.Context,
	models []mongo.IndexModel,
	opts ...options.Lister[options.CreateIndexesOptions],
) ([]string, error) {
	return i.iv.CreateMany(ctx, models, opts...)
}

// CreateOne is a wrapper for `mongo.IndexView.CreateOne` method
func (i *indexView) CreateOne(
	ctx context.Context,
	model mongo.IndexModel,
	opts ...options.Lister[options.CreateIndexesOptions],
) (string, error) {
	return i.iv.CreateOne(ctx, model, opts...)
}

// DropAll is a wrapper for `mongo.IndexView.DropAll` method
func (i *indexView) DropAll(
	ctx context.Context,
	opts ...options.Lister[options.DropIndexesOptions],
) error {
	return i.iv.DropAll(ctx, opts...)
}

// DropOne is a wrapper for `mongo.IndexView.DropOne` method
func (i *indexView) DropOne(
	ctx context.Context,
	name string,
	opts ...options.Lister[options.DropIndexesOptions],
) error {
	return i.iv.DropOne(ctx, name, opts...)
}

// DropWithKey is a wrapper for `mongo.IndexView.DropWithKey` method
func (i *indexView) DropWithKey(
	ctx context.Context,
	keySpecDocument any,
	opts ...options.Lister[options.DropIndexesOptions],
) error {
	return i.iv.DropWithKey(ctx, keySpecDocument, opts...)
}

// List is a wrapper for `mongo.IndexView.List` method
func (i *indexView) List(
	ctx context.Context,
	opts ...options.Lister[options.ListIndexesOptions],
) (Cursor, error) {
	cr, err := i.iv.List(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

// ListSpecifications is a wrapper for `mongo.IndexView.ListSpecifications` method
func (i *indexView) ListSpecifications(
	ctx context.Context,
	opts ...options.Lister[options.ListIndexesOptions],
) ([]mongo.IndexSpecification, error) {
	return i.iv.ListSpecifications(ctx, opts...)
}

func wrapIndexView(iv *mongo.IndexView) IndexView {
	return &indexView{iv: iv}
}
