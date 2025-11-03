package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// SearchIndexView is an interface for `mongo.SearchIndexView` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#SearchIndexView
type SearchIndexView interface {
	CreateMany(
		ctx context.Context,
		models []mongo.SearchIndexModel,
		opts ...options.Lister[options.CreateSearchIndexesOptions],
	) ([]string, error)
	CreateOne(
		ctx context.Context,
		model mongo.SearchIndexModel,
		opts ...options.Lister[options.CreateSearchIndexesOptions],
	) (string, error)
	DropOne(
		ctx context.Context,
		name string,
		opts ...options.Lister[options.DropSearchIndexOptions],
	) error
	List(
		ctx context.Context,
		searchIdxOpts options.Lister[options.SearchIndexesOptions],
		opts ...options.Lister[options.ListSearchIndexesOptions],
	) (Cursor, error)
	UpdateOne(
		ctx context.Context,
		name string,
		definition any,
		opts ...options.Lister[options.UpdateSearchIndexOptions],
	) error
}

type searchIndexView struct {
	siv *mongo.SearchIndexView
}

// CreateMany is a wrapper for `mongo.SearchIndexView.CreateMany` method
func (i *searchIndexView) CreateMany(
	ctx context.Context,
	models []mongo.SearchIndexModel,
	opts ...options.Lister[options.CreateSearchIndexesOptions],
) ([]string, error) {
	return i.siv.CreateMany(ctx, models, opts...)
}

// CreateOne is a wrapper for `mongo.SearchIndexView.CreateOne` method
func (i *searchIndexView) CreateOne(
	ctx context.Context,
	model mongo.SearchIndexModel,
	opts ...options.Lister[options.CreateSearchIndexesOptions],
) (string, error) {
	return i.siv.CreateOne(ctx, model, opts...)
}

// DropOne is a wrapper for `mongo.SearchIndexView.DropOne` method
func (i *searchIndexView) DropOne(
	ctx context.Context, name string,
	opts ...options.Lister[options.DropSearchIndexOptions],
) error {
	return i.siv.DropOne(ctx, name, opts...)
}

// List is a wrapper for `mongo.SearchIndexView.List` method
func (i *searchIndexView) List(
	ctx context.Context,
	searchIdxOpts options.Lister[options.SearchIndexesOptions],
	opts ...options.Lister[options.ListSearchIndexesOptions],
) (Cursor, error) {
	cr, err := i.siv.List(ctx, searchIdxOpts, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

// UpdateOne is a wrapper for `mongo.SearchIndexView.UpdateOne` method
func (i *searchIndexView) UpdateOne(
	ctx context.Context,
	name string,
	definition any,
	opts ...options.Lister[options.UpdateSearchIndexOptions],
) error {
	return i.siv.UpdateOne(ctx, name, definition, opts...)
}

func wrapSearchIndexView(siv *mongo.SearchIndexView) SearchIndexView {
	return &searchIndexView{siv: siv}
}
